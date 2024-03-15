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
        <template #k_pass>
          <q-btn @click="prompt = true" size='sm' color='secondary'>修改</q-btn>
        </template>
      </crud>
    </div>
    <q-dialog v-model="prompt" persistent>
      <q-card style="min-width: 350px">
        <q-card-section>
          <div class="text-h6">请输入新密码</div>
        </q-card-section>

        <q-card-section class="q-pt-none">
          <q-input type='password' dense v-model="pass[0]" autofocus @keyup.enter="prompt = false">
          </q-input>
          <q-input type='password' dense v-model="pass[1]" autofocus @keyup.enter="prompt = false">
          </q-input>
        </q-card-section>

        <q-card-actions align="right" class="text-primary">
          <q-btn flat label="取消" v-close-popup />
          <q-btn flat label="确定" @click="reset" v-close-popup />
        </q-card-actions>
      </q-card>
    </q-dialog>

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
let prompt = ref(false)
let u = useUserStore()
let table = ref()
const keys = ref<any>([
  { name: 'id', label: 'ID', editable: false },
  { name: 'username', label: '用户名' },
  { name: 'pass', label: '密码' },
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
let pass = ref(['', ''])
const reset = () => {
  if (pass.value[0] == '' || pass.value[1] == '') {
    msg.Warn('密码不能为空')
    return
  }
  if (pass.value[0] != pass.value[1]) {
    msg.Warn('两次密码不一致')
    return
  }
  api.user.reset(u.id, pass.value[0]).then((e) => {
    msg.Info('密码重置成功')
  }).catch(e => {
    msg.Warn('密码重置失败 ' + e)
  })
}
</script>

<style scoped></style>

