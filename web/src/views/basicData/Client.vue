<script setup lang="ts">
import { reactive } from "vue";
import { NCard, NPagination, NDataTable } from "naive-ui";
const page = 1;
const center = {
  display: "flex",
  "align-items": "center",
  "justify-content": "center",
};

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
  <div class="container">
    <n-card title="客户" :segmented="true" :footer-style="center">
      <n-data-table
        :columns="columns"
        :data="data"
        :pagination="paginationReactive"
      ></n-data-table>
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
