<style>
#app {
  @apply h-full w-full flex justify-center items-center;
}
</style>

<template>
  <Frame>
    <router-view v-slot="{ Component }">
      <transition
        mode="out-in"
        enter-active-class="animate__fadeIn"
        leave-active-class="animate__fadeOut"
      >
        <component class="animate__animated animate__faster" :is="Component"></component>
      </transition>
    </router-view>
  </Frame>
</template>

<script setup lang="ts">
import Frame from '@/components/frame.vue'
import { useUserStore } from '@/store/user'
import { onBeforeMount } from 'vue'

let user = useUserStore()
user.setUser()
onBeforeMount(() => {
  let loader = document.getElementById('loader-wrapper')
  if (loader && loader.parentElement) {
    loader.parentElement.removeChild(loader)
  }
  // store.dispatch('fetchSelf')
  // store.dispatch('user/fetchUserData')
})
</script>
