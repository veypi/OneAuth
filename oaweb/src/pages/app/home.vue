 <!--
 * AppHome.vue
 * Copyright (C) 2023 veypi <i@veypi.com>
 * 2023-10-04 15:34
 * Distributed under terms of the MIT license.
 -->
<template>
  <div class="-mt-16">
    <q-page-sticky position="top-right" style="z-index: 20" :offset="[27, 27]">
      <q-btn v-if="preview_mode" @click="preview_mode = false" round icon="save_as" class="" />
    </q-page-sticky>
    <Editor style="" v-if="app.id" :eid="app.id + '.des'" v-model="preview_mode" :content="content" @save="save"></Editor>
  </div>
</template>

<script lang="ts" setup>
import { computed, inject, onMounted, ref, Ref, watch } from 'vue';
import { modelsApp } from 'src/models';
import api from 'src/boot/api';
import msg from '@veypi/msg';
import Editor from 'src/components/editor'
import { oafs } from '@veypi/oaer'




let preview_mode = ref(true)

let app = inject('app') as Ref<modelsApp>
let content = ref()

watch(computed(() => app.value.id), () => {
  sync()
})

const sync = () => {
  if (app.value.des) {
    oafs.get(app.value.des).then(e => content.value = e)
  }
}

const save = (des: string) => {
  let a = new File([des], app.value.name + '.md');
  oafs.upload([a], app.value.id).then(url => {
    api.app.update(app.value.id, { des: url[0] }).then(e => {
      preview_mode.value = true
      app.value.des = url[0]
    }).catch(e => {
      msg.Warn("更新失败: " + e)
    })
  }).catch(e => {
    msg.Warn("更新失败: " + e)
  })
}


onMounted(() => {
  sync()
})
</script>

<style scoped></style>

