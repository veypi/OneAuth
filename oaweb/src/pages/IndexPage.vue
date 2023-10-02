<template>
  <div>
    <div v-if="ofApps.length > 0">
      <div class="flex justify-between">
        <h1 class="page-h1">我的应用</h1>
        <div class="my-5 mr-10">
          <q-btn @click="new_flag = true" v-if="user.auth.Get(R.App, '').CanCreate()">创建应用
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
    <q-dialog :square="false" v-model="new_flag">
      <q-card class="w-4/5 md:w-96 rounded-2xl">
        <q-card-section>
          <div class="text-h6">创建应用</div>
        </q-card-section>
        <q-separator></q-separator>
        <q-card-section>
          <q-form @submit="create_new">
            <q-input label="应用名" v-model="temp_app.name" :rules="rules.name"></q-input>
            <q-field label="icon" stack-label>
              <template v-slot:control>
                <uploader url="test.ico" @success="(e) => {
                  temp_app.icon = e;
                }
                  ">
                  <q-avatar size="xl" round>
                    <img :src="temp_app.icon">
                  </q-avatar>
                </uploader>
              </template>
            </q-field>
            <q-separator></q-separator>
            <div class="flex justify-end mt-8">
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
import { AUStatus, modelsApp, modelsAppUser } from 'src/models';
import AppCard from 'components/app.vue'
import { useUserStore } from 'src/stores/user';
import { R } from 'src/models';
import uploader from 'components/uploader'

let user = useUserStore()


let apps = ref<modelsApp[]>([]);
let ofApps = ref<modelsApp[]>([]);

function getApps() {
  api.app.list().then(
    (e: modelsApp[]) => {
      apps.value = e;
      api.app.user('-').list(user.id).then((aus: modelsAppUser[]) => {
        for (let i in aus) {
          let ai = apps.value.findIndex(a => a.id === aus[i].app_id)
          if (ai >= 0) {
            apps.value[ai].au = aus[i]
            if (aus[i].status === AUStatus.OK) {
              ofApps.value.push(apps.value[ai])
              apps.value.splice(ai, 1)
            }
          }
        }
      })
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
let rules = {
  name: [
    (v: string) => (v && v.length >= 2 && v.length <= 16) || "长度要求2~16"
  ],
};

function create_new() {
  api.app.create(temp_app.value.name, temp_app.value.icon).then((e:
    modelsApp) => {
    console.log(e)
    // e.Status = "ok";
    // ofApps.value.push(e);
    msg.Info("创建成功");
    new_flag.value = false;
  }).catch(e => {
    msg.Warn("创建失败: " + e);
  })
}

</script>
