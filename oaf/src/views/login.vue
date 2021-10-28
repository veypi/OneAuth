<template>
  <div class="flex items-center justify-center">
    <div class="p-3" style="">
      <n-form ref="formRef" label-placement="left">
        <n-form-item required label="username" :validation-status="rules.username[0]" :feedback="rules.username[1]">
          <n-input v-model:value="data.username"></n-input>
        </n-form-item>
        <n-form-item required label="username" :validation-status="rules.username[0]" :feedback="rules.username[1]">
          <n-input v-model:value="data.username"></n-input>
        </n-form-item>
        <n-button @click="login">登录</n-button>
      </n-form>
    </div>
  </div>
</template>

<script lang="ts" setup>
import {computed, ref} from "vue";

let formRef = ref(null)
let data = ref({
  username: null,
  password: null
})
let ruleInline = {
  username: [
    (v: string) => !!v || 'required',
    (v: string) => (v && v.length >= 3 && v.length <= 16) || '长度要求3~16'
  ],
  password: [
    (v: string) => !!v || 'required',
    (v: string) => (v && v.length >= 6 && v.length <= 16) || '长度要求6~16'
  ]
}

function check(rs: [], v: any) {
  for (let r of rs) {
    let res = r(v)
    if (res !== true) {
      return ['error', res]
    }
  }
  return ['', '']
}

let rules = ref({
  username: computed(() => {
    return check(ruleInline.username, data.value.username)
  })
})

function login() {
  formRef.value.validate(e => {
    console.log(e)
  })
}
</script>

<style scoped>

</style>
