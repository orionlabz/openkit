package selfupdate

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"time"
)

type State struct {
	LastCheckedUnix int64  `json:"last_checked_unix"`
	LatestTag       string `json:"latest_tag"`
	ETag            string `json:"etag"`
}

func (s State) LastChecked() time.Time {
	if s.LastCheckedUnix <= 0 {
		return time.Time{}
	}
	return time.Unix(s.LastCheckedUnix, 0)
}

func LoadState(path string) (State, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return State{}, nil
		}
		return State{}, err
	}
	var s State
	if err := json.Unmarshal(b, &s); err != nil {
		return State{}, err
	}
	return s, nil
}

func SaveState(path string, s State) error {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}
	b, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, b, 0644)
}
