<template>
  <n-config-provider :theme-overrides="Theme.overrides" :locale="zhCN" :date-locale="dateZhCN"
                     :theme="Theme">
    <n-message-provider>
      <n-layout class="font-sans select-none" style="height: 100vh">
        <transition enter-active-class="animate__slideInDown" leave-active-class="animate__slideOutUp">
          <one-icon class="header-down animate__animated" @click="$store.commit('hideHeader', false)"
                    v-if="$store.state.hideHeader">down
          </one-icon>
        </transition>
        <n-layout style="height: calc(100vh - 24px)">
          <transition enter-active-class="animate__slideInDown" leave-active-class="animate__slideOutUp">
            <n-layout-header class="animate__animated" v-if="!$store.state.hideHeader" bordered
                             style="height: 64px;line-height: 64px;">
              <div class="flex h-full">
                <div class="h-full">
                  <one-icon color="#000" class="inline-block" @click="$router.push('/')"
                            style="font-size: 48px;margin:8px;color:aqua">
                    glassdoor
                  </one-icon>
                </div>
                <div class="h-full" style="margin-left: 10px">
                  <n-h6 prefix="bar" align-text>
                    <n-text type="primary">统一认证系统</n-text>
                  </n-h6>
                </div>
                <div class="flex-grow flex justify-center">
                  <span class="text-2xl" style="line-height: 64px">{{ $store.state.title }}</span>
                </div>
                <div class="h-full px-3">
                  <fullscreen v-model="isFullScreen" class="header-icon">fullscreen</fullscreen>
                  <div class="header-icon">
                    <one-icon @click="ChangeTheme">
                      {{ IsDark ? 'Daytimemode' : 'nightmode-fill' }}
                    </one-icon>
                  </div>
                  <div class="header-icon" @click="$store.commit('hideHeader', true)">
                    <one-icon>up</one-icon>
                  </div>
                </div>
                <div v-if="$store.state.user.ready"
                     class="h-full flex justify-center items-center mr-5">
                  <OAer @logout="$store.commit('user/logout')" :is-dark="IsDark"></OAer>
                </div>
              </div>
            </n-layout-header>
          </transition>
          <n-layout :native-scrollbar="false">
            <n-loading-bar-provider>
              <n-dialog-provider>
                <slot></slot>
              </n-dialog-provider>
            </n-loading-bar-provider>
            <n-back-top>
            </n-back-top>
          </n-layout>
        </n-layout>
        <n-layout-footer bordered style="height: 24px;line-height: 24px"
                         class="flex justify-around px-3 text-gray-500 text-xs">
          <span class="hover:text-black cursor-pointer" @click="$router.push({name: 'about'})">关于OA</span>
          <span class="hover:text-black cursor-pointer">使用须知</span>
          <span class="hover:text-black cursor-pointer" @click="util.goto('https://veypi.com')">
      ©2021 veypi
      </span>
        </n-layout-footer>
      </n-layout>
    </n-message-provider>
  </n-config-provider>
</template>

<script lang="ts" setup>
import {Theme, IsDark, ChangeTheme} from '@/theme'
import {zhCN, dateZhCN} from 'naive-ui'
import fullscreen from './fullscreen'
import avatar from './avatar'
import {onMounted, ref} from 'vue'
import {useStore} from '@/store'
import {useRouter} from 'vue-router'
import util from '@/libs/util'
import {OAer, Cfg} from '@/oaer'
Cfg.token.value = util.getToken()

let store = useStore()
let router = useRouter()

let isFullScreen = ref(false)
onMounted(() => {
})

onMounted(() => {
})
</script>

<style scoped>
.header-down {
  font-size: 24px;
  position: fixed;
  right: 76px;
  top: 5px;
  z-index: 100;
}
</style>
