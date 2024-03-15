 <!--
 * AppUser.vue
 * Copyright (C) 2023 veypi <i@veypi.com>
 * 2023-10-06 20:44
 * Distributed under terms of the MIT license.
 -->
<template>
  <div class="">
    <q-table title="用户管理" :rows="rows" :columns="columns" row-key="name">
      <template #body-cell-created="props">
        <q-td :props="props">
          {{ util.datetostr(props.value) }}
        </q-td>
      </template>
      <template v-slot:body-cell-status="props">
        <q-td :props="props">
          <div>
            <q-chip outline :color="auOpts[props.row.au][1]" text-color="white" icon="event">
              {{ auOpts[props.row.au][0] }}
            </q-chip>
          </div>
          <q-popup-edit v-model="props.row.au" v-slot="scope" buttons
            @save="update_status(props.row.id, $event, props.row.au)" label-set="确定" label-cancel="取消">
            <div class="mt-4 mb-2">切换状态至</div>
            <div class="q-gutter-sm">

              <q-radio :key="i" v-for="i in [0, 1, 2, 3]" keep-color v-model="scope.value" :val="i" :label="auOpts[i][0]"
                :color="auOpts[i][1]" />

            </div>
          </q-popup-edit>
        </q-td>
      </template>
      <template #body-cell-action="props">

        <q-td :props="props">
          <q-btn v-if='cfg.id === app.id' size='sm' color='secondary' @click='reset(props.row.id)'>重置密码</q-btn>
        </q-td>
      </template>
    </q-table>
  </div>
</template>

<script lang="ts" setup>
import { computed, inject, onMounted, Ref, ref, watch } from 'vue';
import { AUStatus, modelsAppUser, modelsUser, modelsApp } from 'src/models';
import api from 'src/boot/api';
import msg from '@veypi/msg';
import { util } from 'src/libs';
import cfg from 'src/cfg';
import { useQuasar } from 'quasar';

const auOpts: { [index: number]: any } = {
  [AUStatus.OK]: ['正常', 'positive'],
  [AUStatus.Deny]: ['拒绝', 'warning'],
  [AUStatus.Applying]: ['申请中', 'primary'],
  [AUStatus.Disabled]: ['禁用', 'warning'],
}
let $q = useQuasar()
let app = inject('app') as Ref<modelsApp>
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
  { name: 'created', field: 'updated', align: 'center', label: '加入时间', sortable: true },
  { name: 'status', field: 'status', align: 'center', label: '账号状态', sortable: true },
  { name: 'action', field: 'action', align: 'center', label: '操作' },
] as any

const rows = ref([] as modelsUser[])

const update_status = (id: string, n: number, old: number) => {
  api.app.user(app.value.id).update(id, n).then(e => {
    msg.Info('修改成功')
  }).catch(e => {
    let a = rows.value.find(a => a.id = id) || {} as any
    a.status = old
  })

  console.log([id, n, old])
}

const reset = (id: string) => {
  $q.dialog({
    title: '是否重置密码',
    message: '请谨慎操作',
    cancel: true,
    persistent: true
  }).onOk(() => {
    api.user.reset(id).then(e => {
      msg.Info('重置成功 ')
    }).catch(e => {
      msg.Warn('失败 ' + e)
    })
  })
}

const sync = () => {
  if (!app.value.id) {
    return
  }
  api.app.user(app.value.id).list('-', { user: true }).then((e:
    modelsAppUser[]) => {
    rows.value = e.map(i => {
      i.user.au = i.status
      return i.user
    })
  })
}
watch(computed(() => app.value.id), () => sync())

onMounted(() => {
  sync()
})
</script>

<style scoped></style>

