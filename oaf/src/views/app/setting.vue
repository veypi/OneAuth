<template>
  <div class="flex justify-center">
    <div style="line-height: 48px" class="inline-block mt-16 grid grid-cols-5 w-1/3 text-center gap-4">
      <div>应用名</div>
      <div class="col-span-4">
        <n-input v-model:value="data.Name" @blur="update('Name')"></n-input>
      </div>
      <div>UUID</div>
      <div class="col-span-4 select-all">
        {{ data.UUID }}
      </div>
      <div>Key</div>
      <div class="col-span-4">
        <n-popconfirm
          @positive-click="getKey"
        >
          <template #trigger>
          <n-button>获取</n-button>
          </template>
          获取key将导致之前的key失效 是否获取?
        </n-popconfirm>
      </div>
      <div>logo</div>
      <div class="col-span-4">
        <uploader
          :url="uuid + '.ico'"
          @success="handleFinish"
        >
          <n-avatar size="large" round :src="data.Icon">
          </n-avatar>
        </uploader>
      </div>
      <div>自主注册</div>
      <div class='col-span-4'>
        <n-switch @update:value="update('EnableRegister', $event)" v-model:value='data.EnableRegister'>
          <template #checked> 允许</template>
          <template #unchecked> 禁止</template>
        </n-switch>
      </div>
      <div>应用简介</div>
      <div class='col-span-4 text-left'>
        <n-input :autosize="{minRows: 3}" type="textarea" @blur="update('Des')" maxlength="256"
                 v-model:value="data.Des"></n-input>
      </div>
      <div>项目首页</div>
      <div class="col-span-4 text-left">
        <n-input v-model:value="data.Host" @blur="update('Host')"></n-input>
      </div>
      <div>跳转地址</div>
      <div class="col-span-3 text-left">
        <n-input v-model:value="data.UserRefreshUrl" @blur="update('UserRefreshUrl')"></n-input>
      </div>
      <div class="col-span-1">
        <span class="text-blue-500" @click="util.goto('/login?uuid='+uuid)">GO</span>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import {inject, watch, ref, onMounted} from 'vue'
import api from '@/api'
import {useDialog, useMessage} from 'naive-ui'
import util from '@/libs/util'
import {modelsApp} from '@/models'
import Uploader from '@/components/uploader'

let msg = useMessage()
let dialog = useDialog()
let app = inject<{ value: modelsApp }>('app')
let uuid = inject('uuid')
let data = ref<modelsApp>({} as modelsApp)

function handleFinish(e: string) {
  data.value.Icon = e
  console.log(e)
  update('Icon')
}

function update(key: string, v?: any) {
  // @ts-ignore
  if (v === undefined) {
    v = data.value[key]
  }
  api.app.update(app.value.UUID, {[key]: v}).Start(e => {
    msg.success('更新成功')
    app.value[key] = v
  }, e => {
    data.value[key] = app.value[key]
  })
}

function sync() {
  Object.assign(data.value, app.value)
}

watch(app, sync)
onMounted(sync)

function getKey() {
  api.app.getKey(data.value.UUID).Start(e => {
    dialog.success({
      title: '请保存好秘钥',
      content: e,
    })
  })
}

</script>

<style scoped>

</style>
