<template>
  <div>
    <slot>
    </slot>
    <n-modal @after-leave="emit('update:modelValue', false)" v-model:show="modelValue">
      <n-card class="w-4/5 md:w-1/2 rounded-2xl" :title="role.Name" :bordered="false"
              size="huge">
        <template #header-extra>
        <UserSelect @selected="tmp=$event"></UserSelect>
        <n-button @click="add">添加用户</n-button>
        </template>
        <div class="grid grid-cols-4 gap-1 gap-y-8" style="line-height: 34px">
          <div>ID</div>
          <div>昵称</div>
          <div>用户名</div>
          <div></div>
          <template :key="key" v-for="(item, key) in users">
          <div>{{ item.ID }}</div>
          <div>{{ item.Nickname }}</div>
          <div>{{ item.Username }}</div>
          <div>
            <n-button @click="del(item.ID, key)">删除</n-button>
          </div>
          </template>
        </div>
        <template #footer>
        </template>
      </n-card>
    </n-modal>
  </div>
</template>

<script lang="ts" setup>
import {modelsRole, modelsUser} from '@/models'
import {computed, ref, watch} from 'vue'
import api from '@/api'
import {useMessage} from 'naive-ui'
import UserSelect from '@/components/userSelect.vue'

let msg = useMessage()
let props = withDefaults(defineProps<{
  uuid: string
  role: modelsRole
  modelValue: boolean
}>(), {})
let emit = defineEmits<{
  (e: 'update:modelValue', v: boolean): void
}>()
let id = computed(() => {
  return props.role.ID || 0
})
let value = computed(() => {
  return props.modelValue
})
let users = ref<modelsUser[]>([])

function del(uid: number, index: number) {
  props.role.UserCount --
  api.role(props.uuid).user(id.value).delete(uid).Start(e => {
    users.value.splice(index, 1)
    msg.success('删除成功')
  })
}

let tmp = ref<modelsUser>(null)

function add() {
  if (tmp.value && tmp.value.ID > 0) {
    api.role(props.uuid).user(id.value).create(tmp.value.ID).Start(e => {
      let added = false
      for (let u of users.value) {
        if (u.ID === tmp.value.ID) {
          added = true
        }
      }
      if (!added) {
        users.value.push(tmp.value)
        props.role.UserCount++
      }
    })
  }
}

watch(value, () => {
  if (id.value > 0 && props.modelValue) {
    api.role(props.uuid).user(id.value).list().Start(e => {
      users.value = e
    })
  }
})
</script>

<style scoped>

</style>
