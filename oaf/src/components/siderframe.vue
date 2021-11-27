<template>
  <n-layout has-sider>
    <n-layout-sider
      collapse-mode="transform"
      :collapsed-width="0"
      :width="150"
      show-trigger="bar"
      content-style="padding: 4px;"
      bordered
      :default-collapsed="true"
      :native-scrollbar="false"
      :style="{'height': $store.state.height}"
    >
      <slot name="sider"></slot>
    </n-layout-sider>
    <n-layout :style="{'height': $store.state.height}" :native-scrollbar="false">
      <div class="mx-5" :style="{'min-height': $store.state.height}">
        <n-page-header @back="back">
          <template #title>
          <slot name="title"></slot>
          </template>
          <template #subtitle>
          <slot name="subtitle"></slot>
          </template>
          <template #header>
          <n-breadcrumb>
            <n-breadcrumb-item @click="go(item)"
                               :key="key"
                               v-for="(item, key) in breads">
              <one-icon class="inline-block" v-if="item.Type==='icon'">{{ item.Name }}</one-icon>
              <span v-else>{{ item.Name }}</span>
            </n-breadcrumb-item>
          </n-breadcrumb>
          </template>
          <template #avatar>
          <slot name="avatar"></slot>
          </template>
          <template #extra>
          <slot name="extra"></slot>
          </template>
          <template #footer>
          <slot name="footer"></slot>
          </template>
        </n-page-header>
        <slot></slot>
      </div>
      <n-back-top>
      </n-back-top>
    </n-layout>
  </n-layout>
</template>

<script lang="ts" setup>
import {modelsBread} from '@/models'
import {useRoute, useRouter} from 'vue-router'
import {computed} from 'vue'
import {useStore} from '@/store'

let router = useRouter()
let route = useRoute()
let store = useStore()

let breads = computed(() => {
  let list: modelsBread[] = []
  for (let b of store.state.breads) {
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
    router.push({name: b.RName, query: b.RQuery, params: b.RParams})
  }
}
</script>

<style scoped>

</style>
