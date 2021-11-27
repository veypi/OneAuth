<template>
  <div>
    <n-select
      filterable
      placeholder="搜索用户"
      :options="options"
      :loading="loading"
      clearable
      remote
      @search="handleSearch"
      @update-value="select"
    />
  </div>
</template>

<script lang="ts" setup>
import {ref} from 'vue'
import api from '@/api'
import {modelsUser} from '@/models'

let emits = defineEmits<{
  (e: 'selected', v:modelsUser): void
}>()
let options = ref([])
let loading = ref(false)

function select(v, o) {
  emits('selected', o.user)
}

function handleSearch(query: string) {
  if (!query.length) {
    options.value = []
    return
  }
  loading.value = true
  api.user.search(query).Start((e: modelsUser[]) => {
    let l = []
    for (let u of e) {
      l.push({
        label: u.Username,
        value: u.ID,
        user: u,
      })
    }
    options.value = l
    loading.value = false
  })
}
</script>

<style scoped>

</style>
