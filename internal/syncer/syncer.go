package syncer

import (
	"fmt"
	"io/fs"
	"os"
	pathpkg "path"
	"path/filepath"
	"sort"
	"strings"

	"github.com/openkit-devtools/openkit/internal/managedstate"
)

type DesiredFile struct {
	OutputPath string
	Bytes      []byte
	ArtifactID string
	Mode       string
}

type Action string

const (
	ActionCreate    Action = "create"
	ActionUpdate    Action = "update"
	ActionOverwrite Action = "overwrite"
	ActionSkip      Action = "skip"
	ActionConflict  Action = "conflict"
	ActionDelete    Action = "delete"
)

type PlanEntry struct {
	Action     Action
	Path       string
	Reason     string
	ArtifactID string
}

type Plan struct {
	Entries []PlanEntry

	Create    []string
	Update    []string
	Overwrite []string
	Skip      []string
	Conflicts []string
	Delete    []string
	Orphaned  []string
}

type Options struct {
	DryRun    bool
	Overwrite bool
	Prune     bool
}

type ApplyResult struct {
	Plan       Plan
	BackupsDir string
}

func DesiredFromEmbeddedBase(base fs.FS, agentFolder string) ([]DesiredFile, error) {
	return DesiredFromEmbeddedSubdir(base, "base", agentFolder)
}

func DesiredFromEmbeddedSubdir(efs fs.FS, embeddedDir string, outputPrefix string) ([]DesiredFile, error) {
	var out []DesiredFile

	embeddedDir = pathpkg.Clean(strings.TrimSuffix(strings.TrimSpace(embeddedDir), "/"))
	if embeddedDir == "." || embeddedDir == "" {
		return nil, fmt.Errorf("invalid embeddedDir")
	}

	err := fs.WalkDir(efs, embeddedDir, func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		rel := strings.TrimPrefix(p, embeddedDir+"/")
		rel = pathpkg.Clean(rel)
		b, err := fs.ReadFile(efs, p)
		if err != nil {
			return err
		}
		output := filepath.ToSlash(filepath.Join(outputPrefix, filepath.FromSlash(rel)))
		out = append(out, DesiredFile{
			OutputPath: output,
			Bytes:      b,
			ArtifactID: "embedded/" + embeddedDir + "/" + rel,
			Mode:       "copy",
		})
		return nil
	})
	if err != nil {
		return nil, err
	}

	sort.Slice(out, func(i, j int) bool { return out[i].OutputPath < out[j].OutputPath })
	return out, nil
}

func DesiredRootFile(root fs.FS, rootPath string, outputPath string, artifactID string) (DesiredFile, error) {
	b, err := fs.ReadFile(root, filepath.ToSlash(rootPath))
	if err != nil {
		return DesiredFile{}, err
	}
	return DesiredFile{OutputPath: filepath.ToSlash(outputPath), Bytes: b, ArtifactID: artifactID, Mode: "copy"}, nil
}

func NormalizeRelOutputPath(p string) (string, error) {
	p = filepath.ToSlash(filepath.Clean(strings.TrimSpace(p)))
	p = strings.TrimPrefix(p, "./")
	if p == "." || p == "" {
		return "", fmt.Errorf("invalid output path")
	}
	if strings.Contains(p, "..") {
		return "", fmt.Errorf("output path contains '..'")
	}
	if strings.HasPrefix(p, "/") {
		return "", fmt.Errorf("output path must be relative")
	}
	return p, nil
}

func SafeAbsPath(projectRootAbs string, relOutputPath string) (string, error) {
	rel, err := NormalizeRelOutputPath(relOutputPath)
	if err != nil {
		return "", err
	}
	rootAbs, err := filepath.Abs(projectRootAbs)
	if err != nil {
		return "", err
	}
	targetAbs := filepath.Join(rootAbs, filepath.FromSlash(rel))
	targetAbs, err = filepath.Abs(targetAbs)
	if err != nil {
		return "", err
	}

	rootClean := filepath.Clean(rootAbs)
	targetClean := filepath.Clean(targetAbs)
	sep := string(os.PathSeparator)
	if targetClean != rootClean && !strings.HasPrefix(targetClean, rootClean+sep) {
		return "", fmt.Errorf("refusing to write outside project root")
	}
	return targetAbs, nil
}

