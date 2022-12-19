<template>
    <div class="flex items-center justify-center">
        <div :style="{ background: Theme.me.lightBox, 'box-shadow': Theme.me.lightBoxShadow }"
            class="px-10 pb-9 pt-28 rounded-xl w-96">
            <n-form label-width="70px" label-align="left" :model="data" ref="form_ref" label-placement="left"
                :rules="rules">
                <n-form-item required label="用户名" path="username">
                    <n-input @keydown.enter="divs[1].focus()" :ref="el => { if (el) divs[0] = el }"
                        v-model:value="data.username"></n-input>
                </n-form-item>
                <n-form-item required label="密码" path="password">
                    <n-input @keydown.enter="login" :ref="el => { if (el) divs[1] = el }" v-model:value="data.password"
                        type="password"></n-input>
                </n-form-item>
                <div class="flex justify-around mt-4">
                    <n-button @click="login">登录</n-button>
                    <n-button @click="router.push({ name: 'register' })">注册</n-button>
                </div>
            </n-form>
        </div>
    </div>
</template>
<script lang="ts" setup>
import { computed, onMounted, ref, watch } from 'vue'
import { Theme } from '@/theme'
import api from '@/api'
import { useRoute, useRouter } from 'vue-router'
import { store } from '@/store'
import { modelsApp } from '@/models'
import util from '@/libs/util'
import msg from '@veypi/msg'

const route = useRoute()
const router = useRouter()

const divs = ref([])
let form_ref = ref(null)
let data = ref({
    username: '',
    password: '',
})
let rules = {
    username: [
        {
            required: true,
            validator(r: any, v: any) {
                return (v && v.length >= 3 && v.length <= 16) || new Error('长度要求3~16')
            },
            trigger: ['input', 'blur'],
        },
    ],
    password: [{
        required: true,
        validator(r: any, v: any) {
            return (v && v.length >= 6 && v.length <= 16) || new Error('长度要求6~16')
        },
    }],
}

let uuid = computed(() => {
    return route.query.uuid
})
let ifLogOut = computed(() => {
    return route.query.logout === '1'
})

function login() {
    // @ts-ignore
    form_ref.value.validate((e: any) => {
        if (!e) {
            api.user.login(data.value.username, data.value.password).Start((headers: any) => {
                if ('auth_token' in headers) {
                    localStorage.auth_token = headers.auth_token
                    store.commit('user/refreshToken', localStorage.auth_token)
                    msg.Info('登录成功')
                    store.dispatch('user/fetchUserData')
                    let url = route.query.redirect || headers.redirect || '/'
                    redirect(url)
                } else {
                    msg.Info('正在申请加入，请等待管理员审批')
                }
                console.log(headers)
            }, e => {
                console.log(e)
                msg.Warn('登录失败：' + e.headers.error)
            })
        }
    })
}

function redirect(url?: string) {
    if (uuid.value && uuid.value !== store.state.oauuid) {
        api.app.get(uuid.value as string).Start((app: modelsApp) => {
            api.token(uuid.value as string).get().Start(e => {
                if (!url) {
                    url = app.UserRefreshUrl
                }
                e = encodeURIComponent(e)
                url = url.replaceAll('$token', e)
                window.location.href = url
            })
        }, e => {
        })
    } else if (util.checkLogin()) {
        if (url) {
            router.push(url)
        } else {
            router.push({ name: 'home' })
        }
    }
}

onMounted(() => {
    if (!ifLogOut.value) {
        redirect()
    }
    if (divs.value[0]) {
        // @ts-ignore
        divs.value[0].focus()
    }
})
</script>

<style scoped>
</style>
