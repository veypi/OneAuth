<template>
  <div class="core rounded-2xl p-3">
    <div class="grid gap-4 grid-cols-5">
      <div class="col-span-2">
        <n-avatar @click="$router.push({name: 'app.main', params: {uuid: core.uuid}})" round :size="80" :src="core.icon">
        </n-avatar>
      </div>
      <div class="col-span-3 grid grid-cols-1 items-center text-left">
        <div class="h-10 flex items-center text-2xl italic font-bold">{{ core.name }}</div>
        <div class="select-all">{{ core.uuid }}</div>
      </div>
    </div>
    <textarea disabled style="background: none;border: none" class="focus:outline-none w-full">{{core.des}}</textarea>
  </div>
</template>
<script setup lang='ts'>
import {defineProps, withDefaults} from "vue";
interface App {
  uuid: ''
  [key: string]: any
}
let props = withDefaults(defineProps<{
  core?: App
}>(), {
  // @ts-ignore
  core: {}
})
</script>
<style scoped>
.core {
  width: 256px;
  background: rgba(146, 145, 145, 0.1);
}
</style>
