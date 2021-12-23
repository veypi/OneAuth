<template>
  <div @click="click">
    <input ref="file" type="file" hidden @change="upload">
    <slot></slot>
  </div>
</template>

<script lang="ts" setup>
import {createClient} from '@/libs/webdav'
import {ref} from 'vue'
import {useStore} from '@/store'
import {useMessage} from 'naive-ui'

let store = useStore()
let msg = useMessage()
let file = ref(null)
let emits = defineEmits<{
  (e: 'success', v: string): void
  (e: 'failed'): void
}>()
let props = withDefaults(defineProps<{
  url: string
}>(), {
  url: '',
})

function click() {
  file.value.dispatchEvent(new MouseEvent('click'))
}

let prefix = '/file/public/app/' + store.state.oauuid + '/'

let client = createClient(prefix,
  {headers: {auth_token: localStorage.getItem('auth_token')}})

async function upload() {
  let list = file.value.files
  if (list.length) {
    let reader = new FileReader()
    reader.onload = function (event) {
      var res = event.target.result
      // let data = new Blob([res])
      let url = props.url.replaceAll('.',  '.'+new Date().getTime().toString() + '.')
      client.putFileContents(url, res).then(e => {
        if (e) {
          emits('success', prefix + url)
        } else {
          emits('failed')
        }
      })
    }
    reader.readAsArrayBuffer(list[0])
  } else {
  }
}
</script>

<style scoped>

</style>
