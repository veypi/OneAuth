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

    <template v-for="link in items" :key="link.title">

      <q-item class="flex items-center" v-ripple clickable tag="a" :to="link.to">
        <q-icon size="1.5rem" class="mr-2" :name="link.icon" />
        <q-item-section>
          <q-item-label>{{ link.title }}</q-item-label>
        </q-item-section>
      </q-item>
    </template>
  </q-list>
</template>

<script lang="ts" setup>
import { computed } from 'vue';
import menus from './menus'

const items = computed(() => menus.items.value)


</script>

<style scoped></style>