func BuildPlan(projectRootAbs string, agentID string, desired []DesiredFile, st *managedstate.State, opts Options) (Plan, error) {
	var plan Plan
	if st == nil {
		st = &managedstate.State{SchemaVersion: managedstate.SchemaVersion}
	}
	ag := managedstate.EnsureAgent(st, agentID)

	desiredSet := map[string]DesiredFile{}
	for _, df := range desired {
		norm, err := NormalizeRelOutputPath(df.OutputPath)
		if err != nil {
			return Plan{}, fmt.Errorf("desired output path %q: %w", df.OutputPath, err)
		}
		df.OutputPath = norm
		desiredSet[df.OutputPath] = df
	}

	for path, df := range desiredSet {
		abs, err := SafeAbsPath(projectRootAbs, path)
		if err != nil {
			return Plan{}, fmt.Errorf("path %q: %w", path, err)
		}

		_, statErr := os.Stat(abs)
		entry := ag.Files[path]
		if statErr != nil {
			if os.IsNotExist(statErr) {
				plan.Entries = append(plan.Entries, PlanEntry{Action: ActionCreate, Path: path, Reason: "missing", ArtifactID: df.ArtifactID})
				plan.Create = append(plan.Create, path)
				continue
			}
			return Plan{}, statErr
		}

		currentBytes, err := os.ReadFile(abs)
		if err != nil {
			return Plan{}, err
		}
		currentSha := managedstate.Sha256HexBytes(currentBytes)
		desiredSha := managedstate.Sha256HexBytes(df.Bytes)

		if entry == nil {
			if opts.Overwrite {
				plan.Entries = append(plan.Entries, PlanEntry{Action: ActionOverwrite, Path: path, Reason: "unmanaged-exists", ArtifactID: df.ArtifactID})
				plan.Overwrite = append(plan.Overwrite, path)
			} else {
				plan.Entries = append(plan.Entries, PlanEntry{Action: ActionConflict, Path: path, Reason: "unmanaged-exists", ArtifactID: df.ArtifactID})
				plan.Conflicts = append(plan.Conflicts, path)
			}
			continue
		}

		if currentSha != entry.InstalledSHA256 {
			if opts.Overwrite {
				plan.Entries = append(plan.Entries, PlanEntry{Action: ActionOverwrite, Path: path, Reason: "checksum-drift", ArtifactID: df.ArtifactID})
				plan.Overwrite = append(plan.Overwrite, path)
			} else {
				plan.Entries = append(plan.Entries, PlanEntry{Action: ActionConflict, Path: path, Reason: "checksum-drift", ArtifactID: df.ArtifactID})
				plan.Conflicts = append(plan.Conflicts, path)
			}
			continue
		}

		if currentSha == desiredSha {
			plan.Entries = append(plan.Entries, PlanEntry{Action: ActionSkip, Path: path, Reason: "up-to-date", ArtifactID: df.ArtifactID})
			plan.Skip = append(plan.Skip, path)
			continue
		}

		plan.Entries = append(plan.Entries, PlanEntry{Action: ActionUpdate, Path: path, Reason: "managed-unchanged", ArtifactID: df.ArtifactID})
		plan.Update = append(plan.Update, path)
	}

	for path, entry := range ag.Files {
		if entry == nil {
			continue
		}
		if _, ok := desiredSet[path]; ok {
			continue
		}
		plan.Orphaned = append(plan.Orphaned, path)
		if opts.Prune {
			abs, err := SafeAbsPath(projectRootAbs, path)
			if err != nil {
				continue
			}
			b, err := os.ReadFile(abs)
			if err != nil {
				continue
			}
			sha := managedstate.Sha256HexBytes(b)
			if sha == entry.InstalledSHA256 {
				plan.Entries = append(plan.Entries, PlanEntry{Action: ActionDelete, Path: path, Reason: "orphaned-managed-unchanged", ArtifactID: entry.ArtifactID})
				plan.Delete = append(plan.Delete, path)
			}
		}
	}

	sort.Strings(plan.Create)
	sort.Strings(plan.Update)
	sort.Strings(plan.Overwrite)
	sort.Strings(plan.Skip)
	sort.Strings(plan.Conflicts)
	sort.Strings(plan.Delete)
	sort.Strings(plan.Orphaned)

	sort.Slice(plan.Entries, func(i, j int) bool {
		if plan.Entries[i].Path == plan.Entries[j].Path {
			return plan.Entries[i].Action < plan.Entries[j].Action
		}
		return plan.Entries[i].Path < plan.Entries[j].Path
	})

	return plan, nil
}

