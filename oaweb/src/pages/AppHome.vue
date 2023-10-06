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
    <div v-if="edit_mode" id="vditor"></div>
    <div v-else>
      {{ app }}
    </div>
  </div>
</template>

<script lang="ts" setup>
import { inject, onMounted, ref } from 'vue';
import { modelsApp } from 'src/models';
import Vditor from 'vditor';
import api from 'src/boot/api';

let edit_mode = ref(false)

const vditor = ref<Vditor | null>(null);



let app = inject('app') as modelsApp

const sync_editor = () => {
  if (edit_mode.value) {
    api.app.update(app.id, { des: "" }).then(e => {
      edit_mode.value = false
      console.log(e)
    })
    return
  }
  edit_mode.value = true
  setTimeout(() => {
    vditor.value = new Vditor('vditor', {
      toolbarConfig: {
        hide: true
      },
      after: () => {
        // vditor.value is a instance of Vditor now and thus can be safely used here
        vditor.value!.setValue('Vue Composition API + Vditor + TypeScript Minimal Example');
      },
    });
  }, 0)
}

onMounted(() => {
})
</script>

<style scoped></style>

