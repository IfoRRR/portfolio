<template>
  <div class="blobs" aria-hidden="true">
    <div class="blob blob-1"></div>
    <div class="blob blob-2"></div>
    <div class="blob blob-3"></div>
  </div>
</template>

<style scoped>
.blobs {
  position: fixed;
  inset: 0;
  z-index: -1;
  overflow: hidden;
  pointer-events: none;
}

.blob {
  position: absolute;
  border-radius: 50%;
  filter: blur(80px);
  opacity: 0.35;
  animation: float 12s ease-in-out infinite;
}

.blob-1 {
  width: 500px;
  height: 500px;
  background: var(--accent3);
  top: -100px;
  left: -100px;
}

.blob-2 {
  width: 400px;
  height: 400px;
  background: var(--accent1);
  bottom: 10%;
  right: -80px;
  animation-delay: -4s;
}

.blob-3 {
  width: 350px;
  height: 350px;
  background: var(--accent2);
  top: 40%;
  left: 30%;
  animation-delay: -8s;
  opacity: 0.2;
}
</style>
