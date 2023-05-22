<script setup lang="ts">
import Container from "@/components/Container.vue";
import { h } from "vue";
import { NDataTable, NButton, type DataTableColumns } from "naive-ui";
import { useInventory } from "@/stores/useInventory";
const inventory = useInventory();
const columns: DataTableColumns = [
  {
    title: "序号",
    key: "index",
    align: "center",
    render(rowData, rowIndex) {
      return rowIndex + 1;
    },
  },
  {
    title: "仓库",
    key: "warehouse_id",
    align: "center",
  },
  {
    title: "产品名称",
    key: "name",
    align: "center",
  },
  {
    title: "产品编号",
    key: "product_id",
    align: "center",
  },
  {
    title: "流水类型",
    key: "flow",
    align: "center",
  },
  {
    title: "变化之前数量",
    key: "previous",
    align: "center",
  },
  {
    title: "变化数量",
    key: "change",
    align: "center",
  },
  {
    title: "变化之后数量",
    key: "present",
    align: "center",
  },
  {
    title: "创建日期",
    key: "created_at",
    align: "center",
  },
  {
    title: "操作",
    key: "action",
    align: "center",
    render(rowData, rowIndex) {
      return h(
        NButton,
        {
          size: "small",
        },
        { default: () => "详细" }
      );
    },
  },
];
</script>
<template>
  <container title="库存流水">
    <NDataTable :columns="columns" :data="inventory.flow" />
  </container>
</template>
<style scoped></style>
