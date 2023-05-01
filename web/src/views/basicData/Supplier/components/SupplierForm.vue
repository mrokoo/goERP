<script setup lang="ts">
import { inject } from "vue";
import {
  NInputNumber,
  NForm,
  NFormItem,
  NInput,
  NSelect,
  FormRules,
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
  name: [
    {
      required: true,
      message: "请输入供应商名称",
      trigger: ["input", "blur"],
    },
  ],
  id: [
    {
      required: true,
      message: "请输入供应商编号",
      trigger: ["input", "blur"],
    },
  ],
  state: [
    {
      message: "请选择状态",
      required: true,
      trigger: ["blur"],
    },
  ],
  debt: [
    {
      message: "请输入供应商欠款余额",
      type: "number",
      required: true,
      trigger: ["input", "blur"],
    },
  ],
};

const { model, formRef } = inject("form") as any;
</script>
<template>
  <n-form ref="formRef" :model="model" :rules="rules">
    <n-form-item path="name" label="供应商名称">
      <n-input type="text" v-model:value="model.name" placeholder="" />
    </n-form-item>
    <n-form-item path="id" label="供应商编号">
      <n-input type="text" v-model:value="model.id" placeholder="" />
    </n-form-item>
    <n-form-item path="contact" label="联系人">
      <n-input type="text" v-model:value="model.contact" placeholder="" />
    </n-form-item>
    <n-form-item path="email" label="邮箱">
      <n-input type="text" v-model:value="model.email" placeholder="" />
    </n-form-item>
    <n-form-item path="address" label="地址">
      <n-input type="text" v-model:value="model.address" placeholder="" />
    </n-form-item>
    <n-form-item path="account" label="银行账户">
      <n-input type="text" v-model:value="model.account" placeholder="" />
    </n-form-item>
    <n-form-item path="bank" label="开户行">
      <n-input type="text" v-model:value="model.bank" placeholder="" />
    </n-form-item>
    <n-form-item path="note" label="备注">
      <n-input type="text" v-model:value="model.note" placeholder="" />
    </n-form-item>
    <n-form-item path="state" label="状态">
      <n-select v-model:value="model.state" :options="options" placeholder="" />
    </n-form-item>
    <n-form-item path="debt" label="初期欠款">
      <NInputNumber
        type="text"
        v-model:value="model.debt"
        placeholder=""
        style="width: 100%"
      />
    </n-form-item>
  </n-form>
</template>
<style scoped></style>
