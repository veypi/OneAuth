 <!--
 * stats.vue
 * Copyright (C) 2023 veypi <i@veypi.com>
 * 2023-10-20 22:56
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
    <div class="flex flex-nowrap">
      <div class="grow" style="min-height: 50vh;">
        <tschart :item="querys[idx]"></tschart>
      </div>
      <div class="flex flex-col gap-5">
        <q-chip :color="idx === i ? 'primary' : ''" class="select-none" v-for="(q, i) in querys" :key="i" @click="idx = i"
          clickable>{{
            q.name }} </q-chip>
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
import tschart from 'src/components/tschart';

const idx = ref(0)
const querys = ref<{
  name: string, query: string[] | string, label?: any,
  valueFormatter?: any
}[]>([
  {
    name: 'cpu占用',
    // query: `100 - avg (irate(node_cpu_seconds_total{mode="idle"}[3s]))
    // by(id) * 100`,
    query: `avg by
(id)(irate(node_cpu_seconds_total{mode=~"sytem|user|iowait|irq|softirq|nice|steal|guest"}[3s]))
* 100`,
    label: (d: any) => d.id as string,
    valueFormatter: (value: number) => value.toFixed(2) + "%",
  },
  {
    name: 'linux内存使用率',
    query: [
      `((node_memory_Buffers_bytes + node_memory_Cached_bytes +
      node_memory_MemFree_bytes) / node_memory_MemTotal_bytes) * 100`,
    ],
    label: (d: any) => d.id as string,
    valueFormatter: (value: number) => value.toFixed(2) + "%",
  },
  {
    name: '磁盘',
    query: `(1 - avg(node_filesystem_avail_bytes /
node_filesystem_size_bytes[3s]) by (device, id)) * 100 `,
    label: (d: any) => `${d.id}: ${d.device}` as string,
    valueFormatter: (value: number) => value.toFixed(2) + "%",
  },
  {
    name: '磁盘IOPS',
    query: [
      `sum by (id) (rate(node_disk_reads_completed_total[3s]))`,
      `sum by (id) (rate(node_disk_writes_completed_total[3s]))`,
    ],
    label: [
      (d: any) => `${d.id} 读`,
      (d: any) => `${d.id} 写`,
    ],
  },
  {
    name: '网络带宽',
    query: [
      `sum by(id)(irate(node_network_receive_bytes_total{device!~"bond.*?|lo"}[3s])) / 1048576`,
      `sum by(id)(irate(node_network_transmit_bytes_total{device!~"bond.*?|lo"}[3s])) / 1048576`
    ],
    label: [
      (d: any) => `${d.id} 下行`,
      (d: any) => `${d.id} 上行`,
    ],
    valueFormatter: (value: number) => value.toFixed(2) + "MB/s",

  },
  {
    name: '内存',
    query: [
      `(node_memory_Buffers_bytes + node_memory_Cached_bytes +
      node_memory_MemFree_bytes) / 1024 / 1024 / 1024`,
      `node_memory_MemTotal_bytes / 1024 /1024 / 1024`
    ],
    label: [(d: any) => `${d.id}使用内`, (d: any) => `${d.id}总内存`],
    valueFormatter: (value: number) => value.toFixed(2) + "GB",
  },
  // {
  //   name: 'Mac cpu频率',
  //   query: 'node_cpu_seconds_total',
  //   label: (d: any) => `cpu: ${d.cpu} mode: ${d.mode}` as string
  // },
  // {
  //   name: 'mem',
  //   query: 'node_memory_free_bytes / 1024 / 1024 / 1024'
  // },
  // {
  //   name: 'Mac swap',
  //   query: 'node_memory_swap_used_bytes / 1024 / 1024 '
  // },
])
</script>

<style scoped></style>

