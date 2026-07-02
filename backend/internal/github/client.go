package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"portfolio/internal/models"
)

type repoResponse struct {
	Name        string   `json:"name"`
	Description *string  `json:"description"`
	HTMLURL     string   `json:"html_url"`
	Language    *string  `json:"language"`
	Stargazers  int      `json:"stargazers_count"`
	Forks       int      `json:"forks_count"`
	Topics      []string `json:"topics"`
	UpdatedAt   string   `json:"updated_at"`
}

type Client struct {
	username string
	http     *http.Client
}

func NewClient(username string) *Client {
	return &Client{
		username: username,
		http: &http.Client{
			Timeout: 15 * time.Second,
		},
	}
}

func (c *Client) FetchRepos() ([]models.Project, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/repos?sort=updated&per_page=100", c.username)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("User-Agent", "portfolio-app")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("github request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("github api status: %d", resp.StatusCode)
	}

	var raw []repoResponse
	if err := json.NewDecoder(resp.Body).Decode(&raw); err != nil {
		return nil, fmt.Errorf("decode repos: %w", err)
	}

	projects := make([]models.Project, 0, len(raw))
	for _, r := range raw {
		desc := ""
		if r.Description != nil {
			desc = *r.Description
		}
		lang := ""
		if r.Language != nil {
			lang = *r.Language
		}

		projects = append(projects, models.Project{
			Name:        r.Name,
			Description: desc,
			URL:         r.HTMLURL,
			Language:    lang,
			Stars:       r.Stargazers,
			Forks:       r.Forks,
			Topics:      r.Topics,
			UpdatedAt:   r.UpdatedAt,
		})
	}

	return projects, nil
}

func (c *Client) AvatarURL() string {
	return fmt.Sprintf("https://github.com/%s.png", c.username)
}
