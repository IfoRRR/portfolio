package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"portfolio/internal/config"
	"portfolio/internal/database"
	"portfolio/internal/github"
	"portfolio/internal/handlers"
)

func main() {
	cfgPath := envOr("CONFIG_PATH", "config.json")
	cfg, err := config.Load(cfgPath)
	if err != nil {
		log.Fatalf("config: %v", err)
	}

	dbPath := envOr("DB_PATH", "analytics.json")
	db, err := database.Open(dbPath)
	if err != nil {
		log.Fatalf("database: %v", err)
	}
	defer db.Close()

	gh := github.NewClient(cfg.GithubUsername)
	h := handlers.New(cfg, gh, db)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/health", h.Health)
	mux.HandleFunc("GET /api/about", h.About)
	mux.HandleFunc("GET /api/skills", h.Skills)
	mux.HandleFunc("GET /api/contact", h.Contact)
	mux.HandleFunc("GET /api/projects", h.Projects)
	mux.HandleFunc("GET /api/analytics", h.Analytics)
	mux.HandleFunc("POST /api/analytics/click", h.TrackClick)
	mux.HandleFunc("POST /api/analytics/download", h.TrackDownload)
	mux.HandleFunc("GET /api/resume", h.Resume)

	// Serve Vue build in production
	distDir := envOr("STATIC_DIR", "../frontend/dist")
	if info, err := os.Stat(distDir); err == nil && info.IsDir() {
		mux.Handle("/", spaHandler(distDir))
		log.Printf("serving frontend from %s", distDir)
	}

	addr := envOr("PORT", "8080")
	log.Printf("portfolio API running on http://localhost:%s", addr)
	log.Fatal(http.ListenAndServe(":"+addr, handlers.CORS(mux)))
}

func envOr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func spaHandler(distDir string) http.Handler {
	fileServer := http.FileServer(http.Dir(distDir))
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := filepath.Join(distDir, r.URL.Path)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			http.ServeFile(w, r, filepath.Join(distDir, "index.html"))
			return
		}
		fileServer.ServeHTTP(w, r)
	})
}
