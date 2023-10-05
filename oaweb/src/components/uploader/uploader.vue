<template>
  <div @click="click">
    <input enctype="multipart/form-data" ref="file" name="files" multiple type="file" hidden @change="upload">
    <slot></slot>
  </div>
</template>

<script lang="ts" setup>
// import { createClient } from "webdav";
import { ref } from 'vue';
import axios from "axios";

let file = ref<HTMLInputElement>()
let emits = defineEmits<{
  (e: 'success', v: string): void
  (e: 'failed'): void
}>()
let props = withDefaults(defineProps<{
  multiple?: boolean,
}>(), {
  multiple: false
})

function click() {
  file.value?.dispatchEvent(new MouseEvent('click'))
}

const upload = (evt: Event) => {
  evt.preventDefault()
  let f = (evt.target as HTMLInputElement).files as any
  var data = new FormData();
  console.log(f)
  for (let i of f) {
    console.log(i)
    data.append('files', i, i.data)
  }
  axios.post("/api/upload/", data, {
    headers: {
      "Content-Type": 'multipart/form-data',
      'auth_token': localStorage.getItem('auth_token')
    }
  }).then(e => {
    console.log(e.data)
    emits('success', props.multiple ? e.data : e.data[0])
  })
  // var token = sessionStorage.getItem('token')
  // const config = {
  //   headers: {
  //     'Content-Type': 'multipart/form-data'
  //   }
  // }
  // window.API.post('https://110.10.56.10:8000/images/?token=' + token, data, config)
  //   .then(response => this.$router.push('/listImage'))
  //   .catch((error) => {
  //     console.log(JSON.stringify(error))
  //   })
}




// async function dav_upload() {
//   let prefix = '/file/public/app/' + app.id + '/'
//   let client = createClient(prefix,
//     { headers: { auth_token: localStorage.getItem('auth_token') || '' } })
//   let list = file.value?.files || []
//   if (list.length) {
//     let reader = new FileReader()
//     reader.onload = function (event) {
//       var res = event.target?.result
//       // let data = new Blob([res])
//       let url = props.url.replaceAll('.', '.' + new Date().getTime().toString() + '.')
//       client.putFileContents(url, res).then(e => {
//         if (e) {
//           emits('success', prefix + url)
//         } else {
//           emits('failed')
//         }
//       })
//     }
//     reader.readAsArrayBuffer(list[0])
//   } else {
//   }
// }
</script>

<style scoped></style>
