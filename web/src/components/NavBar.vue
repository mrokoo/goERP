<script setup lang="ts">
import { ref, watch } from "vue";
import { NAvatar, NDropdown, NIcon, NButton } from "naive-ui";
import {
  Sunny as SunnyIcon,
  Moon as MoonIcon,
  LogoGithub as GithubIcon,
  HelpCircle as HelpIcon,
} from "@vicons/ionicons5";
// import { WeatherMoon24Filled as Moon, WeatherSunny24Filled as Sun } from "@vicons/fluent";
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
    <n-button text style="font-size: 28px; margin-right: 10px" @click="">
      <n-icon>
        <HelpIcon />
      </n-icon>
    </n-button>
    <n-button text style="font-size: 24px; margin-right: 10px" @click="">
      <n-icon>
        <GithubIcon />
      </n-icon>
    </n-button>
    <n-button text style="font-size: 24px; margin-right: 10px">
      <n-icon v-if="!active" @click="active = true">
        <SunnyIcon />
      </n-icon>
      <n-icon v-else @click="active = false">
        <MoonIcon />
      </n-icon>
    </n-button>
    <!-- <n-switch v-model:value="active"  size="small">
      <template #checked-icon>
        <n-icon :component="MoonIcon" />
      </template>
      <template #unchecked-icon>
        <n-icon :component="SunnyIcon" />
      </template>
    </n-switch> -->
    <n-dropdown :options="options" trigger="click">
      <n-avatar
        round
        size="large"
        src="https://note-image-1302477034.cos.ap-chengdu.myqcloud.com/pic/202304201641061.png"
      >
      </n-avatar>
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
