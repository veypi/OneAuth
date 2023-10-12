 <!--
 * docItem.vue
 * Copyright (C) 2023 veypi <i@veypi.com>
 * 2023-10-07 22:16
 * Distributed under terms of the MIT license.
 -->
<template>
  <div class="w-full h-full">
    <h1 class="page-h1">文档中心</h1>
    <q-inner-loading :showing="!visible" label="Please wait..." label-class="text-teal" label-style="font-size: 1.1em" />
    <div class="w-full px-8">
      <Editor v-if='doc' eid='doc' preview :content="doc"></Editor>
    </div>
  </div>
</template>

<script lang="ts" setup>
import msg from '@veypi/msg';
import Editor from 'src/components/editor'
import oafs from 'src/libs/oafs';
import { computed, watch, onMounted, ref } from 'vue';
import { useRoute, useRouter } from 'vue-router';
let doc = ref('')


const visible = ref(false)
let route = useRoute()
let router = useRouter()

let url = computed(() => {
  if (route.params.typ === 'public') {
    return '/doc/' + route.params.url
  }
  return route.params.url
})

watch(url, u => {
  render(u as string)
})


const render = (url: string) => {
  console.log(url)
  if (!url) {
    return
  }
  oafs.get(url).then((value) => {
    doc.value = value
    visible.value = true
  }).catch(e => {
    console.warn(e)
    msg.Warn('访问文档地址不存在')
    router.back()
  });
}
onMounted(() => {
  render(url.value as string)
})
</script>

<style scoped></style>


