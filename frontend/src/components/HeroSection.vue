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
      <img
        v-if="about.avatar_url"
        :src="about.avatar_url"
        :alt="about.name"
        class="avatar"
      />

      <p class="hero-kicker mono">Go / backend development</p>
      <h1 class="hero-name">{{ about.name }}</h1>

      <p class="hero-title">{{ about.title }}</p>
      <p class="hero-description">{{ about.about }}</p>

      <div class="hero-actions">
        <a href="#projects" class="btn btn-primary">Проекты</a>
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
  font-size: clamp(2.25rem, 6vw, 4.5rem);
  font-weight: 800;
  line-height: 1.1;
  margin-bottom: 0.5rem;
}

.hero-kicker {
  color: var(--accent2);
  margin-bottom: 1rem;
}

.hero-title {
  font-size: 1.3rem;
  color: var(--muted);
  margin-bottom: 0.5rem;
}

.hero-description {
  max-width: 650px;
  margin: 0 auto 2rem;
  color: var(--muted);
  line-height: 1.7;
}

.hero-actions {
  display: flex;
  gap: 1rem;
  justify-content: center;
  flex-wrap: wrap;
}
</style>
