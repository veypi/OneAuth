<template>
  <div>
    <div @click="setValue(true)">
      <slot>
      </slot>
    </div>
    <div @click.self="setValue(false)" class="core" style="height: 100vh;width: 100vw;" v-if="props.modelValue">
      <div style="height: 100%; width: 300px" class="core-right">
        <transition appear enter-active-class=" animate__slideInRight">
          <div class="flex right-title animate__animated animate__faster px-3">
            <div class="flex-grow text-left" style="font-size: 1.2rem">
              <slot name="title"></slot>
            </div>
            <div class="flex-grow-0 flex items-center h-full">
              <OneIcon @click="setValue(false)" color="#fff" style="font-size: 24px">close</OneIcon>
            </div>
          </div>
        </transition>
        <div class="right-main">
          <transition appear enter-active-class="animate__slideInDown">
            <div class="right-main-core animate__animated animate__faster" :style="{ 'background': backgound }">
              <slot name="main"></slot>
            </div>
          </transition>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { OneIcon } from '@veypi/one-icon'
import { computed, watch } from 'vue'

let emits = defineEmits<{
  (e: 'update:modelValue', v: boolean): void
}>()
let props = withDefaults(defineProps<{
  isDark?: boolean,
  modelValue?: boolean
}>(), {})

let backgound = computed(() => {
  return props.isDark ? '#222' : '#eee'
})
watch(props, () => {
})

function setValue(b: boolean) {
  emits('update:modelValue', b)
}

</script>

<style scoped>
.core {
  position: fixed;
  left: 0;
  top: 0;
  background: rgba(0, 0, 0, 0.4);
  z-index: 100;
}

.core-right {
  position: absolute;
  right: 0;
  top: 0;
}

.right-main {
  width: 100%;
  height: calc(100% - 50px);
  overflow: hidden;
}

.right-main-core {
  height: 100%;
  width: 100%;
  -webkit-animation-delay: 0.4s;
  animation-delay: 0.4s;
  --animate-duration: 400ms;
}

.right-title {
  width: 100%;
  height: 50px;
  line-height: 50px;
  background: linear-gradient(90deg, #f74d22, #fa9243);
}
</style>
