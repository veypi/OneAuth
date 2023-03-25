 <!--
 * index.vue
 * Copyright (C) 2022 veypi <i@veypi.com>
 * 2022-08-15 19:56
 * Distributed under terms of the Apache license.
 -->
<template>
    <img class="v-avator" :src="props.src || defsrc" alt="" :onerror="onerror"
        :style="{ height: props.size, width: props.size }">
</template>

<script lang="ts" setup>
import { ref } from 'vue';

const defsrc = ref('/media/icon/404.svg')

const props = withDefaults(defineProps<{
    src?: string
    size?: string
}>(), {
    size: '2rem'
})
const onerror = (e: any) => {
    let img = e.srcElement;
    img.src = defsrc.value;
    img.onerror = null;
}
</script>

<style scoped>
.v-avator {
    display: flex;
    border-radius: 50%;
    align-items: center;
    justify-content: center;
    overflow: hidden;
}
</style>

