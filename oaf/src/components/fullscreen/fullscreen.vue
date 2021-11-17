<template>
  <div @click="handleFullscreen">
    <one-icon>{{ props.modelValue ? 'fullscreen-exit' : 'fullscreen' }}</one-icon>
  </div>
</template>

<script lang="ts" setup>
// @ts-nocheck
import {onMounted} from "vue";

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

onMounted(() => {
  let isFullscreen =
    document.fullscreenElement ||
    document.mozFullScreenElement ||
    document.webkitFullscreenElement ||
    document.fullScreen ||
    document.mozFullScreen ||
    document.webkitIsFullScreen
  isFullscreen = !!isFullscreen
  document.addEventListener('fullscreenchange', () => {
    emit('update:modelValue', !props.modelValue)
  })
  document.addEventListener('mozfullscreenchange', () => {
    emit('update:modelValue', !props.modelValue)
  })
  document.addEventListener('webkitfullscreenchange', () => {
    emit('update:modelValue', !props.modelValue)
  })
  document.addEventListener('msfullscreenchange', () => {
    emit('update:modelValue', !props.modelValue)
  })
  emit('update:modelValue', isFullscreen)
})

</script>

<style>
</style>
