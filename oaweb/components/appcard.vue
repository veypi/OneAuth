 <!--
 * appcard.vue
 * Copyright (C) 2024 veypi <i@veypi.com>
 * 2024-06-07 16:48
 * Distributed under terms of the MIT license.
 -->

<template>
  <div class="core rounded-2xl p-3">
    <div class="grid gap-4 grid-cols-5">
      <div class="col-span-2">
        <img :src="core.icon">
      </div>
      <div class="col-span-3 grid grid-cols-1 items-center text-left">
        <div class="truncate h-10 flex items-center text-xl italic font-bold">
          {{ core.name }}
        </div>
        <span class="truncate">{{ }}</span>
      </div>
    </div>
  </div>
</template>
<script setup lang="ts">
import msg from "@veypi/msg";


const router = useRouter()


let props = withDefaults(defineProps<{
  core: modelsApp,
  is_part: boolean
}>(),
  {}
)


const u = useUserStore()

function Go() {
  if (props.is_part) {
    router.push({ name: "app.home", params: { id: props.core.id } });
    return
  }
  // $q.dialog({
  //   title: '确认',
  //   message: '是否确定申请加入应用 ' + props.core.name,
  //   cancel: true,
  // }).onOk(() => {
  api.app.user(props.core.id).add(u.id).then(e => {
    switch (e.status) {
      case AUStatus.OK:
        msg.Info('加入成功')
        router.push({ name: "app.home", params: { id: props.core.id } });
        return;
      case AUStatus.Applying:
        msg.Info("请等待管理员审批进入");
        return;
      case AUStatus.Deny:
        msg.Warn("进入申请未通过");
        return;
      case AUStatus.Disabled:
        msg.Warn("已被禁止使用");
        return;
    }

  }).catch(e => {
    msg.Warn("加入失败" + e)
  })
}
</script>
<style scoped>
.core {
  width: 256px;
  background: rgba(146, 145, 145, 0.1);
}
</style>
