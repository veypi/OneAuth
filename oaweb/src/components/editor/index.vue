 <!--
 * index.vue
 * Copyright (C) 2023 veypi <i@veypi.com>
 * 2023-10-07 18:57
 * Distributed under terms of the MIT license.
 -->
<template>
  <div class="w-full h-full">
    <!-- <div class="absolute bg-red-400 left-0 top-0 w-full h-full"></div> -->
    <div class="w-full h-full" :id="eid"></div>
  </div>
</template>

<script lang="ts" setup>
import Cherry from 'cherry-markdown';
import options from './options'
import { computed, onMounted, ref, watch } from 'vue';
import { oafs } from '@veypi/oaer'

let editor = {} as Cherry;
let emits = defineEmits<{
  (e: 'save', v: string): void
  (e: 'update:modelValue', v: boolean): void
}>()
let props = withDefaults(defineProps<{
  modelValue?: boolean,
  eid?: string,
  content?: string,
  static_dir?: string,
}>(),
  {
    eid: 'v-editor',
    content: '',
    modelValue: true,
  }
)

watch(computed(() => props.modelValue), (e) => {
  set_mode(e)
})
watch(computed(() => props.content), (e) => {
  if (e) {
    editor.setValue(e)
  }
})


const set_mode = (preview: boolean) => {
  editor.switchModel(preview ? 'previewOnly' : 'edit&preview')
}

const fileUpload = (f: File, cb: (url: string, params: any) => void) => {
  /**
   * @param params.name 回填的alt信息
   * @param params.poster 封面图片地址（视频的场景下生效）
   * @param params.isBorder 是否有边框样式（图片场景下生效）
   * @param params.isShadow 是否有阴影样式（图片场景下生效）
   * @param params.isRadius 是否有圆角样式（图片场景下生效）
   * @param params.width 设置宽度，可以是像素、也可以是百分比（图片、视频场景下生效）
   * @param params.height 设置高度，可以是像素、也可以是百分比（图片、视频场景下生效）
   */
  oafs.upload([f], props.static_dir).then((e: any) => {
    cb(e[0], {
      name: f.name, isBorder: false, isShadow: false, isRadius: false, width: '', height: '',
    })
  })
}

const saveMenu = Cherry.createMenuHook('保存', {
  onClick: function () {
    let des = editor.getValue()
    emits('save', des)
    return
  }
});

const backMenu = Cherry.createMenuHook('返回', {
  onClick: function () {
    emits('update:modelValue', true)
    return
  }
})

const init = () => {
  let config = {
    value: props.content,
    id: props.eid,
    // isPreviewOnly: true,
    callback: {
    },
    fileUpload: fileUpload,
  };
  // @ts-ignore
  options.toolbars.customMenu.saveMenu = saveMenu
  // @ts-ignore
  options.toolbars.customMenu.backMenu = backMenu
  editor = new Cherry(Object.assign({}, options, config));
  set_mode(props.modelValue)
}


onMounted(() => {
  init()
})
</script>

<style>
iframe.cherry-dialog-iframe {
  width: 100%;
  height: 100%;
}

.cherry {
  background: none;
  box-shadow: none;
  height: 100%;
}

.cherry-previewer {
  background: none;
  border: none;
}

.cherry-toolbar {
  box-shadow: none;
}
</style>

