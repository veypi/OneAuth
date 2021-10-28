<template>
  <div id="wx_reg"></div>
</template>
<script setup lang='ts'>
import {onMounted} from 'vue'

function goto(id: string, app: string, url: string, state?: number, href?: string) {
  // eslint-disable-next-line
  // @ts-ignore
  window.WwLogin({
    id: 'wx_reg',
    appid: id,
    agentid: app,
    redirect_uri: encodeURIComponent(url),
    state: state,
    href: href
  })
}

let aid = ''

let app = ''

let url = ''

onMounted(() => {
  goto(aid, app, url, new Date().getTime())
})

</script>
<style scoped>

</style>
