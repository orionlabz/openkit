package selfupdate

import "testing"

func TestArtifactFilename(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		goos   string
		goarch string
		want   string
	}{
		{
			name:   "darwin arm64",
			goos:   "darwin",
			goarch: "arm64",
			want:   "openkit_Darwin_arm64.tar.gz",
		},
		{
			name:   "linux amd64",
			goos:   "linux",
			goarch: "amd64",
			want:   "openkit_Linux_x86_64.tar.gz",
		},
		{
			name:   "windows amd64",
			goos:   "windows",
			goarch: "amd64",
			want:   "openkit_Windows_x86_64.zip",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := artifactFilename(tt.goos, tt.goarch)
			if err != nil {
				t.Fatalf("artifactFilename() error = %v", err)
			}
			if got != tt.want {
				t.Fatalf("artifactFilename() = %q; want %q", got, tt.want)
			}
		})
	}
}

func TestArtifactFilename_UnsupportedPlatform(t *testing.T) {
	t.Parallel()

	if _, err := artifactFilename("plan9", "amd64"); err == nil {
		t.Fatalf("expected error for unsupported OS")
	}
	if _, err := artifactFilename("darwin", "386"); err == nil {
		t.Fatalf("expected error for unsupported arch")
	}
}
