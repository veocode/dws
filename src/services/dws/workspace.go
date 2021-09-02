package dws

import (
	"path/filepath"

	"github.com/veocode/dws/src/config"
	"github.com/veocode/dws/src/models"
	"github.com/veocode/dws/src/services/jsonio"
	"github.com/veocode/dws/src/services/shell"
)

type WorkspaceManager struct {
}

func NewWorkspaceManager() *WorkspaceManager {
	return new(WorkspaceManager)
}

func (w *WorkspaceManager) CreateWorkspace(targetDir string) error {
	hub, err := NewHub()
	if err != nil {
		return err
	}

	err = hub.ExtractPath("init/scratch", targetDir)
	if err != nil {
		return err
	}

	workspace := models.NewWorkspace()
	workspace.Name = filepath.Base(targetDir)
	workspace.Subnet, err = shell.Ask("Docker Subnet for this workspace", "192.171.0.0/16")
	if err != nil {
		return err
	}

	err = jsonio.NewSaver().SaveToFile(workspace, filepath.Join(targetDir, config.FileMetaWorkspace))
	if err != nil {
		return err
	}

	return nil
}
