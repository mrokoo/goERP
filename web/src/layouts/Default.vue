<script lang="ts" setup>
import {
  NLayout,
  NLayoutSider,
  NLayoutContent,
  NLayoutHeader,
  NScrollbar,
} from "naive-ui";
import SideBar from "@/components/SideBar.vue";
import NavBar from "@/components/NavBar.vue";
</script>

<template>
  <n-layout has-sider style="height: 100vh">
    <n-layout-sider
      content-style="padding: 24px;"
      bordered
      :native-scrollbar="false"
    >
      <!-- <n-scrollbar> -->
      <SideBar />
      <!-- </n-scrollbar> -->
    </n-layout-sider>
    <n-layout>
      <n-layout-header bordered>
        <NavBar />
      </n-layout-header>
      <n-layout-content bordered content-style="padding: 10px;" embedded>
        <slot></slot>
      </n-layout-content>
    </n-layout>
  </n-layout>
</template>
