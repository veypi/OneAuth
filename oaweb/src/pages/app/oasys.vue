 <!--
 * oasys.vue
 * Copyright (C) 2023 veypi <i@veypi.com>
 * 2023-10-18 23:45
 * Distributed under terms of the MIT license.
 -->
<template>
  <div>
    <div v-if="data.server?.id">
      <div class="text-2xl mb-4">
        消息服务
        <div class="float-right text-sm">{{ new Date(data.server.time).toLocaleString() }}</div>
      </div>
      <div class="">
        <div class="w-full">ID: {{ data.server?.id }}</div>
        <div class="flex gap-8">
          <div>CPU占用: {{ data.statsz.cpu }}%</div>
          <div>内存占用: {{ (data.statsz.mem / 1024 / 1024).toFixed(2) }}M</div>
          <div>连接数: {{ data.statsz.connections }}</div>
        </div>
        <div>发送: {{ (send_received[1] / 1024).toFixed(2) }} KB/s</div>
        <div>收到: {{ (send_received[0] / 1024).toFixed(2) }} KB/s</div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { computed, ref, watch, onUnmounted } from 'vue';
import { nats } from '@veypi/oaer'


const data = ref({} as any)
const id = computed(() => data.value.server?.id)
const subs: any[] = []
const timer = ref()

let old_data = [0, 0]
const send_received = computed(() => {
  if (!id.value) {
    return [0, 0]
  }
  let os = data.value.statsz.sent.bytes
  let or = data.value.statsz.received.bytes
  let res = [os - old_data[0], or - old_data[1]]
  old_data = [os, or]
  return res
})

watch(id, (_) => {
  timer.value = setInterval(() => {
    nats.request('$SYS.REQ.SERVER.PING').then((m) => {
      data.value = JSON.parse(m)
    })
  }, 1000)
})

watch(computed(() => nats.ready.value), e => {
  if (e) {
    nats.request('$SYS.REQ.SERVER.PING').then((m) => {
      data.value = JSON.parse(m)
      let os = data.value.statsz.sent.bytes
      let or = data.value.statsz.received.bytes
      old_data = [os, or]
    })
  }
}, { immediate: true })

onUnmounted(() => {
  clearInterval(timer.value)
  for (let i of subs) {
    i.unsubscribe()
  }
})

</script>

<style scoped></style>

