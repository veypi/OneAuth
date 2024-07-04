 <!--
 * menu.vue
 * Copyright (C) 2024 veypi <i@veypi.com>
 * 2024-06-06 16:18
 * Distributed under terms of the MIT license.
 -->
<template>
  <div class="select-none items flex flex-col">
    <template v-for="(v, i) in menus">
      <div class="item flex items-center justify-start px-3 gap-1 py-4" :active='is_equal(v.path, route.fullPath)'
        @click="$router.push(v.path)">
        <slot :name="'L' + i" @click='$router.push(v.path)'>
          <div class='ico' v-if="show_ico">
            <OneIcon :name='v.ico' />
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

let menu_handler = useMenuStore()
let menus = computed(() => menu_handler.menus)
let route = useRoute()

const is_equal = (p1: string, p2: string) => {
  if (!p1.endsWith('/')) {
    p1 = p1 + '/'
  }
  if (!p2.endsWith('/')) {
    p2 = p2 + '/'
  }
  return p1 === p2
}

withDefaults(defineProps<{
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
    color: var(--color-warning);
    cursor: default;
  }

  .item:hover {
    background: var(--base-bg-1);
  }
}
</style>

