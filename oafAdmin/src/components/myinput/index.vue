<template>
    <div :vtype="type"
        :class="[hideBorder ? 'hide-hr' : '', tidy ? '' : 'no-tidy', flexy ? 'flex-col' : '', disabled ? 'cursor-not-allowed' : '']"
        ref="all" class="my-input center flex justify-center items-center relative" :style="dy_style">
        <div class="flex-shrink" :style="{ 'width': labelWidth }">
            <slot name="label"></slot>
        </div>
        <template v-if="type === ArgType.Number">
            <input :type="type" :disabled="disabled" @input="check()" :value="value" @focusout="update"
                @focusin="change('input')" class="main w-full" :style="dy_style" style="font-weight: inherit"
                ref="inputRef" @blur="update" @keyup.enter="update; unblur()">
        </template>
        <template v-else-if="type === ArgType.Text">
            <input :type="type" :disabled="disabled" @input="check()" :value="value" @focusout="update"
                @focusin="change('input');" class="main w-full" :style="dy_style" style="font-weight: inherit"
                ref="inputRef" @blur="update" @keyup.enter="update">
        </template>
        <template v-else-if="type === ArgType.Password">
            <input :type="type" :disabled="disabled" @input="check()" :value="value" @focusout="update"
                @focusin="change('input');" class="main w-full" :style="dy_style" style="font-weight: inherit"
                ref="inputRef" @blur="update" @keyup.enter="update">
        </template>

        <template v-else-if="type === ArgType.File">
            <!--    <FormKit v-model="value" @change="update" class="main" type="file" outer-class="w-full" />-->
            <!--    <FormKit class="main" type="file" outer-class="w-full" />-->
            <div class="div-center rounded-md w-full relative" style="">
                <slot name='file'>
                    {{ value == "" ? "no file chosen" : value }}
                </slot>
                <input class="absolute w-full h-full" type="file"
                    style="color:white;font-size: large;opacity: 0%;top:0;left:0" @change="choose_file($event)" />
            </div>
            <!--          @click="setParameter(v.key,vv.key, c.value)"-->
        </template>
        <template v-else-if="type === ArgType.Radio">
            <template v-for="(ov, ok) in transDic">
                <div :class="[value === ok ? 'radio-btn-active' : 'div-btn', flexy ? 'w-full' : '']"
                    @click="setSelect(ok)" style="color:white;height: 3rem"
                    class="div-center font-bold truncate radio-btn mx-8 rounded-md p-2 my-4 transition duration-500">
                    {{ ov }}
                </div>
            </template>

        </template>

        <template v-else-if="type === ArgType.Select">
            <div class="main cursor-pointer w-full overflow-x-auto whitespace-nowrap" @click="showSelect"
                :title="title">
                <span v-if="!value">未选择</span>
                <span v-else-if="!Array.isArray(value)">{{ transDic[value] || value }}</span>
                <template v-else>
                    <span class="mx-2" v-for="iv in value">{{ transDic[iv] || iv }}</span>
                </template>
            </div>
            <div @mouseleave="showSelectOpt = false"
                :style="{ left: selectPos[0] + 'px', top: selectPos[1] + 'px', height: showSelectOpt ? '20rem' : '0rem' }"
                class="select-opt text-base text-white rounded-md overflow-y-auto" style="min-width: 10rem;"
                :title="title">
                <div class="m-2 p-2" v-if="!options">暂无选项</div>
                <div :class="[ok === value ? 'bg-gray-500' : 'bg-gray-800']"
                    class="cursor-pointer m-2 p-2 rounded-md hover:bg-gray-500" @click="setSelect(ok)"
                    v-for="(ov, ok) in transDic">
                    {{ ov }}
                </div>
                <div class="w-full h-32"></div>
            </div>
        </template>
        <template v-else-if="type === ArgType.Region">
            <div class="flex items-center justify-center">
                <template v-if="value[0] !== '∞'">
                    <one-icon class="div-btn" @click="updateIndex(0, '∞')">kuohao</one-icon>
                    <input type="number" :disabled="disabled" @input="check()" v-model="value[0]" @focusout="update"
                        @focusin="change('input')" class="main w-1/3 text-center" @blur="update" @keyup.enter="update">
                </template>
                <template v-else>
                    <one-icon class="div-btn" @click="updateIndex(0, 0)">zuokuohao</one-icon>
                    <div class="w-1/3 flex justify-center items-center">
                        <one-icon>minus</one-icon>
                        <one-icon>infinite</one-icon>
                    </div>
                </template>
                <div>,</div>

                <template v-if="value[1] !== '∞'">
                    <input type="number" :disabled="disabled" v-model="value[1]" @focusout="update"
                        @focusin="change('input')" class="main w-1/3 text-center" @blur="update" @keyup.enter="update">
                    <one-icon class="div-btn" @click="updateIndex(1, '∞')">kuohao-r</one-icon>
                </template>
                <template v-else>
                    <div class="w-1/3 flex justify-center items-center">
                        <one-icon>plus</one-icon>
                        <one-icon>infinite</one-icon>
                    </div>
                    <one-icon class="div-btn" @click="updateIndex(1, 1)">youkuohao</one-icon>
                </template>
            </div>
        </template>
        <template v-else-if="type === ArgType.Bool">
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
        <hr>
    </div>
