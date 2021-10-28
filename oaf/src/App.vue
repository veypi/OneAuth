<script setup lang="ts">
// This starter template is using Vue 3 <script setup> SFCs
import { onBeforeMount, ref} from 'vue'
import util from './libs/util'
import {store} from "./store";
import {useOsTheme} from 'naive-ui'

const osThemeRef = useOsTheme()

onBeforeMount(() => {
  util.title("统一认证")
  store.dispatch('fetchSelf')
  store.commit('setTheme', osThemeRef.value)
})
</script>

<template>
  <n-config-provider :theme="store.getters.GetTheme" class="h-full w-full">
    <n-layout class="h-full w-full font-sans select-none">
      <n-layout has-sider style="height: calc(100% - 24px)">
        <n-layout-sider
          class="h-full"
          collapse-mode="transform"
          :collapsed-width="0"
          :width="120"
          show-trigger="bar"
          content-style="padding: 24px;"
          bordered
          default-collapsed
        >
          -
        </n-layout-sider>
        <n-layout>
          <n-layout-header bordered style="height: 64px;line-height: 64px;">
            {{ osThemeRef }}
            <one-icon @click="store.dispatch('changeTheme')" class="float-right" style="font-size: 36px; margin: 14px">
              {{ store.getters.IsDark ? 'Daytimemode' : 'nightmode-fill' }}
            </one-icon>
          </n-layout-header>
          <n-layout style="height: calc(100% - 64px)">
            <router-view class="h-full w-full"></router-view>
          </n-layout>
        </n-layout>
      </n-layout>
      <n-layout-footer style="height: 24px;line-height: 24px" class="flex justify-around px-3 text-gray-500 text-xs">
        <span class="hover:text-black cursor-pointer">关于OA</span>
        <span class="hover:text-black cursor-pointer">使用须知</span>
        <span class="hover:text-black cursor-pointer">
      ©2021 veypi
      </span>
      </n-layout-footer>
    </n-layout>
  </n-config-provider>
</template>

<style>
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

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
  width: 100%;
  height: 100%;
}

::-webkit-scrollbar {
  display: none; /* Chrome Safari */
}
</style>
