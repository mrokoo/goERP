<script setup lang="ts">
import { NCard } from "naive-ui";
const props = defineProps<{
  title: string;
}>();
const center = {
  display: "flex",
  "align-items": "center",
  "justify-content": "center",
};
</script>
<template>
  <div class="container">
    <n-card :title="props.title" :segmented="true" :footer-style="center">
      <slot></slot>
    </n-card>
  </div>
</template>
<style scoped>
.container {
  display: flex;
  height: 100%;
  justify-content: center;
  align-items: center;
}
</style>
