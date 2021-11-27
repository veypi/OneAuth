<template>
  <div class="core rounded-2xl p-3">
    <div class="grid gap-4 grid-cols-5">
      <div class="col-span-2">
        <n-avatar style="--color: none;" @click="Go" round :size="80" :src="core.Icon">
        </n-avatar>
      </div>
      <div class="col-span-3 grid grid-cols-1 items-center text-left">
        <div class="h-10 flex items-center text-2xl italic font-bold">{{ core.Name }}</div>
        <span class="truncate">{{ core.Des }}</span>
      </div>
    </div>
  </div>
</template>
<script setup lang='ts'>
import {withDefaults} from 'vue'
import {useRouter} from 'vue-router'
import {useMessage, useLoadingBar} from 'naive-ui'
import api from '@/api'
import {useStore} from '@/store'
import {modelsApp} from '@/models'

let router = useRouter()
let store = useStore()
let msg = useMessage()
let bar = useLoadingBar()

let props = withDefaults(defineProps<{
  core?: modelsApp
}>(), {
  // @ts-ignore
  core: {},
})

function Go() {
  switch (props.core.UserStatus) {
    case 'ok':
      router.push({name: 'app.main', params: {uuid: props.core.UUID}})
      return
    case 'apply':
      msg.info('请等待管理员审批进入')
      return
    case 'deny':
      msg.warning('进入申请未通过')
      return
    case 'disabled':
      msg.warning('已被禁止使用')
      return
  }
  bar.start()
  api.app.user(props.core.UUID).add(store.state.user.id).Start(e => {
    bar.finish()
    if (e.Status === 'ok') {
      router.push({name: 'app.main', params: {uuid: props.core.UUID}})
      return
    }
    props.core.UserStatus = e.Status
    msg.info('已发起加入申请')
  }, (e) => {
    msg.warning('加入失败: ' + e)
    bar.error()
  })
  return
}
</script>
<style scoped>
.core {
  width: 256px;
  background: rgba(146, 145, 145, 0.1);
}
</style>
