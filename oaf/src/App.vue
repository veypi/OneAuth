<template>
    <base-frame>
        <router-view v-slot="{ Component }">
            <transition mode="out-in" enter-active-class="animate__fadeInLeft"
                leave-active-class="animate__fadeOutRight">
                <component class="animate__animated animate__400ms" :is="Component"
                    :style="{ 'min-height': store.state.height }" style="margin: 10px"></component>
            </transition>
        </router-view>
    </base-frame>
</template>
<script setup lang="ts">
// This starter template is using Vue 3 <script setup> SFCs
import BaseFrame from './components/frame.vue'
import { onBeforeMount, ref } from 'vue'
import { useStore } from "./store";
import msg from '@veypi/msg'

let store = useStore()

onBeforeMount(() => {
    let loader = document.getElementById("loader-wrapper")
    if (loader && loader.parentElement) {
        loader.parentElement.removeChild(loader)
    }
    store.dispatch('fetchSelf')
    store.dispatch('user/fetchUserData')
    msg.Warn('asd')
})

let collapsed = ref(true)

</script>


<style lang="postcss">
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

.header-icon {
    display: inline-block;
    font-size: 24px;
    margin: 20px 10px 20px 10px;
}

#app {
    @apply text-base font-mono text-center h-full w-full;
    color: #2c3e50;
}

::-webkit-scrollbar {
    display: none;
    /* Chrome Safari */
}
</style>
