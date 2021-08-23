package actions

import (
	"fmt"
	"log"

	"github.com/veocode/dws/src/services"
)

type Test struct {
}

func (t Test) Execute() {

	repository := "veocode/dws"
	folder := "src"

	fmt.Printf("Grabbing %s from %s\n", repository, folder)

	grabber := services.NewGithubGrabber()
	err := grabber.DownloadRepositoryFolder(repository, folder)

	if err != nil {
		log.Fatalf("FAILED: %s", err)
	}

}
