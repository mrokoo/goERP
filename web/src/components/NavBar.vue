<script setup lang="ts">
import { ref, watch } from "vue";
import { NAvatar, NDropdown, NIcon, NSwitch } from "naive-ui";
import {
  Sunny as SunnyIcon,
  Moon as MoonIcon,
  PersonCircleOutline as UserIcon,
} from "@vicons/ionicons5";
import { useTheme } from "@/hooks/useTheme";
const active = ref(false);
const { toggleTheme } = useTheme();
watch(active, () => {
  toggleTheme();
});

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
    <n-switch v-model:value="active" style="margin-right: 10px" size="medium">
      <template #checked-icon>
        <n-icon :component="MoonIcon" />
      </template>
      <template #unchecked-icon>
        <n-icon :component="SunnyIcon" />
      </template>
    </n-switch>
    <n-dropdown :options="options" trigger="click">
      <n-avatar round size="large"> 未登入 </n-avatar>
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
