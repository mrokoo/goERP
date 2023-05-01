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
const typeOptions = [
  {
    label: "现金",
    value: "cash",
  },
  {
    label: "微信钱包",
    value: "weipay",
  },

  {
    label: "支付宝",
    value: "alipay",
  },
  {
    label: "其他",
    value: "other",
  },
];
const rules: FormRules = {
  name: [
    {
      required: true,
      message: "请输入账户名称",
      trigger: ["input", "blur"],
    },
  ],
  id: [
    {
      required: true,
      message: "请输入账户编号",
      trigger: ["input", "blur"],
    },
  ],
  type: [
    {
      message: "请选择类型",
      required: true,
      trigger: ["blur"],
    },
  ],
  holder: [
    {
      message: "请输入开户人",
    },
  ],
  number: [
    {
      message: "请输入开户账号",
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
  balance: [
    {
      message: "请输入账户初期余额",
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
    <n-form-item path="name" label="账户名称">
      <n-input type="text" v-model:value="model.name" placeholder="" />
    </n-form-item>
    <n-form-item path="id" label="账户编号">
      <n-input type="text" v-model:value="model.id" placeholder="" />
    </n-form-item>
    <n-form-item path="type" label="账户类型">
      <n-select
        v-model:value="model.type"
        :options="typeOptions"
        placeholder=""
      />
    </n-form-item>
    <n-form-item path="holder" label="开户人">
      <n-input type="text" v-model:value="model.holder" placeholder="" />
    </n-form-item>
    <n-form-item path="number" label="开户账号">
      <n-input type="text" v-model:value="model.number" placeholder="" />
    </n-form-item>
    <n-form-item path="note" label="备注">
      <n-input type="text" v-model:value="model.note" placeholder="" />
    </n-form-item>
    <n-form-item path="state" label="状态">
      <n-select v-model:value="model.state" :options="options" placeholder="" />
    </n-form-item>
    <n-form-item path="balance" label="初期余额">
      <NInputNumber
        type="text"
        v-model:value="model.balance"
        placeholder=""
        style="width: 100%"
      />
    </n-form-item>
  </n-form>
</template>
<style scoped></style>
