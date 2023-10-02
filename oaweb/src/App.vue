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
  if (iconName.startsWith('app:') === true) {
    // we strip the "app:" part
    const name = iconName.substring(4)

    return {
      cls: 'my-app-icon ' + name
    }
  }
}


onBeforeMount(() => {
  useUserStore().fetchUserData()
})

</script>
