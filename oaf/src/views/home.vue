<template>
  <div>
    <div v-if="ofApps.length > 0">
      <div class="flex justify-between">
        <h1 class="page-h1">我的应用</h1>
        <div class="my-5 mr-10">
          <n-button
            @click="new_flag = true"
            v-if="store.state.user.auth.Get(R.App, '').CanCreate()"
            >创建应用
          </n-button>
        </div>
      </div>
      <div
        class="grid gap-4 grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 2xl:grid-cols-5 text-center"
      >
        <div
          v-for="(item, k) in ofApps"
          class="flex items-center justify-center"
          :key="k"
        >
          <AppCard :core="item"></AppCard>
        </div>
      </div>
    </div>
    <div class="mt-20" v-if="apps.length > 0">
      <h1 class="page-h1">应用中心</h1>
      <div
        class="grid gap-4 grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 2xl:grid-cols-5 text-center"
      >
        <div
          v-for="(item, k) in apps"
          class="flex items-center justify-center"
          :key="k"
        >
          <AppCard :core="item"></AppCard>
        </div>
      </div>
    </div>
    <n-modal v-model:show="new_flag">
      <n-card
        class="w-4/5 md:w-96 rounded-2xl"
        title="创建应用"
        :bordered="false"
        size="huge"
      >
        <n-form
          label-width="70px"
          label-align="left"
          :model="temp_app"
          ref="form_ref"
          label-placement="left"
          :rules="rules"
        >
          <n-form-item required label="应用名" path="name">
            <n-input v-model:value="temp_app.name"></n-input>
          </n-form-item>
          <n-form-item required label="icon" path="icon">
            <uploader
              url="test.ico"
              @success="
                (e) => {
                  temp_app.icon = e;
                }
              "
            >
              <n-avatar size="large" round :src="temp_app.icon"> </n-avatar>
            </uploader>
          </n-form-item>
        </n-form>
        <template #footer>
          <div class="flex justify-end">
            <n-button class="mx-3" @click="new_flag = false">取消</n-button>
            <n-button @click="create_new">创建</n-button>
          </div>
        </template>
      </n-card>
    </n-modal>
  </div>
</template>

<script lang="ts" setup>
import { onMounted, ref } from "vue";
import api from "@/api";
import AppCard from "@/components/app.vue";
import { useStore } from "@/store";
import { R } from "@/auth";
import { useMessage, useLoadingBar } from "naive-ui";
import { modelsApp } from "@/models";
import Uploader from "@/components/uploader";

let msg = useMessage();
let bar = useLoadingBar();
let store = useStore();
let apps = ref<modelsApp[]>([]);
let ofApps = ref<modelsApp[]>([]);

function getApps() {
  bar.start();
  api.app.list().Start(
    (e) => {
      apps.value = e;
      api.app
        .user("")
        .list(store.state.user.id)
        .Start(
          (e) => {
            bar.finish();
            ofApps.value = [];
            for (let i in e) {
              let ai = apps.value.findIndex((a) => a.UUID === e[i].AppUUID);
              if (ai >= 0) {
                apps.value[ai].UserStatus = e[i].Status;
                if (e[i].Status === "ok") {
                  ofApps.value.push(apps.value[ai]);
                  apps.value.splice(ai, 1);
                }
              }
            }
          },
          () => {
            bar.error();
          }
        );
    },
    () => bar.error()
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
          (v && v.length >= 2 && v.length <= 16) || new Error("长度要求2~16")
        );
      },
      trigger: ["input", "blur"],
    },
  ],
};

function create_new() {
  // @ts-ignore
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

<style scoped></style>
