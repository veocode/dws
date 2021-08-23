package services

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type GithubFileList = []GithubFileDescription

type GithubFileDescription struct {
	Name        string `json:"name"`
	Path        string `json:"path"`
	Sha         string `json:"sha"`
	Size        int    `json:"size"`
	URL         string `json:"url"`
	HTMLURL     string `json:"html_url"`
	GitURL      string `json:"git_url"`
	DownloadURL string `json:"download_url"`
	Type        string `json:"type"`
	Links       struct {
		Self string `json:"self"`
		Git  string `json:"git"`
		HTML string `json:"html"`
	} `json:"_links"`
}

type GithubGrabber struct {
	httpClient *http.Client
}

func NewGithubGrabber() *GithubGrabber {
	grabber := new(GithubGrabber)
	grabber.httpClient = &http.Client{}
	return grabber
}

func (g *GithubGrabber) DownloadRepositoryFolder(repository string, folder string) (string, error) {
	targetDir, err := ioutil.TempDir("", "dwsrepo")
	if err != nil {
		return "", err
	}

	err = g.downloadRecursive(targetDir, repository, folder)
	if err != nil {
		return "", err
	}

	return targetDir, nil
}

func (g *GithubGrabber) downloadRecursive(targetDir string, repository string, path string) error {
	fileList, err := g.getContentList(repository, path)
	if err != nil {
		return err
	}

	for _, fileInfo := range *fileList {
		if fileInfo.Type == "dir" {
			fmt.Println("/" + fileInfo.Name)
		}
	}

	return nil
}

func (g *GithubGrabber) getContentList(repository string, path string) (*GithubFileList, error) {
	list := new(GithubFileList)
	url := fmt.Sprintf("https://api.github.com/repos/%s/contents/%s", repository, path)

	err := NewJSONLoader().Load(url).ToStruct(list)
	if err != nil {
		return nil, err
	}

	return list, err
}
