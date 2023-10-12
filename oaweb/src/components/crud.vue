 <!--
 * crud.vue
 * Copyright (C) 2023 veypi <i@veypi.com>
 * 2023-10-10 19:29
 * Distributed under terms of the MIT license.
 -->
<template>
  <div>
    <div class="v-crud" :vertical='modeV'>
      <div class="v-crud-keys">
        <template v-for="k of keys" :key="k.name">
          <div class="v-crud-cell" :style="Object.assign({},
            cstyle.width[k.name], cstyle.k, cstyle.kv[k.name])">
            {{ k.label || k.name }}
          </div>
        </template>
      </div>
      <div class="v-crud-values">
        <div class="v-crud-line rounded-3xl" :class="cstyle.line" v-for=" (item, idx)  in  items" :key="idx">
          <template v-for="k of keys" :key="k.name">
            <div class="v-crud-cell" :changed="item.__changed[k.name]" :selected="selected
              === `${item.__idx}.${k.name}`" @click="ifselect ?
    selected = `${item.__idx}.${k.name}` : ''" :style="Object.assign({},
    cstyle.width[k.name], cstyle.v, cstyle.kv[k.name])">
              <slot :name="`k_${k.name}`" :row="item" :value="item[k.name]" :set="setv(item, k.name)">
                <template v-if="k.editable === undefined ? editable :
                  k.editable">
                  <vinput :align="valign" :model-value="item[k.name] === undefined ?
                    k.default : item[k.name]" :type="k.typ" :options="k.options" @update:model-value="setv(item,
    k.name)($event)"></vinput>
                </template>
                <template v-else>
                  <span class="truncate">
                    {{ item[k.name] === undefined ? k.default :
                      item[k.name] }}
                  </span>
                </template>
                <slot :name="`k_${k.name}_append`" :row="item" :value="item[k.name]" :set="setv(item, k.name)">
                </slot>
              </slot>
            </div>
          </template>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { computed, onMounted, ref, watch } from 'vue';
import vinput from 'src/components/vinput'
import { ArgType, Dict } from 'src/models';

interface keyProp {
  name: string,
  label?: string,
  default?: any,
  width?: number,
  typ?: ArgType,
  editable?: boolean,
  options?: any,
  style?: { [k: string]: string },
}
let emits = defineEmits<{
  (e: 'update', v: any): void
}>()

let props = withDefaults(defineProps<{
  vertical?: boolean
  keys?: keyProp[],
  data?: any[]
  hover?: boolean,
  kstyle?: { [k: string]: string },
  vstyle?: { [k: string]: string },
  cstyle?: { [k: string]: string },
  kalign?: 'center' | 'left' | 'right',
  valign?: 'center' | 'left' | 'right',
  editable?: boolean
  ifselect?: boolean

}>(),
  {
    vertical: false,
    data: [] as any,
    hover: false,
    kvstyle: {} as any,
    kstyle: {} as any,
    vstyle: {} as any,
    cstyle: {} as any,
    kalign: 'center',
    valign: 'center',
  }
)
const modeV = ref(props.vertical)
watch(computed(() => props.vertical), (v) => modeV.value = v)

let items = ref<any[]>([])
watch(computed(() => JSON.stringify(props.data)), (_) => {
  syncItems()
  // if (JSON.stringify(v) !== JSON.stringify(o)) {
  //   console.log(JSON.stringify(v))
  //   syncItems()
  // }
}, {})
const syncItems = () => {
  let res = props.data?.map((v: any, i: any) => {
    return Object.assign({ __idx: i, __changed: {} }, v)
  }) as any
  items.value.splice(0, items.value.length)
  items.value.push(...res)
  // Object.assign(items, res)
}


const selected = ref()

let alignDic = { 'center': 'center', 'left': 'start', 'right': 'end' }
const cstyle = computed(() => {
  let res = { line: [], width: {}, kv: {} } as any
  let l = props.keys?.length || 0
  let w = 100
  let style = modeV.value ? 'flex-basis' : 'height'
  props.keys?.forEach((k, i) => {
    if (k.width && k.width > 0 && k.width < 100) {
      res.width[k.name] = { [style]: (k.width || 1) + '%' }
      w = w - k.width
      l = l - 1
    }
    if (k.style) {
      res.kv[k.name] = k.style
    }
  })
  props.keys?.forEach((k, i) => {
    if (k.width && k.width > 0 && k.width < 100) {
    } else {
      res.width[k.name] = { [style]: w / l + '%' }
    }
  })
  res.k = Object.assign({ 'justify-content': alignDic[props.kalign] }, props.cstyle, props.kstyle)
  res.v = Object.assign({ 'justify-content': alignDic[props.valign] }, props.cstyle, props.vstyle)
  if (props.hover) {
    res.line.push('hover:bg-gray-200');
  }
  return res
})

const updatedItems = ref<Dict[]>([])
const setv = (item: any, k: string) => {
  return (v: any) => {
    item[k] = v
    item.__changed[k] = true
    if (!updatedItems.value[item.__idx]) {
      updatedItems.value[item.__idx] = {}
    }
    updatedItems.value[item.__idx][k] = v
    emits('update', updatedItems.value)
    console.log(`update ${k} to ${v}`)
  }
}

const save = () => {
}

const reload = () => {
  console.log('reload')
  syncItems()
}

onMounted(() => {
  syncItems()
})

defineExpose({
  reload
})
</script>

<style lang="scss" scoped>
.v-crud {
  display: flex;
}


.v-crud-keys {
  display: flex;
}

.v-crud-values {
  display: flex;
}

.v-crud-line {
  display: flex;
}

.v-crud-cell {
  min-width: 8rem;
  min-height: 4rem;
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 1rem;
  overflow: auto;
  overflow-y: auto;

  &[selected=true] {
    background: rgba($color: #888, $alpha: .1);
  }

  &[changed=true] {
    background: rgba($color: #888, $alpha: .1);
  }
}



.v-crud[vertical=true] {
  flex-direction: column;

  .v-crud-keys {
    flex-wrap: nowrap;
  }

  .v-crud-values {
    flex-direction: column;

    .v-crud-line {
      flex-wrap: nowrap;
    }
  }
}

.v-crud[vertical=false] {
  flex-wrap: nowrap;

  .v-crud-keys {
    flex-direction: column;
  }

  .v-crud-values {
    flex-wrap: nowrap;

    .v-crud-line {
      flex-direction: column;
    }
  }
}
</style>

