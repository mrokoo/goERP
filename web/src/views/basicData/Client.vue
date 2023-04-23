<script setup lang="ts">
import { reactive } from "vue";
import { NDataTable } from "naive-ui";
import Container from "@/components/Container.vue";

const columns = [
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
];

const data = Array.from({ length: 46 }).map((_, index) => ({
  key: index,
  name: `Edward King ${index}`,
  age: 32,
  address: `London, Park Lane no. ${index}`,
}));
const paginationReactive = reactive({
  page: 1,
  pageSize: 12,
  onChange: (page: number) => {
    paginationReactive.page = page;
  },
});
</script>
<template>
  <container title="客户">
    <n-data-table
      :columns="columns"
      :data="data"
      :pagination="paginationReactive"
    ></n-data-table>
  </container>
</template>
<style scoped>
.container {
  display: flex;
  height: 100%;
  justify-content: center;
  align-items: center;
}
</style>
