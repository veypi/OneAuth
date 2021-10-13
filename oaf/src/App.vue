<template>
  <v-app>
    <v-app-bar
      app
      color="primary"
      dark
    >
      <div class="d-flex align-center">
      </div>
      <v-spacer></v-spacer>
      <v-btn text class="font-italic" style="font-size: 20px">
        统一认证
      </v-btn>
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

  mounted() {
    util.title('统一认证')
  }
})
</script>
