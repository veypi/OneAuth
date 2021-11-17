<template>
  <div>
    <div v-if="ofApps.length > 0">
      <h1 class="page-h1">已绑定应用</h1>
      <div class="grid gap-4 grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 2xl:grid-cols-5 text-center">
        <div v-for="(item, k) in ofApps" class="flex items-center justify-center" :key="k">
          <AppCard :core="item"></AppCard>
        </div>
        <div class="flex items-center justify-center" v-for="(item) in '123456789'.split('')"
             :key="item">
          <AppCard :core2="{}"></AppCard>
        </div>
      </div>
    </div>
    <div v-if="apps.length > 0">
      <h1 class="page-h1">应用中心</h1>
      <div class="grid gap-4 grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 2xl:grid-cols-5 text-center">
        <div v-for="(item, k) in apps" class="flex items-center justify-center" :key="k">
          <AppCard :core="item"></AppCard>
        </div>
        <div class="flex items-center justify-center" v-for="(item) in '123456'.split('')"
             :key="item">
          <AppCard :core2="{}"></AppCard>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import {onMounted, ref} from "vue";
import api from "@/api";
import AppCard from '@/components/app.vue'
import {useStore} from "@/store";

let store = useStore()
let apps = ref([])
let ofApps = ref([])

function getApps() {
  api.app.list().Start(e => {
    apps.value = e
    api.app.user('').list(store.state.user.id).Start(e => {
      console.log(e)
      ofApps.value = []
      for (let i in e) {
        let ai = apps.value.findIndex(a => a.id === e[i].app_id)
        if (ai >= 0) {
          apps.value[ai].user_status = e[i].status
          ofApps.value.push(apps.value[ai])
          apps.value.splice(ai, 1)
        }
      }
    })
  })
}

onMounted(() => {
  getApps()
})

</script>

<style scoped>

</style>
