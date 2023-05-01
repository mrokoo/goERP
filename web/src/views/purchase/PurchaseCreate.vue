<script setup lang="ts">
import { reactive } from "vue";
import { NForm, NFormItem, NInput, NSelect, NDivider } from "naive-ui";
import Container from "@/components/Container.vue";
import { useBasic } from "@/stores/useBasic";
import { useSystem } from "@/stores/useSystem";
const basic = useBasic();
const system = useSystem();

type PurchaseOrder = Partial<{
  id: string;
  warehouse_id: string;
  supplier_id: string;
  user_id: string;
}>;
const model = reactive<PurchaseOrder>({});
const formValue = reactive({});
</script>
<template>
  <container title="采购开单">
    <n-form
      inline
      :model="formValue"
      :label-placement="'left'"
      :label-width="100"
    >
      <n-form-item label="采购编号:" path="id">
        <NInput v-model:value="model.id" style="width: 200px" placeholder="" />
      </n-form-item>
      <n-form-item label="仓库:" path="warehouse">
        <NSelect
          v-model:value="model.warehouse_id"
          :options="basic.warehouseOptions"
          style="width: 200px"
          placeholder=""
        />
      </n-form-item>
      <n-form-item label="供应商:" path="supplier">
        <NSelect
          v-model:value="model.supplier_id"
          :options="basic.supplierOptions"
          style="width: 200px"
          placeholder=""
        />
      </n-form-item>
      <n-form-item label="经手人:" path="user">
        <NSelect
          v-model:value="model.user_id"
          :options="system.userOptions"
          style="width: 200px"
          placeholder=""
        />
      </n-form-item>
    </n-form>
    <n-divider title-placement="left">产品信息</n-divider>
  </container>
</template>
<style scoped></style>
