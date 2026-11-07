<script setup>
import { api } from '../api/client.js'

const props = defineProps({
  projects: { type: Array, required: true },
})

const emit = defineEmits(['stats-updated'])

const petProjectNames = new Set(['WEB-socket', 'URL-SHORTENER', 'meeting-room-mvp'])
const companyCaseNames = new Set(['vk-quiz', 'vk-kanban'])

function projectGroup(project) {
  if (companyCaseNames.has(project.name)) return 'company'
  if (petProjectNames.has(project.name)) return 'pet'
  return 'pet'
}

function projectLabel(project) {
  return projectGroup(project) === 'company' ? 'VK case' : 'Pet project'
}

const langColors = {
  Go: '#00ADD8',
  Vue: '#42b883',
  JavaScript: '#f7df1e',
  HTML: '#e34c26',
  CSS: '#264de4',
}

function langColor(lang) {
  return langColors[lang] || varAccent()
}

function varAccent() {
  return '#F72585'
}

async function openProject(project) {
  try {
    await api.trackClick(project.name, 'project')
    emit('stats-updated')
  } catch {
    // ignore
  }
  window.open(project.url, '_blank')
}

function formatDate(iso) {
  if (!iso) return ''
  return new Date(iso).toLocaleDateString('ru-RU', {
    year: 'numeric',
    month: 'short',
  })
}

const petProjects = () => props.projects.filter((project) => projectGroup(project) === 'pet')
const companyCases = () => props.projects.filter((project) => projectGroup(project) === 'company')
</script>

<template>
  <section id="projects">
    <div class="container">
      <h2 class="section-title">Проекты</h2>
      <p class="section-sub">Избранные pet-проекты и кейсы для компаний.</p>

      <div class="project-group">
        <div class="group-heading">
          <span class="group-kicker mono">01</span>
          <h3>Pet-проекты</h3>
        </div>
        <div class="projects-grid">
          <article
            v-for="(project, i) in petProjects()"
            :key="project.name"
            class="project-card glass"
            :style="{ animationDelay: `${i * 0.1}s` }"
            @click="openProject(project)"
          >
            <span class="project-label mono">{{ projectLabel(project) }}</span>
            <div class="card-top">
              <h3 class="project-name">{{ project.name }}</h3>
              <span
                v-if="project.language"
                class="lang-badge mono"
                :style="{ '--lang-color': langColor(project.language) }"
              >
                {{ project.language }}
              </span>
            </div>
            <p class="project-desc">{{ project.description || 'Pet-проект на Go' }}</p>
            <div class="project-meta mono">
              <span v-if="project.stars">★ {{ project.stars }}</span>
              <span v-if="project.forks">⑂ {{ project.forks }}</span>
              <span>{{ formatDate(project.updated_at) }}</span>
            </div>
          </article>
        </div>
      </div>

      <div class="project-group">
        <div class="group-heading">
          <span class="group-kicker mono">02</span>
          <h3>Кейсы для компаний</h3>
        </div>
        <div class="projects-grid">
          <article
            v-for="(project, i) in companyCases()"
            :key="project.name"
            class="project-card company-card glass"
            :style="{ animationDelay: `${i * 0.1}s` }"
            @click="openProject(project)"
          >
            <span class="project-label mono">{{ projectLabel(project) }}</span>
            <div class="card-top">
              <h3 class="project-name">{{ project.name }}</h3>
              <span
                v-if="project.language"
                class="lang-badge mono"
                :style="{ '--lang-color': langColor(project.language) }"
              >
                {{ project.language }}
              </span>
            </div>
            <p class="project-desc">{{ project.description || 'VK case' }}</p>
            <div class="project-meta mono">
              <span v-if="project.stars">★ {{ project.stars }}</span>
              <span v-if="project.forks">⑂ {{ project.forks }}</span>
              <span>{{ formatDate(project.updated_at) }}</span>
            </div>
          </article>
          <article class="project-card placeholder-card glass">
            <span class="project-label mono">Planned</span>
            <h3 class="project-name">Т-Банк</h3>
            <p class="project-desc">Заготовка для будущего кейса: задача, решение, результат и стек.</p>
            <span class="placeholder-note mono">case study in progress</span>
          </article>
        </div>
      </div>
    </div>
  </section>
</template>

<style scoped>
.project-group + .project-group {
  margin-top: 4rem;
}

.group-heading {
  display: flex;
  align-items: baseline;
  gap: 0.75rem;
  margin-bottom: 1.25rem;
}

.group-heading h3 {
  font-size: 1.35rem;
}

.group-kicker {
  color: var(--accent1);
  font-size: 0.75rem;
}

.projects-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(320px, 1fr));
  gap: 1.5rem;
}

.project-card {
  padding: 1.75rem;
  cursor: pointer;
  position: relative;
  transition: transform 0.2s, border-color 0.2s;
  animation: fade-up 0.6s ease forwards;
  opacity: 0;
}

.project-card:hover {
  transform: translateY(-3px);
  border-color: var(--accent1);
}

.company-card {
  border-color: rgba(247, 37, 133, 0.35);
}

.card-top {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  gap: 0.5rem;
  margin-bottom: 0.75rem;
}

.project-name {
  font-size: 1.2rem;
  font-weight: 700;
  color: var(--text);
}

.lang-badge {
  font-size: 0.7rem;
  padding: 0.2rem 0.6rem;
  border-radius: 999px;
  background: color-mix(in srgb, var(--lang-color) 20%, transparent);
  color: var(--lang-color);
  border: 1px solid color-mix(in srgb, var(--lang-color) 40%, transparent);
  white-space: nowrap;
}

.project-desc {
  color: var(--muted);
  font-size: 0.95rem;
  margin-bottom: 1.25rem;
  min-height: 3rem;
}

.project-meta {
  display: flex;
  gap: 1rem;
  font-size: 0.75rem;
  color: var(--muted);
  opacity: 0.7;
}

.project-label {
  display: inline-block;
  margin-bottom: 0.8rem;
  color: var(--accent1);
  font-size: 0.68rem;
  text-transform: uppercase;
  letter-spacing: 0.08em;
}

.placeholder-card {
  border-style: dashed;
  cursor: default;
}

.placeholder-note {
  color: var(--muted);
  font-size: 0.72rem;
}
</style>
