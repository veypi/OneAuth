 <!--
 * cfg.vue
 * Copyright (C) 2024 veypi <i@veypi.com>
 * 2024-06-12 18:38
 * Distributed under terms of the MIT license.
 -->
<template>
  <div>
    <div class="flex justify-center pt-10">
      <Crud ref="table" :keys="keys" :data="core.id ? [core] : []" kalign="left" valign="left" editable
        :vstyle="{ 'width': '50vw' }" @update="newApp = $event[0]" :kstyle="{ 'width': '10rem' }">
        <template #k_icon="{ value, set }">
          <div class="w-full flex justify-center">
            <Uploader class="" @success="set" dir="app_icon">
              <img class="w-12 h-12 rounded-full" :src="value">
            </Uploader>
          </div>
        </template>
        <template #k_key>
          <div class="w-full div-center">
            <div class='vbtn' @click="getkey">获取秘钥</div>
            <span class="mx-2 select-all" v-if="tmpkey">{{ tmpkey }}</span>
          </div>
        </template>
        <template #k_redirect_append>
          <div class="mx-8 vbtn" @click="$router.push('/login?uuid=' +
            core.id)">Go</div>
        </template>
      </Crud>
    </div>
    <div v-if="newApp" class="flex justify-center gap-8 mt-6">
      <div class="vbtn" style="background: var(--color-primary);" @click="table.reload">回退</div>
      <div class="vbtn" style="background: var(--color-warning);" @click="save">保存</div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import msg from '@veypi/msg';

const keys = ref<any>([
  {
    name: 'name',
    label: '应用名',
  },
  { name: 'id', label: 'uuid', editable: false },
  { name: 'key', label: '秘钥Key' },
  { name: 'icon', label: 'logo' },
  {
    name: 'join_method', label: '用户注册', typ: ArgType.Radio,
    options: [{ key: 0, name: '允许' }, { key: 1, name: '禁止' },
    { key: 2, name: '申请' }]
  },
  { name: 'host', label: '项目首页' },
  { name: 'redirect', label: '跳转地址' },
])
let props = withDefaults(defineProps<{
  core: modelsApp,
}>(),
  {}
)

const newApp = ref(null)
const table = ref()

const save = () => {
  api.app.update(props.core.id, newApp.value).then(e => {
    msg.Info('更新成功')
    Object.assign(props.core, newApp.value)
    newApp.value = null
  }).catch(e => {
    msg.Warn('更新失败 ' + e)
  })
}

const tmpkey = ref('')
const getkey = () => {
  msg.Prompt('请输入应用名确认', '').then((s) => {
    if (s === props.core.name) {
      api.app.getKey(props.core.id).then(e => {
        tmpkey.value = e
        copyToClipboard(e).then(e => {
          msg.Info('已复制到剪贴板')
        }).catch(e => {
          tmpkey.value = e
        })
      }).catch(e => {
        msg.Warn('获取失败 ' + e)
      })
    } else {
      msg.Info('输入错误')
    }
  })
}


onMounted(() => {
})

</script>

<style scoped></style>

