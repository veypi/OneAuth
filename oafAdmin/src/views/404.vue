<style scoped>
</style>
<template>
  <div>404</div>
</template>
<script lang="ts" setup>
import {useRouter} from 'vue-router'
import {onMounted, ref} from 'vue'

const router = useRouter()
let count = ref(1)
onMounted(() => {
  let timer = setInterval(() => {
    count.value--
    if (count.value === 0) {
      router.push('/')
      clearInterval(timer)
    }
  }, 1000)
})
</script>
