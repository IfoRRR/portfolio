<script setup>
defineProps({
  skills: { type: Array, required: true },
})
</script>

<template>
  <section id="skills">
    <div class="container">
      <h2 class="section-title">Навыки</h2>
      <p class="section-sub">Инструменты и технологии</p>

      <div class="skills-grid">
        <div
          v-for="(skill, i) in skills"
          :key="skill.name"
          class="skill-card glass"
          :style="{ animationDelay: `${i * 0.08}s` }"
        >
          <span class="skill-name">{{ skill.name }}</span>
        </div>
      </div>
    </div>
  </section>
</template>

<style scoped>
.skills-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(180px, 1fr));
  gap: 0.75rem;
}

.skill-card {
  padding: 1rem 1.1rem;
}

.skill-name {
  font-weight: 600;
}
</style>
