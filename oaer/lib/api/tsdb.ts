/*
 * tsdb.ts
 * Copyright (C) 2023 veypi <i@veypi.com>
 * 2023-10-20 23:21
 * Distributed under terms of the MIT license.
 */


import cfg from '../cfg'
import ajax from './axios'

export default {
  local: () => cfg.BaseUrl() + '/ts/',
  range(query: string, props?: { start?: string, end?: string, step?: string, query?: string }) {
    if (query !== undefined) {
      // @ts-ignore
      props.query = query
    } else {
      props = { query: query }
    }
    return ajax.get(this.local() + 'query_range', props)
  },
  query(query: string) {
    return ajax.get(this.local() + 'query', { query: query })
  }
}
