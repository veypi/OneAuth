<template>
  <div>
    <slot></slot>
    <n-modal @after-leave="emit('update:modelValue',false)" v-model:show="modelValue">
      <n-card class="w-4/5 md:w-1/2 rounded-2xl" :title="res.ID > 0 ? res.Name:' '" :bordered="false"
              size="huge">
        <template #header-extra>{{ res.ID > 0 ? '编辑' : '创建' }}</template>
        <div class="grid grid-cols-5 gap-1 gap-y-8" style="line-height: 34px">
          <div>角色名</div>
          <div class="col-span-4">
            <n-input :disabled="res.ID> 0" v-model:value="res.Name"></n-input>
          </div>
          <div>角色标签</div>
          <div class="col-span-4">
            <n-input type="textarea" v-model:value="res.Tag"></n-input>
          </div>
        </div>
        <template #footer>
        <div class="flex justify-end">
          <n-button class="mx-3" @click="emit('update:modelValue', false)">取消</n-button>
          <n-button @click="update">{{ res.ID > 0 ? '更新' : '创建' }}</n-button>
        </div>
        </template>
      </n-card>
    </n-modal>
  </div>
</template>

<script lang="ts" setup>
import {modelsRole} from '@/models'
import api from '@/api'
import {useMessage} from 'naive-ui'

let props = withDefaults(defineProps<{
  res: modelsRole
  modelValue: boolean
  uuid: string
}>(), {
  res: {} as any,
  modelValue: false,
  uuid: '',
})

let msg = useMessage()
let emit = defineEmits<{
  (e: 'update:modelValue', v: boolean): void
  (e: 'ok', v: modelsRole)
}>()

function update() {
  if (props.res.ID > 0) {
    emit('update:modelValue', false)
    api.role(props.uuid).update(props.res.ID, props.res).Start(e => {
      msg.success('更新成功')
    })
    return
  }
  api.role(props.uuid).create(props.res.Name, props.res.Tag).Start(e => {
    msg.success('添加成功')
    emit('ok', e)
    emit('update:modelValue', false)
  })
}
</script>

<style scoped>

</style>
