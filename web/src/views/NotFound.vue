<script setup lang="ts">
import { NResult, NButton } from "naive-ui";
import { useRouter } from "vue-router";
const router = useRouter();
const home = {
  name: "home",
};
</script>
<template>
  <n-result status="404" title="404 资源不存在">
    <template #footer>
      <n-button @click="router.push(home)">返回首页</n-button>
    </template>
  </n-result>
</template>
<style scoped></style>
