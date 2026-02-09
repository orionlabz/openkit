package targets

import (
	"strings"
	"testing"
	"testing/fstest"
)

func TestDesiredClaudeCommands_CopiesMarkdownFiles(t *testing.T) {
	fsys := fstest.MapFS{
		"base/commands/specify.md": {Data: []byte("---\ndescription: Spec\n---\n\nHello $ARGUMENTS\n")},
		"base/commands/plan.md":    {Data: []byte("---\ndescription: Plan\n---\n\nCreate plan for $ARGUMENTS\n")},
		"base/commands/README.md":  {Data: []byte("ignore")},
	}

	files, err := desiredClaudeCommands(fsys)
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	if len(files) != 2 {
		t.Fatalf("files = %d, want 2", len(files))
	}

	// Verify output paths
	expectedPaths := map[string]bool{
		".claude/commands/plan.md":    true,
		".claude/commands/specify.md": true,
	}
	for _, f := range files {
		if !expectedPaths[f.OutputPath] {
			t.Errorf("unexpected OutputPath = %q", f.OutputPath)
		}
	}
}

func TestDesiredClaudeCommands_PreservesContent(t *testing.T) {
	content := "---\ndescription: Test\n---\n\nContent with $ARGUMENTS\n"
	fsys := fstest.MapFS{
		"base/commands/test.md": {Data: []byte(content)},
	}

	files, err := desiredClaudeCommands(fsys)
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	if len(files) != 1 {
		t.Fatalf("files = %d, want 1", len(files))
	}

	f := files[0]
	if f.OutputPath != ".claude/commands/test.md" {
		t.Fatalf("OutputPath = %q", f.OutputPath)
	}

	// Claude should preserve $ARGUMENTS (unlike Gemini which converts to {{args}})
	if !strings.Contains(string(f.Bytes), "$ARGUMENTS") {
		t.Fatalf("expected $ARGUMENTS to be preserved: %q", string(f.Bytes))
	}
}

func TestDesiredClaudeCommands_SkipsReadme(t *testing.T) {
	fsys := fstest.MapFS{
		"base/commands/README.md": {Data: []byte("# Commands\n\nThis is ignored.")},
	}

	files, err := desiredClaudeCommands(fsys)
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	if len(files) != 0 {
		t.Fatalf("files = %d, want 0 (README should be skipped)", len(files))
	}
}

func TestAdaptCommandForClaude_PreservesContent(t *testing.T) {
	input := []byte("Some content with $ARGUMENTS placeholder")
	output := adaptCommandForClaude(input)

	if string(output) != string(input) {
		t.Fatalf("expected content to be preserved, got %q", string(output))
	}
}
