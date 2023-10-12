 <!--
 * AppLayout.vue
 * Copyright (C) 2023 veypi <i@veypi.com>
 * 2023-10-04 10:46
 * Distributed under terms of the MIT license.
 -->
<template>
  <div class="p-4 w-full h-full">
    <div class="flex items-center mb-8">
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
import api from 'src/boot/api';
import { MenuLink, modelsApp } from 'src/models';
import { useMenuStore } from 'src/stores/menu';
import { computed, watch, ref, onMounted, provide, onBeforeUnmount } from 'vue';
import { useRoute } from 'vue-router';
import { RouteLocationNamedRaw } from 'vue-router';
let route = useRoute();
let menu = useMenuStore()

let id = computed(() => route.params.id)
let app = ref({} as modelsApp)

provide('app', app)

const sync_app = () => {
  api.app.get(id.value as string).then((e: modelsApp) => {
    app.value = e
    Links.value[1].title = e.name
    for (let i in Links.value) {
      let l: RouteLocationNamedRaw = Links.value[i].to as any
      if (l.params) {
        l.params.id = e.id
      }
    }
  }).catch(e => {
    msg.Warn('sync app data failed: ' + e)
  })
}
const Links = ref([
  {
    title: '应用中心',
    caption: '',
    icon: 'v-apps',
    to: { name: 'home' }
  },
  {
    title: '',
    caption: '',
    icon: 'v-home',
    to: { name: 'app.home', params: { id: id.value } }
  },
  {
    title: '用户管理',
    icon: 'v-team',
    to: { name: 'app.user', params: { id: id.value } }
  },
  {
    title: '权限管理',
    icon: 'v-key',
    to: { name: 'app.auth', params: { id: id.value } }
  },
  {
    title: '应用设置',
    caption: '',
    icon: 'v-setting',
    to: { name: 'app.settings', params: { id: id.value } }
  },
] as MenuLink[])
watch(id, (e) => {
  if (e) {
    sync_app()
  }
})

onMounted(() => {
  sync_app()
  menu.set(Links.value)
})
onBeforeUnmount(() => {
  menu.load_default()

})
</script>

<style scoped></style>

