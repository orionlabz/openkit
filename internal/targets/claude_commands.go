package targets

import (
	"io/fs"
	"path"
	"sort"
	"strings"

	"github.com/openkit-devtools/openkit/internal/syncer"
)

// desiredClaudeCommands converts base/commands/*.md to .claude/commands/*.md format.
// Claude Code supports custom commands via .claude/commands/ directory.
// Unlike Gemini (TOML), Claude uses markdown files directly.
func desiredClaudeCommands(base fs.FS) ([]syncer.DesiredFile, error) {
	entries, err := fs.ReadDir(base, "base/commands")
	if err != nil {
		return nil, err
	}

	sort.Slice(entries, func(i, j int) bool { return entries[i].Name() < entries[j].Name() })

	var out []syncer.DesiredFile
	for _, ent := range entries {
		if ent.IsDir() {
			continue
		}
		name := ent.Name()
		if name == "README.md" {
			continue
		}
		if !strings.HasSuffix(name, ".md") {
			continue
		}

		embeddedPath := path.Join("base/commands", name)
		b, err := fs.ReadFile(base, embeddedPath)
		if err != nil {
			return nil, err
		}

		// Claude uses markdown files directly, but we adapt the format slightly
		// to work better with Claude's command system
		cmdContent := adaptCommandForClaude(b)

		out = append(out, syncer.DesiredFile{
			OutputPath: ".claude/commands/" + name,
			Bytes:      cmdContent,
			ArtifactID: "generated/claude/commands/" + name,
			Mode:       "copy",
		})
	}

	sort.Slice(out, func(i, j int) bool { return out[i].OutputPath < out[j].OutputPath })
	return out, nil
}

// adaptCommandForClaude adapts a command file for Claude Code.
// Claude expects markdown but we ensure $ARGUMENTS is preserved and
// add any Claude-specific annotations if needed.
func adaptCommandForClaude(content []byte) []byte {
	// Claude Code uses $ARGUMENTS placeholder natively (same as our format)
	// No conversion needed, but we ensure the content is clean
	return content
}
