<script lang="ts" setup>
import { h, Component } from "vue";
import { NIcon, NMenu } from "naive-ui";
import type { MenuOption } from "naive-ui";
import { RouterLink } from "vue-router";
import {
  HomeOutline as HomeIcon,
  BookOutline as BookIcon,
} from "@vicons/ionicons5";

function renderIcon(icon: Component) {
  return () => h(NIcon, null, { default: () => h(icon) });
}

const menuOptions: MenuOption[] = [
  {
    label: () =>
      h(
        RouterLink,
        {
          to: "/",
        },
        { default: () => "回家" }
      ),
    key: "go-back-home",
    icon: renderIcon(HomeIcon),
  },
  {
    key: "divider-1",
    type: "divider",
    props: {
      style: {
        marginLeft: "32px",
      },
    },
  },
  {
    label: () =>
      h(
        RouterLink,
        {
          to: "about",
        },
        { default: () => "关于" }
      ),
    key: "hear-the-wind-sing",
    icon: renderIcon(BookIcon),
  },
];
</script>

<template>
  <n-menu :options="menuOptions" />
</template>
