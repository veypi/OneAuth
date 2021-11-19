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
              <n-input v-model:value="user.Nickname" @blur="update('Nickname')"></n-input>
            </n-form-item>
            <n-form-item label="头像">
              <n-upload
                action="/api/upload"
                @finish="handleFinish"
                :show-file-list="false"
              >
                <n-avatar size="large" round :src="user.Icon">
                </n-avatar>
              </n-upload>
            </n-form-item>
          </n-form>
        </div>
        <div v-else class="animate__animated animate__faster">
          <n-form label-align="left" label-width="80px" label-placement="left">
            <n-form-item label="Username">
              <n-input disabled v-model:value="user.Username"></n-input>
            </n-form-item>
            <n-form-item label="phone">
              <n-input v-model:value="user.Phone" @blur="update('Phone')"></n-input>
            </n-form-item>
            <n-form-item label="email">
              <n-auto-complete :options="emailOptions" v-model:value="user.Email"
                               @blur="update('Email')"></n-auto-complete>
            </n-form-item>
            <n-form-item label="邮件通知">
              <n-switch>
                <template #checked>启用</template>
                <template #unchecked>关闭</template>
              </n-switch>
            </n-form-item>
            <n-form-item label="短信通知">
              <n-switch>
                <template #checked>启用</template>
                <template #unchecked>关闭</template>
              </n-switch>
            </n-form-item>
          </n-form>
        </div>
      </transition>
    </div>
  </div>
</template>

<script lang="ts" setup>
import {ref, computed} from 'vue'
import {IsDark} from '@/theme'
import {useStore} from '@/store'
import api from '@/api'
import {useMessage} from 'naive-ui'
import {modelsUser} from '@/models'

let msg = useMessage()
let store = useStore()

let ifInfo = ref(true)
let user = ref<modelsUser>({
  Username: store.state.user.local.Username,
  Nickname: store.state.user.local.Nickname,
  Icon: store.state.user.local.Icon,
  Email: store.state.user.local.Email,
  Phone: store.state.user.local.Phone,
} as modelsUser)
let emailOptions = computed(() => {
  return ['@qq.com', '@163.com', '@gmail.com', '@outlook.com', '@icloud.com', '@169.com'].map((suffix) => {
    const prefix = user.value.Email.split('@')[0]
    return {
      label: prefix + suffix,
      value: prefix + suffix,
    }
  })
})

function handleFinish(e: any) {
  if (e.event.target.response) {
    let data = JSON.parse(e.event.target.response)
    if (data.status === 1) {
      user.value.Icon = data.content
      update('icon')
      return
    }
  }
  msg.error('上传失败')
  user.value.Icon = store.state.user.local.Icon
}

function update(key: string) {
  // @ts-ignore
  let v = user.value[key]
  if (v === store.state.user.local[key]) {
    return
  }
  api.user.update(store.state.user.id, {[key]: v}).Start(e => {
    msg.success('更新成功')
    store.state.user.local[key] = v
  }, e => {
    msg.error('更新失败: ' + e)
  })
}
</script>

<style scoped>
</style>
