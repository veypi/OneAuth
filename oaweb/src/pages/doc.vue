 <!--
 * doc.vue
 * Copyright (C) 2023 veypi <i@veypi.com>
 * 2023-10-07 22:07
 * Distributed under terms of the MIT license.
 -->
<template>
  <div>
    <h1 class="page-h1">文档中心</h1>
    <div class="mx-8 mt-10">
      <template v-for="(doc, i) in Docs" :key="i">
        <div class="mb-10">
          <div class="text-xl flex items-center mb-4">
            <q-icon class="mx-2" :name="doc.icon"></q-icon>
            <span>{{ doc.name }}</span>
          </div>
          <div class="flex gap-8">
            <template v-for="item in doc.items" :key="item.name">
              <q-chip class="" clickable outline @click="$router.push({
                name: 'doc_item',
                params: { url: item.url, typ: 'public' }
              })" icon="bookmark" color="none">
                <span>{{ item.name }}</span>
              </q-chip>
            </template>
          </div>
          <q-separator class="mt-6" inset></q-separator>

        </div>
      </template>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { DocGroup } from 'src/models';
import { ref } from 'vue';
const Docs = ref<DocGroup[]>([
  {
    name: '用户',
    icon: 'v-team',
    items: [
      { name: '用户注册授权过程', url: '' },
      { name: '用户角色与权限', url: '' },
      { name: 'api文档', url: '' }
    ]
  },
  {
    name: "应用",
    icon: 'v-apps',
    items: [
      { name: '应用创建及基本设置', url: '' },
      { name: '应用权限设置', url: '' },
      { name: '应用对接oa流程', url: '' },
    ],
  },
  {
    name: "系统使用",
    icon: "v-setting",
    items: [
      { name: "编辑器使用及语法", url: 'markdown.md' }
    ]
  }
])
</script>

<style scoped></style>

