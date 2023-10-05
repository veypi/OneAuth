<template>
  <div class="core rounded-2xl p-3">
    <div class="grid gap-4 grid-cols-5">
      <div class="col-span-2">
        <q-avatar class="cursor-pointer" style="--color: none" @click="Go" round size="5rem">
          <img :src="core.icon">
        </q-avatar>
      </div>
      <div class="col-span-3 grid grid-cols-1 items-center text-left">
        <div class="h-10 flex items-center text-2xl italic font-bold">
          {{ core.name }}
        </div>
        <span class="truncate">{{ core.des }}</span>
      </div>
    </div>
  </div>
</template>
<script setup lang="ts">
import msg from "@veypi/msg";
import api from "src/boot/api";
import { AUStatus, modelsApp, modelsAppUser } from "src/models";
import { useUserStore } from "src/stores/user";
import { useRouter } from "vue-router"


const router = useRouter()


let props = withDefaults(defineProps<{
  core: modelsApp
}>(),
  {}
)



function Go() {
  switch (props.core.au.status) {
    case AUStatus.OK:
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
  // api.app.user(props.core.id).add(useUserStore().id).then(e => {
  //   console.log(e)
  // })
  // api.app
  //   .user(props.core.UUID)
  //   .add(store.state.user.id)
  //   .Start(
  //     (e) => {
  //       bar.finish();
  //       if (e.Status === "ok") {
  //         router.push({ name: "app.main", params: { uuid: props.core.UUID } });
  //         return;
  //       }
  //       props.core.UserStatus = e.Status;
  //       msg.info("已发起加入申请");
  //     },
  //     (e) => {
  //       msg.warning("加入失败: " + e);
  //       bar.error();
  //     }
  //   );
  // return;
}
</script>
<style scoped>
.core {
  width: 256px;
  background: rgba(146, 145, 145, 0.1);
}
</style>
