<script setup lang="ts">
import { shallowRef, computed } from "vue";
import Default from "./layouts/Default.vue";
import { NConfigProvider, darkTheme, useOsTheme } from "naive-ui";
import { useTheme } from "@/hooks/useTheme";
const layout = shallowRef(Default);
const { theme, Theme } = useTheme();
const themec = computed(() => {
  switch (theme.value) {
    case Theme.Dark:
      return darkTheme;
    case Theme.Light:
      return null;
    default:
      return useOsTheme().value === "dark" ? darkTheme : null;
  }
});
</script>

<template>
  <n-config-provider :theme="themec">
    <component :is="layout">
      <router-view #default="{ Component }">
        <transition name="fade" mode="out-in" appear>
          <Component
            :is="Component"
            :currentLayout="layout"
            @update:currentLayout="(newLayout:any) => (layout = newLayout)"
          />
        </transition>
      </router-view>
    </component>
  </n-config-provider>
</template>

<style scoped lang="less">
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.2s ease-in-out;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}

/* fade-slide */
.fade-slide-leave-active,
.fade-slide-enter-active {
  transition: all 0.3s;
}

.fade-slide-enter-from {
  opacity: 0;
  transform: translateX(-30px);
}

.fade-slide-leave-to {
  opacity: 0;
  transform: translateX(30px);
}

// ///////////////////////////////////////////////
// Fade Bottom
// ///////////////////////////////////////////////

// Speed: 1x
.fade-bottom-enter-active,
.fade-bottom-leave-active {
  transition: opacity 0.25s, transform 0.3s;
}

.fade-bottom-enter-from {
  opacity: 0;
  transform: translateY(-10%);
}

.fade-bottom-leave-to {
  opacity: 0;
  transform: translateY(10%);
}

// fade-scale
.fade-scale-leave-active,
.fade-scale-enter-active {
  transition: all 0.28s;
}

.fade-scale-enter-from {
  opacity: 0;
  transform: scale(1.2);
}

.fade-scale-leave-to {
  opacity: 0;
  transform: scale(0.8);
}

// ///////////////////////////////////////////////
// Fade Top
// ///////////////////////////////////////////////

// Speed: 1x
.fade-top-enter-active,
.fade-top-leave-active {
  transition: opacity 0.2s, transform 0.25s;
}

.fade-top-enter-from {
  opacity: 0;
  transform: translateY(8%);
}

.fade-top-leave-to {
  opacity: 0;
  transform: translateY(-8%);
}
</style>
