<template>
  <div :vtype="type" :class="[hideBorder ? 'hide-hr' : '', flexy ? 'flex-col' : '', disabled ? 'cursor-not-allowed' : '']"
    ref="all" class="vinput center flex justify-center items-center relative" :style="dy_style">
    <div v-if="props.label" class="flex-shrink" :style="{ 'width': labelWidth }">
      <slot name="label">{{ props.label }}</slot>
    </div>
    <div v-if="clicktoedit" class="flex-grow vinput-body
      cursor-pointer" @click=" clicktoedit = false"> {{ normalwrap ?
        normalwrap(value) : value }}&nbsp;&nbsp;</div>
    <div v-else class="flex-grow vinput-body">
      <template v-if="type === 'number'">
        <input :type="type" :disabled="disabled" @input="check()" :value="value" @focusin="change('input')"
          class="noborder w-full" :style="dy_style" style="font-weight: inherit" ref="inputRef" @blur="update"
          @keyup.enter="update">
        <hr>
      </template>
      <template v-else-if="type === 'text'">
        <input :type="type" :disabled="disabled" @input="check()" :value="value" @focusin="change('input');"
          class="noborder w-full" :style="dy_style" style="font-weight: inherit" ref="inputRef" @blur="update"
          @keyup.enter="update">
        <hr>
      </template>
      <template v-else-if="type === 'password'">
        <input :type="type" :disabled="disabled" @input="check()" :value="value" @focusin="change('input');"
          class="noborder w-full" :style="dy_style" style="font-weight: inherit" ref="inputRef" @blur="update"
          @keyup.enter="update">
        <hr>
      </template>

      <template v-else-if="type === 'file'">
        <!--    <FormKit v-model="value" @change="update" class="noborder" type="file" outer-class="w-full" />-->
        <!--    <FormKit class="noborder" type="file" outer-class="w-full" />-->
        <div class="div-center rounded-md relative" style="">
          <slot name='file'>
            {{ value == "" ? "no file chosen" : value }}
          </slot>
          <input class="absolute w-full h-full" type="file" style="color:white;font-size: large;opacity: 0%;top:0;left:0"
            @change="choose_file($event)" />
        </div>
        <!--          @click="setParameter(v.key,vv.key, c.value)"-->
      </template>
      <template v-else-if="type === 'radio'">
        <div class="flex justify-between gap-4">
          {{ transDic }}
          <!-- <template :key="ok" v-for="(ov, ok) of transDic"> -->
          <!--   <div :class="[value === ok ? 'radio-btn-active' : -->
          <!--     'div-btn']" @click="setSelect(ok)" style="color:white;" -->
          <!--     class="div-center font-bold grow truncate radio-btn rounded-md transition duration-500"> -->
          <!--     {{ ov || ok }} -->
          <!--   </div> -->
          <!-- </template> -->
        </div>

      </template>

      <template v-else-if="type === 'select'">
        <div class="noborder cursor-pointer w-full overflow-x-auto whitespace-nowrap" @click="showSelect" :title="title">
          <span v-if="value === undefined || value === null">未选择</span>
          <span v-else-if="!Array.isArray(value)">{{ transDic[value] || value }}</span>
          <template v-else>
            <span class="mx-2" v-for="  iv   in   value  " :key="iv">{{ transDic[iv] || iv }}</span>
          </template>
        </div>
        <div @mouseleave="showSelectOpt = false"
          :style="{ left: selectPos[0] + 'px', top: selectPos[1] + 'px', height: showSelectOpt ? '20rem' : '0rem' }"
          class="select-opt text-base text-white rounded-md overflow-y-auto" style="min-width: 10rem;" :title="title">
          <div class="m-2 p-2" v-if="!options">暂无选项</div>
          <div :class="[ov.key === value ? 'bg-gray-500' : 'bg-gray-800']" class="cursor-pointer m-2 p-2 rounded-md
            hover:bg-gray-500" @click="setSelect(ov.key)" v-for="ov of options" :key="ov.key">
            {{ ov.name }}
          </div>
          <div class="w-full h-32"></div>
        </div>
      </template>
      <template v-else-if="type === 'region'">
        <div class="flex items-center justify-center">
          <template v-if="value[0] !== '∞'">
            <OneIcon class="div-btn" @click="updateIndex(0, '∞')">kuohao</OneIcon>
            <input type="number" :disabled="disabled" @input="check()" v-model="value[0]" @focusout="update"
              @focusin="change('input')" class="noborder w-1/3 text-center" @blur="update" @keyup.enter="update">
          </template>
          <template v-else>
            <OneIcon class="div-btn" @click="updateIndex(0, 0)">zuokuohao</OneIcon>
            <div class="w-1/3 flex justify-center items-center">
              <OneIcon>minus</OneIcon>
              <OneIcon>infinite</OneIcon>
            </div>
          </template>
          <div>,</div>

          <template v-if="value[1] !== '∞'">
            <input type="number" :disabled="disabled" v-model="value[1]" @focusout="update" @focusin="change('input')"
              class="noborder w-1/3 text-center" @blur="update" @keyup.enter="update">
            <OneIcon class="div-btn" @click="updateIndex(1, '∞')">kuohao-r</OneIcon>
          </template>
          <template v-else>
            <div class="w-1/3 flex justify-center items-center">
              <OneIcon>plus</OneIcon>
              <OneIcon>infinite</OneIcon>
            </div>
            <OneIcon class="div-btn" @click="updateIndex(1, 1)">youkuohao</OneIcon>
          </template>
        </div>
      </template>
      <template v-else-if="type === 'bool'">
        <div class="rounded-full relative overflow-x-hidden transition duration-300 cursor-pointer text-white leading-8"
          @click="value = !value; update()" style='height: 2rem;width: 6rem;'
          :style="{ 'background': value ? '#1467ff' : '#555' }">
          <template v-if="value">
            <slot name="ok"></slot>
          </template>
          <template v-else>
            <slot name="no"></slot>
          </template>
          <div class="bool-bg rounded-full m-1" style="background: #fff;height: 1.5rem;width: 1.5rem;"
            :style="{ 'transform': 'translateX(' + (value ? '4' : '0') + 'rem)' }">
          </div>
        </div>
      </template>
    </div>
  </div>
