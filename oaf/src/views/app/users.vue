<template>
  <div>
    <h1 class="page-h1">用户名单</h1>
    <n-data-table
      :bordered="false"
      :columns="columns"
      :data="users"
    />
  </div>
</template>

<script lang="ts" setup>
import {inject, onMounted, ref, h} from "vue";
import api from "@/api";
import {NTag} from 'naive-ui'

let uuid: any = inject('uuid')
let users = ref([])

onMounted(() => {
  api.app.user(uuid.value as string).list(0).Start(e => {
    users.value = e
    console.log(e)
  })
})

const columns = [
  {
    title: 'ID',
    key: 'user_id',
    width: 50
  },
  {
    title: '用户',
    key: 'user.username',
    width: 200,
    fixed: 'left'
  },
  {
    title: '加入时间',
    key: 'user.created_at',
    fixed: 'left'
  },
  {
    title: 'Status',
    key: 'status',
    render(row) {
      let t = statusTag(row.status)
      return h(NTag,{
          'type': t[1]
        },
        {
          default: () => t[0]
        }
      )
    }
  }
]

function statusTag(s: string) {
  switch (s) {
    case 'ok':
      return ['正常', 'success']
    case "apply":
      return ['申请中', 'info']
    case 'deny':
      return ['拒绝', '']
    case 'disabled':
      return ['禁用', 'warning']
  }
  return ['未知', '']
}

</script>

<style scoped>

</style>
