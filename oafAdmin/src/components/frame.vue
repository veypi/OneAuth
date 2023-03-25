<!--
 * frame.vue
 * Copyright (C) 2022 veypi <i@veypi.com>
 * 2022-12-20 00:10
 * Distributed under terms of the Apache license.
 -->
<template>
  <div class="font-sans select-none" style="height: 100vh; width: 100vw">
    <transition enter-active-class="animate__slideInDown" leave-active-class="animate__slideOutUp">
      <one-icon
        class="header-down animate__animated"
        @click="app.hideHeader = !app.hideHeader"
        v-if="app.hideHeader"
      >
        down
      </one-icon>
    </transition>
    <div style="height: calc(100vh - 24px)">
      <transition
        enter-active-class="animate__slideInDown"
        leave-active-class="animate__slideOutUp"
      >
        <div
          class="animate__animated header"
          v-if="!app.hideHeader"
          bordered
          style="height: 64px; line-height: 64px"
        >
          <div class="flex h-full">
            <div class="h-full">
              <one-icon
                color="#000"
                class="inline-block"
                @click="$router.push('/')"
                style="font-size: 48px; margin: 8px; color: aqua"
              >
                glassdoor
              </one-icon>
            </div>
            <div class="h-full flex gap-1" style="">
              <div style="width: 3px; height: 100%; background: var(--L0)"></div>
              <span style="color: var(--L0)">统一认证系统</span>
            </div>
            <div class="flex-grow flex justify-center">
              <span class="text-2xl" style="line-height: 64px">{{ app.title }}</span>
            </div>
            <div class="h-full px-3">
              <fullscreen v-model="isFullScreen" class="header-icon">fullscreen</fullscreen>
              <div class="header-icon">
                <one-icon @click="app.toggle_theme()">
                  {{ app.isDark ? 'Daytimemode-fill' : 'night' }}
                </one-icon>
              </div>
              <div class="header-icon" @click="app.hideHeader = true">
                <one-icon>up</one-icon>
              </div>
            </div>
            <div vif="$store.state.user.ready" class="h-full flex justify-center items-center mr-5">
              <!-- <OAer @logout="$store.commit('user/logout')" :is-dark="IsDark"></OAer> -->
            </div>
          </div>
        </div>
      </transition>
      <div style="height: calc(100vh - 88px)">
        <slot></slot>
      </div>
    </div>
    <div
      bordered
      style="height: 24px; line-height: 24px"
      class="flex justify-around px-3 text-gray-500 text-xs header"
    >
      <span class="hover:text-black cursor-pointer" @click="$router.push({ name: 'about' })">
        关于OA
      </span>
      <span class="hover:text-black cursor-pointer">使用须知</span>
      <span class="hover:text-black cursor-pointer" @click="util.goto('https://veypi.com')">
        ©2021 veypi
      </span>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { useAppStore } from '@/store/app'
import fullscreen from './fullscreen'
import util from '@/libs/util'

let app = useAppStore()
let isFullScreen = false
</script>

<style scoped>
.header {
  background: var(--base-bg);
}
.header-icon {
  display: inline-block;
  font-size: 24px;
  margin: 20px 10px 20px 10px;
}
.header-down {
  font-size: 24px;
  position: fixed;
  right: 76px;
  top: 5px;
  z-index: 100;
}
</style>
