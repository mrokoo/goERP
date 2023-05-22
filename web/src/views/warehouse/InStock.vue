<script setup lang="ts">
import { ref, h, reactive } from "vue";
import dayjs from "dayjs";
import isLeapYear from "dayjs/plugin/isLeapYear"; // 导入插件
dayjs.extend(isLeapYear); // 使用插件
dayjs.locale("zh-cn"); // 使用本地化语言
import "dayjs/locale/zh-cn"; // 导入本地化语言
import { useInventory } from "@/stores/useInventory";
import {
  NTabs,
  NTabPane,
  NCard,
  NDataTable,
  NButton,
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
    width: "300px",
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
    render(rowData, rowIndex) {
      switch (rowData.kind) {
        case "in_purchase":
          return "采购入库";
        case "in_sale":
          return "销售退货";
        case "in_allocation":
          return "调拨入库";
      }
    },
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
    width: "300px",
    render(rowData, rowIndex) {
      return dayjs(rowData.created_at).format("YYYY-MM-DD");
    },
    sorter: "default",
  },
  {
    title: "操作",
    key: "action",
    align: "center",
    render(rowData, rowIndex) {
      return h(
        "div",
        {
          style: {
            display: "flex",
            justifyContent: "center",
            alignItems: "center",
          },
        },
        [
          h(
            NButton,
            {
              size: "small",
              onClick: () => {
                console.log(rowData);
              },
            },
            "详细"
          ),
          h(
            NButton,
            {
              size: "small",
              type: "primary",
              onClick: () => {
                console.log("切换到入库页面");
              },
            },
            "入库"
          ),
        ]
      );
    },
  },
];
const paginationReactive = reactive({
  pageSize: 12,
});

type RowData2 = {
  id: string;
  warehouse_id: string;
  user_id: string;
  created_at: string;
};

const columns2: DataTableColumns<RowData2> = [
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
    width: "300px",
  },
  {
    title: "仓库",
    key: "warehouse_id",
    align: "center",
  },
  {
    title: "经手人",
    key: "user_id",
    align: "center",
  },
  {
    title: "入库日期",
    key: "created_at",
    align: "center",
    render(rowData, rowIndex) {
      return dayjs(rowData.created_at).format("YYYY-MM-DD");
    },
    sorter: "default",
  },
  {
    title: "操作",
    key: "action",
    align: "center",
    render(rowData, rowIndex) {
      return h(
        "div",
        {
          style: {
            display: "flex",
            justifyContent: "center",
            alignItems: "center",
          },
        },
        [
          h(
            NButton,
            {
              size: "small",
              onClick: () => {
                console.log(rowData);
              },
            },
            "详细"
          ),
          h(
            NButton,
            {
              size: "small",
              type: "error",
              onClick: () => {
                console.log("作废");
              },
            },
            "作废"
          ),
        ]
      );
    },
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
        <NDataTable
          :columns="columns"
          :data="inventory.inTask"
          :pagination="paginationReactive"
        />
      </n-tab-pane>
      <n-tab-pane name="TaskRecord" tab="入库记录" display-directive="show">
        <NDataTable
          :columns="columns2"
          :pagination="paginationReactive"
          :data="inventory.intaskRecord"
        />
      </n-tab-pane>
    </n-tabs>
  </n-card>
</template>
<style scoped></style>
