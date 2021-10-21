<template>
  <v-app>
    <v-app-bar
      app
      color="primary"
      dark
    >
      <div class="d-flex align-center">
        <one-icon style="color: aqua;font-size: 56px">glassdoor</one-icon>
        <span class="font-italic font-weight-bold" style="font-size: 20px">统一认证</span>
      </div>
      <v-spacer></v-spacer>
    </v-app-bar>

    <v-main>
      <router-view></router-view>
    </v-main>
  </v-app>
</template>

<script lang="ts">
import Vue from 'vue'
import util from '@/libs/util'

export default Vue.extend({
  name: 'App',

  components: {},

  data: () => ({
    //
  }),

  beforeCreate() {
    util.title('统一认证')
    this.$store.dispatch('fetchSelf')
  }
})
</script>
