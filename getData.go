package getgithubdata

// 建立githubApi模組化之前

// type Repo struct {
// 	StargazersCount int `json:"stargazers_count"`
// }

// func GetAverageStar(name string) (float64, error) {
// 	res, err := http.Get(fmt.Sprintf("https://api.github.com/users/%s/repos", name))
// 	if err != nil {
// 		return 0, err
// 	}

// 	defer res.Body.Close()

// 	var repos []Repo
// 	if err := json.NewDecoder(res.Body).Decode(&repos); err != nil {
// 		return 0, err
// 	}

// 	if len(repos) == 0 {
// 		return 0, nil
// 	}

// 	var total int
// 	for _, v := range repos {
// 		total += v.StargazersCount
// 	}

// 	return float64(total) / float64(len(repos)), nil

// }

// 模組化之後, 使用githubApi的方法

func GetAverageStar(repoAPI RepoAPI, name string) (float64, error) {

	repos, err := repoAPI.GetRepos(name)

	if len(repos) == 0 {
		return 0, nil
	}

	if err != nil {
		return 0, err
	}

	var total int
	for _, v := range repos {
		total += v.StargazersCount
	}

	return float64(total) / float64(len(repos)), nil
}
