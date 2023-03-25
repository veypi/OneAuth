<template>
  <Menu>
    <template #menu>
      <slot name="sider"></slot>
    </template>
    <div class="mx-5">
      <slot name="title"></slot>
      <slot name="subtitle"></slot>
      <div @click="go(item)" :key="key" v-for="(item, key) in breads">
        <one-icon class="inline-block" v-if="item.Type === 'icon'">
          {{ item.Name }}
        </one-icon>
        <span v-else>{{ item.Name }}</span>
      </div>
      <slot name="avatar"></slot>
      <slot name="extra"></slot>
      <slot name="footer"></slot>
      <slot></slot>
    </div>
  </Menu>
</template>

<script lang="ts" setup>
import Menu from './menu.vue'
import { modelsBread } from '@/models'
import { useRoute, useRouter } from 'vue-router'
import { computed } from 'vue'
import { useAppStore } from '@/store/app'

let router = useRouter()
let route = useRoute()
let store = useAppStore()

let breads = computed(() => {
  let list: modelsBread[] = []
  for (let b of store.breads) {
    list.push(b)
    if (b.RName === route.name) {
      break
    }
  }
  return list
})

function go(b: modelsBread) {
  router.push({
    name: b.RName,
    params: b.RParams,
    query: b.RQuery,
  })
}

function back() {
  if (breads.value.length > 1) {
    let b = breads.value[breads.value.length - 2]
    router.push({ name: b.RName, query: b.RQuery, params: b.RParams })
  }
}
</script>

<style scoped></style>
