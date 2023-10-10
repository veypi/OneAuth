 <!--
 * AppCfg.vue
 * Copyright (C) 2023 veypi <i@veypi.com>
 * 2023-10-10 16:08
 * Distributed under terms of the MIT license.
 -->
<template>
  <div>
    <div>
      <CRUD :keys="keys" :data="data">
        <template #a="">_____</template>
      </CRUD>
    </div>
  </div>
</template>

<script lang="ts" setup>
import CRUD from 'src/components/crud.vue'
import { ref } from 'vue';
const keys = ref<any>([
  {
    name: 'name',
    label: '阿萨德',
  }, { name: 'key' }, { name: 'status' }
])

const data = ref([
  { name: 'asd', key: 'akk' },
  { name: '1sd', key: '1kk' },
  { name: '2sd', key: '2kk' },
  { name: '3sd', key: '3kk' },
  { name: '4sd', key: '4kk' },
])
</script>

<style scoped></style>