</template>

<script lang="ts" setup>

import { onMounted, ref, watch, computed } from 'vue'
import { ArgType, Dict } from '@/models'
import validator from 'validator';


const props = withDefaults(defineProps<{
    modelValue?: any
    type?: ArgType,
    options?: any,
    disabled?: boolean
    hideBorder?: boolean
    tidy?: boolean
    labelWidth?: string
    align?: string
    flexy?: boolean
    require?: boolean
    validator?: any
    //用于超出宽度时鼠标放上去显示值
    title?: string
}>(), {
    modelValue: '',
    type: ArgType.Text,
    disabled: false,
    hideBorder: false,
    tidy: false,
    align: '',
    flexy: false,
    require: false,
    labelWidth: '4rem'
})
const emit = defineEmits<{
    (e: 'update:modelValue', data: any): void
    (e: 'change', data: any): void
    (e: 'upload', data: any): void
}>()

const dy_style = computed(() => `text-align:${props.align}`)

let inputRef = ref<HTMLInputElement>()
let all = ref<HTMLElement>()

const transDic = ref({} as Dict)

const change = (s: string) => {
    if (props.disabled) {
        return
    }
    if (s === 'idle') {
        all.value?.classList.remove('my-input-active')
        all.value?.classList.remove('my-input-error')
        return
    } else if (s === 'input') {
        all.value?.classList.add('my-input-active')
    } else if (s === 'error') {
        all.value?.classList.add('my-input-error')
    }
}

const value = ref(props.modelValue)
const sync = () => {
    if (typeof props.modelValue === 'object') {
        value.value = JSON.parse(JSON.stringify(props.modelValue))
    } else {
        value.value = props.modelValue
    }
    if (props.type === ArgType.Number) {
        let v = parseFloat(props.modelValue) || 0
    }
    if (props.type === ArgType.Radio || props.type === ArgType.Select) {
        transDic.value = {}
        if (Array.isArray(props.options)) {
            for (let i of props.options) {
                if (typeof i === 'string') {
                    transDic.value[i] = i
                } else {
                    transDic.value[i.key] = i.name
                }
            }
        } else {
            for (let i in props.options) {
                transDic.value[i] = props.options[i]
            }
        }
    }
}
watch(props, sync)
const check = (e?: InputEvent) => {
    if (props.type === ArgType.Number) {
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
    } else if (props.type === ArgType.Region) {
        if (value.value[0] !== '∞' && value.value[1] !== '∞' && value.value[0] >= value.value[1]) {
            return false
        }
    } else if (props.type === ArgType.Text || props.type === ArgType.Password) {
        value.value = inputRef.value?.value
        if (!validator.isLength(value.value, props.options)) {
            return false
        }
    }
    if (typeof props.validator === 'function') {
        if (!props.validator(value.value)) {
            return false
        }
    }
    return true
}
const update = () => {
    if (check()) {
        change('idle')
        emit('update:modelValue', value.value)
        emit('change', value.value)
    } else {
        change('error')
    }
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
    var filename = String(e.target.files[0].name)
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

<style lang="less" scoped>
.no-tidy {
    padding: 0.5rem 2rem;
}

.my-input {
    position: relative;

    hr {
        margin: auto;
        position: absolute;
        bottom: -1px;
        width: calc(100% - 4rem);
        left: 2rem;
        border: var(--input-line-default) solid 1px;
        //visibility: hidden;
        transition: all 0.2s linear;
    }

    &:hover hr {
        border: var(--input-line-shine) solid 1px;
        width: 100%;
        left: 0;
    }
}

.hide-hr {
    hr {
        border: none !important;
        width: 0;
        left: 50%;
    }
}

.my-input-active {
    hr {
        border: var(--input-line-shine) solid 1px;
        width: 100%;
        left: 0;
    }
}

.my-input-error {
    hr {
        border: var(--input-line-error) solid 1px !important;
        width: 100%;
        left: 0;
    }
}

.main {
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
    min-height: 2.5rem;
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
