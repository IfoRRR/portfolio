<script setup>
defineProps({
  about: { type: Object, required: true },
})
</script>

<template>
  <section id="about">
    <div class="container">
      <h2 class="section-title">Обо мне</h2>
      <p class="section-sub">Коротко о моём направлении</p>

      <div class="about-card glass">
        <div class="about-grid">
          <div class="about-text">
            <p>{{ about.about }}</p>
          </div>
          <div class="about-highlights">
            <div class="highlight">
              <span class="highlight-num mono">Go</span>
              <span class="highlight-label">основной язык разработки</span>
            </div>
            <div class="highlight">
              <span class="highlight-num mono">API</span>
              <span class="highlight-label">REST и WebSocket-сервисы</span>
            </div>
            <div class="highlight">
              <span class="highlight-num mono">VPS</span>
              <span class="highlight-label">Docker, nginx и Linux</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<style scoped>
.about-card {
  padding: 2.5rem;
}

.about-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 2rem;
  align-items: center;
}

.about-text p {
  font-size: 1.1rem;
  color: var(--muted);
  line-height: 1.8;
}

.about-highlights {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.highlight {
  padding: 1.25rem;
  border-radius: 16px;
  background: rgba(65, 90, 119, 0.18);
  border: 1px solid rgba(247, 37, 133, 0.18);
  display: flex;
  align-items: center;
  gap: 1rem;
  transition: transform 0.2s;
}

.highlight:hover {
  transform: translateX(8px);
}

.highlight-num {
  font-size: 1.5rem;
  font-weight: 700;
  color: var(--accent2);
  min-width: 3rem;
}

.highlight-label {
  color: var(--muted);
  font-size: 0.9rem;
}

@media (max-width: 768px) {
  .about-grid {
    grid-template-columns: 1fr;
  }
}
</style>
