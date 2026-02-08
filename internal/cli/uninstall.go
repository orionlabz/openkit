package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/openkit-devtools/openkit/internal/platform"
	"github.com/spf13/cobra"
)

var uninstallYes bool

var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall OpenKit CLI from this machine",
	Long:  "Removes the OpenKit home directory (~/.openkit) and any symlinks created in user PATH directories.",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runUninstall()
	},
}

func init() {
	uninstallCmd.Flags().BoolVar(&uninstallYes, "yes", false, "Confirm uninstall")
}

func runUninstall() error {
	if !uninstallYes {
		exitWithError("Refusing to uninstall without --yes")
	}

	home, err := platform.OpenKitHome()
	if err != nil {
		exitWithError(fmt.Sprintf("Failed to determine OPENKIT_HOME: %v", err))
	}

	binDir, err := platform.OpenKitBinDir()
	if err != nil {
		exitWithError(fmt.Sprintf("Failed to determine install dir: %v", err))
	}

	target := filepath.Join(binDir, "openkit")
	_ = removeSymlinksTo(target)

	if err := os.RemoveAll(home); err != nil {
		exitWithError(fmt.Sprintf("Failed to remove %s: %v", home, err))
	}

	printSuccess("OpenKit uninstalled")
	return nil
}

func removeSymlinksTo(target string) error {
	candidates := []string{}
	if v := os.Getenv("XDG_BIN_HOME"); v != "" {
		candidates = append(candidates, filepath.Join(v, "openkit"))
	}
	if home, err := os.UserHomeDir(); err == nil {
		candidates = append(candidates,
			filepath.Join(home, ".local", "bin", "openkit"),
			filepath.Join(home, "bin", "openkit"),
		)
	}
	// Best-effort system locations (may fail without permissions)
	candidates = append(candidates,
		"/usr/local/bin/openkit",
		"/opt/homebrew/bin/openkit",
	)

	for _, p := range candidates {
		_ = maybeRemoveSymlink(p, target)
	}
	return nil
}

func maybeRemoveSymlink(path string, target string) error {
	fi, err := os.Lstat(path)
	if err != nil {
		return nil
	}
	if fi.Mode()&os.ModeSymlink == 0 {
		return nil
	}
	link, err := os.Readlink(path)
	if err != nil {
		return nil
	}
	// Normalize relative symlinks
	if !filepath.IsAbs(link) {
		link = filepath.Clean(filepath.Join(filepath.Dir(path), link))
	}
	if filepath.Clean(link) != filepath.Clean(target) {
		return nil
	}
	_ = os.Remove(path)
	return nil
}
