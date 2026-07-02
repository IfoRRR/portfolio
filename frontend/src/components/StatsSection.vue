<script setup>
import { computed } from 'vue'

const props = defineProps({
  stats: { type: Object, default: null },
})

const topLinks = computed(() => {
  if (!props.stats?.links) return []
  return [...props.stats.links].sort((a, b) => b.count - a.count).slice(0, 5)
})
</script>

<template>
  <section id="stats">
    <div class="container">
      <h2 class="section-title">Статистика</h2>
      <p class="section-sub">Переходы по ссылкам и скачивания резюме в реальном времени</p>

      <div v-if="stats" class="stats-grid">
        <div class="stat-card glass stat-main">
          <div class="stat-icon">🔗</div>
          <div class="stat-value mono">{{ stats.total_clicks }}</div>
          <div class="stat-label">переходов по ссылкам</div>
        </div>

        <div class="stat-card glass stat-main">
          <div class="stat-icon">📄</div>
          <div class="stat-value mono">{{ stats.resume_downloads }}</div>
          <div class="stat-label">скачиваний резюме</div>
        </div>

        <div class="stat-card glass stat-details">
          <h3 class="details-title mono">топ ссылок</h3>
          <ul v-if="topLinks.length" class="links-list">
            <li v-for="item in topLinks" :key="item.link + item.type">
              <span class="link-name">{{ item.link }}</span>
              <span class="link-type mono">{{ item.type }}</span>
              <span class="link-count mono">{{ item.count }}</span>
            </li>
          </ul>
          <p v-else class="empty mono">пока нет данных — кликни на проект!</p>
        </div>
      </div>
    </div>
  </section>
</template>

<style scoped>
.stats-grid {
  display: grid;
  grid-template-columns: 1fr 1fr 1.5fr;
  gap: 1.25rem;
}

.stat-card {
  padding: 2rem;
}

.stat-main {
  text-align: center;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  transition: transform 0.2s;
}

.stat-main:hover {
  transform: scale(1.03);
}

.stat-icon {
  font-size: 2rem;
  margin-bottom: 0.5rem;
}

.stat-value {
  font-size: 3rem;
  font-weight: 700;
  background: linear-gradient(135deg, var(--accent2), var(--accent1));
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  line-height: 1;
  margin-bottom: 0.5rem;
}

.stat-label {
  color: var(--muted);
  font-size: 0.9rem;
}

.details-title {
  color: var(--accent2);
  font-size: 0.8rem;
  text-transform: uppercase;
  letter-spacing: 0.08em;
  margin-bottom: 1rem;
}

.links-list {
  list-style: none;
  display: flex;
  flex-direction: column;
  gap: 0.6rem;
}

.links-list li {
  display: grid;
  grid-template-columns: 1fr auto auto;
  gap: 0.75rem;
  align-items: center;
  padding: 0.5rem 0;
  border-bottom: 1px solid var(--border);
  font-size: 0.85rem;
}

.link-name {
  color: var(--text);
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.link-type {
  color: var(--muted);
  font-size: 0.7rem;
  padding: 0.15rem 0.5rem;
  border: 1px solid var(--border);
  border-radius: 999px;
}

.link-count {
  color: var(--accent1);
  font-weight: 700;
}

.empty {
  color: var(--muted);
  font-size: 0.85rem;
}

@media (max-width: 768px) {
  .stats-grid {
    grid-template-columns: 1fr;
  }
}
</style>
