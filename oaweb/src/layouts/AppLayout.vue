 <!--
 * AppLayout.vue
 * Copyright (C) 2023 veypi <i@veypi.com>
 * 2023-10-04 10:46
 * Distributed under terms of the MIT license.
 -->
<template>
  <div class="p-4 w-full h-full">
    <div class="flex items-center mb-8 cursor-pointer" @click="util.goto(app.host)">
      <q-avatar class="mx-2" round size="4rem">
        <img :src="app.icon">
      </q-avatar>
      <h1 class="text-4xl">{{ app.name }}</h1>
    </div>
    <router-view v-slot="{ Component }">
      <transition mode="out-in" enter-active-class="animate__fadeIn" leave-active-class="animate__fadeOut">
        <component class="animate__animated animate__400ms p-10" style="min-height:calc(100% - 96px)" :is="Component">
        </component>
      </transition>
    </router-view>
  </div>
</template>

<script lang="ts" setup>
import msg from '@veypi/msg';
import util from 'src/libs/util';
import api from 'src/boot/api';
import { modelsApp } from 'src/models';
import { computed, watch, ref, onMounted, provide, onBeforeUnmount } from 'vue';
import { useRoute } from 'vue-router';
import { RouteLocationNamedRaw } from 'vue-router';
import menus from './menus'
let route = useRoute();

let id = computed(() => route.params.id)
let app = ref({} as modelsApp)

provide('app', app)

const sync_app = () => {
  let tid = id.value as string
  api.app.get(id.value as string).then((e: modelsApp) => {
    app.value = e
    let links = menus.appLinks.value.concat([])
    links[0].title = e.name
    if (menus.uniqueLinks[tid]?.length) {
      for (let r of menus.uniqueLinks[tid]) {
        links.splice(1, 0, r)
      }
    }
    for (let i in links) {
      let l: RouteLocationNamedRaw = links[i].to as any
      if (l.params) {
        l.params.id = e.id
      } else {
        l.params = { id: e.id }
      }
    }
    menus.items.value = links
  }).catch(e => {
    msg.Warn('sync app data failed: ' + e)
  })
}
watch(id, (e) => {
  if (e) {
    sync_app()
  }
}, { immediate: true })

onMounted(() => {
})
onBeforeUnmount(() => {
  menus.load_default()

})
</script>

<style scoped></style>

