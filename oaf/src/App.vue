<template>
  <base-frame>
    <router-view v-slot="{ Component }">
      <transition mode="out-in" enter-active-class="animate__fadeInLeft"
                  leave-active-class="animate__fadeOutRight">
        <component class="animate__animated animate__400ms" :is="Component"
                   :style="{'min-height': store.state.height}"
                   style="margin: 10px"
        ></component>
      </transition>
    </router-view>
  </base-frame>
</template>
<script setup lang="ts">
// This starter template is using Vue 3 <script setup> SFCs
import BaseFrame from './components/frame.vue'
import {onBeforeMount, ref} from 'vue'
import {useStore} from "./store";

let store = useStore()

onBeforeMount(() => {
  let loader = document.getElementById("loader-wrapper")
  if (loader && loader.parentElement) {
    loader.parentElement.removeChild(loader)
  }
  store.dispatch('fetchSelf')
  store.dispatch('user/fetchUserData')
})

let collapsed = ref(true)

</script>


<style lang="less">
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

html,
body {
  width: 100%;
  height: 100%;
  margin: 0;
  padding: 0;
}

/* 周围滑动留白 */
html {
  overflow: hidden;
  height: 100%;
}

body {
  overflow: auto;
  height: 100%;
}

.header-icon {
  display: inline-block;
  font-size: 24px;
  margin: 20px 10px 20px 10px;
}

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
  width: 100%;
  height: 100%;
}

::-webkit-scrollbar {
  display: none; /* Chrome Safari */
}
</style>
