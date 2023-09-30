<template>
  <div>
    <div v-if="ofApps.length > 0">
      <div class="flex justify-between">
        <h1 class="page-h1">我的应用</h1>
        <div class="my-5 mr-10">
          <q-btn @click="new_flag = true" v-if="store.state.user.auth.Get(R.App, '').CanCreate()">创建应用
          </q-btn>
        </div>
      </div>
      <div class="grid gap-4 grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 2xl:grid-cols-5 text-center">
        <div v-for="(item, k) in ofApps" class="flex items-center justify-center" :key="k">
          <AppCard :core="item"></AppCard>
        </div>
      </div>
    </div>
    <div class="mt-20" v-if="apps.length > 0">
      <h1 class="page-h1">应用中心</h1>
      <div class="grid gap-4 grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 2xl:grid-cols-5 text-center">
        <div v-for="(item, k) in apps" class="flex items-center justify-center" :key="k">
          <AppCard :core="item"></AppCard>
        </div>
      </div>
    </div>
    <q-dialog v-model="new_flag">
      <q-card class="w-4/5 md:w-96 rounded-2xl">
        <q-card-section>
          <div class="text-h6">Our Changing Planet</div>
          <div class="text-subtitle2">by John Doe</div>
        </q-card-section>
        <q-separator></q-separator>
        <q-card-section>
          <q-form @submit="create_new">
            <q-input label="应用名" v-model="temp_app.name"></q-input>
            <!-- <uploader url="test.ico" @success="(e) => { -->
            <!--   temp_app.icon = e; -->
            <!-- } -->
            <!--   "> -->
            <!--   <q-avatar size="large" round :src="temp_app.icon"> </q-avatar> -->
            <!-- </uploader> -->
            <div class="flex justify-end">
              <q-btn class="mx-3" @click="new_flag = false">取消</q-btn>
              <q-btn type="submit">创建</q-btn>
            </div>
          </q-form>
        </q-card-section>
      </q-card>
    </q-dialog>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue';
import api from 'src/boot/api';
import msg from '@veypi/msg';
import { modelsApp, modelsUser } from 'src/models';
import { useQuasar } from 'quasar';
import { useUserStore } from 'src/stores/user';
import AppCard from 'components/app.vue'


let apps = ref<modelsApp[]>([]);
let ofApps = ref<modelsApp[]>([]);
let $q = useQuasar()

function getApps() {
  $q.loadingBar.start()
  api.app.list().then(
    (e: modelsApp[]) => {
      apps.value = e;
      api.app
        .user("")
        .list(useUserStore().id)
        .then(
          (e: modelsUser[]) => {
            $q.loadingBar.stop();
            ofApps.value = [];
            console.log(e)
            // for (let i in e) {
            //   let ai = apps.value.findIndex((a) => a.id === e[i]);
            //   if (ai >= 0) {
            //     apps.value[ai].UserStatus = e[i].Status;
            //     if (e[i].Status === "ok") {
            //       ofApps.value.push(apps.value[ai]);
            //       apps.value.splice(ai, 1);
            //     }
            //   }
            // }
          }
        );
    }
  );
}

onMounted(() => {
  getApps();
});

let new_flag = ref(false);
let temp_app = ref({
  name: "",
  icon: "",
});
let form_ref = ref(null);
let rules = {
  name: [
    {
      required: true,
      validator(r: any, v: any) {
        return (
          (v && v.length >= 2 && v.length <= 16) || "长度要求2~16"
        );
      },
      trigger: ["input", "blur"],
    },
  ],
};

function create_new() {
  form_ref.value.validate((e: any) => {
    if (!e) {
      api.app.create(temp_app.value.name, temp_app.value.icon).Start(
        (e) => {
          e.Status = "ok";
          ofApps.value.push(e);
          msg.success("创建成功");
          new_flag.value = false;
        },
        (e) => {
          msg.warning("创建失败: " + e);
        }
      );
    }
  });
}

</script>
