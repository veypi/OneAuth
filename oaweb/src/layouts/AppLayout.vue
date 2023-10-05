 <!--
 * AppLayout.vue
 * Copyright (C) 2023 veypi <i@veypi.com>
 * 2023-10-04 10:46
 * Distributed under terms of the MIT license.
 -->
<template>
  <div>
    <h1>{{ app.name }}</h1>
    <router-view :data="{ a: 1 }" />
  </div>
</template>

<script lang="ts" setup>
import msg from '@veypi/msg';
import api from 'src/boot/api';
import { modelsApp } from 'src/models';
import { useMenuStore } from 'src/stores/menu';
import { computed, watch, ref, onMounted, provide, onBeforeUnmount } from 'vue';
import { useRoute } from 'vue-router';
let route = useRoute();
let menu = useMenuStore()

let id = computed(() => route.params.id)
let app = ref({} as modelsApp)

provide('app', app)

const sync_app = () => {
  api.app.get(id.value as string).then((e: modelsApp) => {
    app.value = e
  }).catch(e => {
    msg.Warn('sync app data failed: ' + e)
  })
}
watch(id, () => {
  sync_app()
})

onMounted(() => {
  sync_app()
  menu.set([

  ])
})
onBeforeUnmount(() => {
  menu.load_default()

})
</script>

<style scoped></style>

