<template>
  <div @click="click">
    <input enctype="multipart/form-data" ref="file" name="files" :multiple="multiple" type="file" hidden @change="upload">
    <slot></slot>
  </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue';
import {oafs} from '@veypi/oaer'

let file = ref<HTMLInputElement>()
let emits = defineEmits<{
  (e: 'success', v: string): void
  (e: 'failed'): void
}>()
let props = withDefaults(defineProps<{
  multiple?: boolean,
  renames?: string,
  dir?: string,
}>(), {
  multiple: false,
  renames: ''
})

function click() {
  file.value?.dispatchEvent(new MouseEvent('click'))
}

const upload = (evt: Event) => {
  evt.preventDefault()
  let f = (evt.target as HTMLInputElement).files as FileList
  oafs.upload(f, props.dir, props.renames?.split(/[, ]+/)).then((e: any) => {
    console.log(e)
    emits('success', props.multiple ? e : e[0])
  })
}


</script>

<style scoped></style>
