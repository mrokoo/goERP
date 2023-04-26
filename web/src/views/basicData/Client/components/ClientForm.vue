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
const gradeOptions = [
  {
    label: "优",
    value: "high",
  },
  {
    label: "良",
    value: "medium",
  },

  {
    label: "差",
    value: "low",
  },
];
const rules: FormRules = {
  name: [
    {
      required: true,
      message: "请输入客户名称",
      trigger: ["input", "blur"],
    },
  ],
  id: [
    {
      required: true,
      message: "请输入客户编号",
      trigger: ["input", "blur"],
    },
  ],
  grade: [
    {
      message: "请选择等级",
      required: true,
      trigger: ["blur"],
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
      message: "请输入初期欠款金额",
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
    <n-form-item path="name" label="客户名称">
      <n-input type="text" v-model:value="model.name" placeholder="" />
    </n-form-item>
    <n-form-item path="id" label="客户编号">
      <n-input type="text" v-model:value="model.id" placeholder="" />
    </n-form-item>
    <n-form-item path="grade" label="客户等级">
      <n-select
        v-model:value="model.grade"
        :options="gradeOptions"
        placeholder=""
      />
    </n-form-item>
    <n-form-item path="contact" label="联系人">
      <n-input type="text" v-model:value="model.contact" placeholder="" />
    </n-form-item>
    <n-form-item path="phone" label="联系电话">
      <n-input type="text" v-model:value="model.phone" placeholder="" />
    </n-form-item>
    <n-form-item path="email" label="邮箱">
      <n-input type="text" v-model:value="model.email" placeholder="" />
    </n-form-item>
    <n-form-item path="address" label="地址">
      <n-input type="text" v-model:value="model.address" placeholder="" />
    </n-form-item>
    <n-form-item path="note" label="注释">
      <n-input type="text" v-model:value="model.note" placeholder="" />
    </n-form-item>
    <n-form-item path="state" label="状态">
      <n-select v-model:value="model.state" :options="options" placeholder="" />
    </n-form-item>
    <n-form-item path="debt" label="初期欠款金额">
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
