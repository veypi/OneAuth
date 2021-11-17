<template>
  <n-layout has-sider>
    <n-layout-sider
      collapse-mode="transform"
      :collapsed-width="0"
      :width="150"
      show-trigger="bar"
      content-style="padding: 4px;"
      bordered
      default-collapsed
      :native-scrollbar="false"
      style="height: calc(100vh - 108px)"
    >
      <slot name="sider"></slot>
    </n-layout-sider>
    <n-layout style="height: calc(100vh - 108px);padding-left: 20px" :native-scrollbar="false">
      <slot></slot>
      <n-back-top>
      </n-back-top>
    </n-layout>
  </n-layout>
</template>

<script lang="ts" setup>
</script>

<style scoped>

</style>
