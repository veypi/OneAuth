 <!--
 * index.vue
 * Copyright (C) 2023 veypi <i@veypi.com>
 * 2023-10-07 18:57
 * Distributed under terms of the MIT license.
 -->
<template>
  <div class="w-full h-full relative">
    <!-- <div class="absolute bg-red-400 left-0 top-0 w-full h-full"></div> -->
    <div class="w-full h-full" :id="eid"></div>
  </div>
</template>

<script lang="ts" setup>
import Cherry from 'cherry-markdown';
import options from './options'
import { computed, onMounted, ref, watch } from 'vue';
import { CherryOptions } from 'cherry-markdown/types/cherry';

let editor = {} as Cherry;
let emits = defineEmits<{
  (e: 'updated', v: string): void
}>()
let props = withDefaults(defineProps<{
  eid?: string,
  content?: string
  preview?: boolean
}>(),
  {
    eid: 'v-editor',
    content: '',
    preview: false,
  }
)

watch(computed(() => props.preview), (e) => {
  if (e) {
    let des = editor.getValue()
    console.log(des)
    emits('updated', des)
  }
  set_mode(e)
})
watch(computed(() => props.content), (e) => {
  if (e) {
    editor.setValue(e)
  }
})


const set_mode = (preview: boolean) => {
  editor.switchModel(preview ? 'previewOnly' : 'edit&preview')
}

const init = () => {
  let config = {
    value: props.content,
    id: props.eid,
    // isPreviewOnly: props.preview,
    callback: {},
  } as CherryOptions;
  config.callback.afterInit = () => {
  }
  editor = new Cherry(Object.assign({}, options, config));
  set_mode(props.preview)
}


onMounted(() => {
  init()
})
</script>

<style>
iframe.cherry-dialog-iframe {
  width: 100%;
  height: 100%;
}

.cherry {
  background: none;
  box-shadow: none;
}

.cherry-previewer {
  background: none;
  border: none;
}

.cherry-toolbar {
  box-shadow: none;
}
</style>

