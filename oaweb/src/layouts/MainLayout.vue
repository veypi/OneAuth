<template>
  <q-layout view="hHh LpR fFf">
    <q-header elevated class="bg-primary text-white" height-hint="98">
      <q-toolbar>
        <q-icon size="xl" color="aqua" name='svguse:#icon-glassdoor'></q-icon>

        <q-toolbar-title>
          统一认证系统
        </q-toolbar-title>

        <div>OneAuth v2.0.0</div>
      </q-toolbar>
      <q-toolbar class="">
        <q-icon @click="toggleLeftDrawer" class="cursor-pointer" name="menu" size="sm"></q-icon>
        <q-tabs align="left">
          <q-route-tab to="/page1" label="Page One" />
          <q-route-tab to="/page2" label="Page Two" />
          <q-route-tab to="/page3" label="Page Three" />
        </q-tabs>

      </q-toolbar>
    </q-header>

    <q-drawer v-model="leftDrawerOpen" side="left" bordered>
      <q-list>
        <q-item-label header>
          Essential Links
        </q-item-label>

        <EssentialLink v-for="link in essentialLinks" :key="link.title" v-bind="link" />
      </q-list>
    </q-drawer>

    <q-page-container>
      <router-view />
    </q-page-container>
    <q-footer bordered class="bg-grey-8 text-white flex justify-around">
      <span class="hover:text-black cursor-pointer" @click="$router.push({ name: 'about' })">关于OA</span>
      <span class="hover:text-black cursor-pointer">使用须知</span>
      <span class="hover:text-black cursor-pointer" @click="util.goto('https://veypi.com')">
        ©2021 veypi
      </span>
    </q-footer>

  </q-layout>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import EssentialLink, { EssentialLinkProps } from 'components/EssentialLink.vue';
import { util } from 'src/libs';

const essentialLinks: EssentialLinkProps[] = [
  {
    title: 'Docs',
    caption: 'quasar.dev',
    icon: 'school',
    link: 'https://quasar.dev'
  },
  {
    title: 'Github',
    caption: 'github.com/quasarframework',
    icon: 'code',
    link: 'https://github.com/quasarframework'
  },
]

const leftDrawerOpen = ref(false)

function toggleLeftDrawer() {
  leftDrawerOpen.value = !leftDrawerOpen.value
}
</script>
