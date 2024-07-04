 <!--
 * appcard.vue
 * Copyright (C) 2024 veypi <i@veypi.com>
 * 2024-06-07 16:48
 * Distributed under terms of the MIT license.
 -->

<template>
  <div class="core rounded-2xl" @click="Go">
    <img class="logo rounded-full" :src="core.icon">
    <div class="txt">
      <div class="title truncate italic font-bold">{{ core.name }}</div>
      <div class="des"></div>
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
    router.push('/app/' + props.core.id)
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
<style scoped lang="scss">
.core {
  user-select: none;
  cursor: pointer;
  width: 256px;
  height: 96px;
  padding: 12px;
  background: var(--base-bg-1);
  filter: brightness(1.05);

  .logo {
    vertical-align: top;
    display: inline-block;
    width: 72px;
    height: 72px;
  }

  .txt {
    vertical-align: top;
    display: inline-block;
    width: 160px;
    height: 72px;

    .title {
      padding-left: 12px;
      text-align: left;
      width: 160px;
      height: 48px;
      line-height: 48px;
      font-size: 24px;
    }

    .des {
      height: 24px;
      padding-left: 12px;
    }
  }
}
</style>
