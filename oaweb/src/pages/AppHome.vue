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
    <Editor v-if="app.id" :eid="app.id + '.des'" :preview="!edit_mode" :content="content" @updated="save"></Editor>
  </div>
</template>

<script lang="ts" setup>
import { computed, inject, onMounted, ref, Ref, watch } from 'vue';
import { modelsApp } from 'src/models';
import api from 'src/boot/api';
import msg from '@veypi/msg';
import Editor from 'src/components/editor'
import oafs from 'src/libs/oafs';




let edit_mode = ref(false)

let app = inject('app') as Ref<modelsApp>
let content = ref()

watch(computed(() => app.value.id), () => {
  if (app.value.des) {
    oafs.get(app.value.des).then(e => content.value = e)
  }
})

const save = (des: string) => {
  let a = new File([des], app.value.name + '.md');
  oafs.upload([a]).then(url => {
    api.app.update(app.value.id, { des: url[0] }).then(e => {
      edit_mode.value = false
      app.value.des = url[0]
    }).catch(e => {
      msg.Warn("更新失败: " + e)
    })
  }).catch(e => {
    msg.Warn("更新失败: " + e)
  })
}


const sync_editor = () => {
  edit_mode.value = !edit_mode.value
}



onMounted(() => {

})
</script>

<style scoped></style>

