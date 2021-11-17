<template>
  <v-card class="core" elevation="4">
    <v-system-bar color="info">
      <one-icon>mail</one-icon>
      <v-spacer></v-spacer>
      <one-icon>user-group</one-icon>
      <span>*{{ core.user_count }}</span>
    </v-system-bar>
    <div></div>
  </v-card>
</template>
<script lang='ts'>
import {Component, Prop, Vue} from 'vue-property-decorator'

@Component({
  components: {}
})
export default class AppCard extends Vue {
  @Prop({default: {}})
  core: any

  mounted() {
    console.log(this.core)
  }
}
</script>
<style scoped>
.core {
  width: 256px;
  background: #2c3e50;
  height: 128px;
}
</style>
