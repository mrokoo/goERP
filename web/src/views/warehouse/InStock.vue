<script setup lang="ts">
import { ref } from "vue";
import { useInventory } from "@/stores/useInventory";
import {
  NTabs,
  NTabPane,
  NCard,
  NDataTable,
  type DataTableColumns,
} from "naive-ui";

const inventory = useInventory();
type RowData = {
  id: string;
  warehouse_id: string;
  kind: string;
  state: string;
  created_at: string;
};
const columns: DataTableColumns<RowData> = [
  {
    title: "序号",
    key: "index",
    align: "center",
    render(rowData, rowIndex) {
      return rowIndex + 1;
    },
  },
  {
    title: "单号",
    align: "center",
    key: "id",
  },
  {
    title: "仓库",
    key: "warehouse_id",
    align: "center",
  },
  {
    title: "入库类型",
    key: "kind",
    align: "center",
  },
  {
    title: "入库完成状态",
    key: "status",
    align: "center",
    render: () => "待入库",
  },
  {
    title: "状态",
    key: "state",
    align: "center",
    render(rowData, rowIndex) {
      if (rowData.state === "normal") {
        return "正常";
      } else {
        return "已作废";
      }
    },
  },
  {
    title: "处理日期",
    key: "created_at",
    align: "center",
  },
  {
    title: "操作",
    key: "action",
    align: "center",
  },
];
</script>
<template>
  <n-card>
    <n-tabs type="line" animated size="large">
      <n-tab-pane
        name="TaskNotification"
        tab="入库通知单"
        display-directive="show"
      >
        <NDataTable :columns="columns" :data="inventory.inTask" />
      </n-tab-pane>
      <n-tab-pane name="TaskRecord" tab="入库记录" display-directive="show">
        Hey Jude
      </n-tab-pane>
    </n-tabs>
  </n-card>
</template>
<style scoped></style>
