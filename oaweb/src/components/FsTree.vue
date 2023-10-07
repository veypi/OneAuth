 <!--
 * FsTree.vue
 * Copyright (C) 2023 veypi <i@veypi.com>
 * 2023-10-06 15:35
 * Distributed under terms of the MIT license.
 -->
<template>
  <div>
    <div :style="{ paddingLeft: depth * 2 + 'rem' }"
      class="cursor-pointer rounded-full h-8 pr-4 flex items-center hover:bg-gray-100" @click="toggle">
      <q-icon class="transition-all mx-2" :class="[expand ? 'rotate-90' :
        '']" style="font-size: 24px;" :name="root.type ===
    'directory' ? 'v-caret-right' : 'v-file'"> </q-icon>
      <div>
        {{ root.filename }}
      </div>
      <div class="grow"></div>
      <div>{{ new Date(root.lastmod).toLocaleString() }}</div>
    </div>
    <div v-if="expand">
      <template v-for="(s, si) of subs" :key="si">
        <FsTree :root="s" :depth="depth + 1"></FsTree>
      </template>
    </div>
  </div>
</template>

<script lang="ts" setup>
import FsTree from './FsTree.vue'
import { ref } from 'vue';
import oafs, { fileProps } from 'src/libs/oafs';
import { util } from 'src/libs';




let expand = ref(false)
let subs = ref([] as fileProps[])

let props = withDefaults(defineProps<{
  root: fileProps,
  depth?: number,
}>(),
  {
    depth: 0
  }
)

const toggle = () => {
  if (props.root.type === 'file') {
    util.goto('/file' + props.root.filename)
    return
  }
  if (!expand.value) {
    oafs.dav().dir(props.root.filename).then(
      (e: any) => {
        subs.value = e
        expand.value = true
      })
    return
  }
  expand.value = !expand.value
}
</script>

<style scoped></style>

