package cli

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/openkit-devtools/openkit/internal/platform"
	"github.com/openkit-devtools/openkit/internal/selfupdate"
	"github.com/spf13/cobra"
)

var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Download and install the latest OpenKit CLI release",
	Long:  "Downloads the latest release from GitHub, verifies checksums, and installs into your OpenKit home directory.",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runUpgrade(cmd.Context())
	},
}

func runUpgrade(ctx context.Context) error {
	installDir, err := platform.OpenKitBinDir()
	if err != nil {
		exitWithError(fmt.Sprintf("Failed to determine install dir: %v", err))
	}

	client := &http.Client{Timeout: 60 * time.Second}
	latestURL := "https://api.github.com/repos/openkit-devtools/openkit/releases/latest"
	tag, _, _, err := selfupdate.FetchLatestTag(ctx, client, latestURL, "")
	if err != nil {
		exitWithError(fmt.Sprintf("Failed to fetch latest release: %v", err))
	}

	printInfo(fmt.Sprintf("Upgrading to %s...", tag))
	path, err := selfupdate.Upgrade(ctx, client, selfupdate.UpgradeOptions{
		RepoOwner:  "openkit-devtools",
		RepoName:   "openkit",
		Tag:        tag,
		InstallDir: installDir,
	})
	if err != nil {
		exitWithError(fmt.Sprintf("Upgrade failed: %v", err))
	}

	printSuccess(fmt.Sprintf("Installed to %s", path))
	printInfo("Restart your terminal session if needed")

	// Best-effort: update state so next runs don't recheck immediately.
	statePath, err := platform.OpenKitStatePath()
	if err == nil {
		_ = selfupdate.SaveState(statePath, selfupdate.State{LastCheckedUnix: time.Now().Unix(), LatestTag: tag})
	}

	return nil
}
