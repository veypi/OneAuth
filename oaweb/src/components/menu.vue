 <!--
 * main.vue
 * Copyright (C) 2023 veypi <i@veypi.com>
 * 2023-10-04 10:56
 * Distributed under terms of the MIT license.
 -->
<template>
  <q-list>
    <!-- <q-item-label header> -->
    <!--   Essential Links -->
    <!-- </q-item-label> -->

    <EssentialLink v-for="link in menu.list" :key="link.title" v-bind="link" />
  </q-list>
</template>

<script lang="ts" setup>
import EssentialLink from 'src/components/EssentialLink.vue';
import { useMenuStore } from 'src/stores/menu';
import { onMounted } from 'vue';
let menu = useMenuStore()


onMounted(() => {
  console.log('loading main menu')
})
</script>

<style scoped></style>