</template>

<script lang="ts" setup>

import { OneIcon } from '@veypi/one-icon'

type ArgType = 'text' | 'password' | 'bool' | 'select' | 'radio' | 'number' | 'region' | 'numList' | 'strList' | 'table' | 'grid' | 'file' | 'img'

// const ArgTypesTrans = {
//   [ArgType.Text]: '文本',
//   [ArgType.Password]: '密码',
//   [ArgType.Select]: '选择器',
//   [ArgType.Radio]: '单选框',
//   [ArgType.Number]: '数字',
//   [ArgType.Region]: '区间',
//   [ArgType.NumList]: '数组',
//   [ArgType.StrList]: '文本集合',
//   [ArgType.Table]: '表格',
//   [ArgType.Grid]: '矩阵',
//   [ArgType.File]: '文件',
//   [ArgType.Img]: '图片',
//   [ArgType.Bool]: '开关',
// }

const props = withDefaults(defineProps<{
  modelValue?: any
  type?: ArgType,
  options?: any,
  disabled?: boolean
  hideBorder?: boolean
  tidy?: boolean
  label?: string
  labelWidth?: string
  align?: string
  flexy?: boolean
  require?: boolean
  clicktoedit?: boolean,
  normalwrap?: (s: string) => string
  validate?: any
  //用于超出宽度时鼠标放上去显示值
  title?: string
}>(), {
  modelValue: '',
  type: 'text',
  disabled: false,
  hideBorder: false,
  tidy: false,
  align: '',
  flexy: false,
  require: false,
  clicktoedit: false,
  labelWidth: '4rem'
})

const clicktoedit = ref(props.clicktoedit)

const emit = defineEmits<{
  (e: 'update:modelValue', data: any): void
  (e: 'change', data: any): void
  (e: 'upload', data: any): void
}>()

const dy_style = computed(() => `text-align:${props.align}`)

let inputRef = ref<HTMLInputElement>()
let all = ref<HTMLElement>()

const transDic = ref({})

const change = (s: string) => {
  if (props.disabled) {
    return
  }
  if (s === 'idle') {
    all.value?.classList.remove('vinput-active')
    all.value?.classList.remove('vinput-error')
    return
  } else if (s === 'input') {
    all.value?.classList.add('vinput-active')
  } else if (s === 'error') {
    all.value?.classList.add('vinput-error')
  }
}

