<script setup lang="ts">
import { inject } from "vue";
import { NForm, NFormItem, NInput, NSelect, FormRules } from "naive-ui";

const typeOptions = [
  {
    label: "收入",
    value: "in",
  },
  {
    label: "支出",
    value: "out",
  },
];
const rules: FormRules = {
  name: [
    {
      required: true,
      message: "请输入收支项目名称",
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
};

const { model, formRef } = inject("form") as any;
</script>
<template>
  <n-form ref="formRef" :model="model" :rules="rules">
    <n-form-item path="name" label="收支项目">
      <n-input type="text" v-model:value="model.name" placeholder="" />
    </n-form-item>
    <n-form-item path="type" label="收支类型">
      <n-select
        v-model:value="model.type"
        :options="typeOptions"
        placeholder=""
      />
    </n-form-item>
    <n-form-item path="note" label="备注">
      <n-input type="text" v-model:value="model.note" placeholder="" />
    </n-form-item>
  </n-form>
</template>
<style scoped></style>
