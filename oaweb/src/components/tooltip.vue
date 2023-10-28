 <!--
 * tooltip.vue
 * Copyright (C) 2023 veypi <i@veypi.com>
 * 2023-10-23 21:49
 * Distributed under terms of the MIT license.
 -->
<template>
  <div class="v-tooltip">
    <slot></slot>
    <div class="v-tooltip-text">
      <slot name="text">
        {{ text }}
      </slot>
    </div>
  </div>
</template>

<script lang="ts" setup>

withDefaults(defineProps<{
  text?: string
}>(),
  {}
)
</script>

<style scoped>
.v-tooltip {
  position: relative;
  display: inline-block;
}

.v-tooltip .v-tooltip-text {
  visibility: hidden;
  visibility: hidden;
  background-color: rgba(0, 0, 0, 0.8);
  color: #fff;
  text-align: center;
  padding: 5px 5px;
  border-radius: 6px;
  position: absolute;
  z-index: 1;
}

.v-tooltip:hover .v-tooltip-text {
  visibility: visible;
}
</style>
