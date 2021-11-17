<template>
  <siderframe>
    <template v-slot:sider>
    <div class="grid grid-cols-1">
      <div class="cursor-pointer" v-for="(item, key) in navRouter" :key="key">
        <div class="pl-4 text-lg my-4" :style="{color: isEqualRoute(item) ? '#88baea': '#888'}"
             @click="router.push(Object.assign({}, item, {params: route.params, query: route.query}))">
          {{ item.meta.title }}
        </div>
        <template v-if="isEqualRoute(item) && nav && nav.length>0">
        <transition-group appear enter-active-class="animate__zoomIn" leave-active-class="animate__zoomOut">
          <div @click="goAnchor(tt)" class="pl-8 rounded my-0.5 animate__animated animate__400ms"
               v-for="(tt, kk) in nav" :key="kk">
            {{ tt.innerText }}
          </div>
        </transition-group>
        </template>
      </div>
    </div>
    </template>
    <router-view v-slot="{ Component }">
      <transition mode="out-in" enter-active-class="animate__fadeInLeft"
                  leave-active-class="animate__fadeOutRight">
        <component class="animate__animated animate__400ms" :is="Component" ref="main"
        ></component>
      </transition>
    </router-view>
  </siderframe>
</template>

<script lang="ts" setup>
import {elementScrollIntoView} from "seamless-scroll-polyfill";
import {useRoute, useRouter} from "vue-router";
import {computed, onMounted, ref, provide} from "vue";
import api from "@/api";
import Siderframe from "@/components/siderframe.vue";
import {useMessage} from "naive-ui";
import {useStore} from "@/store";


let store = useStore()
let mgs = useMessage()
let route = useRoute()
let router = useRouter()
let uuid = computed(() => route.params.uuid)
let app = ref({})
provide('app', app)
provide('uuid', uuid)
let main = ref(null)
// @ts-ignore
let nav = computed(() => main.value ? main.value.nav : [])
let navRouter = ref(buildRouter())

function buildRouter() {
  let navs = []
  for (let n of router.getRoutes()) {
    if (n.name && (n.name as string).startsWith('app')) {
      if (n.meta.checkAuth) {
        if (n.meta.checkAuth(store.state.user.auth)) {
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
    router.push({name: '404', params: {path: route.path}})
    return
  }
  api.app.get(uuid.value as string).Start(e => {
    app.value = e
  }, e => {
    mgs.error('获取应用数据失败: ' + (e.err || e))
    router.push({name: '404', params: {path: route.path}})
  })
})

function goAnchor(element: any) {
  // safari not support
  // element.scrollIntoView({
  //   behavior: "smooth"
  // })
  elementScrollIntoView(element, {behavior: "smooth"});
}


</script>

<style scoped>

</style>
