 <!--
 * tschart.vue
 * Copyright (C) 2023 veypi <i@veypi.com>
 * 2023-10-20 22:50
 * Distributed under terms of the MIT license.
 -->
<template>
  <div class="w-full h-full">
    <div v-if="enable_mode" class="h-16 flex justify-start items-center">
      <div :color="enable_sync ? 'primary' : ''" @click="enable_sync = !enable_sync">{{ enable_sync ? '关闭同步'
        :
        '开启同步' }}</div>
      <div :color="mode === k ? 'primary' : ''" v-for="(v, k) in mode_label" :key="k" @click="change_mode(k)">{{
        v }}</div>
    </div>
    <div class="v-chart w-full" :style="{
      height:
        enable_mode ? 'calc(100% - 4rem)' : '100%'
    }" ref="chartdom"></div>
  </div>
</template>

<script lang="ts" setup>
import * as echart from 'echarts'
import use_params from './params'
import { onMounted, onUnmounted, computed, ref, watch, markRaw } from 'vue';

let { params, mode, change_mode, mode_label } = use_params()
interface Item {
  name: string
  query: string | string[]
  valueFormatter?: (s: number) => string
  label?: string | string[] | ((s: any) => string)
}
let props = withDefaults(defineProps<{
  item: Item,
  sync?: boolean,
  enable_zoom?: boolean,
  time_mode?: number,
  enable_mode?: boolean,
}>(),
  {
  }
)
let count = 0
let timer = ref<any[]>([])
let enable_sync = ref(false)
let chartdom = ref()
let options = ref<{ [key: string]: any }>({})
let chart: echart.ECharts = {} as any
let tooltip = {
  trigger: 'axis',
  axisPointer: {
    type: 'cross',
    label: {},
  },
  valueFormatter: (value: number) => value.toFixed(2),
  className: 'v-echarts-tooltip',
}

const init_chart = () => {
  count++
  if (chart.clear) {
    chart.clear()
  }
  timer.value.forEach(e => {
    clearInterval(e)
  })
  options.value = {
    title: { text: props.item.name, x: 'center' },
    animationThreshold: 200,
    tooltip: Object.assign({}, tooltip),
    axisPointer: {
      link: { xAxisIndex: 'all' },
      label: {
        backgroundColor: '#777'
      }
    },
    xAxis: {
      type: 'time',
    },
    yAxis: {},
    series: []
  }
  if (props.enable_zoom) {
    options.value.dataZoom = [
      {
        type: 'slider',
        xAxisIndex: [0],
        filterMode: 'filter'
      },
    ]
  }
  if (props.item.valueFormatter) {
    options.value.tooltip.valueFormatter = props.item.valueFormatter
  }
  let tmp = {
    start: params.value.start.toISOString(),
    step: params.value.step,
  }
  let querys: string[] = Array.isArray(props.item.query) ? props.item.query :
    [props.item.query]
  let labels = Array.isArray(props.item.label) ? props.item.label :
    [props.item.label]
  for (let q = 0; q < querys.length; q++) {
    let query = querys[q]
    api.tsdb.range(query, tmp).then(e => {
      if (e.status == 'success') {
        let data = e.data.result as any[]
        if (data.length == 0) {
          console.warn('not get data')
          return
        }
        let idx = options.value.series.length || 0
        data.forEach(d => {
          let name = props.item.name
          let label = labels[q]
          if (typeof label === 'string') {
            name = label
          } else if (typeof label === 'function') {
            name = label(d.metric)
          }
          options.value.series.push({
            name: name,
            data: d.values.map((e: any) =>
              [e[0] * 1000, Number(e[1])]),
            metric: d.metric,
            metric_str: JSON.stringify(d.metric),
            origin: query,
            symbol: 'none',
            smooth: true,
            type: 'line',
          })
        })
        chart.setOption(options.value)
        let t = setInterval(() => {
          sync_chart(idx, query, count)
        }, 1000)
        timer.value.push(t)
      }
    })
  }
  // let query = props.query
}
const sync_chart = (idx: number, query: string, c: number) => {
  if (!enable_sync.value) {
    return
  }
  api.tsdb.query(query).then(e => {
    if (e.status == 'success') {
      let data = e.data.result as any[]
      if (data.length == 0) {
        console.warn('not get data')
        timer.value.forEach(e => {
          clearInterval(e)
        })
        return
      }
      if (count === c) {
        data.forEach((d, i) => {
          let sidx = idx + i
          if (d.metric) {
            let ti = options.value.series.findIndex((s: any) =>
              query === s.origin && JSON.stringify(d.metric) === s.metric_str)
            if (ti >= 0) {
              sidx = ti
            }
          }
          options.value.series[sidx].data.push([d.value[0] * 1000,
          Number(d.value[1])])
        })
        chart.setOption(options.value)
      }
    }
  })
}



watch(computed(() => props.item), q => {
  if (q) {
    init_chart()
  }
})

watch(mode, q => {
  init_chart()
})

onMounted(() => {
  enable_sync.value = props.sync
  if (props.time_mode) {
    change_mode(props.time_mode)
  }
  chart = markRaw(echart.init(chartdom.value, null, { renderer: 'svg' }))
  init_chart()
})
onUnmounted(() => {
  timer.value.forEach(e => {
    clearInterval(e)
  })
})
</script>

<style>
.v-chart {
  min-width: 20rem;
  min-height: 15rem;
}

.v-echarts-tooltip {
  /* height: 5rem; */
  /* width: 10rem; */
}
</style>

