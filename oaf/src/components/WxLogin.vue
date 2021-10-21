<template>
  <div id="wx_reg"></div>
</template>
<script lang='ts'>
import {Component, Vue, Prop} from 'vue-property-decorator'
import '@/libs/wwLogin.js'

@Component({
  components: {}
})
export default class WxLogin extends Vue {
  goto(id: string, app: string, url: string, state?: number, href?: string) {
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

  @Prop({default: ''})
  aid = ''

  @Prop({default: ''})
  app = ''

  @Prop({default: ''})
  url = ''

  mounted() {
    this.goto(this.aid, this.app, this.url, new Date().getTime())
  }

  created() {
  }
}
</script>
<style scoped>

</style>
