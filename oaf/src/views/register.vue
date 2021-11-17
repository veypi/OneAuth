<template>
  <div class="flex items-center justify-center">
    <div
      :style="{background:Theme.me.lightBox, 'box-shadow': Theme.me.lightBoxShadow}"
      class="px-10 pb-9 pt-28 rounded-xl w-96">
      <n-form label-width="70px" label-align="left" :model="data" ref="form_ref" label-placement="left" :rules="rules">
        <n-form-item required label="用户名" path="username">
          <n-input @keydown.enter="divs[1].focus()" :ref="el => {if (el)divs[0]=el}"
                   v-model:value="data.username"></n-input>
        </n-form-item>
        <n-form-item required label="密码" path="password">
          <n-input @keydown.enter="divs[2].focus()" :ref="el => {if (el) divs[1]=el}" v-model:value="data.password"
                   type="password"></n-input>
        </n-form-item>
        <n-form-item required label="重复密码" path="pass">
          <n-input @keydown.enter="register" :ref="el => {if (el) divs[2]=el}" v-model:value="data.pass"
                   type="password"></n-input>
        </n-form-item>
        <div class="flex justify-around mt-4">
          <n-button @click="register">注册</n-button>
        </div>
      </n-form>
    </div>
  </div>
</template>
<script lang="ts" setup>
import {onMounted, ref} from "vue";
import {Theme} from "@/theme";
import {useMessage} from 'naive-ui'
import api from "@/api"
import {useRouter} from "vue-router";

let msg = useMessage()
const router = useRouter()

const divs = ref([])
let form_ref = ref(null)
let data = ref({
  username: '',
  password: '',
  pass: ''
})
let rules = {
  username: [
    {
      required: true,
      validator(r: any, v: any) {
        return (v && v.length >= 3 && v.length <= 16) || new Error('长度要求3~16')
      },
      trigger: ['input', 'blur']
    }
  ],
  password: [{
    required: true,
    validator(r: any, v: any) {
      return (v && v.length >= 6 && v.length <= 16) || new Error('长度要求6~16')
    },
    trigger: ['input', 'blur']
  }],
  pass: [
    {
      required: true,
      validator(r: any, v: any) {
        return (v && v === data.value.password) || new Error('密码不正确')
      },
      trigger: ['input', 'blur']
    }
  ]
}

function register() {
  // @ts-ignore
  form_ref.value.validate((e: any) => {
    if (!e) {
      api.user.register(data.value.username, data.value.password).Start((url: string) => {
        msg.success('注册成功')
        router.push({name: 'login'})
      }, e => {
        console.log(e)
        msg.warning('注册失败：' + e)
      })
    }
  })
}

onMounted(() => {
  if (divs.value[0]) {
    // @ts-ignore
    divs.value[0].focus()
  }
})
</script>

<style scoped>
</style>
