<template>
  <NuxtLoadingIndicator />
  <NuxtLayout>
    <NuxtPage />
  </NuxtLayout>
</template>
<script setup lang="ts">


let app = useAppConfig()

onMounted(() => {
  app.layout.size = [document.body.clientWidth, document.body.clientHeight]
  window.onresize = () => {
    app.layout.size = [document.body.clientWidth, document.body.clientHeight]
  }
})

</script>

<style >
body,
html {
  height: 100vh;
  width: 100vw;
  margin: 0;
  padding: 0;
  background: var(--base-bg);
  color: var(--base-txt);
}

/* 周围滑动留白 */
html {
  overflow: hidden;
}

body {
  overflow: auto;
}

#__nuxt {
  height: 100vh;
  width: 100vw;
}
</style>
