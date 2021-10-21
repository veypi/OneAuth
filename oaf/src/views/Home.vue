<style>
.home {
  height: 100%;
  width: 100%;
}
</style>
<template>
  <div class='home d-flex justify-center align-center'>
    <one-icon style="color: aqua;font-size: 50px">glassdoor</one-icon>
  </div>
</template>

<script lang='ts'>
import {Component, Vue} from 'vue-property-decorator'
import util from '@/libs/util'

@Component({
  components: {}
})
export default class Home extends Vue {
  apps = []

  getApps() {
    this.$api.app.list().Start(d => {
      console.log(d)
      this.apps = d
    })
  }

  mounted() {
    this.getApps()
  }

  created() {
  }

  beforeCreate() {
    if (!util.checkLogin()) {
      this.$router.push({name: 'login', query: this.$route.query, params: this.$route.params})
    }
  }
}
</script>
