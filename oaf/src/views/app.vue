<template>
  <div>
    {{ uuid }}
  </div>
</template>

<script lang="ts" setup>
import {useRoute, useRouter} from "vue-router";
import {computed, onMounted} from "vue";
import api from "../api";

let route = useRoute()
let router = useRouter()
let uuid = computed(() => route.params.uuid)
onMounted(() => {
  if (uuid.value === '') {
    router.push({name: '404', params: {path: route.path}})
    return
  }
  api.app.get(uuid.value as string).Start(e => {
    console.log(e)
  })
})
</script>

<style scoped>

</style>
