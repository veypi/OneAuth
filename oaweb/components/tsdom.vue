 <!--
 * tsdom.vue
 * Copyright (C) 2024 veypi <i@veypi.com>
 * 2024-07-02 16:54
 * Distributed under terms of the MIT license.
 -->
<template>
  <div>
    <template v-if="data.status === 'success' &&
      data.data?.result?.length">
      <slot v-if="data.data.result[0]?.value?.length" name="single" :data="data.data.result[0].value"></slot>
      <slot name="list" :stats="data.stats" :data="data.data.result"></slot>
    </template>
    <slot v-if="data.status === 'success' && data.data?.result?.length" name="success" :stats="data.stats"
      :data="data.data.result"></slot>
    <slot v-else name="fail">未获取数据</slot>
  </div>
</template>

<script lang="ts" setup>
import axios from 'axios';
let url = window.location.protocol + '//' + window.location.host + '/api/ts/'

let props = withDefaults(defineProps<{
  is_range?: boolean,
  query: string,
  start?: string,
  end?: string,
  step?: string,
  delta?: string
  sync_delta?: number
}>(),
  {
    is_range: false,
  }
)
const parse_delta = (d: string) => {
  let re = /^(\d*)([smhMy])$/.exec(d)
  if (re?.length) {
    let n = parseInt(re[1])
    let u = re[2]
    let m = { s: 1, m: 60, h: 3600, d: 86400, M: 2592000, y: 31536000 }
    // @ts-ignore
    return n * m[u] * 1000
  }
  return undefined
}
let qdata = computed(() => {
  let res = { query: props.query } as any
  if (props.delta) {
    let d = parse_delta(props.delta)
    if (d) {
      res.start = new Date().getTime() - d
      res.step = Math.ceil(d / 1000000) + 's'
    }
  }
  if (props.start) {
    res.start = props.start
  }
  if (props.end) {
    res.end = props.end
  }
  if (props.step) {
    res.step = props.step
  }
  return res
})
let data = ref<any>({})


const sync_data = () => {
  let u = url + 'query'
  if (props.is_range) {
    u = url + 'query_range'
  }
  axios.get(u, { params: qdata.value }).then(e => {
    data.value = e.data
  })
  if (props.sync_delta && props.sync_delta > 0) {
    setTimeout(sync_data, props.sync_delta * 1000)
  }
}

watch(() => props, sync_data, { immediate: true })
</script>

<style scoped></style>

