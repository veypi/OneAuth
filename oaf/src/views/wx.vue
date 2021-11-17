<template>
  <div class='home d-flex justify-center align-center'>
    <wx-login v-if="enable" :aid="aid" :app="agentID" :url="url"></wx-login>
  </div>
</template>
<script setup lang='ts'>
import WxLogin from '@/components/WxLogin.vue'
import {computed, onMounted} from "vue";
import {useRoute} from 'vue-router'
import api from '@/api'

let route = useRoute()

let aid = ''
let agentID = ''
let url = ''
let uuid = computed(() => {
  return route.query.uuid
})
let enable = computed(() => {
  return uuid && aid && agentID && url
})
let code = computed(() => {
  return route.query.code
})

let state = computed(() => {
  return route.query.state
})

let msg = computed(() => {
  return route.query.msg
})
onMounted(() => {
  if (msg) {
    console.log(msg)
    alert(msg)
  }
})

if (uuid) {
  api.app.get(uuid.value as string).Start(e => {
    url = e.wx.url + '/api/wx/login/' + uuid
    aid = e.wx.corp_id
    agentID = e.wx.agent_id
  })
}
</script>
<style scoped>

</style>
