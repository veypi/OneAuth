<template>
  <div>
  </div>
</template>

<script lang="ts" setup>
import {ref} from "vue";
import api from "../api";

let apps = ref([])

function getApps() {
  api.app.list().Start(e => {
    apps.value = e
  })
}

getApps()

</script>

<style scoped>

</style>
