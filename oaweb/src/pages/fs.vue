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
import oafs, { fileProps } from 'src/libs/oafs';
import { onMounted, ref } from 'vue';
let root = ref({} as fileProps)

onMounted(() => {
  oafs.dav().stat('/').then(e => {
    console.log(e)
    root.value = e as fileProps
  })
})
</script>

<style scoped></style>

