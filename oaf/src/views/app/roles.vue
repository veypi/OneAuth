<template>
  <div>
    <h1 class="page-h1">角色管理</h1>
    <n-data-table
      :bordered="false"
      :columns="columns"
      :data="roles"
    />
    <div class="flex justify-between">
      <h1 class="page-h1">资源管理</h1>
      <div class="my-5 mr-10">
        <n-button @click="tmp_res={};tr_flag=true">添加资源</n-button>
      </div>
    </div>
    <n-data-table class="mb-96" :bordered="false" :data="resources" :columns="resCols">
    </n-data-table>
    <n-modal v-model:show="tr_flag">
      <n-card class="w-4/5 md:w-96 rounded-2xl" :title="tmp_res.index >= 0 ? tmp_res.Name:' '" :bordered="false"
              size="huge">
        <template #header-extra>{{ tmp_res.index >= 0 ? '编辑' : '创建' }}</template>
        <div class="grid grid-cols-5 gap-1 gap-y-8" style="line-height: 34px">
          <div>资源名</div>
          <div class="col-span-4">
            <n-input v-model:value="tmp_res.Name"></n-input>
          </div>
          <div>资源描述</div>
          <div class="col-span-4">
            <n-input type="textarea" v-model:value="tmp_res.Des"></n-input>
          </div>
        </div>
        <template #footer>
        <div class="flex justify-end">
          <n-button class="mx-3" @click="tr_flag=false">取消</n-button>
          <n-button @click="add_res">创建</n-button>
        </div>
        </template>
      </n-card>
    </n-modal>
  </div>
</template>

<script lang="ts" setup>
import {h, inject, onMounted, Ref, ref} from 'vue'
import api from '@/api'
import {NButton, useMessage} from 'naive-ui'
import {modelsBread, modelsResource, modelsRole} from '@/models'
import {useStore} from '@/store'
import {useRoute} from 'vue-router'

let store = useStore()
let route = useRoute()
let msg = useMessage()
let roles = ref<modelsRole[]>([])
let uuid = inject<Ref<string>>('uuid')
const columns = [
  {title: 'ID', key: 'ID', width: 50},
  {title: '角色名', key: 'Name', width: 100, fixed: 'left'},
  {title: '标签', key: 'Tag', width: 100},
  {title: '创建时间', key: 'CreatedAt', fixed: 'left'},
  {title: '绑定用户数', key: 'UserCount', fixed: 'left'},
  {
    title: '操作',
    key: '',
    render(row) {
      return [
        h(NButton, {
            class: 'mr-1',
            size: 'small',
            onClick: () => console.log(row),
          },
          {default: () => '查看权限'}),
        h(NButton, {
            class: 'mr-1',
            size: 'small',
            onClick: () => console.log(row),
          },
          {default: () => '查看用户'},
        ),
      ]
    },
  },
]
onMounted(() => {
  store.commit('setBreads', {
    Index: 2,
    Name: '权限',
    RName: route.name,
    RParams: route.params,
    RQuery: route.query,
  } as modelsBread)
  api.role(uuid.value).list().Start(e => {
    roles.value = e
  })
  api.resource(uuid.value).list().Start(e => {
    resources.value = e
  })
})

let resources = ref<modelsResource[]>([])
const resCols = [
  {title: 'ID', key: 'ID', width: 50},
  {title: 'Name', key: 'Name', width: 200, fixed: 'left'},
  {title: '描述', key: 'Des'},
  {
    title: '操作',
    key: '',
    width: 100,
    fixed: 'right',
    render(row, i) {
      return h(NButton, {
          class: 'mr-1',
          size: 'small',
          onClick: () => {
            api.resource(uuid.value).delete(row.ID).Start(e => {
              resources.value.splice(i, 1)
              msg.success('删除成功')
            })
          },
        }, {default: () => '删除'},
      )
    },
  },
]

let tmp_res = ref({
  index: -1,
  Name: '',
  Des: '',
})
let tr_flag = ref(false)

function add_res() {
  if (tmp_res.value.index >= 0) {
    return
  }
  api.resource(uuid.value).create(tmp_res.value.Name, tmp_res.value.Des).Start(e => {
    resources.value.push(e)
    tr_flag.value = false
  })
}
</script>

<style scoped>

</style>
