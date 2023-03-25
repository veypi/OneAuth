<template>
  <div class="pt-10">
    <div class="flex justify-center">
      <div
        class="relative rounded-xl text-lg text-black"
        :style="{ background: IsDark ? '#555' : '#d5d5d5' }"
      >
        <div
          @click="ifInfo = true"
          class="inline-block px-5 rounded-xl"
          :style="{ background: ifInfo ? '#fc0005' : '' }"
        >
          个人信息
        </div>
        <div
          @click="ifInfo = false"
          class="inline-block px-5 rounded-xl"
          :style="{ background: ifInfo ? '' : '#fc0005' }"
        >
          账户管理
        </div>
      </div>
    </div>

    <div class="inline-block flex justify-center mt-10">
      <transition
        mode="out-in"
        enter-active-class="animate__fadeInLeft"
        leave-active-class="animate__fadeOutRight"
      >
        <div v-if="ifInfo" class="animate__animated animate__faster">
          <n-form label-placement="left" label-width="80px" label-align="left">
            <n-form-item label="昵称">
              <n-input v-model:value="user.Nickname" @blur="update('Nickname')"></n-input>
            </n-form-item>
            <n-form-item label="头像">
              <uploader :url="user.ID + '.ico'" @success="handleFinish">
                <n-avatar size="large" round :src="user.Icon"></n-avatar>
              </uploader>
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
              <n-auto-complete
                :options="emailOptions"
                v-model:value="user.Email"
                @blur="update('Email')"
              ></n-auto-complete>
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
import { ref, computed } from 'vue'
import api from '@/api'
import { modelsUser } from '@/models'
import Uploader from '@/components/uploader'
import { useUserStore } from '@/store/user'
import msg from '@/msg'

let store = useUserStore()

let ifInfo = ref(true)
let user = ref<modelsUser>({
  ID: store.id,
  Username: store.local.Username,
  Nickname: store.local.Nickname,
  Icon: store.local.Icon,
  Email: store.local.Email,
  Phone: store.local.Phone,
} as modelsUser)
let emailOptions = computed(() => {
  return ['@qq.com', '@163.com', '@gmail.com', '@outlook.com', '@icloud.com', '@169.com'].map(
    (suffix) => {
      const prefix = user.value.Email.split('@')[0]
      return {
        label: prefix + suffix,
        value: prefix + suffix,
      }
    },
  )
})

function handleFinish(e: string) {
  console.log(e)
  user.value.Icon = e
  update('Icon')
  return
}

function update(key: string) {
  // @ts-ignore
  let v = user.value[key]
  // @ts-ignore
  if (v === store.local[key]) {
    return
  }
  api.user.update(store.id, { [key]: v }).Start(
    (e) => {
      msg.Info('更新成功')
      // @ts-ignore
      store.local[key] = v
    },
    (e) => {
      msg.Warn('更新失败: ' + e)
    },
  )
}
</script>

<style scoped></style>
