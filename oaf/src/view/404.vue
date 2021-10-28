<template>
  <div class="flex justify-center items-center">
    <div class="text-center text-xl">
      <one-icon style="font-size: 200px">404</one-icon>
      <span>
      路径失效啦! {{count}}秒
      </span>
    </div>
  </div>
</template>

<script lang="ts" setup>
import {useRouter, useRoute} from 'vue-router'
import {onMounted, ref} from "vue";

const route = useRoute()
const router = useRouter()
let count = ref(5)
onMounted(() => {
  console.log([route.path, route.params])
  let timer = setInterval(()=> {
    count.value--
    if (count.value === 0) {
      router.push('/')
      clearInterval(timer)
    }
  }, 1000)
})
</script>
<style scoped>

</style>
