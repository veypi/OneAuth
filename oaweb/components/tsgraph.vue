 <!--
 * tsgraph.vue
 * Copyright (C) 2024 veypi <i@veypi.com>
 * 2024-07-02 18:11
 * Distributed under terms of the MIT license.
 -->
<template>
  <div class="v-graph" ref="chartdom"></div>
</template>

<script lang="ts" setup>
import * as echart from 'echarts'

let props = withDefaults(defineProps<{
  title?: string,
  options?: any,
  data: any[]
  dataFormatter?: (data: any[]) => any
}>(),
  {}
)
let options = computed(() => {
  let opt: echarts.EChartsOption = {
    // @ts-ignore
    title: { text: props.title, x: 'center' },
    animationThreshold: 200,
    tooltip: {
      trigger: 'axis',
      axisPointer: {
        type: 'cross',
        label: {},
      },
      className: 'v-echarts-tooltip',
      valueFormatter: (value: any) => value.toFixed(2),
    },
    axisPointer: {
      link: [{ xAxisIndex: 'all' }],
      label: {
        backgroundColor: '#777'
      }
    },
    xAxis: {
      type: 'time',
    },
    dataZoom: [
      {
        type: 'inside',
        xAxisIndex: [0],
        // yAxisIndex: [0],
      },
    ],
    yAxis: {},
    series: []
  }
  if (props.options) {
    opt = Object.assign(opt, props.options)
  }
  return opt
})

let chartdom = ref()
let chart: echart.ECharts = {} as any
watch(() => props.data, (d) => {
  let opts = options.value
  opts.series = props.dataFormatter ? props.dataFormatter(d) : d
  chart.setOption(options.value)
}, { deep: true })
onMounted(() => {
  chart = markRaw(echart.init(chartdom.value, null, { renderer: 'svg' }))
  let opts = options.value
  opts.series = props.dataFormatter ? props.dataFormatter(props.data)
    : props.data
  chart.setOption(options.value)
})
</script>

<style scoped>
.v-graph {
  min-width: 20rem;
  min-height: 15rem;
  width: 100%;
  height: 100%;
}
</style>

