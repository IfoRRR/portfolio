<script setup>
import { api } from '../api/client.js'

defineProps({
  projects: { type: Array, required: true },
})

const emit = defineEmits(['stats-updated'])

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
  return '#9b5de5'
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
</script>

<template>
  <section id="projects">
    <div class="container">
      <h2 class="section-title">Проекты</h2>
      <p class="section-sub">Мои репозитории с GitHub — клик засчитывается в статистику</p>

      <div class="projects-grid">
        <article
          v-for="(project, i) in projects"
          :key="project.name"
          class="project-card glass"
          :class="{ featured: project.featured }"
          :style="{ animationDelay: `${i * 0.1}s` }"
          @click="openProject(project)"
        >
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

          <p class="project-desc">
            {{ project.description || 'Pet-проект на Go' }}
          </p>

          <div class="project-meta mono">
            <span v-if="project.stars">★ {{ project.stars }}</span>
            <span v-if="project.forks">⑂ {{ project.forks }}</span>
            <span>{{ formatDate(project.updated_at) }}</span>
          </div>

          <div v-if="project.featured" class="featured-tag">★ featured</div>
        </article>
      </div>
    </div>
  </section>
</template>

<style scoped>
.projects-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(320px, 1fr));
  gap: 1.5rem;
}

.project-card {
  padding: 1.75rem;
  cursor: pointer;
  position: relative;
  transition: transform 0.3s, box-shadow 0.3s, border-color 0.3s;
  animation: fade-up 0.6s ease forwards;
  opacity: 0;
}

.project-card:hover {
  transform: translateY(-6px) rotate(-1deg);
  box-shadow: 0 12px 40px rgba(155, 93, 229, 0.25);
  border-color: rgba(155, 93, 229, 0.4);
}

.project-card.featured {
  border-color: rgba(0, 245, 212, 0.3);
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
  color: var(--accent2);
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

.featured-tag {
  position: absolute;
  top: -0.5rem;
  right: 1rem;
  background: linear-gradient(135deg, var(--accent2), var(--accent3));
  color: var(--bg);
  font-size: 0.65rem;
  font-family: 'Space Mono', monospace;
  padding: 0.2rem 0.6rem;
  border-radius: 999px;
  font-weight: 700;
}
</style>
