 <!--
 * AppHome.vue
 * Copyright (C) 2023 veypi <i@veypi.com>
 * 2023-10-04 15:34
 * Distributed under terms of the MIT license.
 -->
<template>
  <div>
    {{ app }}
  </div>
</template>

<script lang="ts" setup>
import { inject } from 'vue';
import { modelsApp } from 'src/models';



let app = inject('app') as modelsApp
</script>

<style scoped></style>

