 <!--
 * fsFile.vue
 * Copyright (C) 2023 veypi <i@veypi.com>
 * 2023-10-08 05:12
 * Distributed under terms of the MIT license.
 -->
<template>
  <div>
    <h1 class="page-h1">云文件中心</h1>
    <div class="px-4">
      <FsTree v-if="root.filename" :root="root"></FsTree>
    </div>
  </div>
</template>

<script lang="ts" setup>
import FsTree from 'src/components/FsTree.vue';
import { oafs, fileProps } from '@veypi/oaer';
import { onMounted, ref, watch } from 'vue';
let root = ref({} as fileProps)

watch(oafs.ready, e => {
  if (e) {
    oafs.dav().stat('/').then(e => {
      root.value = e as fileProps
    })
  }
}, { immediate: true })
onMounted(() => {
})
</script>

<style scoped></style>

