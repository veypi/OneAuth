 <!--
 * app.[id].vue
 * Copyright (C) 2024 veypi <i@veypi.com>
 * 2024-06-07 17:46
 * Distributed under terms of the MIT license.
 -->
<template>
  <div>
    <div class="flex justify-start items-center gap-4">
      <img class="w-12 h-12 rounded-full" :src="core.icon">
      <div class="text-3xl">
        {{ core.name }}
      </div>
    </div>
    <NuxtPage keepalive :core='core' :page-key="route => route.fullPath"></NuxtPage>
  </div>
</template>

<script lang="ts" setup>
import msg from '@veypi/msg';

const route = useRoute()
const router = useRouter()

let menu = useMenuStore()
let core = ref({} as modelsApp)
const set_menu = () => {
  let p = '/app/' + core.value.id
  menu.set([
    { ico: 'home', name: core.value.name, path: p },
    { ico: 'user', name: '用户管理', path: p + '/user' },
    { ico: 'team', name: '权限设置', path: p + '/auth' },
    { ico: 'setting', name: '应用设置', path: p + '/cfg' },
  ])
}

onMounted(() => {
  api.app.get(route.params.id as string).then((e) => {
    core.value = e
    set_menu()
  }).catch(e => {
    msg.Warn('获取数据失败: ' + e)
    router.push('/')
  })
})


onActivated(() => {
  set_menu()
})
onBeforeRouteLeave(() => {
  menu.default()
})

</script>

<style scoped></style>

