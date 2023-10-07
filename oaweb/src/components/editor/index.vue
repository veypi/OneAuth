 <!--
 * index.vue
 * Copyright (C) 2023 veypi <i@veypi.com>
 * 2023-10-07 18:57
 * Distributed under terms of the MIT license.
 -->
<template>
  <div :id="eid"></div>
</template>

<script lang="ts" setup>
import Cherry from 'cherry-markdown';
import options from './options'
import { onMounted } from 'vue';

let editor = {} as Cherry;
let emits = defineEmits<{
  (e: 'updated', v: string): void
}>()
let props = withDefaults(defineProps<{
  eid?: string,
  content?: string
}>(),
  {
    eid: 'v-editor',
    content: ''
  }
)

onMounted(() => {
  let config = Object.assign({}, options, {
    value: props.content, id:
      props.eid
  });
  editor = new Cherry(config);
})
</script>

<style>
iframe.cherry-dialog-iframe {
  width: 100%;
  height: 100%;
}
</style>

