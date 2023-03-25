<style>
#app {
  @apply h-full w-full flex justify-center items-center;
  color: var(--base-color);
  background: var(--base-bg-3);
}
.animate__400ms {
  --animate-duration: 400ms;
}
.page-h1 {
  font-size: 1.5rem;
  line-height: 2rem;
  margin-left: 2.5rem;
  margin-top: 1.25rem;
  margin-bottom: 1.25rem;
}
::-webkit-scrollbar {
  display: none;
  /* Chrome Safari */
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
onBeforeMount(() => {
  let loader = document.getElementById('loader-wrapper')
  if (loader && loader.parentElement) {
    loader.parentElement.removeChild(loader)
  }
  // store.dispatch('fetchSelf')
  // store.dispatch('user/fetchUserData')
  user.fetchUserData()
})
</script>
