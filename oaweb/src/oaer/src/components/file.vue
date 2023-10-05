<template>
  <div class="w-full px-3">
    <div class="h-16 flex justify-between items-center">
      <span style="color: #777">
        我的云盘
      </span>
      <span class="cursor-pointer" style="color:#f36828">文件中心</span>
    </div>
    <div class="">
      {{ usr.Used }} KB / {{ usr.Space }} GB
      <n-progress type="line" color="#0f0" rail-color="#fff" :percentage="1" indicator-text-color="#f00" />
    </div>
    <div class="flex justify-center">
      <n-button @click="showModal = true" type="primary">获取挂载链接</n-button>
    </div>
    <n-modal v-model:show="showModal">
      <n-card style="width: 600px;" title="云盘挂载地址" :bordered="false" size="huge">
        <template #header-extra>复制</template>
        {{ Cfg.userFileUrl() }}
        <template #footer> 挂载说明</template>
      </n-card>
    </n-modal>
    <hr class="mt-10" style="border:none;border-top:1px solid #777;">
  </div>
</template>
<script lang="ts" setup>
import { createClient } from 'webdav'
import { modelsUser } from '../models'
import { onMounted, ref } from 'vue'
import { Cfg } from '../api'

let showModal = ref(false)
let props = withDefaults(defineProps<{
  usr: modelsUser
}>(), {})

let client = createClient('http://127.0.0.1:4001/file/usr/',
  { headers: { auth_token: localStorage.getItem('auth_token') as string } })
onMounted(() => {
  client.stat('').then((e) => {
    console.log(e)
  }).catch((e) => {
    console.log(e)
  })
})
</script>
<style scoped></style>
