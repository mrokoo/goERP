<script setup lang="ts">
import { inject } from "vue";
import { useGoods } from "@/stores/useGoods";
// import StockTable from "./StockTable.vue";
const goods = useGoods();
import {
  NForm,
  NFormItem,
  NInput,
  FormRules,
  NDivider,
  NSelect,
  NUpload,
  NInputNumber,
} from "naive-ui";

const options = [
  {
    label: "激活",
    value: "active",
  },
  {
    label: "冻结",
    value: "freeze",
  },
];
const rules: FormRules = {
  id: [
    {
      required: true,
      message: "请输入产品编号",
    },
  ],
  name: [
    {
      required: true,
      message: "请输入产品名称",
      trigger: ["input", "blur"],
    },
  ],

  state: [
    {
      required: true,
      message: "请选择状态",
      trigger: ["blur"],
    },
  ],
};

const { model, formRef } = inject("form") as any;
</script>
<template>
  <n-form ref="formRef" :model="model" :rules="rules">
    <n-divider title-placement="left" style="margin-top: 0">
      基本信息
    </n-divider>
    <n-form-item path="id" label="产品编号">
      <n-input type="text" v-model:value="model.id" placeholder="" />
    </n-form-item>
    <n-form-item path="name" label="产品名称">
      <n-input type="text" v-model:value="model.name" placeholder="" />
    </n-form-item>
    <n-form-item path="category" label="分类">
      <!-- <n-input type="text" v-model:value="model.category_id" placeholder="" /> -->
      <NSelect
        v-model:value="model.category_id"
        :options="goods.categoryOptions"
        placeholder=""
      />
    </n-form-item>
    <n-form-item path="unit" label="单位">
      <!-- <n-input type="text" v-model:value="model.unit_id" placeholder="" /> -->
      <NSelect
        v-model:value="model.unit_id"
        :options="goods.unitOptions"
        placeholder=""
      />
    </n-form-item>
    <n-form-item path="state" label="状态">
      <n-select v-model:value="model.state" :options="options" placeholder="" />
    </n-form-item>
    <n-form-item path="note" label="备注">
      <n-input type="text" v-model:value="model.note" placeholder="" />
    </n-form-item>
    <n-divider title-placement="left" style="margin-top: 0">
      价格管理
    </n-divider>
    <n-form-item path="purchase" label="采购价(元)">
      <NInputNumber
        type="text"
        v-model:value="model.purchase"
        :default-value="0"
        placeholder=""
        style="width: 100%"
      />
    </n-form-item>
    <n-form-item path="retail" label="零售价(元)">
      <NInputNumber
        type="text"
        v-model:value="model.retail"
        :default-value="0"
        placeholder=""
        style="width: 100%"
      />
    </n-form-item>
    <n-form-item path="grade1" label="等级价1(元)">
      <NInputNumber
        type="text"
        v-model:value="model.grade1"
        placeholder=""
        style="width: 100%"
      />
    </n-form-item>
    <n-form-item path="grade2" label="等级价2(元)">
      <NInputNumber
        type="text"
        v-model:value="model.grade2"
        placeholder=""
        style="width: 100%"
      />
    </n-form-item>
    <n-form-item path="grade3" label="等级价3(元)">
      <NInputNumber
        type="text"
        v-model:value="model.grade3"
        placeholder=""
        style="width: 100%"
      />
    </n-form-item>
    <n-divider title-placement="left" style="margin-top: 0">
      图文信息
    </n-divider>
    <n-form-item path="img" label="图片">
      <n-upload
        action="https://www.mocky.io/v2/5e4bafc63100007100d8b70f"
        list-type="image-card"
      />
    </n-form-item>
    <n-form-item path="intro" label="产品描述">
      <n-input
        type="textarea"
        v-model:value="model.intro"
        placeholder="关于产品的详细描述..."
      />
    </n-form-item>
    <n-divider title-placement="left" style="margin-top: 0; color: #ff0000">
      注：期初库存功能(已废弃，迁移使用库存模块盘点 功能)
    </n-divider>
    <!-- <StockTable /> -->
  </n-form>
</template>
<style scoped></style>