const value = ref(props.modelValue)
const sync = () => {
  if (props.modelValue === value.value && props.clicktoedit) {
    clicktoedit.value = true
  }
  if (typeof props.modelValue === 'object') {
    value.value = JSON.parse(JSON.stringify(props.modelValue))
  } else {
    value.value = props.modelValue
  }
  if (props.type === 'number') {
    let v = parseFloat(props.modelValue) || 0
  }
  if (props.type === 'radio' || props.type === 'select') {
    transDic.value = {}
    if (Array.isArray(props.options)) {
      for (let i of props.options) {
        if (typeof i === 'string') {
          transDic.value[i] = i
        } else {
          transDic.value[i.key] = i.name
        }
      }
      console.log(transDic.value)
    } else {
      for (let i in props.options) {
        transDic.value[i] = props.options[i]
      }
    }
  }
}
watch(computed(() => props), sync)
const check = (e?: InputEvent) => {
  if (props.type === 'number') {
    let v = inputRef.value?.valueAsNumber
    if (v !== 0 && !v) {
      return false
    }
    if (typeof props.options?.max === 'number' && v > props.options.max) {
      return false
    }
    if (typeof props.options?.min === 'number' && v < props.options.min) {
      return false
    }
    value.value = v
  } else if (props.type === 'region') {
    if (value.value[0] !== '∞' && value.value[1] !== '∞' && value.value[0] >= value.value[1]) {
      return false
    }
  } else if (props.type === 'text' || props.type === 'password') {
    value.value = inputRef.value?.value
  }
  if (typeof props.validate === 'function') {
    if (!props.validate(value.value)) {
      return false
    }
  } else if (props.validate && typeof props.validate.test ===
    'function') {
    if (!props.validate.test(value.value)) {
      return false
    }
  }
  return true
}
const update = () => {
  // if (value.value === props.modelValue) {
  //   return
  // }
  if (check()) {
    change('idle')
    emit('update:modelValue', value.value)
  } else {
    change('error')
  }
  emit('change', value.value)
}

const updateIndex = (index: number, v: any) => {
  if (props.disabled) {
    return
  }
  value.value[index] = v
  update()
}


onMounted(() => {
  sync()
})

const showSelectOpt = ref(false)
const selectPos = ref([0, 0])

const showSelect = (e: MouseEvent) => {
  if (props.disabled) {
    return
  }
  selectPos.value[0] = e.clientX - 20
  selectPos.value[1] = e.clientY - 20
  showSelectOpt.value = true
}

const setSelect = (e: any) => {
  showSelectOpt.value = false
  if (Array.isArray(value.value)) {
    for (let i in value.value) {
      if (value.value[i] === e) {
        value.value.splice(i, 1)
        update()
        return
      }
    }
    value.value.push(e)
  } else {
    value.value = e
  }
  update()
}


function choose_file(e: any) {
  let filename = String(e.target.files[0].name)
  const h = filename.substring(filename.lastIndexOf('.') + 1)

  if (filename.length > 25) {
    value.value = filename.slice(0, 15) + "...\xa0\xa0\xa0." + h
  }
  else {
    value.value = filename
  }
  emit('upload', e.target.files[0])
  // if (resultFile) {
  //     var reader = new FileReader();
  //     reader.readAsText(resultFile);
  //     reader.onload = function (e) {
  //         let d = this.result
  //     };
  //
  // }
}



function unblur() {
  inputRef.value?.blur()
}

</script>

<style lang="scss" scoped>
.no-tidy {
  padding: 0.5rem 2rem;
}

.vinput {
  position: relative;

  width: 100%;

  hr {
    margin: auto;
    position: absolute;
    bottom: -1px;
    width: calc(100% - 0rem);
    left: 0rem;
    border: var(--input-line-default) solid 1px;
    //visibility: hidden;
    transition: all 0.2s linear;
  }

  &:hover hr {
    border: var(--input-line-shine) solid 1px;
    width: 100%;
    left: 0rem;
  }

}

.vinput-body {}

.hide-hr {
  hr {
    border: none !important;
    width: 0;
    left: 50%;
  }
}

.vinput-active {
  hr {
    border: var(--input-line-shine) solid 1px;
  }
}

.vinput-error {
  hr {
    border: var(--input-line-error) solid 1px !important;
  }
}

.noborder {
  border: none;
  outline: none;
  background: none;
}

select {
  -webkit-appearance: none;
  -moz-appearance: none;
}

.select-opt {
  z-index: 10;
  position: fixed;
  left: 0;
  top: 0;
  background: #333;
  transform-origin: top;
  transition: height 0.3s linear;
}

.radio-btn {
  background: #A8A8A8;
  min-height: 2rem;
}

.radio-btn-active {
  background: #EF857D;
}


.bool-bg {
  position: absolute;
  height: 100%;
  left: 0px;
  bottom: 0px;
  /* 渐变背景 ,自左到右 */
  /* background: linear-gradient(135deg, #FF9D6C, #BB4E75); */
  /* background: linear-gradient(to right, #f09819, #ff5858); */
  /* 添加动画过渡.贝塞尔曲线 */
  transition: 0.3s cubic-bezier(1, 0.05, 0.9, 0.9);
  /* transition: left 3s linear; */
}
</style>
