<template>
  <router-view />
</template>

<script setup lang="ts">
import { useQuasar } from 'quasar';
import { onBeforeMount } from 'vue';
import { useUserStore } from './stores/user';

const $q = useQuasar()

$q.iconMapFn = (iconName) => {
  // iconName is the content of QIcon "name" prop

  // your custom approach, the following
  // is just an example:
  if (iconName.startsWith('v-') === true) {
    // we strip the "app:" part
    const name = iconName.substring(2)
    console.log(name)

    return {
      icon: 'svguse:#icon-' + name
    }
  }
}


onBeforeMount(() => {
  useUserStore().fetchUserData()
})

</script>

<style>
html,
body,
#q-app {
  @apply font-mono h-full w-full select-none;
}

.page-h1 {
  font-size: 2.5rem;
  line-height: 2.5rem;
  margin-left: 2.5rem;
  margin-top: 1.5rem;
  margin-bottom: 2rem;
}
</style>
