<!--
 * index.vue
 * Copyright (C) 2022 veypi <i@veypi.com>
 * 2022-12-20 23:45
 * Distributed under terms of the Apache license.
 -->
<template>
  <Teleport to="body">
    <transition
      mode="out-in"
      enter-active-class="animate__fadeInUpBig"
      leave-active-class="animate__fadeOutDownBig"
    >
      <div
        v-if="modelValue"
        class="vmodal flex justify-center items-center animate__animated"
        @click.self="emit('update:modelValue', false)"
      >
        <div class="vcore rounded-lg">
          <slot></slot>
        </div>
      </div>
    </transition>
  </Teleport>
</template>

<script lang="ts" setup>
let emit = defineEmits<{
  (e: 'update:modelValue', v: boolean): void
}>()

let props = defineProps<{
  modelValue: boolean
}>()
</script>

<style scoped>
.vmodal {
  position: fixed;
  left: 0;
  top: 0;
  height: 100vh;
  width: 100vw;
}
.vcore {
  background: var(--base-bg-2);
  width: 30rem;
  height: 40rem;
}
</style>
