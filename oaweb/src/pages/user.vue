 <!--
 * user.vue
 * Copyright (C) 2023 veypi <i@veypi.com>
 * 2023-10-08 05:31
 * Distributed under terms of the MIT license.
 -->
<template>
  <div>
    <h1 class="page-h1">账号设置</h1>
    <div class="flex justify-center pt-10">
      <crud ref="table" editable :keys="keys" :data="[u.local]" @update="newData = $event[0]">
        <template #k_icon="{ value, set }">
          <div class="w-full flex justify-center">
            <uploader class="" @success="set" dir="user_icon">
              <q-avatar>
                <img :src="value">
              </q-avatar>
            </uploader>
          </div>
        </template>
      </crud>
    </div>

    <div v-if="newData" class="flex justify-center gap-8 mt-6">
      <q-btn color="brown-5" label="回退" @click="table.reload" />
      <q-btn color="deep-orange" glossy label="保存" @click="save" />
    </div>
  </div>
</template>

<script lang="ts" setup>
import msg from '@veypi/msg';
import api from 'src/boot/api';
import crud from 'src/components/crud.vue';
import uploader from 'src/components/uploader';
import { useUserStore } from 'src/stores/user';
import { ref } from 'vue';
let u = useUserStore()
let table = ref()
const keys = ref<any>([
  { name: 'id', label: 'ID', editable: false },
  { name: 'username', label: '用户名' },
  { name: 'nickname', label: '昵称' },
  { name: 'icon', label: 'logo' },
  { name: 'email', label: '邮箱' },
  { name: 'phone', label: '手机号' },
])

const newData = ref()
const save = () => {
  api.user.update(u.id, newData.value).then(e => {
    msg.Info('更新成功')
    Object.assign(u.local, newData.value)
    newData.value = null
  }).catch(e => {
    msg.Warn('更新失败 ' + e)
  })
}
</script>

<style scoped></style>

