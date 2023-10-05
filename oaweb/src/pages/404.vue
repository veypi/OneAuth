<template>
  <div class="flex justify-center items-center w-full h-full">
    <div class="text-center text-xl">
      <q-icon style="font-size: 200px" name="svguse:#icon-404"></q-icon>
      <div>
        路径失效啦! {{ count }}秒
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { useRouter, useRoute } from 'vue-router'
import { onBeforeUnmount, onMounted, ref } from "vue";

const route = useRoute()
const router = useRouter()
let count = ref(5)
let timer = ref()
onMounted(() => {
  timer.value = setInterval(() => {
    count.value--
    if (count.value === 0) {
      router.back()
      // router.push('/')
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
