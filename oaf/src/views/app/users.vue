<template>
  <div>
    <div class="flex justify-between">
      <h1 class="page-h1">用户名单</h1>
      <div class="my-5 mr-10">
        <n-button @click="temp_user = {}; tu_flag = true">添加用户</n-button>
      </div>
    </div>
    <n-data-table :bordered="false" :columns="columns" :scroll-x="980" :data="users" />
    <n-modal v-model:show="tu_flag">
      <n-card class="w-4/5 md:w-96 rounded-2xl" :title="temp_user.Index >= 0 ? temp_user.Username : ' '" :bordered="false"
        size="huge">
        <template #header-extra>{{ temp_user.Index >= 0 ? '编辑' : '创建' }}</template>
        <div class="grid grid-cols-5 gap-1 gap-y-8" style="line-height: 34px">
          <div>用户名</div>
          <div class="col-span-4">
            <n-input v-model:value="temp_user.Username"></n-input>
          </div>
        </div>
        <template #footer>
          <div class="flex justify-end">
            <n-button class="mx-3" @click="tu_flag = false">取消</n-button>
            <n-button>更新</n-button>
          </div>
        </template>
      </n-card>
    </n-modal>
  </div>
</template>

<script lang="ts" setup>
import { inject, onMounted, ref, h, computed, Ref } from 'vue'
import api from '@/api'
import { NTag as ntag, NButton as nbtn, useDialog } from 'naive-ui'
import { useStore } from '@/store'
import { R } from '@/auth'
import { modelsBread, modelsUser } from '@/models'
import { useRoute } from 'vue-router'

let store = useStore()
let route = useRoute()
let dialog = useDialog()
let uuid = inject<Ref<string>>('uuid')
let users = ref<modelsUser[]>([])
let isOA = computed(() => uuid.value === store.state.oauuid)

onMounted(() => {
  if (isOA) {
    columns.value.push({
      title: '操作',
      key: '',
      width: 200,
      render(row, index) {
        return [
          h(nbtn, {
            onClick: () => {
              temp_user.value = Object.assign({ Index: index }, row.User)
              tu_flag.value = true
            },
          }, {
            default: () => '编辑',
          }),
        ]
      },
    },
    )
  }

  api.app.user(uuid.value as string).list('-').Start(e => {
    users.value = e
  })
})

let columns = ref([
  {
    title: 'ID',
    key: 'UserID',
    width: 100,
  },
  {
    title: '用户',
    key: 'User.Username',
    width: 100,
    fixed: 'left',
  },
  {
    title: '加入时间',
    key: 'User.CreatedAt',
  },
  {
    title: 'Status',
    key: 'Status',
    width: 100,
    render(row) {
      let t = statusTag(row.Status)
      // @ts-ignore
      return h(ntag, {
        'type': t[1],
        onClick: () => {
          changeStatus(row)
        },
      },
        {
          default: () => t[0],
        },
      )
    },
  },
])

function statusTag(s: string) {
  switch (s) {
    case 'ok':
      return ['正常', 'success']
    case 'apply':
      return ['申请中', 'info']
    case 'deny':
      return ['拒绝', '']
    case 'disabled':
      return ['禁用', 'warning']
  }
  return ['未知', '']
}

function changeStatus(u) {
  if (store.state.user.auth.Get(R.User, uuid.value).CanUpdate()) {
    dialog.warning({
      title: '请选择切换状态',
      content: () => {
        let tags = []
        for (let s of ['ok', 'apply', 'deny', 'disabled']) {
          let t = statusTag(s)
          if (u.Status !== s) {
            // @ts-ignore
            tags.push(h(ntag, {
              'type': t[1],
              onClick: () => {
                api.app.user(uuid.value).update(u.UserID, s).Start(e => {
                  u.Status = s
                  dialog.destroyAll()
                })
              },
            }, {
              default: () => t[0],
            }))
          }
        }
        return h('div', {
          class: 'flex  justify-between mx-16 mt-10',
        }, {
          default: () => tags,
        })
      },
    })
  }
}

let temp_user = ref<modelsUser>({} as modelsUser)
let tu_flag = ref(false)

function add_user() {
}

store.commit('setBreads', {
  Index: 2,
  Name: '用户',
  RName: route.name,
  RParams: route.params,
  RQuery: route.query,
} as modelsBread)
</script>

<style scoped></style>
