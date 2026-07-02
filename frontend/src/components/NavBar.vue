<script setup>
const links = [
  { href: '#about', label: 'Обо мне' },
  { href: '#projects', label: 'Проекты' },
  { href: '#skills', label: 'Навыки' },
  { href: '#stats', label: 'Статистика' },
  { href: '#contact', label: 'Контакты' },
]
</script>

<template>
  <nav class="nav glass">
    <div class="container nav-inner">
      <a href="#" class="logo mono">&lt;AA /&gt;</a>
      <ul class="nav-links">
        <li v-for="link in links" :key="link.href">
          <a :href="link.href">{{ link.label }}</a>
        </li>
      </ul>
    </div>
  </nav>
</template>

<style scoped>
.nav {
  position: fixed;
  top: 1rem;
  left: 50%;
  transform: translateX(-50%);
  z-index: 100;
  width: calc(100% - 2rem);
  max-width: 900px;
  border-radius: 999px;
}

.nav-inner {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0.75rem 1.5rem;
}

.logo {
  font-weight: 700;
  color: var(--accent2);
  text-decoration: none;
  font-size: 1rem;
}

.nav-links {
  display: flex;
  list-style: none;
  gap: 1.5rem;
}

.nav-links a {
  color: var(--muted);
  text-decoration: none;
  font-size: 0.9rem;
  font-weight: 500;
  transition: color 0.2s;
}

.nav-links a:hover {
  color: var(--text);
}

@media (max-width: 640px) {
  .nav-links {
    gap: 0.75rem;
  }
  .nav-links a {
    font-size: 0.75rem;
  }
}
</style>
