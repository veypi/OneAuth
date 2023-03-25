<template>
  <div class="flex items-center justify-center h-full w-full">
    <div class="box px-10 pb-9 pt-16 rounded-xl w-96 relative">
      <Myinput
        :type="ArgType.Text"
        label-width="8rem"
        v-model="data.username"
        :options="{ min: 5, max: 16 }"
      >
        <template #label>{{ $t('a.username') }}</template>
      </Myinput>
      <Myinput
        :type="ArgType.Password"
        label-width="8rem"
        v-model="data.password"
        :options="{ min: 6, max: 16 }"
      >
        <template #label>{{ $t('a.password') }}</template>
      </Myinput>
      <div class="flex justify-around mt-4">
        <div
          class="div-btn px-4 py-1 rounded"
          style="background: var(--color-primary)"
          @click="login"
        >
          {{ $t('a.login') }}
        </div>
        <div
          class="div-btn px-4 py-1 rounded"
          style="background: var(--color-primary)"
          @click="router.push({ name: 'register' })"
        >
          {{ $t('a.register') }}
        </div>
      </div>
    </div>
  </div>
</template>
<script lang="ts" setup>
import { onMounted, ref, computed } from 'vue'
import api from '@/api'
import { useRoute, useRouter } from 'vue-router'
import msg from '@/msg'
import Myinput from '@/components/myinput'
import { ArgType } from '@/models'
import { useI18n } from 'vue-i18n'
import { useUserStore } from '@/store/user'
import { util } from '@/libs'

const route = useRoute()
const router = useRouter()
const { t } = useI18n()
const u = useUserStore()

let data = ref({
  username: '',
  password: '',
})

function login2() {
  // @ts-ignore
  api.user.login(data.value.username, data.value.password).Start(
    (headers: any) => {
      if ('auth_token' in headers) {
        localStorage.auth_token = headers.auth_token
        u.init_user()
        // store.commit('user/refreshToken', localStorage.auth_token)
        msg.Info(t('a.success'))
        // store.dispatch('user/fetchUserData')
        let url = route.query.redirect || headers.redirect || '/'
        router.push(url)
      } else {
        msg.Info('正在申请加入，请等待管理员审批')
      }
    },
    (e) => {
      console.log(e)
      msg.Warn(t('a.failed') + '：' + e.headers.error)
    },
  )
}

function login() {
  // @ts-ignore
  api.user.login(data.value.username, data.value.password).Start(
    (headers: any) => {
      if ('auth_token' in headers) {
        localStorage.auth_token = headers.auth_token
        // store.commit('user/freshToken', localStorage.auth_token)
        msg.Info('登录成功')
        // store.dispatch('user/fetchUserData')
        let url = route.query.redirect || headers.redirect || '/'
        redirect(url)
      } else {
        msg.Info('正在申请加入，请等待管理员审批')
      }
      console.log(headers)
    },
    (e) => {
      console.log(e)
      msg.Warn('登录失败：' + e.headers.error)
    },
  )
}

function redirect(url?: string) {
  if (uuid.value && uuid.value !== local.id) {
    api.app.get(uuid.value as string).Start(
      (app: modelsApp) => {
        api
          .token(uuid.value as string)
          .get()
          .Start((e) => {
            if (!url) {
              url = app.UserRefreshUrl
            }
            e = encodeURIComponent(e)
            url = url.replaceAll('$token', e)
            window.location.href = url
          })
      },
      (e) => {},
    )
  } else if (util.checkLogin()) {
    if (url) {
      router.push(url)
    } else {
      router.push({ name: 'home' })
    }
  }
}
let uuid = computed(() => {
  return route.query.uuid
})
let ifLogOut = computed(() => {
  return route.query.logout === '1'
})

onMounted(() => {
  if (!ifLogOut.value) {
    redirect()
  }
  // if (divs.value[0]) {
  //   // @ts-ignore
  //   divs.value[0].focus()
  // }
})
</script>

<style scoped>
.box {
  background: linear-gradient(145deg, var(--base-bg-3), var(--base-bg-2));
  box-shadow: 20px 20px 60px var(--base-bg-3), -20px -20px 60px var(--base-bg-2);
}
</style>
