<script setup lang="ts">
import {
  NForm,
  NFormItem,
  NSelect,
  NButton,
  NInputNumber,
  NDataTable,
  DataTableColumns,
  useMessage,
} from "naive-ui";
import { type Product } from "@/@types/goods";
import { h, inject, reactive, ref } from "vue";
import { useBasic } from "@/stores/useBasic";

const message = useMessage();
const basic = useBasic();
type RowData = {
  warehouse_id: string;
  amount: number;
};

const columns: DataTableColumns<RowData> = [
  {
    title: "仓库",
    key: "warehouse",
    render(row) {
      return row.warehouse_id;
    },
  },
  {
    title: "初期库存",
    key: "amount",
    render(row) {
      return h(NInputNumber, {
        value: row.amount,
      });
    },
  },
];

const { model } = inject("form") as any;

const stock = reactive({
  warehouse_id: "",
  amount: 0,
});
const options = basic.warehouseOptions;
const addStock = () => {
  // 检查model.value.openStock是否存在stock的仓库id
  const index = model.value.openStock.findIndex((item: any) => {
    return item.warehouse_id === stock.warehouse_id;
  });

  if (index != -1) {
    message.error("该仓库已存在");
    return;
  }

  model.value.openStock.push({
    ...stock,
  });
  stock.amount = 0;
  stock.warehouse_id = "";
  message.success("添加成功");
};
</script>

<template>
  <n-form>
    <n-form-item label="选择仓库(不能重复)">
      <NSelect
        placeholder=""
        :options="options"
        v-model:value="stock.warehouse_id"
      />
    </n-form-item>
    <n-form-item label="数量">
      <NInputNumber
        placeholder=""
        v-model:value="stock.amount"
        :default-value="0"
      />
      <n-button
        type="primary"
        style="margin-left: 50px; width: 100px"
        @click="addStock"
      >
        添加
      </n-button>
    </n-form-item>
  </n-form>

  <n-data-table :columns="columns" :data="model.openStock" />
</template>

<style scoped></style>
