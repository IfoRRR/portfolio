const BASE = '/api'

async function request(path, options = {}) {
  const res = await fetch(`${BASE}${path}`, {
    headers: { 'Content-Type': 'application/json', ...options.headers },
    ...options,
  })
  if (!res.ok) {
    const err = await res.json().catch(() => ({}))
    throw new Error(err.error || `HTTP ${res.status}`)
  }
  return res.json()
}

export const api = {
  getAbout: () => request('/about'),
  getProjects: () => request('/projects'),
  getSkills: () => request('/skills'),
  getContact: () => request('/contact'),
  getAnalytics: () => request('/analytics'),
  trackClick: (link, type = 'link') =>
    request('/analytics/click', {
      method: 'POST',
      body: JSON.stringify({ link, type }),
    }),
  trackDownload: () =>
    request('/analytics/download', { method: 'POST' }),
  resumeUrl: () => `${BASE}/resume`,
}

export function trackedLink(link, type = 'link') {
  return async (event) => {
    try {
      await api.trackClick(link, type)
    } catch {
      // не блокируем переход при ошибке аналитики
    }
  }
}
