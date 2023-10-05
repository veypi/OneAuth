<template>
  <div class="w-full px-3">
    <div class="h-16 flex justify-between items-center">
      <span style="">我的应用</span>
      <span @click="Cfg.goto('/')" class="cursor-pointer" style="color:#f36828">应用中心</span>
    </div>
    <div class="grid grid-cols-5">
      <template v-for="(ap, ai) of apps" :key="ai">
        <div class="mx-2" @click="Cfg.goto(ap.redirect)" v-if="ap.id !== Cfg.uuid.value">
          <img class="oa_avator" v-if="ap.icon" :src="Cfg.host.value +
            ap.icon" alt="Avatar" />
        </div>
      </template>
    </div>
    <hr class="mt-10" style="border:none;border-top:1px solid #777;">
  </div>
</template>
<script lang="ts" setup>
import { modelsApp } from '../models'
import { Cfg } from '../api'

withDefaults(defineProps<{
  apps: modelsApp[]
}>(), {})
</script>
<style scoped></style>
