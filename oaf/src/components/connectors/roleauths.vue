<template>
  <div>
    <slot></slot>
    <n-modal @after-leave="emit('update:modelValue', false)" v-model:show="modelValue">
      <n-card class="w-4/5 md:w-1/2 rounded-2xl" :title="role.Name" :bordered="false"
              size="huge">
        <template #header-extra>
        <n-button @click="auths.push({edit: true})">添加权限</n-button>
        </template>
        <div class="grid grid-cols-5 gap-1 gap-y-8" style="line-height: 34px">
          <div>ID</div>
          <div>作用资源</div>
          <div>作用ID</div>
          <div>权限等级</div>
          <div></div>
          <template :key="key" v-for="(item, key) in auths">
          <template v-if="item.edit">
          <div>{{ item.ID }}</div>
          <div>
            <n-select v-model:value="item.ResourceID" :options="RIDOptions"/>
          </div>
          <div>
            <n-input v-model:value="item.RUID"></n-input>
          </div>
          <div>
            <n-select v-model:value="item.Level" :options="levelOptions()"></n-select>
          </div>
          <div>
            <n-button @click="update(item)">保存</n-button>
          </div>
          </template>
          <template v-else>
          <div>{{ item.ID }}</div>
          <div>{{ item.RID }}</div>
          <div>{{ item.RUID }}</div>
          <div>{{ item.Level }}</div>
          <div>
            <n-button @click="item.edit=true">编辑</n-button>
            <n-button @click="del(item.ID, key)">删除</n-button>
          </div>
          </template>
          </template>
        </div>
        <template #footer>
        <n-alert type="warning">
          请谨慎操作， 操作不当可能使所有用户无法正常使用
        </n-alert>
        </template>
      </n-card>
    </n-modal>
  </div>
</template>

<script lang="ts" setup>
import {computed, ref, watch} from 'vue'
import {modelsAuth, modelsResource, modelsRole} from '@/models'
import api from '@/api'
import {useMessage} from 'naive-ui'

let props = withDefaults(defineProps<{
  uuid: string
  role: modelsRole
  res: modelsResource[]
  modelValue: boolean
}>(), {})
let emit = defineEmits<{
  (e: 'update:modelValue', v: boolean): void
}>()

let msg = useMessage()
let id = computed(() => {
  return props.role.ID || 0
})
let value = computed(() => {
  return props.modelValue
})
let auths = ref<modelsAuth[]>([])
let RIDOptions = computed(() => {
  let l = []
  for (let r of props.res) {
    l.push({
      label: r.Name,
      value: r.ID,
    })
  }
  return l
})

let levelOptions = () => {
  let l = []
  for (let i = 0; i < 7; i++) {
    l.push({
      label: i,
      value: i,
    })
  }
  return l
}

watch(value, () => {
  if (id.value > 0 && props.modelValue) {
    api.auth(props.uuid).listOfRole(id.value).Start(e => {
      auths.value = e
    })
  }
})

function del(id: number, index: number) {
  api.auth(props.uuid).del(id).Start(e => {
    auths.value.splice(index, 1)
    msg.success('删除成功')
  })
}

function update(row: modelsAuth) {
  if (row.ID > 0) {
    api.auth(props.uuid).update(row.ID, row.ResourceID, row.RUID, row.Level).Start(e => {
      Object.assign(row, e)
      msg.success('更新成功')
    })
  } else {
    api.auth(props.uuid).create(row.ResourceID, null, id.value, row.RUID, row.Level).Start(e => {
      Object.assign(row, e)
      msg.success('添加成功')
    })
  }
  // @ts-ignore
  row.edit = false
}
</script>

<style scoped>

</style>
