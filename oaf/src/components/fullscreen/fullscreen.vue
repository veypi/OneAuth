<template>
  <div @click="handleFullscreen">
    <one-icon>{{ props.modelValue ? 'fullscreen-exit' : 'fullscreen' }}</one-icon>
  </div>
</template>

<script lang="ts" setup>
// @ts-nocheck
import {onMounted} from 'vue'

let emit = defineEmits<{
  (e: 'update:modelValue', v: boolean): void
}>()

let props = defineProps<{
  modelValue: boolean
}>()

function handleFullscreen() {
  let main = document.body
  if (props.modelValue) {
    if (document.exitFullscreen) {
      document.exitFullscreen()
    } else if (document.mozCancelFullScreen) {
      document.mozCancelFullScreen()
    } else if (document.webkitCancelFullScreen) {
      document.webkitCancelFullScreen()
    } else if (document.msExitFullscreen) {
      document.msExitFullscreen()
    }
  } else {
    if (main.requestFullscreen) {
      main.requestFullscreen()
    } else if (main.mozRequestFullScreen) {
      main.mozRequestFullScreen()
    } else if (main.webkitRequestFullScreen) {
      main.webkitRequestFullScreen()
    } else if (main.msRequestFullscreen) {
      main.msRequestFullscreen()
    }
  }
}

function sync() {
  let isFullscreen =
    document.fullscreenElement ||
    document.mozFullScreenElement ||
    document.webkitFullscreenElement ||
    document.fullScreen ||
    document.mozFullScreen ||
    document.webkitIsFullScreen
  isFullscreen = !!isFullscreen
  emit('update:modelValue', isFullscreen)
}

onMounted(() => {
  document.addEventListener('fullscreenchange', sync)
  document.addEventListener('mozfullscreenchange', sync)
  document.addEventListener('webkitfullscreenchange', sync)
  document.addEventListener('msfullscreenchange', sync)
  sync()
})

</script>

<style>
</style>
