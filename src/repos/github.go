package services

type GithubRepo struct {
}

func NewGithubRepo() *GithubRepo {
	repo := new(GithubRepo)
	repo.checkCache()
	return repo
}

func (g *GithubRepo) checkCache() {
	// if not already exists
	g.downloadCache()
}

func (g *GithubRepo) downloadCache() error {
	// download from github.com/veocode/dwslib to ~/.dws/lib
	return nil
}

func (g *GithubRepo) IsSelectorExists(selector string) bool {
	// app/go
	//	-f ~/.dws/lib/app/go ?
	return false
}

func (g *GithubRepo) ExtractSelector(selector string, targetDir string) error {
	//
	// better return only path in repo
	// copy with another file service
	//

	// app/php => /path/to/project
	// 		cp -r ~/.dws/lib/app/php/* /path/to/project
	// db/redis => /path/to/project
	// 		cp -r ~/.dws/lib/db/redis/* /path/to/project
	return nil
}
