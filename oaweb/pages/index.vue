 <!--
 * index.vue
 * Copyright (C) 2024 veypi <i@veypi.com>
 * 2024-05-31 17:10
 * Distributed under terms of the MIT license.
 -->

<template>
  <div>

    <div v-if="ofApps.length > 0">
      <div class="flex justify-between">
        <h1 class="page-h1">我的应用</h1>
        <div class="my-5 mr-10">
          <div class='vbtn bg-gray-400' @click="new_flag = true" v-if="user.auth.Get(R.App, '').CanCreate()">创建应用
          </div>
        </div>
      </div>
      <div class="grid gap-4 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 2xl:grid-cols-5 text-center">
        <div v-for="(item, k) in ofApps" class="flex items-center justify-center" :key="k">
          <Appcard :core="item" :is_part="true"></Appcard>
        </div>
      </div>
    </div>
    <div class="mt-20" v-if="apps.length > 0">
      <h1 class="page-h1">应用中心</h1>
      <div class="grid gap-4 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 2xl:grid-cols-5 text-center">
        <div v-for="(item, k) in apps" class="flex items-center justify-center" :key="k">
          <Appcard :core="item" :is_part="false"></Appcard>
        </div>
      </div>
    </div>
    <UModal v-model="new_flag">
      <div class="p-4">
        123
      </div>
    </UModal>
    <!-- <q-dialog :square="false" v-model="new_flag"> -->
    <!--   <q-card class="w-4/5 md:w-96 rounded-2xl"> -->
    <!--     <q-card-section> -->
    <!--       <div class="text-h6">创建应用</div> -->
    <!--     </q-card-section> -->
    <!--     <q-separator></q-separator> -->
    <!--     <q-card-section> -->
    <!--       <q-form @submit="create_new"> -->
    <!--         <q-input label="应用名" v-model="temp_app.name" :rules="rules.name"></q-input> -->
    <!--         <div class="flex justify-center my-4 items-center" label='icon'> -->
    <!--           <uploader @success="temp_app.icon = $event" dir="app_icon"> -->
    <!--             <q-avatar> -->
    <!--               <img :src="temp_app.icon"> -->
    <!--             </q-avatar> -->
    <!--           </uploader> -->
    <!--           <q-icon class="ml-2" size="1rem" name='autorenew' @click="temp_app.icon = rand_icon()"></q-icon> -->
    <!--         </div> -->


    <!--         <q-separator></q-separator> -->
    <!--         <div class="flex justify-end mt-8"> -->
    <!--           <q-btn class="mx-3" @click="new_flag = false">取消</q-btn> -->
    <!--           <q-btn type="submit">创建</q-btn> -->
    <!--         </div> -->
    <!--       </q-form> -->
    <!--     </q-card-section> -->
    <!--   </q-card> -->
    <!-- </q-dialog> -->
  </div>
</template>

<script setup lang="ts">
import type { models } from '#imports';
import msg from '@veypi/msg';

let user = useUserStore()


let apps = ref<models.App[]>([]);
let ofApps = ref<models.App[]>([]);

function getApps() {
  api.app.List({}).then(
    (e) => {
      apps.value = e;
      api.app.AppUserList("*", { user_id: user.id }).then(aus => {
        for (let i in aus) {
          let ai = apps.value.findIndex(a => a.id === aus[i].app_id)
          if (ai >= 0) {
            if (aus[i].status === "ok") {
              ofApps.value.push(apps.value[ai])
              apps.value.splice(ai, 1)
            }
          }
        }
      })
    }
  );
}


const rand_icon = () => {
  return "/media/icon/sign/scenery-" + util.randomNum(1, 20) + ".png"
}
let new_flag = ref(false);
let temp_app = ref({
  name: "",
  icon: rand_icon()
});
let rules = {
  name: [
    (v: string) => (v && v.length >= 2 && v.length <= 16) || "长度要求2~16"
  ],
};

function create_new() {
  api.app.Post({
    name: temp_app.value.name, icon: temp_app.value.icon,
    des: "", participate: "auto"
  }).then((e: models.App) => {
    ofApps.value.push(e);
    msg.Info("创建成功");
    new_flag.value = false;
  }).catch(e => {
    msg.Warn("创建失败: " + e);
  })
}
onMounted(() => {
  getApps();
});

</script>
