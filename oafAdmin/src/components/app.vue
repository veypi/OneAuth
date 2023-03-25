<template>
  <div class="core rounded-2xl p-3">
    <div class="grid gap-4 grid-cols-5">
      <div class="col-span-2">
        <vavator style="--color: none" @click="Go" round size="4rem" :src="core.icon"></vavator>
      </div>
      <div class="col-span-3 grid grid-cols-1 items-center text-left">
        <div class="h-10 flex items-center text-2xl italic font-bold">
          {{ core.name }}
        </div>
        <span class="truncate">{{ core.des }}</span>
      </div>
    </div>
  </div>
</template>
<script setup lang="ts">
import { withDefaults } from 'vue'
import { useRouter } from 'vue-router'
import api from '@/api'
import { modelsApp } from '@/models'
import vavator from '@/components/vavator'
import { useUserStore } from '@/store/user'
import msg from '@/msg'

let router = useRouter()
let store = useUserStore()

let props = withDefaults(
  defineProps<{
    core: modelsApp
  }>(),
  {},
)

function Go() {
  switch (props.core.status) {
    case 'ok':
      router.push({ name: 'app.main', params: { uuid: props.core.UUID } })
      return
    case 'apply':
      msg.Info('请等待管理员审批进入')
      return
    case 'deny':
      msg.Warn('进入申请未通过')
      return
    case 'disabled':
      msg.Warn('已被禁止使用')
      return
  }
  api.app
    .user(props.core.id)
    .add(store.id)
    .Start(
      (e) => {
        if (e.Status === 'ok') {
          router.push({ name: 'app.main', params: { uuid: props.core.id } })
          return
        }
        props.core.status = e.Status
        msg.Info('已发起加入申请')
      },
      (e) => {
        msg.Warn('加入失败: ' + e)
      },
    )
  return
}
</script>
<style scoped>
.core {
  width: 256px;
  background: rgba(146, 145, 145, 0.1);
}
</style>
