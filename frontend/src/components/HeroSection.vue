<script setup>
import { api } from '../api/client.js'

defineProps({
  about: { type: Object, required: true },
})

const emit = defineEmits(['stats-updated'])

async function downloadResume() {
  emit('stats-updated')
  window.open(api.resumeUrl(), '_blank')
}
</script>

<template>
  <section class="hero">
    <div class="container hero-inner fade-up">
      <div class="hero-badge mono">backend developer</div>

      <img
        v-if="about.avatar_url"
        :src="about.avatar_url"
        :alt="about.name"
        class="avatar"
      />

      <h1 class="hero-name">
        {{ about.name.split(' ')[0] }}
        <span class="gradient">{{ about.name.split(' ').slice(1).join(' ') }}</span>
      </h1>

      <p class="hero-title">{{ about.title }}</p>
      <p class="hero-tagline mono">Go · PostgreSQL · Docker · REST API</p>

      <div class="hero-actions">
        <a href="#projects" class="btn btn-primary">Смотреть проекты</a>
        <button class="btn btn-outline" @click="downloadResume">
          ↓ Скачать резюме
        </button>
      </div>
    </div>
  </section>
</template>

<style scoped>
.hero {
  min-height: 100vh;
  display: flex;
  align-items: center;
  padding-top: 5rem;
}

.hero-inner {
  text-align: center;
}

.hero-badge {
  display: inline-block;
  padding: 0.35rem 1rem;
  border: 1px solid var(--accent2);
  border-radius: 999px;
  color: var(--accent2);
  font-size: 0.75rem;
  text-transform: uppercase;
  letter-spacing: 0.1em;
  margin-bottom: 1.5rem;
}

.avatar {
  width: 120px;
  height: 120px;
  border-radius: 50%;
  border: 3px solid transparent;
  background: linear-gradient(var(--bg), var(--bg)) padding-box,
    linear-gradient(135deg, var(--accent2), var(--accent1)) border-box;
  margin-bottom: 1.5rem;
}

.hero-name {
  font-size: clamp(2.5rem, 7vw, 4.5rem);
  font-weight: 800;
  line-height: 1.1;
  margin-bottom: 0.5rem;
}

.gradient {
  display: block;
  background: linear-gradient(135deg, var(--accent2), var(--accent1), var(--accent3));
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.hero-title {
  font-size: 1.3rem;
  color: var(--muted);
  margin-bottom: 0.5rem;
}

.hero-tagline {
  color: var(--accent2);
  font-size: 0.85rem;
  margin-bottom: 2rem;
  opacity: 0.8;
}

.hero-actions {
  display: flex;
  gap: 1rem;
  justify-content: center;
  flex-wrap: wrap;
}
</style>
