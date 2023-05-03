<script setup lang="ts">
import { h, reactive, ref, Component, computed } from "vue";
import Container from "@/components/Container.vue";
import { NButton, NDataTable, type DataTableColumns } from "naive-ui";
import { usePurchaseOrder } from "@/stores/usePurchaseOrder";

const order = usePurchaseOrder();

type RowData = {
  id: string;
  supplier_id: string;
  user_id: string;
  created_at: string;
  total_cost: number;
  actal_payment: number;
  other_cost: number;
  is_validated: boolean;
};
const paginationReactive = reactive({
  pageSize: 12,
});
const tcolumns: DataTableColumns<RowData> = [
  {
    title: "序号",
    key: "index",
    align: "center",
    render(rowData, rowIndex) {
      return rowIndex + 1;
    },
  },
  {
    title: "采购退货编号",
    key: "id",
    align: "center",
  },
  {
    title: "供应商",
    key: "supplier_id",
    align: "center",
  },
  {
    title: "经手人",
    key: "user_id",
    align: "center",
  },
  {
    title: "采购日期",
    key: "created_at",
    align: "center",
  },
  {
    title: "采购总额",
    key: "total_cost",
    align: "center",
  },
  {
    title: "实际支付",
    key: "actal_payment",
    align: "center",
  },
  {
    title: "其他费用",
    key: "other_cost",
    align: "center",
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
            alignItems: "center",
            justifyContent: "center",
          },
        },
        [
          h(
            NButton,
            {
              size: "small",
              onClick: () => {
                console.log("oo");
              },
            },
            {
              default: () => "详细",
            }
          ),
          h(
            NButton,
            {
              size: "small",
              type: rowData.is_validated ? "tertiary" : "error",
              disabled: rowData.is_validated,
              onClick: async () => {
                const a = await order.invalidatePurchaseReturnOrder(rowData.id);
                if (a === true) {
                  rowData.is_validated = true;
                }
              },
            },
            {
              default: () => (rowData.is_validated ? "已作废" : "作废"),
            }
          ),
        ]
      );
    },
  },
];
</script>
<template>
  <container title="采购退货订单">
    <n-data-table
      :columns="tcolumns"
      :data="order.returnOrder"
      :pagination="paginationReactive"
    />
  </container>
</template>
<style scoped></style>
