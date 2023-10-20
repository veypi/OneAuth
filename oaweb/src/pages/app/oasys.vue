 <!--
 * oasys.vue
 * Copyright (C) 2023 veypi <i@veypi.com>
 * 2023-10-18 23:45
 * Distributed under terms of the MIT license.
 -->
<template>
  <div>
    <div v-if="id">
      <div class="text-2xl mb-4">
        消息服务
        <div class="float-right text-sm">{{ new Date(data.now).toLocaleString() }}</div>
      </div>
      <div class="">
        <div class="w-full">ID: {{ id }}</div>
        <div class="flex gap-8">
          <div>CPU占用: {{ data.cpu }}%</div>
          <div>内存占用: {{ (data.mem / 1024 / 1024).toFixed(2) }}M</div>
          <div>连接数: {{ data.connections }}</div>
        </div>
        <div>发送: {{ (send_received[0] / 1024).toFixed(2) }} KB/s</div>
        <div>收到: {{ (send_received[1] / 1024).toFixed(2) }} KB/s</div>
      </div>
      <div class="grid grid-cols-4 gap-4 mt-10" v-if="conns.length">
        <div>ID</div>
        <div>Name</div>
        <div>运行时间</div>
        <div>订阅主题</div>
        <template v-for="c of conns" :key="c.cid">
          <div>{{ c.cid }}</div>
          <div>{{ c.name || '无' }}</div>
          <div>{{ new Date(c.start).toLocaleString() }}</div>
          <div>{{ c.subscriptions_list ?
            c.subscriptions_list.sort().join(' ') : '' }}</div>
        </template>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { computed, ref, watch, onUnmounted } from 'vue';
import { nats } from '@veypi/oaer'
import api from 'src/boot/api'


const data = ref({} as any)
const conns = ref<any[]>([])
const id = computed(() => data.value.server_id)
const subs: any[] = []
const timer = ref()

let old_data = [0, 0]
const send_received = computed(() => {
  if (!id.value) {
    return [0, 0]
  }
  let os = data.value.out_bytes
  let or = data.value.in_bytes
  let res = [os - old_data[0], or - old_data[1]]
  old_data = [os, or]
  return res
})

watch(id, (_) => {
  timer.value = setInterval(() => {
    api.nats.general().then(e => {
      data.value = e
    })
    api.nats.conns().then(e => {
      conns.value = e.connections
    })
    // nats.request('$SYS.REQ.SERVER.PING').then((m) => {
    //   data.value = JSON.parse(m)
    // })
  }, 1000)
})

watch(computed(() => nats.ready.value), e => {
  if (e) {
    api.nats.general().then(e => {
      old_data = [e.out_bytes, e.in_bytes]
      data.value = e
    })
    // nats.request('$SYS.REQ.SERVER.PING').then((m) => {
    //   data.value = JSON.parse(m)
    //   let os = data.value.statsz.sent.bytes
    //   let or = data.value.statsz.received.bytes
    //   old_data = [os, or]
    // })
  }
}, { immediate: true })

onUnmounted(() => {
  if (timer.value) {
    clearInterval(timer.value)
  }
  for (let i of subs) {
    i.unsubscribe()
  }
})

</script>

<style scoped></style>

