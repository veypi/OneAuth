<template>
  <q-layout view="hHh LpR fFf">
    <q-header elevated class="bg-primary text-white" height-hint="98">
      <q-toolbar class="pl-0">

        <q-toolbar-title class="flex items-center cursor-pointer" @click="router.push({ name: 'home' })">
          <q-icon size="3rem" class="mx-1" color="#0ff" name='v-glassdoor' style="color: aqua;"></q-icon>
          <q-separator dark vertical inset />
          <span class="ml-3">
            统一认证系统
          </span>
        </q-toolbar-title>

        <q-icon class="mx-2" size="2rem" @click="$q.fullscreen.toggle()"
          :name="$q.fullscreen.isActive ? 'v-compress' : 'v-expend'" />

        <q-icon class="mx-2" size="2rem" @click="$q.dark.toggle" :name="$q.dark.mode ? 'v-light' : 'v-dark'"></q-icon>
        <OAer v-if="user.ready" @logout="user.logout" :is-dark="$q.dark.mode as boolean">
          123
        </OAer>
      </q-toolbar>
      <q-toolbar class="">
        <q-icon @click="toggleLeftDrawer" class="cursor-pointer" name="menu" size="sm"></q-icon>
        <q-tabs align="left">
          <!-- <q-route-tab to="/page1" label="Page One" /> -->
          <!-- <q-route-tab to="/page2" label="Page Two" /> -->
          <!-- <q-route-tab to="/page3" label="Page Three" /> -->
        </q-tabs>
      </q-toolbar>
    </q-header>

    <q-drawer show-if-above :mini="miniState" :width="140" :breakpoint="500" bordered side="left" class="pt-4">
      <Menu></Menu>
    </q-drawer>

    <q-page-container class="flex">
      <q-page class="w-full">
        <router-view v-slot="{ Component }">
          <transition mode="out-in" enter-active-class="animate__fadeIn" leave-active-class="animate__fadeOut">
            <component class="animate__animated animate__400ms py-8
              px-8 h-full
              w-full" :is="Component"></component>
          </transition>
        </router-view>
      </q-page>
    </q-page-container>
    <q-footer style="z-index: 1;" bordered class="bg-grey-8 text-white flex justify-around">
      <span class="hover:text-black cursor-pointer" @click="$router.push({ name: 'doc' })">关于OA</span>
      <span class="hover:text-black cursor-pointer" @click="$router.push({ name: 'doc' })">使用须知</span>
      <span class="hover:text-black cursor-pointer" @click="util.goto('https://veypi.com')">
        ©2021 veypi
      </span>
    </q-footer>

  </q-layout>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import Menu from './menu.vue'
import { useUserStore } from 'src/stores/user';
import { OAer } from "@veypi/oaer";
import { util } from 'src/libs';

const user = useUserStore()
const router = useRouter()




const miniState = ref(false)



function toggleLeftDrawer() {
  miniState.value = !miniState.value
}
</script>

<style>
.animate__400ms {
  --animate-duration: 300ms;
}
</style>
