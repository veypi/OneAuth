<template>
  <div class="flex justify-center">
    <div style="line-height: 48px" class="inline-block mt-16 grid grid-cols-5 w-1/3 text-center gap-4">
      <div>应用名</div>
      <div class="col-span-4">
        <n-input v-model:value="data.name" @blur="update('name')"></n-input>
      </div>
      <div>logo</div>
      <div class="col-span-4">
        <n-upload
          action="/api/upload"
          @finish="handleFinish"
          :show-file-list="false"
        >
          <n-avatar size="large" round :src="data.icon">
          </n-avatar>
        </n-upload>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import {inject, watch, ref, onMounted} from "vue";
import api from "@/api";
import {useMessage} from "naive-ui";

let msg = useMessage()
let app: any = inject('app')
let data = ref({
  name: '',
  icon: ''
})

function handleFinish(e: any) {
  if (e.event.target.response) {
    let d = JSON.parse(e.event.target.response)
    if (d.status === 1) {
      data.value.icon = d.content
      update('icon')
      return
    }
  }
  msg.error('上传失败')
  data.value.icon = app.value.icon
}

function update(key: string) {
  // @ts-ignore
  let v = data.value[key]
  if (v === app.value[key]) {
    return
  }
  api.app.update(app.value.uuid, {[key]: v}).Start(e => {
    msg.success('更新成功')
    app.value[key] = v
  }, e => {
    data.value[key] = app.value[key]
  })
}


function sync() {
  data.value.name = app.value.name
  data.value.icon = app.value.icon
}
watch(app, sync)
onMounted(sync)

</script>

<style scoped>

</style>
