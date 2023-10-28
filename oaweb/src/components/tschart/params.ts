/*
 * params.ts
 * Copyright (C) 2023 veypi <i@veypi.com>
 * 2023-10-22 05:13
 * Distributed under terms of the MIT license.
 */
import { ref } from 'vue'


let mode = ref(0)
let mode_label = ['近5分钟', '近1小时', '近24小时', '近7天', '近30天']
let change_mode = (m: number) => {
  mode.value = m
  let now = new Date()
  switch (m) {
    case 0: {
      now.setMinutes(now.getMinutes() - 5)
      params.value.start = now
      params.value.step = "2s"
      break
    }
    case 1: {
      now.setHours(now.getHours() - 1)
      params.value.start = now
      params.value.step = "10s"
      break
    }
    case 2: {
      now.setHours(now.getHours() - 24)
      params.value.start = now
      params.value.step = "20s"
      break
    }
    case 3: {
      now.setHours(now.getHours() - 24 * 7)
      params.value.start = now
      params.value.step = "30s"
      break
    }
    case 4: {
      now.setHours(now.getHours() - 24 * 29)
      params.value.start = now
      params.value.step = "1h"
      break
    }
    case 5: {
      break
    }
  }
}
let params = ref<{ start: Date, end: Date, step: string }>({
  start: new Date(),
  end: new Date(),
  step: '2s'
})

change_mode(0)

const set_delta = (start?: Date, end?: Date) => {
  if (start) {
    params.value.start = start
  }
  if (end) {
    params.value.end = end
  }
  let delta = params.value.end.getTime() -
    params.value.start.getTime()
  console.log(delta)
}

export { params, change_mode, mode, mode_label }
