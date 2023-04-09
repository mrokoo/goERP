<script lang="ts" setup>
import {
  NLayout,
  NLayoutSider,
  NSpace,
  NLayoutContent,
  NLayoutHeader,
} from "naive-ui";
import SideMenu from "../components/SideMenu.vue";
import TopMenu from "../components/TopMenu.vue";
import { RouterView } from "vue-router";
</script>

<template>
  <n-layout has-sider style="height: 100%">
    <n-layout-sider content-style="padding: 24px;" bordered>
      <side-menu></side-menu>
    </n-layout-sider>
    <n-layout>
      <n-layout-header bordered>
        <top-menu></top-menu>
      </n-layout-header>
      <n-layout-content content-style="padding: 24px;margin-top:15px;" embedded>
        <router-view></router-view>
      </n-layout-content>
    </n-layout>
  </n-layout>
</template>

<style scoped></style>
