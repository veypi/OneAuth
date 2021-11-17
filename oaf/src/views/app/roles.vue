<template>
  <div>
    <h1 class="page-h1">角色管理</h1>
    <n-data-table
      :bordered="false"
      :columns="columns"
      :data="roles"
    />
  </div>
</template>

<script lang="ts" setup>
import {h, inject, onMounted, Ref, ref} from 'vue'
import api from '@/api'
import {NButton} from 'naive-ui'

let roles = ref([])
let uuid = inject<Ref>('uuid')
const columns = [
  {
    title: 'ID',
    key: 'id',
    width: 50,
  },
  {
    title: '角色名',
    key: 'name',
    width: 100,
    fixed: 'left',
  },
  {
    title: '创建时间',
    key: 'created_at',
    fixed: 'left',
  },
  {
    title: '绑定用户数',
    key: 'user_count',
    fixed: 'left',
  },
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
  api.role(uuid.value).list().Start(e => {
    roles.value = e
  })
})

</script>

<style scoped>

</style>
