 <!--
 * crud.vue
 * Copyright (C) 2023 veypi <i@veypi.com>
 * 2023-10-10 19:29
 * Distributed under terms of the MIT license.
 -->
<template>
  <div>

    <q-page-sticky position="top-right" style="z-index: 20" :offset="[27, 27]">
      <q-btn @click="modeV = !modeV" round icon="save_as" class="" />
    </q-page-sticky>
    <div class="grid" :class="[modeV ? '' : 'grid-cols-2']">
      <div class="grid" :style="modeV ? grid_len : ''">
        <div :class="styles.k" v-for="k of keys" :key="k.name">
          {{ k.label || k.name }}
        </div>
      </div>
      <div class="grid" :style="modeV ? '' : grid_len">
        <div class="grid hover:bg-gray-200" :style="modeV ? grid_len : ''" v-for="( item, idx ) in  data " :key="idx">
          <div :class="styles.v" v-for="k of keys" :key="k.name">
            {{ item[k.name] }}
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { computed, ref, watch } from 'vue';


interface itemProp {
  name: '',
  label?: '',
  value: any,
  type?: '',
}

interface keyProp {
  name: string,
  label?: string,
  default?: any,
  typ?: string,
}

const grid_len = computed(() => {
  return {
    'grid-template-columns': 'repeat(' +
      (modeV.value ? props.keys?.length : props.data?.length)
      + ', minmax(0, 1fr))'
  }
})


let props = withDefaults(defineProps<{
  vertical?: boolean
  keys?: keyProp[],
  data?: any[]
  kclass?: Array<string>,
  vclass?: Array<string>,
  cclass?: Array<string>,
}>(),
  {
    vertical: false,
    data: [] as any,
    kclass: [] as any,
    vclass: [] as any,
    cclass: ['w-40', 'h-40'] as any,
  }
)

const styles = computed(() => {
  let k = [];
  let v = [];
  k.push(...props.kclass, ...props.cclass)
  v.push(...props.vclass, ...props.cclass)
  return {
    k, v
  }
})
const modeV = ref(props.vertical)

watch(computed(() => props.vertical), (v) => modeV.value = v)

</script>

<style scoped></style>

