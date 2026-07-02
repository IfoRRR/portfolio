<script setup>
import { api } from '../api/client.js'

defineProps({
  contact: { type: Object, required: true },
})

const emit = defineEmits(['stats-updated'])

async function trackAndOpen(link, type, url) {
  try {
    await api.trackClick(link, type)
    emit('stats-updated')
  } catch {
    // ignore
  }
  window.open(url, '_blank')
}

function openEmail(email) {
  api.trackClick('email', 'contact').then(() => emit('stats-updated')).catch(() => {})
  window.location.href = `mailto:${email}`
}

function openTelegram(tg) {
  const username = tg.replace('@', '')
  trackAndOpen('telegram', 'contact', `https://t.me/${username}`)
}

function openGithub(url) {
  trackAndOpen('github', 'contact', url)
}
</script>

<template>
  <section id="contact">
    <div class="container">
      <h2 class="section-title">Контакты</h2>
      <p class="section-sub">Свяжись со мной — каждый клик учитывается</p>

      <div class="contact-grid">
        <button class="contact-card glass" @click="openEmail(contact.email)">
          <span class="contact-icon">✉</span>
          <span class="contact-label">Email</span>
          <span class="contact-value mono">{{ contact.email }}</span>
        </button>

        <button class="contact-card glass" @click="openTelegram(contact.telegram)">
          <span class="contact-icon">✈</span>
          <span class="contact-label">Telegram</span>
          <span class="contact-value mono">{{ contact.telegram }}</span>
        </button>

        <button class="contact-card glass" @click="openGithub(contact.github)">
          <span class="contact-icon">⌥</span>
          <span class="contact-label">GitHub</span>
          <span class="contact-value mono">IfoRRR</span>
        </button>
      </div>
    </div>
  </section>
</template>

<style scoped>
.contact-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
  gap: 1.25rem;
}

.contact-card {
  padding: 2rem;
  text-align: center;
  cursor: pointer;
  transition: transform 0.25s, border-color 0.25s, box-shadow 0.25s;
  background: none;
  color: inherit;
  font: inherit;
  width: 100%;
}

.contact-card:hover {
  transform: translateY(-4px) rotate(1deg);
  border-color: rgba(255, 107, 203, 0.4);
  box-shadow: 0 8px 32px rgba(255, 107, 203, 0.15);
}

.contact-icon {
  font-size: 2rem;
  display: block;
  margin-bottom: 0.75rem;
}

.contact-label {
  display: block;
  font-weight: 600;
  font-size: 1.1rem;
  margin-bottom: 0.25rem;
}

.contact-value {
  color: var(--muted);
  font-size: 0.85rem;
}
</style>
