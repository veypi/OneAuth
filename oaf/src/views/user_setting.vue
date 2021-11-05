<template>
  <div class="pt-10">
    <div class="flex justify-center">
      <div class="relative rounded-xl text-lg text-black" :style="{background: IsDark?'#555': '#d5d5d5'}">
        <div @click="ifInfo=true" class="inline-block px-5 rounded-xl" :style="{background: ifInfo ? '#fc0005': ''}">
          个人信息
        </div>
        <div @click="ifInfo=false" class="inline-block px-5 rounded-xl" :style="{background: ifInfo ? '': '#fc0005'}">
          账户管理
        </div>
      </div>
    </div>

    <div class="inline-block flex justify-center mt-10">
      <transition mode="out-in" enter-active-class="animate__fadeInLeft" leave-active-class="animate__fadeOutRight">
        <div v-if="ifInfo" class="animate__animated animate__faster">
          <n-form label-placement="left" label-width="80px" label-align="left">
            <n-form-item label="昵称">
              <n-input v-model:value="user.nickname" @blur="update('nickname')"></n-input>
            </n-form-item>
            <n-form-item label="头像">
              <n-upload
                action=""
                :headers="{'': ''}"
                :data="{}"
              >
                <n-avatar size="large" round :src="user.icon">
                </n-avatar>
              </n-upload>
            </n-form-item>
          </n-form>
        </div>
        <div v-else class="animate__animated animate__faster">
          <n-form label-align="left" label-width="80px" label-placement="left">
            <n-form-item label="username">
              <n-input disabled v-model:value="user.username"></n-input>
            </n-form-item>
            <n-form-item label="phone">
              <n-input v-model:value="user.phone" @blur="update('phone')"></n-input>
            </n-form-item>
            <n-form-item label="email">
              <n-auto-complete :options="emailOptions" v-model:value="user.email"
                               @blur="update('email')"></n-auto-complete>
            </n-form-item>
          </n-form>
        </div>
      </transition>
    </div>
  </div>
</template>

<script lang="ts" setup>
import {ref, computed} from "vue";
import {IsDark} from "../theme";
import {useStore} from "../store";
import api from "../api";
import {useMessage} from "naive-ui";

let msg = useMessage()
let store = useStore()

let ifInfo = ref(true)
let user = ref({
  username: store.state.user.username,
  nickname: store.state.user.nickname,
  icon: store.state.user.icon,
  email: store.state.user.email,
  phone: store.state.user.phone,
})
let emailOptions = computed(() => {
  return ['@gmail.com', '@163.com', '@qq.com'].map((suffix) => {
    const prefix = user.value.email.split('@')[0]
    return {
      label: prefix + suffix,
      value: prefix + suffix
    }
  })
})

function update(key: string) {
  // @ts-ignore
  let v = user.value[key]
  if (v === store.state.user[key]) {
    return
  }
  api.user.update(store.state.user.id, {[key]: v}).Start(e => {
    msg.success('更新成功')
    store.state.user[key] = v
  }, e => {
    msg.error('更新失败: ' + e.err)
  })
}
</script>

<style scoped>
</style>
