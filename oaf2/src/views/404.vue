<style>
</style>
<template>
  <div class='home d-flex justify-center align-center'>
    <one-icon style="font-size: 100px">404</one-icon>
  </div>
</template>

<script lang='ts'>
import {Component, Vue} from 'vue-property-decorator'
import util from '@/libs/util'

@Component({
  components: {}
})
export default class NotFound extends Vue {
  mounted() {
  }

  created() {
    util.title('404')
  }
}
</script>
