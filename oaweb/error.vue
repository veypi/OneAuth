 <!--
 * error.vue
 * Copyright (C) 2024 veypi <i@veypi.com>
 * 2024-06-06 17:59
 * Distributed under terms of the MIT license.
 -->

<template>
  <div class="flex justify-center items-center w-full h-full">
    <div class="text-center text-xl">
      <OneIcon name="404" style="font-size: 200px"></OneIcon>
      <div v-if='error && error.statusCode === 404'>
        路径失效啦! {{ count }}秒
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { OneIcon } from '@veypi/one-icon'
import type { NuxtError } from '#app'

const props = defineProps({
  error: Object as () => NuxtError
})

const router = useRouter()
let count = ref(5)
let timer = ref()
onMounted(() => {
  timer.value = setInterval(() => {
    count.value--
    if (count.value === 0) {
      router.push('/')
      clearInterval(timer.value)
    }
  }, 1000)
})
onBeforeUnmount(() => {
  if (timer.value) {
    clearInterval(timer.value)
  }

})
</script>
<style scoped></style>
