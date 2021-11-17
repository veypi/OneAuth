<template>
  <v-app @mousewheel.native.prevent="">
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

<style lang="less">
@import './assets/common';

html,
body {
  .full_size;
  margin: 0;
  padding: 0;
}

#app {
  .full_size;
  //.none_select;
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
}

::-webkit-scrollbar {
  display: none; /* Chrome Safari */
}
</style>
