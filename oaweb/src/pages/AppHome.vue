 <!--
 * AppHome.vue
 * Copyright (C) 2023 veypi <i@veypi.com>
 * 2023-10-04 15:34
 * Distributed under terms of the MIT license.
 -->
<template>
  <div>
    <q-page-sticky position="top-right" :offset="[27, 27]">
      <q-btn @click="sync_editor" :style="{
        color: edit_mode ? 'red' :
          ''
      }" round icon="save_as" class="" />
    </q-page-sticky>
    <Editor v-if="app.id" :eid="app.id + '.des'" :content="app.des" @updated="save"></Editor>
  </div>
</template>

<script lang="ts" setup>
import { inject, onMounted, ref, Ref } from 'vue';
import { modelsApp } from 'src/models';
import api from 'src/boot/api';
import msg from '@veypi/msg';
import Editor from 'src/components/editor'




let edit_mode = ref(false)

let app = inject('app') as Ref<modelsApp>

const save = (des: string) => {
  api.app.update(app.value.id, { des: des }).then(e => {
    edit_mode.value = false
    app.value.des = des as string
  }).catch(e => {
    msg.Warn("更新失败: " + e)
  })
}


const sync_editor = () => {
  if (edit_mode.value) {
    // console.log(editor.getHtml())
    // let des = editor.getValue()
    // return
  }
  edit_mode.value = true
}



onMounted(() => {

})
</script>

<style scoped></style>

