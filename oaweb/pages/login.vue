 <!--
 * login.vue
 * Copyright (C) 2024 veypi <i@veypi.com>
 * 2024-05-31 17:10
 * Distributed under terms of the MIT license.
 -->
<template>
  <div class="h-full w-full flex items-center justify-center">
    <div class="px-10 pb-9 pt-16 rounded-xl w-96 bg-gray-50 relative">
      <img class='vico' :src="'/favicon.ico'">
      <Vinput class="mb-8" v-model="data.username" label="用户名" :validate="/^[\w]{5,}$/" />
      <Vinput class='mb-8' v-model="data.password" label='密码' :validate="/^[\w@_#]{6,}$/" type="password" />
      <div class="flex justify-around mt-4">
        <div class='vbtn bg-green-300' @click="$router.push({
          name:
            'register'
        })">注册</div>
        <div class='vbtn bg-green-300' @click="onSubmit">登录</div>
        <div class='vbtn bg-gray-300' @click="onReset">重置 </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>

definePageMeta({
  layout: false,
})

const app = useAppConfig()
const route = useRoute()
const router = useRouter()


let data = ref({
  username: '',
  password: '',
})


const onSubmit = () => {
  api.user.login(data.value.username,
    data.value.password).then((data: any) => {
      util.setToken(data.auth_token)
      // msg.Info('登录成功')
      // user.fetchUserData()
      let url = route.query.redirect || data.redirect || ''
      console.log([url])
      redirect(url)
    }).catch(e => {
      let m = e === '1' ? '被禁止登录' : e === '2' ? '正在申请中' : e
        === '3' ?
        '申请被拒绝' : '登录失败:' + e
      // msg.Warn(m)
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
    api.app.get(uuid.value as string).then((app) => {
      api.token(uuid.value as string).get({
        token: util.getToken(),
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
    router.push('/')
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

<style scoped>
.vico {
  width: 120px;
  height: 120px;
  position: absolute;
  top: -50px;
  left: -60px;
}
</style>

