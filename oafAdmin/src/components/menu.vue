 <!--
 * menu.vue
 * Copyright (C) 2022 veypi <i@veypi.com>
 * 2022-10-12 10:00
 * Distributed under terms of the Apache license.
 -->
<template>
    <div class="w-full h-full relative">
        <div class="left-menu px-2 py-4 h-full absolute" :style="dy_style.menu">
            <div class="lm-conent overflow-hidden h-full w-full">
                <slot name='menu'>
                </slot>
            </div>
            <div @click="toggle" class='lm-icon'>
                <div class="icon-top"></div>
                <div class="icon-bot"></div>
            </div>
        </div>
        <div class="main h-full absolute" :style="dy_style.main">
            <slot>
            </slot>
        </div>
    </div>
</template>

<script lang="ts" setup>
import { computed, ref } from 'vue';



const props = withDefaults(defineProps<{
    modelValue?: boolean
    width?: number,
}>(), {
    modelValue: true,
    width: 8,
})
const emit = defineEmits<{
    (e: 'update:modelValue', data: boolean): void
}>()

const value = ref(props.modelValue)
const dy_style = computed(() => {
    if (value.value) {
        return {
            menu: {
                left: '0',
                width: props.width + 'rem'
            },
            main: {
                width: 'calc(100% - ' + props.width + 'rem)',
                left: props.width + 'rem'
            },
            top: '12deg',
            bot: '-12deg'
        }
    }
    return {
        menu: {
            left: -props.width + 'rem',
            width: props.width + 'rem'
        },
        main: {
            width: '100%',
            left: '0',
        },
        top: '-12deg',
        bot: '12deg'
    }
})

const toggle = () => {
    value.value = !value.value
    emit('update:modelValue', value.value)
}

</script>

<style scoped>
.left-menu {
    transition: all 0.2s linear;
}

.lm-conent {
    background: var(--base-bg-3);
    position: relative;
}

.lm-icon {
    cursor: pointer;
    height: 72px;
    width: 32px;
    position: absolute;
    top: calc(50% - 36px);
    right: -28px;
    transition: all 0.2s linear;
    z-index: 10;
}

.lm-icon div {
    background: #999;
    transition: all 0.2s linear;
    position: absolute;
    width: 4px;
    border-radius: 2px;
    height: 38px;
    left: 14px;
}

.lm-icon:hover div {
    background: #777;
}


.lm-icon .icon-bot {
    position: absolute;
    top: 34px;
}

.lm-icon:hover .icon-top {
    transform: rotate(v-bind('dy_style.top')) scale(1.15) translateY(-2px);
}


.lm-icon:hover .icon-bot {
    transform: rotate(v-bind('dy_style.bot')) scale(1.15) translateY(2px);
}

.main {
    transition: all 0.2s linear;

}
</style>

