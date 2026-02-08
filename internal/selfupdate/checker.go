package selfupdate

import (
	"context"
	"net/http"
	"time"

	"golang.org/x/mod/semver"
)

type CheckResult struct {
	Current       string
	Latest        string
	HasUpdate     bool
	CheckedOnline bool
}

type Checker struct {
	Client    *http.Client
	LatestURL string
	StatePath string
	TTL       time.Duration
}

func (c Checker) Check(ctx context.Context, currentVersion string) (CheckResult, error) {
	res := CheckResult{Current: currentVersion}

	state, err := LoadState(c.StatePath)
	if err != nil {
		return res, err
	}

	now := time.Now()
	if c.TTL > 0 && !state.LastChecked().IsZero() && now.Sub(state.LastChecked()) < c.TTL {
		res.Latest = state.LatestTag
		res.CheckedOnline = false
		res.HasUpdate = hasUpdate(currentVersion, state.LatestTag)
		return res, nil
	}

	tag, nextETag, notModified, err := FetchLatestTag(ctx, c.Client, c.LatestURL, state.ETag)
	if err != nil {
		return res, err
	}

	if notModified {
		res.Latest = state.LatestTag
		res.CheckedOnline = true
		state.LastCheckedUnix = now.Unix()
		if nextETag != "" {
			state.ETag = nextETag
		}
		_ = SaveState(c.StatePath, state)
		res.HasUpdate = hasUpdate(currentVersion, state.LatestTag)
		return res, nil
	}

	res.Latest = tag
	res.CheckedOnline = true
	state.LastCheckedUnix = now.Unix()
	state.LatestTag = tag
	if nextETag != "" {
		state.ETag = nextETag
	}
	_ = SaveState(c.StatePath, state)

	res.HasUpdate = hasUpdate(currentVersion, tag)
	return res, nil
}

func hasUpdate(current, latest string) bool {
	if current == "" || latest == "" {
		return false
	}
	if !semver.IsValid(current) {
		current = "v" + current
	}
	if !semver.IsValid(latest) {
		latest = "v" + latest
	}
	if !semver.IsValid(current) || !semver.IsValid(latest) {
		return false
	}
	return semver.Compare(latest, current) > 0
}
