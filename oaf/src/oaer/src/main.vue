<template>
  <BaseFrame v-model="shown" :is-dark="isDark">
    <template #title>
      {{self.Name}}
    </template>
    <slot>
      <div class="flex justify-center items-center">
        <n-avatar style="--color: none" :src="Cfg.host.value + usr.Icon"
                  round></n-avatar>
      </div>
    </slot>
    <template v-slot:main>
    <div style="height: 100%">
      <div style="height: calc(100% - 50px)">
        <div class="w-full px-3">
          <div class="h-16 flex justify-between items-center">
            <span style="color: #777">我的账户</span>
            <span @click="$router.push({name: 'user_setting'});shown=false" class="cursor-pointer"
                  style="color:#f36828">账户中心</span>
          </div>
          <div class="grid grid-cols-4 gap-4 h-20">
            <div class="flex items-center justify-center">
              <n-avatar size="50" :src="Cfg.host.value+ usr.Icon" round></n-avatar>
            </div>
            <div class="col-span-2 text-xs grid grid-cols-1 items-center text-left" style="">
              <span>昵称: &ensp;&ensp; {{ usr.Nickname }}</span>
              <span>账户: &ensp;&ensp; {{ usr.Username }}</span>
              <span>邮箱: &ensp;&ensp; {{ usr.Email }}</span>
            </div>
            <div class="">123</div>
          </div>
          <hr class="mt-10" style="border:none;border-top:1px solid #777;">
        </div>
        <File :usr="usr"></File>
        <Apps :apps="ofApps"></Apps>
      </div>
      <hr style="border:none;border-top:2px solid #777;">
      <div style="height: 48px">
        <div @click="evt.emit('logout')"
             class="w-full h-full flex justify-center items-center cursor-pointer transition duration-500 ease-in-out transform hover:scale-125">
          <OneIcon :color="isDark?'#eee': '#333'" class="inline-block" style="font-size: 24px;">
            logout
          </OneIcon>
          <div>
            退出登录
          </div>
        </div>
      </div>
    </div>
    </template>
  </BaseFrame>
</template>
<script lang="ts" setup>
import BaseFrame from './frame.vue'
import Apps from './components/app.vue'
import File from './components/file.vue'
import {OneIcon} from '@veypi/one-icon'
import {computed, onMounted, ref, watch} from 'vue'
import {decode} from 'js-base64'
import {api, Cfg} from './api'
import evt from './evt'
import {modelsApp, modelsUser} from './models'

let shown = ref(false)
let emits = defineEmits<{
  (e: 'logout'): void
  (e: 'load', u: modelsUser): void
}>()
let props = withDefaults(defineProps<{
  isDark?: boolean
}>(), {
  isDark: false,
})
onMounted(() => {
  fetchUserData()
})

let usr = ref<modelsUser>({} as modelsUser)
let ofApps = ref<modelsApp[]>([])
let self = ref<modelsApp>({} as modelsApp)

let token = computed(() => Cfg.token.value)
watch(token, () => {
  fetchUserData()
})

function fetchUserData() {
  let token = Cfg.token.value?.split('.')
  if (!token || token.length !== 3) {
    return false
  }
  let data = JSON.parse(decode(token[1]))
  if (data.ID > 0) {
    api.user.get(data.ID).Start(e => {
      usr.value = e
      console.log(e)
      ofApps.value = []
      for (let v of e.Apps) {
        if (v.Status === 'ok') {
          ofApps.value.push(v.App)
        }
        if (v.App.UUID === Cfg.uuid.value) {
          self.value = v.App
        }
      }
      emits('load', e)
    }, e => {
      console.log(e)
      evt.emit('logout')
    })
  } else {
    evt.emit('logout')
  }
}



evt.on('logout', () => {
  emits('logout')
})
</script>

<style scoped>

</style>
