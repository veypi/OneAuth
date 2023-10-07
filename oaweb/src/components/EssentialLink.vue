<template>
  <q-item class="flex items-center" v-ripple clickable tag="a" :href="link" :to="to">
    <!-- <q-item-section v-if="icon" avatar> -->
    <!-- </q-item-section> -->
    <q-icon size="1.5rem" class="mr-2" :name="icon" />

    <q-item-section>
      <q-item-label>{{ title }}</q-item-label>
      <!-- <q-item-label caption>{{ caption }}</q-item-label> -->
    </q-item-section>
  </q-item>
</template>

<script setup lang="ts">

import { MenuLink } from 'src/models'

withDefaults(defineProps<MenuLink>(), {
  caption: '',
  icon: '',
});
</script>
