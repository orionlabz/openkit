package selfupdate

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type GitHubRelease struct {
	TagName string `json:"tag_name"`
}

func FetchLatestTag(ctx context.Context, client *http.Client, url string, etag string) (tag string, nextETag string, notModified bool, err error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", "", false, err
	}
	if etag != "" {
		req.Header.Set("If-None-Match", etag)
	}
	req.Header.Set("Accept", "application/vnd.github+json")

	resp, err := client.Do(req)
	if err != nil {
		return "", "", false, err
	}
	defer closeWithErr(&err, resp.Body)

	if resp.StatusCode == http.StatusNotModified {
		return "", resp.Header.Get("ETag"), true, nil
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		b, _ := io.ReadAll(io.LimitReader(resp.Body, 8<<10))
		return "", "", false, fmt.Errorf("latest release request failed: %s: %s", resp.Status, string(b))
	}

	var r GitHubRelease
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return "", "", false, err
	}
	if r.TagName == "" {
		return "", "", false, fmt.Errorf("latest release response missing tag_name")
	}

	return r.TagName, resp.Header.Get("ETag"), false, nil
}
