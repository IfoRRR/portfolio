<script setup>
defineProps({
  skills: { type: Array, required: true },
})
</script>

<template>
  <section id="skills">
    <div class="container">
      <h2 class="section-title">Навыки</h2>
      <p class="section-sub">Стек, с которым работаю</p>

      <div class="skills-grid">
        <div
          v-for="(skill, i) in skills"
          :key="skill.name"
          class="skill-card glass"
          :style="{ animationDelay: `${i * 0.08}s` }"
        >
          <div class="skill-header">
            <span class="skill-name">{{ skill.name }}</span>
            <span class="skill-pct mono">{{ skill.level }}%</span>
          </div>
          <div class="skill-bar">
            <div
              class="skill-fill"
              :style="{ width: skill.level + '%', '--i': i }"
            ></div>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<style scoped>
.skills-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
  gap: 1.25rem;
}

.skill-card {
  padding: 1.5rem;
  animation: fade-up 0.5s ease forwards;
  opacity: 0;
}

.skill-header {
  display: flex;
  justify-content: space-between;
  margin-bottom: 0.75rem;
}

.skill-name {
  font-weight: 600;
  font-size: 1.05rem;
}

.skill-pct {
  color: var(--accent2);
  font-size: 0.85rem;
}

.skill-bar {
  height: 6px;
  background: rgba(255, 255, 255, 0.06);
  border-radius: 999px;
  overflow: hidden;
}

.skill-fill {
  height: 100%;
  border-radius: 999px;
  background: linear-gradient(90deg, var(--accent3), var(--accent1), var(--accent2));
  background-size: 200% 100%;
  animation: shimmer 3s ease infinite;
  transition: width 1s ease;
}

@keyframes shimmer {
  0% { background-position: 200% 0; }
  100% { background-position: -200% 0; }
}
</style>
