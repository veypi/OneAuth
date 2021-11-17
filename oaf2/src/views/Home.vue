<style>
.home {
  height: 100%;
  width: 100%;
}
</style>
<template>
  <div class='full_size'>
    <v-row no-gutters class="pa-8">
      <v-col v-for="(item, key) in apps" :key="key" class="mx-4 my-2">
        <AppCard :core="item"></AppCard>
      </v-col>
    </v-row>
  </div>
</template>

<script lang='ts'>
import {Component, Vue} from 'vue-property-decorator'
import util from '@/libs/util'
import AppCard from '@/components/app.vue'

@Component({
  components: {
    AppCard
  }
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
