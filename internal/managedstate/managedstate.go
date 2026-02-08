package managedstate

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const SchemaVersion = "1"

type State struct {
	SchemaVersion string                 `json:"schema_version"`
	Agents        map[string]*AgentState `json:"agents"`
}

type AgentState struct {
	Pack  PackState             `json:"pack"`
	Files map[string]*FileEntry `json:"files"`
}

type PackState struct {
	ID      string `json:"id"`
	Version string `json:"version"`
}

type FileEntry struct {
	ArtifactID      string `json:"artifact_id"`
	InstalledSHA256 string `json:"installed_sha256"`
	InstalledAt     string `json:"installed_at"`
	Mode            string `json:"mode"` // copy|render|template
}

func DefaultPath(projectRootAbs string) string {
	return filepath.Join(projectRootAbs, ".openkit", "managed.json")
}

func Load(pathAbs string) (*State, error) {
	b, err := os.ReadFile(pathAbs)
	if err != nil {
		return nil, err
	}

	var st State
	if err := json.Unmarshal(b, &st); err != nil {
		return nil, fmt.Errorf("parse managed state: %w", err)
	}
	if st.SchemaVersion == "" {
		return nil, fmt.Errorf("managed state missing schema_version")
	}
	if st.SchemaVersion != SchemaVersion {
		return nil, fmt.Errorf("unsupported managed state schema_version=%s", st.SchemaVersion)
	}
	if st.Agents == nil {
		st.Agents = map[string]*AgentState{}
	}
	return &st, nil
}

func Save(pathAbs string, st *State) error {
	if st == nil {
		return fmt.Errorf("nil state")
	}
	if st.SchemaVersion == "" {
		st.SchemaVersion = SchemaVersion
	}
	if st.SchemaVersion != SchemaVersion {
		return fmt.Errorf("unsupported state schema_version=%s", st.SchemaVersion)
	}
	if st.Agents == nil {
		st.Agents = map[string]*AgentState{}
	}

	parent := filepath.Dir(pathAbs)
	if err := os.MkdirAll(parent, 0755); err != nil {
		return err
	}

	b, err := json.MarshalIndent(st, "", "  ")
	if err != nil {
		return err
	}
	b = append(b, '\n')

	tmp := pathAbs + ".tmp"
	if err := os.WriteFile(tmp, b, 0644); err != nil {
		return err
	}
	return os.Rename(tmp, pathAbs)
}

func EnsureAgent(st *State, agentID string) *AgentState {
	if st.Agents == nil {
		st.Agents = map[string]*AgentState{}
	}
	key := strings.ToLower(strings.TrimSpace(agentID))
	if key == "" {
		key = "unknown"
	}
	if st.Agents[key] == nil {
		st.Agents[key] = &AgentState{Files: map[string]*FileEntry{}}
	}
	if st.Agents[key].Files == nil {
		st.Agents[key].Files = map[string]*FileEntry{}
	}
	return st.Agents[key]
}

func NowRFC3339() string {
	return time.Now().UTC().Format(time.RFC3339)
}

func Sha256HexBytes(b []byte) string {
	sum := sha256.Sum256(b)
	return hex.EncodeToString(sum[:])
}

func Sha256HexReader(r io.Reader) (string, error) {
	h := sha256.New()
	if _, err := io.Copy(h, r); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}
