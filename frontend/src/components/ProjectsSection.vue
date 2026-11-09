<script setup>
import { api } from '../api/client.js'

const props = defineProps({
  projects: { type: Array, required: true },
})

const emit = defineEmits(['stats-updated'])

const projectDefinitions = [
  {
    key: 'URL-SHORTENER',
    aliases: ['url-shortener'],
    group: 'Pet-проекты',
    name: 'URL Shortener',
    description: 'Сервис коротких ссылок с API, Vue-интерфейсом и хранением данных.',
    language: 'Go + Vue',
    demo: '/shortener/',
    github: 'https://github.com/IfoRRR/URL-SHORTENER',
  },
  {
    key: 'WEB-socket',
    aliases: ['web-socket'],
    group: 'Pet-проекты',
    name: 'WebSocket Chat',
    description: 'Чат с комнатами, real-time сообщениями и REST API.',
    language: 'Go + WebSocket',
    demo: '/chat/',
    github: 'https://github.com/IfoRRR/WEB-socket',
  },
  {
    key: 'vk-quiz',
    aliases: ['quiz'],
    group: 'VK-кейсы',
    name: 'Quiz',
    description: 'Многопользовательский квиз с WebSocket-событиями.',
    language: 'Go + WebSocket',
    demo: '/quiz/',
    github: 'https://github.com/IfoRRR/vk-quiz',
  },
  {
    key: 'vk-kanban',
    aliases: ['kanban', 'cases'],
    group: 'VK-кейсы',
    name: 'Kanban / Cases',
    description: 'Доска задач и демонстрация работы с комнатами.',
    language: 'Go',
    demo: '/board/',
    github: 'https://github.com/IfoRRR/vk-kanban',
  },
  {
    key: 'meeting-room-mvp',
    aliases: ['meeting'],
    group: 'VK-кейсы',
    name: 'Переговорные',
    description: 'Бронирование переговорных комнат без двойных броней.',
    language: 'Go + PostgreSQL',
    demo: '/meeting/',
    github: 'https://github.com/IfoRRR/meeting-room-mvp',
  },
]

function findApiProject(definition) {
  const names = [definition.key, ...definition.aliases]
  return props.projects.find((project) =>
    names.some((name) => project.name?.toLowerCase() === name.toLowerCase()),
  )
}

function trackAndOpen(project, url, type) {
  api.trackClick(project.name, type).then(() => emit('stats-updated')).catch(() => {})
  window.open(url, '_blank', 'noopener,noreferrer')
}
</script>

<template>
  <section id="projects">
    <div class="container">
      <h2 class="section-title">Проекты</h2>
      <p class="section-sub">Два pet-проекта и три практических кейса.</p>

      <div v-for="group in ['Pet-проекты', 'VK-кейсы']" :key="group" class="project-group">
        <h3 class="group-title">{{ group }}</h3>
        <div class="projects-list">
          <article
            v-for="definition in projectDefinitions.filter((item) => item.group === group)"
            :key="definition.key"
            class="project-card"
          >
            <div class="project-copy">
              <p class="project-type mono">{{ definition.language }}</p>
              <h3 class="project-name">{{ definition.name }}</h3>
              <p class="project-desc">{{ definition.description }}</p>
            </div>
            <div class="project-actions">
              <button
                class="project-link project-link-primary"
                @click="trackAndOpen(definition, definition.demo, 'demo')"
              >
                Демо
              </button>
              <button
                class="project-link"
                @click="trackAndOpen(definition, findApiProject(definition)?.url || definition.github, 'project')"
              >
                GitHub
              </button>
            </div>
          </article>
        </div>
      </div>
    </div>
  </section>
</template>

<style scoped>
.project-group + .project-group { margin-top: 3.5rem; }
.group-title {
  margin-bottom: 1rem;
  color: var(--muted);
  font-size: 0.8rem;
  font-weight: 600;
  letter-spacing: 0.08em;
  text-transform: uppercase;
}
.projects-list { display: grid; gap: 0.75rem; }
.project-card {
  display: flex;
  justify-content: space-between;
  gap: 1.5rem;
  align-items: center;
  padding: 1.4rem 1.5rem;
  border: 1px solid var(--border);
  border-radius: 12px;
  background: rgba(255, 255, 255, 0.025);
}
.project-card:hover { border-color: rgba(0, 245, 212, 0.45); }
.project-type {
  margin-bottom: 0.45rem;
  color: var(--accent2);
  font-size: 0.7rem;
}
.project-name { margin-bottom: 0.45rem; font-size: 1.15rem; }
.project-desc { max-width: 680px; color: var(--muted); font-size: 0.9rem; }
.project-actions { display: flex; gap: 0.55rem; flex-shrink: 0; }
.project-link {
  padding: 0.55rem 0.8rem;
  border: 1px solid var(--border);
  border-radius: 7px;
  background: transparent;
  color: var(--text);
  font: inherit;
  font-size: 0.82rem;
  cursor: pointer;
}
.project-link:hover { border-color: var(--accent2); color: var(--accent2); }
.project-link-primary { background: var(--accent2); border-color: var(--accent2); color: var(--bg); }
.project-link-primary:hover { color: var(--bg); opacity: 0.85; }
@media (max-width: 680px) {
  .project-card { display: block; }
  .project-actions { margin-top: 1rem; }
}
</style>