func Apply(projectRootAbs string, agentID string, packID string, packVersion string, desired []DesiredFile, st *managedstate.State, opts Options) (ApplyResult, *managedstate.State, error) {
	plan, err := BuildPlan(projectRootAbs, agentID, desired, st, opts)
	if err != nil {
		return ApplyResult{}, nil, err
	}
	res := ApplyResult{Plan: plan}
	if opts.DryRun {
		return res, st, nil
	}
	if st == nil {
		st = &managedstate.State{SchemaVersion: managedstate.SchemaVersion}
	}
	ag := managedstate.EnsureAgent(st, agentID)
	ag.Pack = managedstate.PackState{ID: packID, Version: packVersion}

	backupsDir := ""
	ensureBackupDir := func() (string, error) {
		if backupsDir != "" {
			return backupsDir, nil
		}
		backupsDir = filepath.Join(projectRootAbs, ".openkit", "backups", strings.ReplaceAll(managedstate.NowRFC3339(), ":", "-"))
		if err := os.MkdirAll(backupsDir, 0755); err != nil {
			return "", err
		}
		return backupsDir, nil
	}

	desiredMap := map[string]DesiredFile{}
	for _, df := range desired {
		p, err := NormalizeRelOutputPath(df.OutputPath)
		if err != nil {
			return ApplyResult{}, nil, err
		}
		df.OutputPath = p
		desiredMap[p] = df
	}

	for _, e := range plan.Entries {
		switch e.Action {
		case ActionCreate, ActionUpdate, ActionOverwrite:
			df, ok := desiredMap[e.Path]
			if !ok {
				return ApplyResult{}, nil, fmt.Errorf("internal: desired missing for %s", e.Path)
			}
			abs, err := SafeAbsPath(projectRootAbs, e.Path)
			if err != nil {
				return ApplyResult{}, nil, err
			}

			if e.Action == ActionOverwrite {
				if _, err := os.Stat(abs); err == nil {
					bd, err := ensureBackupDir()
					if err != nil {
						return ApplyResult{}, nil, err
					}
					backupAbs := filepath.Join(bd, filepath.FromSlash(e.Path))
					if err := os.MkdirAll(filepath.Dir(backupAbs), 0755); err != nil {
						return ApplyResult{}, nil, err
					}
					b, err := os.ReadFile(abs)
					if err != nil {
						return ApplyResult{}, nil, err
					}
					if err := os.WriteFile(backupAbs, b, 0644); err != nil {
						return ApplyResult{}, nil, err
					}
				}
			}

			if err := os.MkdirAll(filepath.Dir(abs), 0755); err != nil {
				return ApplyResult{}, nil, err
			}
			tmp := abs + ".tmp"
			if err := os.WriteFile(tmp, df.Bytes, 0644); err != nil {
				return ApplyResult{}, nil, err
			}
			if err := os.Rename(tmp, abs); err != nil {
				return ApplyResult{}, nil, err
			}

			sha := managedstate.Sha256HexBytes(df.Bytes)
			ag.Files[e.Path] = &managedstate.FileEntry{
				ArtifactID:      df.ArtifactID,
				InstalledSHA256: sha,
				InstalledAt:     managedstate.NowRFC3339(),
				Mode:            df.Mode,
			}

		case ActionDelete:
			abs, err := SafeAbsPath(projectRootAbs, e.Path)
			if err != nil {
				continue
			}
			if _, err := os.Stat(abs); err != nil {
				continue
			}
			bd, err := ensureBackupDir()
			if err != nil {
				return ApplyResult{}, nil, err
			}
			backupAbs := filepath.Join(bd, filepath.FromSlash(e.Path))
			if err := os.MkdirAll(filepath.Dir(backupAbs), 0755); err != nil {
				return ApplyResult{}, nil, err
			}
			b, err := os.ReadFile(abs)
			if err == nil {
				_ = os.WriteFile(backupAbs, b, 0644)
			}
			_ = os.Remove(abs)
			delete(ag.Files, e.Path)
		}
	}

	res.BackupsDir = backupsDir
	return res, st, nil
}
