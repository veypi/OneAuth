<template>
  <div>
    <div>
      <div class="flex justify-between">
        <h1 class="page-h1">我的应用</h1>
        <div class="my-5 mr-10">
          <div @click="new_flag = true" v-if="store.auth.Get(R.App, '').CanCreate()">创建应用</div>
        </div>
      </div>
      <div
        class="grid gap-4 grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 2xl:grid-cols-5 text-center"
      >
        <div v-for="(item, k) in ofApps" class="flex items-center justify-center" :key="k">
          <AppCard :core="item"></AppCard>
        </div>
      </div>
    </div>
    <div class="mt-20" v-if="apps.length > 0">
      <h1 class="page-h1">应用中心</h1>
      <div
        class="grid gap-4 grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 2xl:grid-cols-5 text-center"
      >
        <div v-for="(item, k) in apps" class="flex items-center justify-center" :key="k">
          <AppCard :core="item"></AppCard>
        </div>
      </div>
    </div>

    <vmodal v-model="new_flag">
      <div class="h-full w-full flex py-8 px-4 flex-col">
        <div>
          <span>应用名</span>
          <myinput v-model="temp_app.name" />
        </div>
        <div>
          <span>头像</span>
          <uploader
            url="test.ico"
            @success="
              (e) => {
                temp_app.icon = e
              }
            "
          >
            <vavator size="3rem" round :src="temp_app.icon"></vavator>
          </uploader>
        </div>
        <div class="grow"></div>
        <div class="flex justify-end">
          <div class="mx-3" @click="new_flag = false">取消</div>
          <div @click="create_new">创建</div>
        </div>
      </div>
    </vmodal>
  </div>
</template>

<script lang="ts" setup>
import { onMounted, ref } from 'vue'
import api from '@/api'
import AppCard from '@/components/app.vue'
import { R } from '@/auth'
import { modelsApp } from '@/models'
import Uploader from '@/components/uploader'
import { useUserStore } from '@/store/user'
import msg from '@/msg'
import vmodal from '@/components/vmodal'
import myinput from '@/components/myinput'
import vavator from '@/components/vavator'

let store = useUserStore()
let apps = ref<modelsApp[]>([])
let ofApps = ref<modelsApp[]>([])

function getApps() {
  api.app.list().Start((e) => {
    apps.value = e
    api.app
      .user('')
      .list(store.id)
      .Start(
        (e) => {
          ofApps.value = []
          for (let i in e) {
            let ai = apps.value.findIndex((a) => a.UUID === e[i].AppUUID)
            if (ai >= 0) {
              apps.value[ai].UserStatus = e[i].Status
              if (e[i].Status === 'ok') {
                ofApps.value.push(apps.value[ai])
                apps.value.splice(ai, 1)
              }
            }
          }
        },
        () => {},
      )
  })
}

onMounted(() => {
  getApps()
})

let new_flag = ref(false)
let temp_app = ref({
  name: '',
  icon: '',
})
let form_ref = ref(null)
let rules = {
  name: [
    {
      required: true,
      validator(r: any, v: any) {
        return (v && v.length >= 2 && v.length <= 16) || new Error('长度要求2~16')
      },
      trigger: ['input', 'blur'],
    },
  ],
}

function create_new() {
  // @ts-ignore
  form_ref.value.validate((e: any) => {
    if (!e) {
      api.app.create(temp_app.value.name, temp_app.value.icon).Start(
        (e) => {
          e.Status = 'ok'
          ofApps.value.push(e)
          msg.Info('创建成功')
          new_flag.value = false
        },
        (e) => {
          msg.Warn('创建失败: ' + e)
        },
      )
    }
  })
}
</script>

<style scoped></style>
