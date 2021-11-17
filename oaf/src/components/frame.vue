<template>
  <n-config-provider :theme-overrides="Theme.overrides" :locale="zhCN" :date-locale="dateZhCN"
                     :theme="Theme">
    <n-message-provider>
      <n-layout class="font-sans select-none">
        <n-layout>
          <n-layout-header class="pr-5" bordered style="height: 64px;line-height: 64px;">
            <div class="inline-block float-left h-full">
              <one-icon color="#000" class="inline-block" @click="$router.push('/')"
                        style="font-size: 48px;margin:8px;color:aqua">
                glassdoor
              </one-icon>
            </div>
            <div class="inline-block float-left h-full" style="margin-left: 10px">
              <n-h6 prefix="bar" align-text>
                <n-text type="primary">统一认证系统</n-text>
              </n-h6>
            </div>
            <div v-if="$store.state.user.ready"
                 class="inline-block h-full float-right flex justify-center items-center">
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
          <n-layout style="height: calc(100vh - 88px)" :native-scrollbar="false">
            <slot></slot>
            <n-back-top>
            </n-back-top>
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
    </n-message-provider>
  </n-config-provider>
</template>

<script lang="ts" setup>
import {Theme, IsDark, ChangeTheme} from "@/theme";
import {zhCN, dateZhCN} from 'naive-ui'
import fullscreen from './fullscreen'
import avatar from './avatar'
import {onMounted, ref} from "vue";


let goto = (url: any) => {
  window.open(url, "_blank")
}

let isFullScreen = ref(false)
onMounted(() => {
})
</script>

<style scoped>

</style>
