 <!--
 * menu.vue
 * Copyright (C) 2024 veypi <i@veypi.com>
 * 2024-06-06 16:18
 * Distributed under terms of the MIT license.
 -->
<template>
  <div class="select-none items flex flex-col">
    <template v-for="(v, i) in list">
      <div class="item flex items-center justify-center gap-2 py-4" :active='v.path === route.fullPath'
        @click="$router.push(v.path)">
        <slot :name="'L' + i" @click='$router.push(v.path)'>
          <div class='ico' v-if="show_ico">
            <OneIcon>{{ v.ico }}</OneIcon>
          </div>
          <div class="text-nowrap" v-if="show_name">
            {{ v.name }}
          </div>
        </slot>
      </div>
    </template>
  </div>
</template>

<script lang="ts" setup>
import { OneIcon } from '@veypi/one-icon'

let route = useRoute()
interface item {
  ico: string
  name: string
  path: string
  label?: string
  subs?: item[]
}

withDefaults(defineProps<{
  list: item[],
  vertical?: boolean,
  show_ico?: boolean,
  show_name?: boolean,
}>(),
  {
    vertical: true,
    show_ico: true,
    show_name: true
  }
)
</script>

<style scoped lang="scss">
.items {
  .item {
    opacity: 0.8;
    cursor: pointer;
    color: var(--base-txt);
  }

  .item[active=true] {
    opacity: 1;
    background: var(--base-bg-1);
    cursor: default;
  }

  .item:hover {
    background: var(--base-bg-1);
  }
}
</style>

