<template>
  <div class="w-full px-3">
    <div class="h-16 flex justify-between items-center">
      <span style="color: #777">
        我的云盘
      </span>
      <span class="cursor-pointer"
            style="color:#f36828">文件中心</span>
    </div>
    <div class="">
      {{ usr.Used }} KB / {{ usr.Space }} GB
      <n-progress
        type="line"
        color="#0f0"
        rail-color="#fff"
        :percentage="1"
        indicator-text-color="#f00"
      />
    </div>
    <div class="flex justify-center">
      <n-button @click="showModal=true" type="primary">获取挂载链接</n-button>
    </div>
    <n-modal v-model:show="showModal">
      <n-card style="width: 600px;" title="模态框" :bordered="false" size="huge">
        <template #header-extra> 噢！ </template>
        内容
        <template #footer> 尾部 </template>
      </n-card>
    </n-modal>
    <hr class="mt-10" style="border:none;border-top:1px solid #777;">
  </div>
</template>
<script lang="ts" setup>
import {modelsUser} from '../models'
import {ref} from 'vue'
let showModal = ref(false)
let props = withDefaults(defineProps<{
  usr: modelsUser
}>(), {})
</script>
<style scoped>
</style>
