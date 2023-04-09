<script setup lang="ts">
import { NAvatar } from "naive-ui";
import { NDropdown } from "naive-ui";
const options = [
  {
    label: "修改密码",
    key: "passwords",
  },
  {
    label: "退出登入",
    key: "exit",
  },
];
</script>

<template>
  <div class="container">
    <n-dropdown :options="options" trigger="click">
      <n-avatar
        round
        size="large"
        :style="{
          color: 'yellow',
          backgroundColor: 'red',
        }"
        >张</n-avatar
      >
    </n-dropdown>
  </div>
</template>

<style scoped>
.container {
  box-sizing: border-box;
  display: flex;
  justify-content: end;
  align-items: center;
  width: 100%;
  height: 64px;
  padding-right: 20px;
  box-shadow: 14.7px 18.9px 14.7px rgba(0, 0, 0, 0.07),
    25.8px 33.2px 40.7px rgba(0, 0, 0, 0.046),
    37.2px 47.9px 98px rgba(0, 0, 0, 0.035),
    87px 112px 325px rgba(0, 0, 0, 0.024);
}
</style>
