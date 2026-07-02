<script setup>
import { ref, onMounted } from 'vue'
import { api } from './api/client.js'
import HeroSection from './components/HeroSection.vue'
import AboutSection from './components/AboutSection.vue'
import ProjectsSection from './components/ProjectsSection.vue'
import SkillsSection from './components/SkillsSection.vue'
import ContactSection from './components/ContactSection.vue'
import StatsSection from './components/StatsSection.vue'
import NavBar from './components/NavBar.vue'
import BackgroundBlobs from './components/BackgroundBlobs.vue'

const about = ref(null)
const projects = ref([])
const stats = ref(null)
const loading = ref(true)
const error = ref(null)

async function loadData() {
  try {
    const [aboutData, projectsData, statsData] = await Promise.all([
      api.getAbout(),
      api.getProjects(),
      api.getAnalytics(),
    ])
    about.value = aboutData
    projects.value = projectsData
    stats.value = statsData
  } catch (e) {
    error.value = e.message
  } finally {
    loading.value = false
  }
}

async function refreshStats() {
  try {
    stats.value = await api.getAnalytics()
  } catch {
    // ignore
  }
}

onMounted(loadData)
</script>

<template>
  <BackgroundBlobs />
  <NavBar />

  <main v-if="loading" class="loading-screen">
    <div class="loader mono">loading portfolio...</div>
  </main>

  <main v-else-if="error" class="loading-screen">
    <p class="error-msg">⚠ {{ error }}</p>
    <p class="hint mono">Запусти Go-бэкенд: cd backend && go run .</p>
  </main>

  <main v-else>
    <HeroSection :about="about" @stats-updated="refreshStats" />
    <AboutSection :about="about" />
    <ProjectsSection :projects="projects" @stats-updated="refreshStats" />
    <SkillsSection :skills="about.skills" />
    <StatsSection :stats="stats" />
    <ContactSection :contact="about.contact" @stats-updated="refreshStats" />
  </main>

  <footer class="footer mono">
    <span>© 2026 {{ about?.name }}</span>
    <span class="footer-dot">·</span>
    <span>Go + Vue</span>
  </footer>
</template>

<style scoped>
.loading-screen {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 1rem;
}

.loader {
  color: var(--accent2);
  animation: pulse-glow 1.5s ease infinite;
}

.error-msg {
  color: var(--accent1);
  font-size: 1.1rem;
}

.hint {
  color: var(--muted);
  font-size: 0.85rem;
}

.footer {
  text-align: center;
  padding: 2rem;
  color: var(--muted);
  font-size: 0.8rem;
  display: flex;
  justify-content: center;
  gap: 0.5rem;
}

.footer-dot {
  opacity: 0.4;
}
</style>
