<template>
  <div>
    <h1 class="page-h1">首页</h1>
    {{ uuid }}
    {{ app }}
    <div :ref="el => nav[6]=el" class="my-80">123123</div>
    <div :ref="el => nav[k-1]=el" class="mb-64 text-center" v-for="(k) in [1,2,3,4,5,6]" :key="k">
      {{ k }}
    </div>
  </div>
</template>

<script lang="ts" setup>
import {inject, ref} from "vue";

let uuid = inject('uuid')
let app = inject('app')
let nav = ref([])

defineExpose({
  nav
})
</script>

<style scoped>

</style>
