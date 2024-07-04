 <!--
 * index.vue
 * Copyright (C) 2024 veypi <i@veypi.com>
 * 2024-06-07 17:46
 * Distributed under terms of the MIT license.
 -->
<template>
  <div>
    <div class="vbtn" v-if="preview_mode" @click="preview_mode = false">
      ss
    </div>
    <Editor style="" v-if="core.id" :eid="core.id + '.des'" v-model="preview_mode" :content="content" @save="save">
    </Editor>
  </div>
</template>

<script lang="ts" setup>
import { oafs } from '@veypi/oaer'

let props = withDefaults(defineProps<{
  core: modelsApp,
}>(),
  {}
)

let preview_mode = ref(true)

let content = ref()

watch(computed(() => props.core.id), () => {
  sync()
})

const sync = () => {
  if (props.core.des) {
    oafs.get(props.core.des).then(e => content.value = e)
  }
}

const save = (des: string) => {
  let a = new File([des], props.core.name + '.md');
  oafs.upload([a], props.core.id).then(url => {
    api.app.update(props.core.id, { des: url[0] }).then(e => {
      preview_mode.value = true
      props.core.des = url[0]
    }).catch(e => {
      // msg.Warn("更新失败: " + e)
    })
  }).catch(e => {
    // msg.Warn("更新失败: " + e)
  })
}


onMounted(() => {
  sync()
})

</script>

<style scoped></style>

