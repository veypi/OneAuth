<template>
  <q-layout view="hHh LpR fFf">
    <q-header elevated class="bg-primary text-white" height-hint="98">
      <q-toolbar class="h-16 pl-0">

        <q-toolbar-title class="flex items-center cursor-pointer" @click="router.push({ name: 'home' })">
          <q-icon size="3.5rem" color="aqua" name='svguse:#icon-glassdoor' style="color: aqua;"></q-icon>
          <q-separator dark vertical inset />
          <span class="ml-3">
            统一认证系统
          </span>
        </q-toolbar-title>

        <q-icon class="mx-2" size="2rem" @click="$q.fullscreen.toggle()"
          :name="$q.fullscreen.isActive ? 'fullscreen_exit' : 'fullscreen'" />

        <q-icon class="mx-2" size="1.5rem" @click="$q.dark.toggle"
          :name="$q.dark.mode ? 'light_mode' : 'dark_mode'"></q-icon>
        <OAer @logout="user.logout" :is-dark="$q.dark.mode as boolean"></OAer>
      </q-toolbar>
      <!-- <q-toolbar class=""> -->
      <!--   <q-icon @click="toggleLeftDrawer" class="cursor-pointer" name="menu" size="sm"></q-icon> -->
      <!--   <q-tabs align="left"> -->
      <!--     <q-route-tab to="/page1" label="Page One" /> -->
      <!--     <q-route-tab to="/page2" label="Page Two" /> -->
      <!--     <q-route-tab to="/page3" label="Page Three" /> -->
      <!--   </q-tabs> -->
      <!-- </q-toolbar> -->
    </q-header>

    <q-drawer show-if-above :mini="miniState" @mouseover="miniState = false" @mouseout="miniState = true" mini-to-overlay
      :width="200" :breakpoint="500" bordered v-model="leftDrawerOpen" side="left" class="pt-4">
      <Menu></Menu>
    </q-drawer>

    <q-page-container class="flex">
      <q-page class="w-full">
        <router-view />
      </q-page>
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
import { util } from 'src/libs';
import { useRouter } from 'vue-router';
import Menu from 'src/components/menu.vue'
import { useAppStore } from 'src/stores/app';
import { useUserStore } from 'src/stores/user';
import { OAer } from "@veypi/oaer";

const app = useAppStore()
const user = useUserStore()
const router = useRouter()




const leftDrawerOpen = ref(false)
const miniState = ref(true)



function toggleLeftDrawer() {
  leftDrawerOpen.value = !leftDrawerOpen.value
}
</script>
