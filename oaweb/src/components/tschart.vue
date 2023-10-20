 <!--
 * tschart.vue
 * Copyright (C) 2023 veypi <i@veypi.com>
 * 2023-10-20 22:50
 * Distributed under terms of the MIT license.
 -->
<template>
  <div class="w-full h-full">
    <div class="v-chart w-full h-full" ref="chartdom"></div>
  </div>
</template>

<script lang="ts" setup>
import * as echart from 'echarts'
import api from 'src/boot/api';
import { onMounted, onUnmounted, computed, ref, watch, markRaw } from 'vue';

interface Item {
  name: string
  query: string | string[]
  ext?: string
  valueFormatter?: (s: number) => string
  label?: string | string[] | ((s: any) => string)
}
let props = withDefaults(defineProps<{
  item: Item,
  // start?: string,
  // end?: string,
  // step?: string
}>(),
  {
  }
)
let getparams = ref<any>({
  start: () => {
    let d = new Date()
    d.setMinutes(d.getMinutes() - 3)
    return d.toISOString()
  }, end: undefined, step: '2s'
})
let count = 0
let timer = ref<any[]>([])
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
  options.value = {
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
  if (props.item.valueFormatter) {
    options.value.tooltip.valueFormatter = props.item.valueFormatter
  }
  let tmp = {} as any
  if (getparams.value.start) {
    tmp.start = getparams.value.start()
  }
  if (getparams.value.step) {
    tmp.step = getparams.value.step
  }
  let query: string[] = Array.isArray(props.item.query) ? props.item.query :
    [props.item.query]
  for (let q = 0; q < query.length; q++) {
    api.tsdb.range(query[q], tmp).then(e => {
      if (e.status == 'success') {
        let data = e.data.result as any[]
        if (data.length == 0) {
          console.warn('not get data')
          return
        }
        let idx = options.value.series.length || 0
        data.forEach(d => {
          let name = props.item.name
          if (typeof props.item.label === 'string') {
            name = props.item.label
          } else if (typeof props.item.label === 'function') {
            name = props.item.label(d.metric)
          } else if (Array.isArray(props.item.label)) {
            name = props.item.label[q]
          }
          options.value.series.push({
            name: name,
            data: d.values.map((e: any) =>
              [e[0] * 1000, Number(e[1])]),
            metric: d.metric,
            origin: query[q],
            symbol: 'none',
            smooth: true,
            type: 'line',
          })
        })
        chart.setOption(options.value)
        let t = setInterval(() => {
          sync_chart(idx, query[q], count)
        }, 1000)
        timer.value.push(t)
      }
    })
  }
  // let query = props.query
}
const sync_chart = (idx: number, query: string, c: number) => {
  api.tsdb.query(query).then(e => {
    if (e.status == 'success') {
      let data = e.data.result as any[]
      if (data.length == 0) {
        console.warn('not get data')
        return
      }
      if (count === c) {
        data.forEach((d, i) => {
          options.value.series[idx + i].data.push([d.value[0] * 1000,
          Number(d.value[1])])
        })
        chart.setOption(options.value)
      }
    }
  })
}

watch(computed(() => props.item), q => {
  timer.value.forEach(e => {
    clearInterval(e)
  })
  if (q) {
    init_chart()
  }
}, { immediate: true })

onMounted(() => {
  chart = markRaw(echart.init(chartdom.value))
})
onUnmounted(() => {
  timer.value.forEach(e => {
    clearInterval(e)
  })
})
</script>

<style>
.v-chart {}

.v-echarts-tooltip {
  /* height: 5rem; */
  /* width: 10rem; */
}
</style>

