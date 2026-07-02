package config

import (
	"encoding/json"
	"os"
)

type Skill struct {
	Name  string `json:"name"`
	Level int    `json:"level"`
}

type Contact struct {
	Email    string `json:"email"`
	Telegram string `json:"telegram"`
	Github   string `json:"github"`
}

type FallbackProject struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Language    string `json:"language"`
}

type Config struct {
	Name             string            `json:"name"`
	Title            string            `json:"title"`
	GithubUsername   string            `json:"github_username"`
	About            string            `json:"about"`
	Skills           []Skill           `json:"skills"`
	Contact          Contact           `json:"contact"`
	FeaturedRepos    []string          `json:"featured_repos"`
	FallbackProjects []FallbackProject `json:"fallback_projects"`
	ResumeFile       string            `json:"resume_file"`
}

func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
