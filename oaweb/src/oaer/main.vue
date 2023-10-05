<template>
  <BaseFrame :class="[isDark ? 'oa_light' : 'oa_dark']" v-model="shown" :is-dark="isDark">
    <template #title>
      {{ self.name }}
    </template>
    <div class="flex justify-center items-center">
      <img class="oa_avatar mx-2" :src="Cfg.host.value + usr.icon" alt="Avatar" />
    </div>
    <template #main>
      <div style="height: 100%">
        <div style="height: calc(100% - 50px)">
          <div class="w-full px-3">
            <div class="h-16 flex justify-between items-center">
              <span style="">我的账户</span>
              <span @click="shown = false" class="cursor-pointer" style="color:#f36828">账户中心</span>
            </div>
            <div class="grid grid-cols-4 gap-4 h-20">
              <div class="flex items-center justify-center">
                <img class="oa_avatar mx-2" :src="Cfg.host.value + usr.icon" alt="Avatar" />
              </div>
              <div class="col-span-2 text-xs grid grid-cols-1 items-center text-left" style="">
                <span>昵称: &ensp;&ensp; {{ usr.nickname }}</span>
                <span>账户: &ensp;&ensp; {{ usr.username }}</span>
                <span>邮箱: &ensp;&ensp; {{ usr.email }}</span>
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
            <OneIcon class="inline-block" style="font-size: 24px;">
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
import { OneIcon } from '@veypi/one-icon'
import { computed, onMounted, ref, watch } from 'vue'
import { decode } from 'js-base64'
import { api, Cfg } from './api'
import evt from './evt'
import { modelsApp, modelsUser } from './models'

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
  console.log(data)
  if (data.id) {
    api.user.get(data.id).then(e => {
      console.log(e)
      usr.value = e
      ofApps.value = []
      for (let v of e.Apps) {
        if (v.Status === 'ok') {
          ofApps.value.push(v.App)
        }
        if (v.App.id === Cfg.uuid.value) {
          self.value = v.App
        }
      }
      emits('load', e)
    }).catch(e => {
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

<style>
.oa_light {
  color: #eee;
}

.oa_dark {
  color: #333;

}

.oa_avatar {
  vertical-align: middle;
  width: 2.5rem;
  height: 2.5rem;
  border-radius: 50%;
}
</style>
