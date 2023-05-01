<script lang="ts" setup>
import {
  NLayout,
  NLayoutSider,
  NLayoutContent,
  NLayoutHeader,
  NBackTop,
} from "naive-ui";
import SideBar from "@/components/SideBar.vue";
import NavBar from "@/components/NavBar.vue";
</script>

<template>
  <n-layout has-sider style="height: 100vh; position: relative">
    <n-layout-sider
      content-style="padding: 24px;"
      bordered
      :native-scrollbar="false"
    >
      <SideBar />
    </n-layout-sider>
    <n-layout :native-scrollbar="false">
      <n-layout-header bordered>
        <NavBar />
      </n-layout-header>
      <n-layout-content
        :native-scrollbar="false"
        bordered
        content-style="padding: 10px;"
        embedded
      >
        <slot></slot>
      </n-layout-content>
      <NBackTop />
    </n-layout>
  </n-layout>
</template>
