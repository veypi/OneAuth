 <!--
 * AppAuth.vue
 * Copyright (C) 2023 veypi <i@veypi.com>
 * 2023-10-09 23:18
 * Distributed under terms of the MIT license.
 -->
<template>
  <div class="">
    <div class="flex justify-between">
      <div class="text-3xl">角色管理</div>
      <q-btn @click="created(0)">创建角色</q-btn>
    </div>
    <div class="w-full">
      <crud vertical :keys="role_keys" :data="role" @update="update(0,
        $event)">
        <template #k_created="{ value }">{{ util.datetostr(value) }}</template>
        <template #k__="{ row }">
          <q-btn color="primary" size="sm" @click="show_dialog(0,
            row.id)">权限</q-btn>
          <q-btn color="secondary" size="sm" @click="show_dialog(1,
            row.id)">用户</q-btn>
          <q-btn color="negative" size="sm" @click="del(0, row.id)">删除</q-btn>
          <template v-if="row.id === app.role_id">
            <q-btn color="positive" disable size="sm">初始角色</q-btn>
          </template>
          <template v-else>
            <q-btn size="sm" @click="api.app.update(app.id, {
              role_id:
                row.id
            }).then(_ => app.role_id = row.id)">设置为初始角色</q-btn>
          </template>
        </template>
      </crud>
    </div>

    <div class="flex mt-16 justify-between">
      <div class="text-3xl">资源管理</div>
      <q-btn @click="created(1)">创建资源</q-btn>
    </div>
    <div class="w-full">
      <crud vertical :keys="resource_keys" :data="resource" @update="update(1,
        $event)">
        <template #k_created="{ value }">{{ util.datetostr(value) }}</template>
        <template #k__="{ row }">
          <q-btn color="negative" size="sm" @click="del(1, row.name)">删除</q-btn>
        </template>
      </crud>
    </div>

    <q-dialog v-model="dialog">
      <q-card v-if="dialog_obj" class="mx-4 mt-4" style="width: 1000px;max-width: 90vw;">
        <q-card-section class="row items-center q-pb-none">
          <div class="text-h6">{{ dialog_obj.name }}
          </div>
          <q-space />
          <div v-if="dialog_mode">
            <q-select filled :model-value="''" @update:model-value="roleuser.add" use-input hide-selected fill-input
              input-debounce="0" label="添加用户" :options="users_cache" @filter="filterFn" style="width: 20rem">
              <template v-slot:no-option>
                <q-item>
                  <q-item-section class="text-grey">
                    无结果
                  </q-item-section>
                </q-item>
              </template>
            </q-select>
          </div>
          <q-btn v-else @click="created(2, { role_id: dialog_obj.id })">添加权限</q-btn>
        </q-card-section>

        <q-card-section v-if="dialog_mode">
          <crud vertical :keys="users_keys" :data="users">
            <template #k__="{ row }">
              <q-btn color="negative" size="sm" @click="roleuser.drop(row)">删除</q-btn>
            </template>
          </crud>
        </q-card-section>
        <q-card-section v-else>
          <crud vertical :keys="access_keys" :data="access" @update="update(2, $event)">
            <template #k__="{ row }">
              <q-btn color="negative" size="sm" @click="del(2, row.id)">删除</q-btn>
            </template>
          </crud>
        </q-card-section>
      </q-card>
    </q-dialog>
  </div>
</template>

<script lang="ts" setup>
import msg from '@veypi/msg';
import { useQuasar } from 'quasar';
import api from 'src/boot/api';
import crud from 'src/components/crud.vue';
import { util } from 'src/libs';
import {
  ArgType, modelsAccess, modelsApp, modelsRole, modelsUser,
  LevelOptions
} from 'src/models';
import { computed, inject, onMounted, Ref, ref, watch } from 'vue';

let $q = useQuasar();
let dialog = ref(false)
let dialog_mode = ref(0)
let dialog_obj = ref<modelsRole>({} as any)
const role_keys = ref<any>([
  {
    name: 'id',
    label: 'ID',
  },
  { name: 'name', label: '标识符' },
  { name: 'des', label: '描述', editable: true },
  { name: 'created', label: '创建时间' },
  { name: 'user_count', label: '绑定用户数' },
  {
    name: '_', label: '操作', style: { 'justify-content': 'start' },
    width: 40
  },
])
const resource_keys = ref<any>([
  { name: 'name', label: '标识符' },
  { name: 'des', label: '描述', editable: true },
  { name: 'created', label: '创建时间' },
  {
    name: '_', label: '操作', style: { 'justify-content': 'start' },
    width: 40
  },
])
const access_keys = ref<any>([
  { name: 'id', label: 'ID' },
  { name: 'name', label: '标识符' },
  { name: 'rid', label: '特定子资源id', editable: true },
  {
    name: 'level', label: '权限等级', editable: true, typ:
      ArgType.Select, options: LevelOptions,
  },
  {
    name: '_', label: '操作', style: { 'justify-content': 'start' },
    width: 40
  },
])
const users_keys = ref<any>([
  { name: 'id', label: 'ID' },
  { name: 'username', label: '用户名' },
  { name: 'nickname', label: '昵称' },
  {
    name: '_', label: '操作', style: { 'justify-content': 'start' },
    width: 40
  },
])

