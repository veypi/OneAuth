 <!--
 * stats.vue
 * Copyright (C) 2023 veypi <i@veypi.com>
 * 2023-10-20 22:56
 * Distributed under terms of the MIT license.
 -->
<template>
  <div class="flex flex-nowrap">
    <div class="grow" style="height: calc(100%);">
      <tschart :item="querys[idx]"></tschart>
    </div>
    <div class="flex flex-col gap-5">
      <q-chip :color="idx === i ? 'primary' : ''" class="select-none" v-for="(q, i) in querys" :key="i" @click="idx = i"
        clickable>{{
          q.name }} </q-chip>
    </div>
  </div>
</template>

<script lang="ts" setup>
import tschart from 'src/components/tschart.vue';
import { ref } from 'vue';

const idx = ref(0)
const querys = ref([
  {
    name: 'cpu占用',
    query: `100 - avg (irate(node_cpu_seconds_total{mode="idle"}[3s]))
    by(id) * 100`,
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
    name: 'linux 内存',
    query: [
      `(node_memory_Buffers_bytes + node_memory_Cached_bytes +
      node_memory_MemFree_bytes) / 1024 / 1024 / 1024`,
      `node_memory_MemTotal_bytes / 1024 /1024 / 1024`
    ],
    label: ['使用内存', '总内存'],
    valueFormatter: (value: number) => value.toFixed(2) + "GB",
  },
  {
    name: 'Mac cpu频率',
    query: 'node_cpu_seconds_total',
    label: (d: any) => `cpu: ${d.cpu} mode: ${d.mode}` as string
  },
  {
    name: 'mem',
    query: 'node_memory_free_bytes / 1024 / 1024 / 1024'
  },
  {
    name: 'Mac swap',
    query: 'node_memory_swap_used_bytes / 1024 / 1024 '
  },
])
</script>

<style scoped></style>

