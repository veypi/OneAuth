 <!--
 * login.vue
 * Copyright (C) 2024 veypi <i@veypi.com>
 * 2024-05-31 17:10
 * Distributed under terms of the MIT license.
 -->
<template>
  <div class="login-page flex items-center justify-center">
    <div class="login box">
      <div class="header flex items-center justify-start gap-2">
        <div class="voa-logo"></div>
        <div class="txt">OneAuth</div>
      </div>
      <Transition name="box" mode="out-in">
        <div class="newbie content" v-if="aOpt === 'newbie'">
          <div :check="checks.u" class="username mt-8">
            <input @change="check" v-model="data.username" autocomplete="username" placeholder="username, phone or Email">
          </div>
          <div :check="checks.p" class="password">
            <input @change="check" v-model="data.password" autocomplete="password" type='password' placeholder="password">
          </div>
          <div :check="checks.p2" class="password">
            <input @change="check" v-model="data.confirm" autocomplete="password" type='password' placeholder="password">
          </div>
          <div class="flex">
            <button @click="aOpt = ''" class='ok voa-btn back'>
              back
            </button>
            <button @click="register" class='ok voa-btn'>
              register
            </button>
          </div>
        </div>
        <div class="oh_no content" v-else-if="aOpt === 'oh_no'">
          <div class="username mt-8">
            <input v-model="data.username" autocomplete="username" placeholder="username, phone or Email">
          </div>
          <div class="flex">
            <button @click="aOpt = ''" class='ok back voa-btn'>
              back
            </button>
            <button @click="reset" class='ok voa-btn'>
              confirm
            </button>
          </div>
        </div>
        <div class="login content flex flex-col justify-between" v-else>
          <div :check="checks.u" class="username mt-8">
            <input @change="check" v-model="data.username" autocomplete="username" placeholder="username, phone or Email">
          </div>
          <div :check="checks.p" class="password">
            <input @change="check" v-model="data.password" autocomplete="password" type='password' placeholder="password">
          </div>
          <button @click="login" class='ok voa-btn'>
            login
          </button>
          <div class="last">
            <div class="icos">
              <div class="github"></div>
              <div class="wechat"></div>
              <div class="google"></div>
            </div>
            <div class="txt">
              <div @click="aOpt = 'newbie'">Create Account</div>
              <div @click="aOpt = 'oh_no'">Forgot Password?</div>
            </div>
          </div>
        </div>
      </Transition>
    </div>
  </div>
  <!-- <div class="h-full w-full flex items-center justify-center"> -->
  <!--   <div class="px-10 pb-9 pt-16 rounded-xl w-96 bg-gray-50 relative"> -->
  <!--     <img class='vico' :src="'/favicon.ico'"> -->
  <!--     <Vinput class="mb-8" v-model="data.username" label="用户名" :validate="" /> -->
  <!--     <Vinput class='mb-8' v-model="data.password" label='密码' :validate="" type="password" /> -->
  <!--     <div class="flex justify-around mt-4"> -->
  <!--       <div class='vbtn bg-green-300' @click="$router.push({ -->
  <!--         name: -->
  <!--           'register' -->
  <!--       })">注册</div> -->
  <!--       <div class='vbtn bg-green-300' @click="onSubmit">登录</div> -->
  <!--       <div class='vbtn bg-gray-300' @click="onReset">重置 </div> -->
  <!--     </div> -->
  <!--   </div> -->
  <!-- </div> -->
</template>

<script lang="ts" setup>
import msg from '@veypi/msg';


definePageMeta({
  layout: false,
})

const app = useAppConfig()
const route = useRoute()
const router = useRouter()


let data = ref({
  username: '',
  password: '',
  confirm: '',
})



