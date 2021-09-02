package hub

import (
	"os"
	"path/filepath"

	"github.com/veocode/dws/src/services/shell"
)

const HubRepoURL = "git@github.com:veocode/dwshub.git"
const HubRepoPath = ".dws"

type Hub struct {
	cacheDir string
}

func NewHub() (*Hub, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	repo := new(Hub)
	repo.cacheDir = filepath.Join(homeDir, HubRepoPath)
	err = repo.checkCache()
	if err != nil {
		return nil, err
	}

	return repo, nil
}

func (h *Hub) checkCache() error {
	if !shell.IsPathExists(h.cacheDir) {
		return h.downloadCache()
	}
	return nil
}

func (h *Hub) downloadCache() error {
	os.MkdirAll(h.cacheDir, os.ModePerm)
	shell.PrintOut("Downloading hub cache to %s...", h.cacheDir)
	err := shell.Execute("git", "clone", HubRepoURL, h.cacheDir)
	return err
}

func (h *Hub) IsSelectorExists(selector string) bool {
	return false
}

func (h *Hub) ExtractPath(path string, targetDir string) error {
	srcDir := filepath.Join(h.cacheDir, path)
	return shell.CopyDirectory(srcDir, targetDir)
}
