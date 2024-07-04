 <!--
 * default.vue
 * Copyright (C) 2024 veypi <i@veypi.com>
 * 2024-05-31 17:09
 * Distributed under terms of the MIT license.
 -->
<template>
  <div class="page">
    <div class="header flex justify-center items-center">
      <div class="ico" @click="router.push('/')"></div>
      <div>统一认证系统</div>
      <div class="grow"></div>
      <OneIcon class="mx-2" @click="toggle_fullscreen" :name="app.layout.fullscreen ? 'compress' : 'expend'"></OneIcon>
      <OneIcon class="mx-2" @click="toggle_theme" :name="app.layout.theme === '' ? 'light' : 'dark'"></OneIcon>
      <OAer class="mx-2" v-if="user.ready" @logout="user.logout" :is-dark="app.layout.theme !== ''">
      </OAer>
    </div>
    <div class="menu">
      <Menu :show_name="menu_mode === 2"></Menu>
    </div>
    <div class="menu-hr"></div>
    <div class="main px-8 py-6">
      <slot />
    </div>
    <div class="footer flex justify-around items-center">
      <div @click="util.goto('https://veypi.com')">© 2024 veypi</div>
      <div>使用说明</div>
      <div>联系我们</div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { OneIcon } from '@veypi/one-icon'
import { OAer } from "@veypi/oaer";
import oaer from '@veypi/oaer'
import '@veypi/oaer/dist/index.css'

let app = useAppConfig()
let router = useRouter()
let user = useUserStore()

app.host = window.location.protocol + '//' + window.location.host
oaer.set({
  token: util.getToken(),
  host: app.host,
  uuid: app.id,
})

bus.on('token', (t: any) => {
  oaer.set({ token: t })
})

if (!util.checkLogin()) {
  router.push('/login')
} else {
  user.fetchUserData()
}
let menu_mode = ref(1)
let toggle_menu = (m: 0 | 1 | 2) => {
  menu_mode.value = m
  if (m == 0) {
    app.layout.menu_width = 0
  } else if (m == 1) {
    app.layout.menu_width = 40
  } else {
    app.layout.menu_width = 108
  }
}
toggle_menu(2)

const toggle_fullscreen = () => {
  app.layout.fullscreen = !app.layout.fullscreen
  if (app.layout.fullscreen) {
    let docElm = document.documentElement;
    docElm.requestFullscreen();
  } else {
    document.exitFullscreen();
  }
}
const toggle_theme = () => {
  app.layout.theme =
    app.layout.theme === '' ? 'dark' : ''
  document.documentElement.setAttribute('theme', app.layout.theme)
}


</script>

<style scoped lang="scss">
.page {
  height: 100vh;
  width: 100vw;

  .header {
    height: v-bind('app.layout.header_height + "px"');
    user-select: none;
    background: var(--header-bg);
    color: var(--header-txt);
    font-size: 24px;

    .ico {
      width: v-bind('app.layout.header_height * 0.8 + "px"');
      height: v-bind('app.layout.header_height * 0.8 + "px"');
      background: url('/favicon.ico') no-repeat;
      background-size: cover;
    }
  }

  .menu {
    overflow: hidden;
    vertical-align: top;
    display: inline-block;
    width: v-bind("app.layout.menu_width + 'px'");
    height: calc(100vh - v-bind('app.layout.header_height + app.layout.footer_height + "px"'));
    transition: width 0.3s linear;
  }

  .menu-hr {
    vertical-align: top;
    display: inline-block;
    width: 1px;
    height: calc(100vh - v-bind('app.layout.header_height + app.layout.footer_height + "px"'));
    background: #999;
  }

  .main {
    vertical-align: top;
    display: inline-block;
    overflow: auto;
    width: calc(100vw - v-bind("app.layout.menu_width + 1 + 'px'"));
    height: calc(100vh - v-bind('app.layout.header_height + app.layout.footer_height + "px"'));
    transition: width 0.3s linear;
  }

  .footer {
    height: v-bind("app.layout.footer_height + 'px'");
    user-select: none;
    background: var(--footer-bg);
    color: var(--footer-txt);
    font-size: 12px;
    line-height: v-bind("app.layout.footer_height + 'px'");

    div {
      cursor: pointer;
      opacity: 0.6;
    }

    div:hover {
      opacity: 1;
    }
  }
}
</style>



