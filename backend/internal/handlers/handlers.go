package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"portfolio/internal/config"
	"portfolio/internal/database"
	"portfolio/internal/github"
	"portfolio/internal/models"
)

type Handler struct {
	cfg    *config.Config
	github *github.Client
	db     *database.DB
}

func New(cfg *config.Config, gh *github.Client, db *database.DB) *Handler {
	return &Handler{cfg: cfg, github: gh, db: db}
}

func (h *Handler) writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

func (h *Handler) writeError(w http.ResponseWriter, status int, msg string) {
	h.writeJSON(w, status, map[string]string{"error": msg})
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	h.writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (h *Handler) About(w http.ResponseWriter, r *http.Request) {
	h.writeJSON(w, http.StatusOK, models.About{
		Name:    h.cfg.Name,
		Title:   h.cfg.Title,
		About:   h.cfg.About,
		Skills:  toModelSkills(h.cfg.Skills),
		Contact: models.Contact(h.cfg.Contact),
		Github:  h.cfg.GithubUsername,
		Avatar:  h.github.AvatarURL(),
	})
}

func (h *Handler) Skills(w http.ResponseWriter, r *http.Request) {
	h.writeJSON(w, http.StatusOK, toModelSkills(h.cfg.Skills))
}

func (h *Handler) Contact(w http.ResponseWriter, r *http.Request) {
	h.writeJSON(w, http.StatusOK, h.cfg.Contact)
}

func (h *Handler) Projects(w http.ResponseWriter, r *http.Request) {
	projects, err := h.github.FetchRepos()
	if err != nil {
		projects = h.fallbackProjects()
	}

	featured := h.cfg.FeaturedRepos
	filtered := make([]models.Project, 0, len(projects))
	for _, p := range projects {
		if len(featured) == 0 || slices.Contains(featured, p.Name) {
			p.Featured = slices.Contains(featured, p.Name)
			filtered = append(filtered, p)
		}
	}

	if len(filtered) == 0 {
		filtered = projects
	}

	h.writeJSON(w, http.StatusOK, filtered)
}

func (h *Handler) fallbackProjects() []models.Project {
	result := make([]models.Project, len(h.cfg.FallbackProjects))
	for i, p := range h.cfg.FallbackProjects {
		result[i] = models.Project{
			Name:        p.Name,
			Description: p.Description,
			URL:         p.URL,
			Language:    p.Language,
			Featured:    slices.Contains(h.cfg.FeaturedRepos, p.Name),
		}
	}
	return result
}

func (h *Handler) TrackClick(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var event models.ClickEvent
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		h.writeError(w, http.StatusBadRequest, "неверный JSON")
		return
	}

	event.Link = strings.TrimSpace(event.Link)
	if event.Link == "" {
		h.writeError(w, http.StatusBadRequest, "link обязателен")
		return
	}
	if event.Type == "" {
		event.Type = "link"
	}

	if err := h.db.RecordClick(event.Link, event.Type); err != nil {
		h.writeError(w, http.StatusInternalServerError, "ошибка записи")
		return
	}

	h.writeJSON(w, http.StatusOK, map[string]string{"status": "recorded"})
}

func (h *Handler) TrackDownload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	if err := h.db.RecordResumeDownload(); err != nil {
		h.writeError(w, http.StatusInternalServerError, "ошибка записи")
		return
	}

	h.writeJSON(w, http.StatusOK, map[string]string{"status": "recorded"})
}

func (h *Handler) Analytics(w http.ResponseWriter, r *http.Request) {
	stats, err := h.db.GetStats()
	if err != nil {
		h.writeError(w, http.StatusInternalServerError, "ошибка статистики")
		return
	}
	h.writeJSON(w, http.StatusOK, stats)
}

func (h *Handler) Resume(w http.ResponseWriter, r *http.Request) {
	resumePath := h.cfg.ResumeFile
	if !filepath.IsAbs(resumePath) {
		resumePath = filepath.Join(".", resumePath)
	}

	f, err := os.Open(resumePath)
	if err != nil {
		h.writeError(w, http.StatusNotFound, "резюме не найдено — положите PDF в backend/assets/resume.pdf")
		return
	}
	defer f.Close()

	if err := h.db.RecordResumeDownload(); err != nil {
		h.writeError(w, http.StatusInternalServerError, "ошибка записи статистики")
		return
	}

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", `attachment; filename="Zaitsev_Artem_Resume.pdf"`)
	_, _ = io.Copy(w, f)
}

func toModelSkills(skills []config.Skill) []models.Skill {
	result := make([]models.Skill, len(skills))
	for i, s := range skills {
		result[i] = models.Skill{Name: s.Name, Level: s.Level}
	}
	return result
}

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
