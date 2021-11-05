<template>
  <div class="grid gap-4 grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 2xl:grid-cols-5 text-center">
    <div class="flex items-center justify-center" v-for="(item, k) in apps" :key="k">
      <AppCard :core="item"></AppCard>
    </div>
    <div class="flex items-center justify-center" v-for="(item) in '1234567890'" :key="item">
      <AppCard :core="{}"></AppCard>
    </div>
  </div>
</template>

<script lang="ts" setup>
import {onMounted, ref} from "vue";
import api from "../api";
import AppCard from '../components/app.vue'

let apps = ref([])

function getApps() {
  api.app.list().Start(e => {
    apps.value = e
  })
}

onMounted(() => {
  getApps()
})

</script>

<style scoped>

</style>
