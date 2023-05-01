<script setup lang="ts">
import { inject } from "vue";
import {
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
      message: "请输入仓库名称",
      trigger: ["input", "blur"],
    },
  ],
  id: [
    {
      required: true,
      message: "请输入仓库编号",
      trigger: ["input", "blur"],
    },
  ],
  admin: [
    {
      message: "请输入管理员",
    },
  ],
  phone: [
    {
      message: "请输入电话",
    },
  ],
  address: [
    {
      message: "请输入仓库地址",
    },
  ],
  note: [
    {
      message: "请输入备注",
    },
  ],
  state: [
    {
      message: "请选择状态",
      required: true,
      trigger: ["blur"],
    },
  ],
};

const { model, formRef } = inject("form") as any;
</script>
<template>
  <n-form ref="formRef" :model="model" :rules="rules">
    <n-form-item path="name" label="仓库名称">
      <n-input type="text" v-model:value="model.name" placeholder="" />
    </n-form-item>
    <n-form-item path="id" label="仓库编号">
      <n-input type="text" v-model:value="model.id" placeholder="" />
    </n-form-item>
    <n-form-item path="admin" label="管理员">
      <n-input type="text" v-model:value="model.admin" placeholder="" />
    </n-form-item>
    <n-form-item path="phone" label="手机号">
      <n-input type="text" v-model:value="model.phone" placeholder="" />
    </n-form-item>
    <n-form-item path="address" label="地址">
      <n-input type="text" v-model:value="model.address" placeholder="" />
    </n-form-item>
    <n-form-item path="note" label="备注">
      <n-input type="text" v-model:value="model.note" placeholder="" />
    </n-form-item>
    <n-form-item path="state" label="状态">
      <n-select v-model:value="model.state" :options="options" placeholder="" />
    </n-form-item>
  </n-form>
</template>

<style scoped></style>
