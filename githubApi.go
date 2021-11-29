package getgithubdata

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Repo struct {
	StargazersCount int `json:"stargazers_count"`
}

// Build an Api for get data from gitgub
type RepoAPI interface {
	GetRepos(name string) ([]Repo, error)
}

// 模擬返回資料
type Mock struct{}

func (m *Mock) GetRepos(name string) ([]Repo, error) {
	return []Repo{
		Repo{
			StargazersCount: 4,
		},
		Repo{
			StargazersCount: 10,
		},
	}, nil
}

// 真正的API
type GitHub struct{}

func (g *GitHub) GetRepos(username string) ([]Repo, error) {
	res, err := http.Get(fmt.Sprintf("https://api.github.com/users/%s/repos", username))
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	repos := []Repo{}
	if err := json.NewDecoder(res.Body).Decode(&repos); err != nil {
		return nil, err
	}

	return repos, nil
}
