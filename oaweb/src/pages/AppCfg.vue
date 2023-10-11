 <!--
 * AppCfg.vue
 * Copyright (C) 2023 veypi <i@veypi.com>
 * 2023-10-10 16:08
 * Distributed under terms of the MIT license.
 -->
<template>
  <div>
    <div class="flex justify-center pt-10">
      <CRUD ref="table" v-if="app.id" :keys="keys" :data="[app]" kalign="left" valign="left" editable
        :vstyle="{ 'width': '50vw' }" @update="newApp = $event[0]" :kstyle="{ 'width': '10rem' }">
        <template #k_icon="{ value, set }">
          <div class="w-full flex justify-center">
            <uploader class="" @success="set" dir="app_icon">
              <q-avatar>
                <img :src="value">
              </q-avatar>
            </uploader>
          </div>
        </template>
        <template #k_key>
          <div class="w-full div-center">
            <q-btn color='primary'>获取秘钥</q-btn>
          </div>
        </template>
      </CRUD>
    </div>
    <div v-if="newApp" class="flex justify-center gap-8 mt-6">
      <q-btn color="brown-5" label="回退" @click="table.reload" />
      <q-btn color="deep-orange" glossy label="保存" @click="save" />
    </div>
  </div>
</template>

<script lang="ts" setup>
import CRUD from 'src/components/crud.vue'
import { ArgType, modelsApp } from 'src/models';
import { inject, onMounted, Ref, ref } from 'vue';
import uploader from 'src/components/uploader';
import api from 'src/boot/api';
import msg from '@veypi/msg';
const keys = ref<any>([
  {
    name: 'name',
    label: '应用名',
  },
  { name: 'id', label: 'uuid', editable: false },
  { name: 'key', label: '秘钥Key' },
  { name: 'icon', label: 'logo' },
  {
    name: 'join_method', label: '用户注册', typ: ArgType.Radio,
    options: [{ key: 0, name: '允许' }, { key: 1, name: '禁止' },
    { key: 2, name: '申请' }]
  },
  { name: 'host', label: '项目首页' },
  { name: 'redirect', label: '跳转地址' },
])
let app = inject('app') as Ref<modelsApp>

const newApp = ref(null)
const table = ref()

const save = () => {
  api.app.update(app.value.id, newApp.value).then(e => {
    msg.Info('更新成功')
    Object.assign(app.value, newApp.value)
    newApp.value = null
  }).catch(e => {
    msg.Warn('更新失败 ' + e)
  })
}


onMounted(() => {
})

</script>

<style scoped></style>