let access = ref<modelsAccess[]>([])
let users = ref<modelsUser[]>([])
let users_cache = ref<any[]>([])
const resource = ref<any[]>([])
const role = ref<any[]>([])
const app = inject('app') as Ref<modelsApp>
watch(computed(() => app.value.id), (v) => {
  sync(v)

})

const sync = (id: any) => {
  if (!id) {
    return
  }
  api.role(id).list().then(e => {
    role.value = e
  })
  api.resource(id).list().then(e => {
    resource.value = e
  })
}

const show_dialog = (mode: number, idx: string) => {
  dialog_obj.value = role.value.find(e => e.id === idx)
  dialog.value = true
  dialog_mode.value = mode
  if (mode) {
    api.user.list({ role_id: idx }).then(e => {
      users.value = e
    })
  } else {
    api.access(app.value.id).list({ role_id: idx }).then(e => {
      access.value = e
    })
  }
}

// 0: role 1: resource 2: access 3: users
const crud_option = (mode: number) => {
  let res = {
    api: api.role(app.value.id),
    lable: '角色',
    obj: role,
  }
  if (mode === 1) {
    res.api = api.resource(app.value.id) as any
    res.lable = '资源'
    res.obj = resource
  } else if (mode === 2) {
    res.api = api.access(app.value.id) as any
    res.lable = '权限'
    res.obj = access
  } else if (mode === 3) {
    res.api = api.user as any
    res.lable = '用户'
    res.obj = users
  }
  return res
}
const created = (k: number, props?: any) => {
  let opt = crud_option(k)
  let options;
  let prompt;
  if (k !== 2) {
    prompt = { model: '', type: 'text' }
  } else {
    console.log(k)
    options = {
      model: '', type: 'radio', items: resource.value.map(e => {
        return {
          label: e.name,
          value: e.name
        }
      })
    }
  }
  $q.dialog({
    title: '创建' + opt.lable,
    message: '请输入标识码',
    prompt: prompt as any,
    options: options as any,
    cancel: true,
    persistent: true
  }).onOk(data => {
    opt.api.create(data, props).then(e => {
      msg.Info('创建成功')
      opt.obj.value.push(e)
    }).catch(e => {
      msg.Warn('创建失败： ' + e)
    })
  })
}

const update = (k: number, props: any[]) => {
  let opt = crud_option(k)
  console.log(props)
  for (let i in props) {
    let id = opt.obj.value[i][k === 1 ? 'name' : 'id']
    console.log(id)
    opt.api.update(id, props[i]).then(() => {
      Object.assign(opt.obj.value[i], props[i])
    }).catch(e => {
      msg.Warn('更新失败： ' + e)
    })
  }
}

const del = (k: number, id: string) => {
  let opt = crud_option(k)
  $q.dialog({
    title: '是否确定删除',
    message: '',
    cancel: true,
    persistent: true
  }).onOk(() => {
    opt.api.del(id).then(e => {
      msg.Info('删除成功')
      opt.obj.value.splice(opt.obj.value.findIndex(e => e.name === id
        || e.id === id), 1)
    }).catch(e => {
      msg.Warn('删除失败： ' + e)
    })
  })
}

const roleuser = {
  add: (u: any) => {
    let idx = users.value.findIndex(e => e.id == u.id)
    if (idx >= 0) {
      return
    }
    api.role(app.value.id).add(dialog_obj.value?.id || '', u.id).then(e => {
      users.value.push(u)
      dialog_obj.value.user_count = dialog_obj.value?.user_count + 1
    })
  },
  drop: (u: any) => {
    api.role(app.value.id).drop(dialog_obj.value?.id || '', u.id).then(e => {
      users.value.splice(users.value.findIndex(e => e.id === u.id), 1)
      dialog_obj.value.user_count = dialog_obj.value?.user_count - 1
    })
  }
}

const filterFn = (val: string, cb: any) => {
  if (val && val.length > 1) {
    api.user.list({ name: val }).then((e: modelsUser[]) => {
      cb(() => {
        users_cache.value = e.map(i => {
          return Object.assign({
            label: i.username,
            value: i.id,
          }, i)
        })
      })
    })
  }
}
onMounted(() => {
  sync(app.value.id)
})
</script>

<style scoped></style>

