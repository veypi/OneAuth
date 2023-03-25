<template>
  <siderframe>
    <template v-slot:avatar>
      <n-avatar
        style="--color: none"
        @click="util.goto(app.host)"
        :src="app.icon"
        round
        size="large"
      ></n-avatar>
    </template>
    <template #title>{{ app.name }}</template>
    <template #subtitle>{{ app.des }}</template>
    <template v-slot:sider>
      <div class="grid grid-cols-1">
        <div class="cursor-pointer" v-for="(item, key) in navRouter" :key="key">
          <div
            class="pl-4 text-lg my-4"
            :style="{ color: isEqualRoute(item) ? '#88baea' : '#888' }"
            @click="
              router.push(Object.assign({}, item, { params: route.params, query: route.query }))
            "
          >
            {{ item.meta.title }}
          </div>
          <template v-if="isEqualRoute(item) && nav && nav.length > 0">
            <transition-group
              appear
              enter-active-class="animate__zoomIn"
              leave-active-class="animate__zoomOut"
            >
              <div
                @click="goAnchor(tt)"
                class="pl-8 rounded my-0.5 animate__animated animate__400ms"
                v-for="(tt, kk) in nav"
                :key="kk"
              >
                {{ tt.innerText }}
              </div>
            </transition-group>
          </template>
        </div>
      </div>
    </template>
    <router-view v-slot="{ Component }">
      <transition
        mode="out-in"
        enter-active-class="animate__fadeInLeft"
        leave-active-class="animate__fadeOutRight"
      >
        <component class="animate__animated animate__400ms" :is="Component" ref="main"></component>
      </transition>
    </router-view>
  </siderframe>
</template>

<script lang="ts" setup>
import { elementScrollIntoView } from 'seamless-scroll-polyfill'
import { useRoute, useRouter, RouteRecord } from 'vue-router'
import { computed, onMounted, ref, provide, onBeforeUnmount } from 'vue'
import api from '@/api'
import Siderframe from '@/components/siderframe.vue'
import { modelsApp, modelsBread } from '@/models'
import util from '@/libs/util'
import { useUserStore, useAppStore } from '@/store'
import msg from '@/msg'

let my = useUserStore()
let local = useAppStore()
let route = useRoute()
let router = useRouter()
let uuid = computed(() => route.params.uuid)
let app = ref<modelsApp>({} as modelsApp)
provide('app', app)
provide('uuid', uuid)
let main = ref(null)
// @ts-ignore
let nav = computed(() => (main.value ? main.value.nav : []))
let navRouter = ref(buildRouter())

function buildRouter(): RouteRecord[] {
  let navs: RouteRecord[] = []
  for (let n of router.getRoutes()) {
    if (n.name && (n.name as string).startsWith('app')) {
      if (n.meta.checkAuth) {
        if (n.meta.checkAuth(my.auth, route)) {
          navs.push(n)
        }
      } else {
        navs.push(n)
      }
    }
  }
  return navs
}

function isEqualRoute(r: any) {
  return r.name === route.name
}

onMounted(() => {
  if (uuid.value === '') {
    router.push({ name: '404', params: { path: route.path } })
    return
  }
  api.app.get(uuid.value as string).Start(
    (e: modelsApp) => {
      app.value = e
      local.title = e.name
      local.setBreads({
        Index: 1,
        Name: e.name,
        RName: 'app.main',
        RParams: { uuid: e.id },
      } as modelsBread)
    },
    (e) => {
      msg.Warn('获取应用数据失败: ' + (e.err || e))
      router.push({ name: '404', params: { path: route.path } })
    },
  )
})

onBeforeUnmount(() => {
  local.title = ''
})

function goAnchor(element: any) {
  // safari not support
  // element.scrollIntoView({
  //   behavior: "smooth"
  // })
  elementScrollIntoView(element, { behavior: 'smooth' })
}
</script>

<style scoped></style>
