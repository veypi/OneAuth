<template>
  <div class="flex items-center justify-center">
    <div class="px-10 pb-9 pt-28 rounded-xl w-96">
      <q-form autofocus @submit="onSubmit" @reset="onReset">
        <q-input v-model="data.username" autocomplete="username" label="用户名" hint="username" lazy-rules
          :rules="data_rules.username" />
        <q-input autocomplete="current-password" v-model="data.password" :type="isPwd ? 'password' :
          'text'" hint="password" :rules="data_rules.password">
          <template v-slot:append>
            <q-icon :name="isPwd ? 'visibility_off' : 'visibility'" class="cursor-pointer" @click="isPwd = !isPwd" />
          </template>
        </q-input>

        <div class="flex justify-around mt-4">
          <q-btn label="注册" @click="router.push({ name: 'register' })" color="info"></q-btn>
          <q-btn label="登录" type="submit" color="primary" />
          <q-btn label="重置" type="reset" color="primary" flat class="q-ml-sm" />
        </div>
      </q-form>
    </div>
  </div>
</template>
<script lang="ts" setup>
import { computed, onMounted, ref, } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import api from 'src/boot/api'
import msg from '@veypi/msg'
import util from 'src/libs/util'
import { useUserStore } from 'src/stores/user'
import { useAppStore } from 'src/stores/app'
import { AUStatus, modelsApp } from 'src/models'
import { Base64 } from 'js-base64'


const app = useAppStore()
const user = useUserStore()
const route = useRoute()
const router = useRouter()


let data = ref({
  username: '',
  password: '',
})
const data_rules = {
  username: [
    (v: string) => v && v.length >= 3 && v.length <= 16 || '长度要求3~16'
  ],
  password: [
    (v: string) => v && v.length >= 6 && v.length <= 16 || '长度要求6~16'
  ]
}
let isPwd = ref(true)

const onSubmit = () => {
  api.user.login(data.value.username,
    data.value.password).then((data: any) => {
      util.setToken(data.auth_token)
      msg.Info('登录成功')
      user.fetchUserData()
      let url = route.query.redirect || data.redirect || ''
      redirect(url)
    }).catch(e => {
      let m = e === '1' ? '被禁止登录' : e === '2' ? '正在申请中' : e
        === '3' ?
        '申请被拒绝' : '登录失败:' + e
      msg.Warn(m)
    })
}
const onReset = () => {
  data.value.password = ''
  data.value.username = ''
}

let uuid = computed(() => {
  return route.query.uuid
})
let ifLogOut = computed(() => {
  return route.query.logout === '1'
})


function redirect(url: string) {
  if (url === 'undefined') {
    url = ''
  }
  if (uuid.value && uuid.value !== app.id) {
    api.app.get(uuid.value as string).then((app: modelsApp) => {
      api.token(uuid.value as string).get({
        token:
          util.getToken()
      }).then(e => {
        url = url || app.redirect
        // let data = JSON.parse(Base64.decode(e.split('.')[1]))
        // console.log(data)
        e = encodeURIComponent(e)
        if (url.indexOf('$token') >= 0) {
          url = url.replaceAll('$token', e)
        } else {
          url = buildURL(url, 'token=' + e)
        }
        window.location.href = url

      })
    })
  } else if (url) {
    router.push(url)
  } else {
    router.push({ name: 'home' })
  }
}
function buildURL(url: string, params?: string) {
  if (!params) {
    return url;
  }

  // params序列化过程略
  var hashmarkIndex = url.indexOf('#');
  if (hashmarkIndex !== -1) {
    url = url.slice(0, hashmarkIndex);
  }

  url += (url.indexOf('?') === -1 ? '?' : '&') + params;

  return url;
}

onMounted(() => {
  if (ifLogOut.value) {
    util.setToken('')
  } else if (util.checkLogin()) {
    redirect(route.query.redirect as string || '')
  }
})
</script>

<style scoped></style>
