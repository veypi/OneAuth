<template>
  <div class='home d-flex justify-center align-center'>
    <wx-login v-if="enable" :aid="aid" :app="agentID" :url="url"></wx-login>
    <v-overlay :value="!enable">
      <v-progress-circular
        indeterminate
        size="64"
      ></v-progress-circular>
    </v-overlay>
  </div>
</template>
<script lang='ts'>
import {Component, Vue} from 'vue-property-decorator'
import WxLogin from '@/components/WxLogin.vue'

@Component({
  components: {
    WxLogin
  }
})
export default class Wx extends Vue {
  aid = ''
  agentID = ''
  url = ''

  get enable() {
    return this.uuid && this.aid && this.agentID && this.url
  }

  get uuid() {
    return this.$route.query.uuid
  }

  get code() {
    return this.$route.query.code
  }

  get state() {
    return this.$route.query.state
  }

  get msg() {
    return this.$route.query.msg
  }

  mounted() {
    if (this.msg) {
      console.log(this.msg)
      alert(this.msg)
    }
  }

  created() {
    if (this.uuid) {
      this.api.app.get(this.uuid).Start(e => {
        this.url = e.wx.url + '/api/wx/login/' + this.uuid
        this.aid = e.wx.corp_id
        this.agentID = e.wx.agent_id
      })
    }
  }
}
</script>
<style scoped>

</style>
