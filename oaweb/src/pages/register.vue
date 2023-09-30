<template>
  <div class="flex items-center justify-center">
    <div class="px-10 pb-9 pt-28 rounded-xl w-96">
      <q-form @submit="register" autofocus>
        <q-input v-model="data.username" label="用户名" hint="username" lazy-rules :rules="rules.username"></q-input>
        <q-input label="密码" v-model="data.password" type="password" lazy-rules :rules="rules.password"></q-input>
        <q-input label="密码" v-model="data.pass" type="password" lazy-rules :rules="rules.pass"></q-input>
        <div class="flex justify-around mt-4">
          <q-btn label="注册" type="submit" color="primary" />
        </div>
      </q-form>
    </div>
  </div>
</template>
<script lang="ts" setup>
import { onMounted, ref } from "vue";
import { useRouter } from "vue-router";
import msg from '@veypi/msg'
import api from "src/boot/api";

const router = useRouter()

let data = ref({
  username: '',
  password: '',
  pass: ''
})
let rules = {
  username: [
    (v: string) => v && v.length >= 3 && v.length <= 16 || '长度要求3~16'
  ],
  password: [
    (v: string) => v && v.length >= 6 && v.length <= 16 || '长度要求6~16'
  ],
  pass: [
    (v: string) => v && v === data.value.password || '密码不正确'
  ]
}

function register() {
  api.user.register(data.value.username, data.value.password).then(u => {
    console.log(u)
    msg.Info('注册成功')
    router.push({ name: 'login' })
  }).catch(e => {
    console.log(e)
    msg.Warn('注册失败：' + e.data)
  })
}

onMounted(() => {
})
</script>

<style scoped></style>
