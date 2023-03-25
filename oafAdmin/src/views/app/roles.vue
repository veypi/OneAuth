<template>
  <div>
    <div class="flex justify-between">
      <h1 class="page-h1">角色管理</h1>
      <div class="my-5 mr-10">
        <EditorRole @ok="roles.push($event)" v-model="roleFlag" :res="tmp" :uuid="uuid">
          <n-button
            @click="
              tmp = {}
              roleFlag = true
            "
          >
            添加角色
          </n-button>
        </EditorRole>
      </div>
    </div>
    <RoleAuths :res="resources" v-model="raFlag" :uuid="uuid" :role="tmp"></RoleAuths>
    <RoleUsers v-model="ruFlag" :uuid="uuid" :role="tmp"></RoleUsers>
    <n-data-table :bordered="false" :columns="columns" :data="roles" />
    <div class="flex justify-between">
      <h1 class="page-h1">资源管理</h1>
      <div class="my-5 mr-10">
        <EditorRes @ok="resources.push($event)" v-model="trFlag" :res="tmp" :uuid="uuid">
          <n-button
            @click="
              tmp = {}
              trFlag = true
            "
          >
            添加资源
          </n-button>
        </EditorRes>
      </div>
    </div>
    <n-data-table
      class="mb-96"
      :bordered="false"
      :data="resources"
      :columns="resCols"
    ></n-data-table>
  </div>
</template>

<script lang="ts" setup>
import { h, inject, onMounted, Ref, ref } from 'vue'
import api from '@/api'
import { modelsBread, modelsResource, modelsRole } from '@/models'
import { useRoute } from 'vue-router'
import EditorRes from '@/components/editor/resource.vue'
import EditorRole from '@/components/editor/role.vue'
import RoleAuths from '@/components/connectors/roleauths.vue'
import RoleUsers from '@/components/connectors/roleusers.vue'
import { useAppStore } from '@/store'

let local = useAppStore()
let route = useRoute()
let roles = ref<modelsRole[]>([])
let uuid = inject<Ref<string>>('uuid')
const columns = [
  { title: 'ID', key: 'ID', width: 50 },
  { title: '角色名', key: 'Name', width: 100, fixed: 'left' },
  { title: '标签', key: 'Tag', width: 100 },
  { title: '创建时间', key: 'CreatedAt', fixed: 'left' },
  { title: '绑定用户数', key: 'UserCount', fixed: 'left' },
  {
    title: '操作',
    key: '',
    render(row: modelsRole, index: number) {
      return [
        h(
          NButton,
          {
            class: 'mr-1',
            size: 'small',
            onClick: () => {
              raFlag.value = true
              tmp.value = row
            },
          },
          { default: () => '权限' },
        ),
        h(
          NButton,
          {
            class: 'mr-1',
            size: 'small',
            onClick: () => {
              ruFlag.value = true
              tmp.value = row
            },
          },
          { default: () => '用户' },
        ),
        h(
          NButton,
          {
            class: 'mr-1',
            size: 'small',
            onClick: () => {
              api
                .role(uuid.value)
                .delete(row.ID)
                .Start((e) => {
                  roles.value.splice(index, 1)
                })
            },
          },
          { default: () => '删除' },
        ),
      ]
    },
  },
]
onMounted(() => {
  local.setBreads({
    Index: 2,
    Name: '权限',
    RName: route.name,
    RParams: route.params,
    RQuery: route.query,
  } as modelsBread)
  api
    .role(uuid.value)
    .list()
    .Start((e) => {
      roles.value = e
    })
  api
    .resource(uuid.value)
    .list()
    .Start((e) => {
      resources.value = e
    })
})

let resources = ref<modelsResource[]>([])
const resCols = [
  { title: 'ID', key: 'ID', width: 50 },
  { title: 'Name', key: 'Name', width: 200, fixed: 'left' },
  { title: '描述', key: 'Des' },
  {
    title: '操作',
    key: '',
    width: 200,
    fixed: 'right',
    render(row, i) {
      return [
        h(
          NButton,
          {
            class: 'mr-1',
            size: 'small',
            onClick: () => {
              trFlag.value = true
              tmp.value = row
            },
          },
          { default: () => '编辑' },
        ),
        h(
          NButton,
          {
            class: 'mr-1',
            size: 'small',
            onClick: () => {
              api
                .resource(uuid.value)
                .delete(row.ID)
                .Start((e) => {
                  resources.value.splice(i, 1)
                  msg.success('删除成功')
                })
            },
          },
          { default: () => '删除' },
        ),
      ]
    },
  },
]

let tmp = ref({})
let trFlag = ref(false)
let roleFlag = ref(false)
let raFlag = ref(false)
let ruFlag = ref(false)
</script>

<style scoped></style>
