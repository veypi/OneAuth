<template>
  <svg class="icon" aria-hidden="true">
    <use :xlink:href="'#icon-'+icon"></use>
  </svg>
</template>
<script lang='ts'>
import {Component, Vue} from 'vue-property-decorator'

@Component({
  components: {}
})
export default class OneIcon extends Vue {
  get icon() {
    if (this.$slots.default) return this.$slots.default[0].text?.trim()
    console.warn('blank icon name')
    return ''
  }
}
</script>
<style scoped>
.icon {
  width: 1em;
  height: 1em;
  vertical-align: -0.15em;
  fill: currentColor;
  overflow: hidden;
}
</style>
