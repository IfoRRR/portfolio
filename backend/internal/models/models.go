package models

type Skill struct {
	Name  string `json:"name"`
	Level int    `json:"level"`
}

type Contact struct {
	Email    string `json:"email"`
	Telegram string `json:"telegram"`
	Github   string `json:"github"`
}

type About struct {
	Name     string  `json:"name"`
	Title    string  `json:"title"`
	About    string  `json:"about"`
	Skills   []Skill `json:"skills"`
	Contact  Contact `json:"contact"`
	Github   string  `json:"github_username"`
	Avatar   string  `json:"avatar_url,omitempty"`
}

type Project struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	URL         string   `json:"url"`
	Language    string   `json:"language"`
	Stars       int      `json:"stars"`
	Forks       int      `json:"forks"`
	Topics      []string `json:"topics"`
	UpdatedAt   string   `json:"updated_at"`
	Featured    bool     `json:"featured"`
}

type ClickEvent struct {
	Link string `json:"link"`
	Type string `json:"type"`
}

type LinkStat struct {
	Link  string `json:"link"`
	Type  string `json:"type"`
	Count int    `json:"count"`
}

type AnalyticsStats struct {
	TotalClicks     int        `json:"total_clicks"`
	ResumeDownloads int        `json:"resume_downloads"`
	Links           []LinkStat `json:"links"`
}
