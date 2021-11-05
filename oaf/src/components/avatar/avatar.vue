<template>
  <base_frame style="line-height:40px" v-model="shown" :isDark="IsDark">
    <div class="flex">
      <n-avatar :src="$store.state.user.icon" round></n-avatar>
    </div>
    <template v-slot:main>
    <div style="height: 100%">
      <div style="height: calc(100% - 50px)">
        <div class="w-full px-3">
          <div class="h-16 flex justify-between items-center">
            <span style="color: #777">我的账户</span>
            <span @click="$router.push({name: 'user_setting'});shown=false" class="cursor-pointer"
                  style="color:#f36828">账户中心</span>
          </div>
          <div class="grid grid-cols-4 gap-4 h-20">
            <div class="flex items-center justify-center">
              <n-avatar size="50" :src="$store.state.user.icon" round></n-avatar>
            </div>
            <div class="col-span-2 text-xs grid grid-cols-1 items-center" style="">
              <span>昵称: &ensp;&ensp; {{ $store.state.user.nickname }}</span>
              <span>账户: &ensp;&ensp; {{ $store.state.user.username }}</span>
              <span>邮箱: &ensp;&ensp; {{ $store.state.user.email }}</span>
            </div>
            <div class="">123</div>
          </div>
          <hr class="mt-10" style="border:none;border-top:1px solid #777;">
        </div>
      </div>
      <hr style="border:none;border-top:2px solid #777;">
      <div style="height: 48px">
        <div @click="$store.commit('user/logout')"
             class="w-full h-full flex justify-center items-center cursor-pointer transition duration-500 ease-in-out transform hover:scale-125">
          <one-icon :color="IsDark?'#eee': '#333'" class="inline-block" style="font-size: 24px;">
            logout
          </one-icon>
          <div>
            退出登录
          </div>
        </div>
      </div>
    </div>
    </template>
  </base_frame>
</template>

<script lang="ts" setup>
import base_frame from './frame.vue'
import {IsDark} from '../../theme'
import {ref} from "vue";

let shown = ref(false)

function asd(e) {
  console.log([e, shown.value])
}

</script>

<style scoped>

</style>
