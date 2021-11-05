<template>
  <n-config-provider :theme-overrides="Theme.overrides" :locale="zhCN" :date-locale="dateZhCN"
                     :theme="Theme">
    <n-layout class="font-sans select-none">
      <n-layout>
        <n-layout-header class="pr-5" bordered style="height: 64px;line-height: 64px;">
          <div class="inline-block float-left h-full">
            <one-icon color="#000" class="inline-block" @click="$router.push('/')" style="font-size: 48px;margin:8px;color:aqua">
              glassdoor
            </one-icon>
          </div>
          <div class="inline-block float-left h-full" style="margin-left: 10px">
            <n-h6 prefix="bar" align-text><n-text type="primary">统一认证系统</n-text></n-h6>
          </div>
          <div v-if="store.state.user.ready" class="inline-block h-full float-right flex justify-center items-center">
            <avatar></avatar>
          </div>
          <div class="inline-block float-right h-full px-3">
            <fullscreen v-model="isFullScreen" class="header-icon">fullscreen</fullscreen>
            <div class="header-icon">
              <one-icon @click="ChangeTheme">
                {{ IsDark ? 'Daytimemode' : 'nightmode-fill' }}
              </one-icon>
            </div>
          </div>
        </n-layout-header>
        <n-layout has-sider style="height: calc(100vh - 88px)">
          <n-layout-sider
            collapse-mode="transform"
            :collapsed-width="0"
            :width="120"
            show-trigger="bar"
            content-style="padding: 24px;"
            bordered
            default-collapsed
            :native-scrollbar="false"
          >
            -
          </n-layout-sider>
          <n-layout class="main" :native-scrollbar="false">
            <n-message-provider>
                <router-view v-slot="{ Component }">
                  <transition mode="out-in" enter-active-class="animate__fadeInLeft" leave-active-class="animate__fadeOutRight">
                    <component class="animate__animated animate__400ms" :is="Component" style="margin: 10px; min-height: calc(100vh - 108px)"
                                 ></component>
                  </transition>
                </router-view>
            </n-message-provider>
          </n-layout>
        </n-layout>
      </n-layout>
      <n-layout-footer bordered style="height: 24px;line-height: 24px"
                       class="flex justify-around px-3 text-gray-500 text-xs">
        <span class="hover:text-black cursor-pointer" @click="$router.push({name: 'about'})">关于OA</span>
        <span class="hover:text-black cursor-pointer">使用须知</span>
        <span class="hover:text-black cursor-pointer" @click="goto('https://veypi.com')">
      ©2021 veypi
      </span>
      </n-layout-footer>
    </n-layout>
  </n-config-provider>
</template>
<script setup lang="ts">
// This starter template is using Vue 3 <script setup> SFCs
import {onBeforeMount, ref} from 'vue'
import util from './libs/util'
import {useStore} from "./store";
import {Theme, IsDark, ChangeTheme} from "./theme";
import {zhCN, dateZhCN} from 'naive-ui'
import avatar from "./components/avatar";
import fullscreen from './components/fullscreen'
import Fullscreen from "./components/fullscreen/fullscreen.vue";

let isFullScreen = ref(false)
let store = useStore()

onBeforeMount(() => {
  util.title("统一认证")
  store.dispatch('fetchSelf')
  store.dispatch('user/fetchUserData')
})

let goto = (url: any) => {
  window.open(url, "_blank")
}

let collapsed = ref(true)

</script>


<style lang="less">
.animate__400ms {
  --animate-duration: 400ms;
}

html,
body {
  width: 100%;
  height: 100%;
  margin: 0;
  padding: 0;
}

/* 周围滑动留白 */
html {
  overflow: hidden;
  height: 100%;
}

body {
  overflow: auto;
  height: 100%;
}

.header-icon {
  display: inline-block;
  font-size: 24px;
  margin: 20px 10px 20px 10px;
}

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
  width: 100%;
  height: 100%;
}

.main {
}

::-webkit-scrollbar {
  display: none; /* Chrome Safari */
}
</style>
