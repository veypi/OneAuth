 <!--
 * AppUser.vue
 * Copyright (C) 2023 veypi <i@veypi.com>
 * 2023-10-06 20:44
 * Distributed under terms of the MIT license.
 -->
<template>
  <div>
    <q-table title="Treats" :rows="rows" :columns="columns" row-key="name">
      <template v-slot:body-cell-status="props">
        <q-td :props="props">
          <div>
            <q-chip outline :color="statusOpts[props.row.status][1]" text-color="white" icon="event">
              {{ statusOpts[props.row.status][0] }}
            </q-chip>
          </div>
          <q-popup-edit v-model="props.row.status" v-slot="scope" buttons
            @save="update_status(props.row.id, $event, props.row.status)" label-set="确定" label-cancel="取消">
            <div class="mt-4 mb-2">切换状态至</div>
            <div class="q-gutter-sm">

              <q-radio :key="i" v-for="i in [0, 1, 2, 3]" keep-color v-model="scope.value" :val="i"
                :label="statusOpts[i][0]" :color="statusOpts[i][1]" />

            </div>
          </q-popup-edit>
        </q-td>
      </template>
    </q-table>
  </div>
</template>

<script lang="ts" setup>
import { onMounted, ref } from 'vue';
import { AUStatus, modelsAppUser } from 'src/models';
import { QTableProps } from 'quasar';
import api from 'src/boot/api';
import { useAppStore } from 'src/stores/app';

const statusOpts: { [index: number]: any } = {
  [AUStatus.OK]: ['正常', 'positive'],
  [AUStatus.Deny]: ['拒绝', 'warning'],
  [AUStatus.Applying]: ['申请中', 'primary'],
  [AUStatus.Disabled]: ['禁用', 'warning'],
}
const columns = [
  {
    name: 'id',
    required: true,
    field: 'id',
    label: 'ID',
    align: 'center',
  },
  {
    name: 'name',
    label: '用户名',
    align: 'center',
    field: (row: any) => row.username +
      (row.nickname ? '(' + row.nickname + ')' : ''),
    sortable: true
  },
  { name: 'created', field: 'created', align: 'center', label: '加入时间', sortable: true },
  { name: 'status', field: 'status', align: 'center', label: '账号状态', sortable: true },
  { name: 'action', field: 'action', align: 'center', label: '操作' },
] as any

const rows = ref([
  {
    id: '1', username: 'name', nickname: 'asd', created: 'asdsss',
    status: 0
  }
])

const update_status = (d, s) => {

  console.log([d, s])
}

onMounted(() => {
  api.app.user(useAppStore().id).list('-', { user: true }).then((e:
    modelsAppUser[]) => {
    rows.value = e.map(i => i.user)
  })
})
</script>

<style scoped></style>

