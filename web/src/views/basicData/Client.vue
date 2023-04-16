<script setup lang="ts">
import { reactive } from "vue";
import { NCard, NPagination, NDataTable } from "naive-ui";
const page = 1;
const center = {
  display: "flex",
  "align-items": "center",
  "justify-content": "center",
};

const column = reactive([
  {
    title: "Name",
    key: "name",
  },
  {
    title: "Age",
    key: "age",
  },
  {
    title: "Address",
    key: "address",
  },
  {
    title: "Tags",
    key: "tags",
  },
  {
    title: "Action",
    key: "actions",
  },
]);

const data = reactive([]);
</script>
<template>
  <n-card title="客户" :segmented="true" :footer-style="center">
    <n-data-table :columns="column" :data="data"></n-data-table>
    <template class="center" #footer>
      <n-pagination v-model:page="page" :page-count="20" />
    </template>
  </n-card>
</template>
<style scoped></style>
