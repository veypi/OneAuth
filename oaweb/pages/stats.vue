
 <!--
 * stats.vue
 * Copyright (C) 2023 veypi <i@veypi.com>
 * 2023-10-20 22:56
 * Distributed under terms of the MIT license.
 -->
<template>
  <div>
    <div class="page-h1">
      服务
    </div>
    <div class="w-40 text-center py-4 start_card">
      <div class="text-3xl"> 已运行 </div>
      <div class="text-2xl mt-2">
        {{ start_time }}
      </div>
    </div>
    <div class="flex flex-nowrap" style="">
      <div class="w-1/2">
        <Tschart :item="querys[0]" :time_mode="1"></Tschart>
      </div>
      <div class="w-1/2">
        <Tschart :item="querys[1]" :time_mode="1"></Tschart>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref, onUnmounted, onMounted } from 'vue';
// import tschart from 'src/components/tschart';

const start_time = ref('')
const timer = ref()
const querys = ref<{
  name: string, query: string[] | string, label?: any,
  valueFormatter?: any
}[]>([
  {
    name: 'cpu',
    query: `srv_cpu{i='oa'}`,
    label: 'cpu',
    valueFormatter: (value: number) => value.toFixed(2) + "%",
  },
  {
    name: '内存',
    query: `srv_mem{i='oa'} / 1048576`,
    label: '内存',
    valueFormatter: (value: number) => value.toFixed(2) + "MB",
  },
])

onMounted(() => {
  api.tsdb.query('srv_start{i="oa"}').then(e => {
    if (e.data.result.length) {
      let s = Number(e.data.result[0].value[1])
      if (s < 60) {
        start_time.value = s + ' 秒'
      } else if (s < 3600) {
        start_time.value = (s / 60).toFixed(1) + ' 分钟'
      } else if (s < 3600 * 24) {
        start_time.value = (s / 60 / 60).toFixed(1) + ' 小时'
      } else {
        start_time.value = (s / 60 / 60 / 24).toFixed(1) + ' 天'
      }
    }
  })
})
onUnmounted(() => {
  if (timer.value) {
    clearInterval(timer.value)
  }
})

</script>

<style lang="scss" scoped>
.start_card {
  border: 1px solid var(--color-primary);

  :first-child {
    color: var(--color-primary)
  }

  :nth-child(2) {}
}
</style>

