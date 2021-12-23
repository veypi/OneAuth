<template>
  <div class="w-full px-3">
    <div class="h-16 flex justify-between items-center">
      <span style="color: #777">我的应用</span>
      <span @click="Cfg.host.value?goto(Cfg.host.value):redirect('/')" class="cursor-pointer"
            style="color:#f36828">应用中心</span>
    </div>
    <div class="grid grid-cols-5">
      <template v-for="(ap,ai) of apps"
                :key="ai">
      <div class="mx-2" @click="redirect(ap.Host)" v-if="ap.UUID !== Cfg.uuid.value">
        <n-avatar v-if="ap.Icon" size="2" :src="Cfg.host.value+ ap.Icon" round></n-avatar>
      </div>
      </template>
    </div>
    <hr class="mt-10" style="border:none;border-top:1px solid #777;">
  </div>
</template>
<script lang="ts" setup>
import {modelsApp} from '../models'
import {Cfg} from '../api'

let props = withDefaults(defineProps<{
  apps: modelsApp[]
}>(), {})

function goto(url: string) {
  window.open(url, '_blank')
}

function redirect(url: string) {
  window.location.href = url
}
</script>
<style scoped>
</style>
