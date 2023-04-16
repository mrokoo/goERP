<script setup lang="ts">
import { NResult, NButton } from "naive-ui";
</script>
<template>
  <n-result status="404" title="404 资源不存在">
    <template #footer>
      <n-button>返回首页</n-button>
    </template>
  </n-result>
</template>
<style scoped></style>
