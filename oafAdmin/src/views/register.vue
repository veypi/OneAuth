<template>
  <div class="flex items-center justify-center h-full w-full">
    <div class="box px-10 pb-9 pt-16 rounded-xl w-96 relative" style="">
      <one-icon class="absolute text-5xl top-4 left-4" @click="$router.push({ name: 'login' })">
        back
      </one-icon>
      <Myinput
        :type="ArgType.Text"
        label-width="8rem"
        v-model="data.username"
        :options="{ min: 5, max: 16 }"
      >
        <template #label>{{ $t('a.username') }}</template>
      </Myinput>
      <Myinput
        :type="ArgType.Password"
        label-width="8rem"
        v-model="data.password"
        :options="{ min: 6, max: 16 }"
      >
        <template #label>{{ $t('a.password') }}</template>
      </Myinput>
      <Myinput
        :type="ArgType.Password"
        label-width="8rem"
        v-model="data.pass"
        :validator="validatefc.pass"
      >
        <template #label>{{ $t('a.repeat_pass') }}</template>
      </Myinput>
      <div class="flex justify-around mt-4">
        <div
          class="div-btn px-4 py-1 rounded"
          style="background: var(--color-primary)"
          @click="register"
        >
          {{ $t('a.register') }}
        </div>
      </div>
    </div>
  </div>
</template>
<script lang="ts" setup>
import { onMounted, ref } from 'vue'
import api from '@/api'
import { useRouter } from 'vue-router'
import msg from '@/msg'
import Myinput from '@/components/myinput'
import { ArgType } from '@/models'
import validator from 'validator'
import { useI18n } from 'vue-i18n'

const { t } = useI18n()

const router = useRouter()

let data = ref({
  username: '',
  password: '',
  pass: '',
})

const validatefc = {
  pass(u: string) {
    return validator.equals(u, data.value.password)
  },
}
function register() {
  if (data.value.username && data.value.password && data.value.password === data.value.pass) {
    // @ts-ignore
    api.user.register(data.value.username, data.value.password).Start(
      (url: string) => {
        msg.Info(t('a.register') + t('a.success'))
        router.push({ name: 'login' })
      },
      (e) => {
        console.log(e)
        msg.Warn('注册失败：' + e.headers.error)
      },
    )
  } else {
    msg.Warn(t('msg.e1'))
  }
}

onMounted(() => {})
</script>

<style scoped>
.box {
  background: linear-gradient(145deg, var(--base-bg-3), var(--base-bg-2));
  box-shadow: 20px 20px 60px var(--base-bg-3), -20px -20px 60px var(--base-bg-2);
}
</style>