let uReg = /^[\w]{5,}$/
let pReg = /^[\w@_#]{6,}$/
let checks = ref({ 'u': true, 'p': true, 'p2': true })
let enable_check = ref(false)
const check = () => {
  if (!enable_check.value) {
    return
  }
  checks.value.u = !Boolean(!data.value.username || !uReg.test(data.value.username))
  checks.value.p = !Boolean(!data.value.password || !pReg.test(data.value.password))
  checks.value.p2 = !Boolean(data.value.confirm !== data.value.password)
}
const login = () => {
  enable_check.value = true
  check()
  if (!checks.value.u || !checks.value.p) {
    return
  }
  api.user.login(data.value.username,
    data.value.password).then((data: any) => {
      util.setToken(data.auth_token)
      // msg.Info('登录成功')
      // user.fetchUserData()
      let url = route.query.redirect || data.redirect || ''
      console.log([url])
      redirect(url)
    }).catch(e => {
      msg.Warn(e)
    })
}
const register = () => {
  enable_check.value = true
  check()
  if (!checks.value.u || !checks.value.p || !checks.value.p2) {
    return
  }
  api.user.register(data.value.username, data.value.password).then(u => {
    console.log(u)
    msg.Info('注册成功')
    aOpt.value = ''
  }).catch(e => {
    console.log(e)
    msg.Warn('注册失败：' + e.data)
  })
}
const reset = () => {
  enable_check.value = true
  check()
}

let uuid = computed(() => {
  return route.query.uuid
})

let ifLogOut = computed(() => {
  return route.query.logout === '1'
})
let aOpt = ref('' as '' | 'newbie' | 'oh_no')


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

<style scoped lang="scss">
.login-page {
  margin: 0;
  padding: 0;
  height: 100vh;
  width: 100vw;
  background-color: #fafafa;
  background-image: url("../assets/img/bg.svg");
  background-size: cover;
  background-position: center;
  /* backdrop-filter: blur(5px); */
}

.box {
  user-select: none;
  position: sticky;
  padding: 2rem;
  width: 50%;
  min-width: 20rem;
  height: 50%;

  &::before {
    content: "";
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    border-radius: 2rem;
    background-color: rgba(200, 200, 200, 0.2);
    backdrop-filter: blur(20px);
    /* 模糊效果 */
    z-index: -1;
  }

  .header {
    line-height: 2rem;
    width: 100%;
    height: 4rem;

    .voa-logo {
      height: 4rem;
      width: 4rem;
    }

    .txt {
      font-size: 1.5rem;
    }
  }

  .content {
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    height: calc(100% - 4rem);

    .username,
    .password {
      position: relative;
      width: 100%;

      input {
        height: 2.5rem;
        line-height: 2.5rem;
        font-size: 1.5rem;
        width: calc(100% - 2rem);
        margin: 0 1rem;
        border: none;
        outline: none;
        background: none;
      }

      input:-webkit-autofill,
      input:autofill,
      input:-webkit-autofill:hover,
      input:-webkit-autofill:focus,
      input:-webkit-autofill:active {
        // box-shadow: inset 0 0 0 100px rgba(200, 200, 200, 0.2) !important;
        background-color: #0f0 !important;
        transition: background-color 15000s ease-in-out 0s;
      }

      &::after {
        content: "";
        position: absolute;
        bottom: 0;
        left: 1rem;
        width: calc(100% - 2rem);
        height: 0.1em;
        background-color: #000;
        transition: all 0.3s;
      }

      &:hover::after {
        left: 0%;
        width: 100%;
        background-color: #00ffff;
      }

      &[check='false']::after {
        background-color: #f00 !important;
      }
    }

    .ok {
      line-height: 3rem;
      font-size: 1.5rem;
      height: 3rem;
      margin: 0rem auto;
      width: 40%;
      background: #73f7ca;
      border-radius: 1.5rem;
    }

    .back {
      background: #ccc;
      opacity: 0.5;

      &:hover {
        opacity: 1;
      }
    }
  }
}

.box-enter-active,
.box-leave-active {
  transition: all 0.3s ease-out;
}

.box-enter-from {
  transform: translateX(-20px);
  opacity: 0;
}

.box-leave-to {
  transform: translateX(20px);
  opacity: 0;
}


.login {



  .last {
    flex-shrink: 0;
    display: flex;
    justify-content: space-between;
    padding: 0 1rem;
    height: 3rem;

    .icos {
      display: flex;
      align-items: center;
      gap: 1rem;

      div {
        opacity: 0.5;
        cursor: pointer;
        height: 2rem;
        width: 2rem;
        background-size: cover;
        background-position: center;

        &:hover {
          opacity: 1;
        }
      }

      .github {
        background-image: url("../assets/img/github.svg");
      }

      .google {
        background-image: url("../assets/img/google.svg");
      }

      .wechat {
        background-image: url("../assets/img/wechat.svg");
      }
    }

    .txt {

      height: 1.5rem;
      line-height: 1.5rem;
      font-size: 1rem;

      div {
        cursor: pointer;
        opacity: 0.5;

        &:hover {
          opacity: 1;
        }
      }
    }
  }
}


.voa-logo {
  background-image: url("../assets/img/favicon.svg");
  background-size: cover;
  background-position: center;
}

.voa-btn {
  position: relative;
  text-align: center;
  display: block;
  border: none;
  cursor: pointer;

  &::after {
    content: "";
    position: absolute;
    inset: 0;
    border-radius: inherit;
    transition: 0.3s;
  }

  &:active::after {
    box-shadow: 0 1px 0px 0px rgba(0, 0, 0, 0.5);
  }

  &:active {
    opacity: 0.8;
  }
}
</style>

