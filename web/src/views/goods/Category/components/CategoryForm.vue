<script setup lang="ts">
import { inject } from "vue";
import { NForm, NFormItem, NInput, FormRules } from "naive-ui";

const rules: FormRules = {
  name: [
    {
      required: true,
      message: "请输入分类名称",
      trigger: ["input", "blur"],
    },
  ],
};

const { model, formRef } = inject("form") as any;
</script>
<template>
  <n-form ref="formRef" :model="model" :rules="rules">
    <n-form-item path="name" label="分类">
      <n-input type="text" v-model:value="model.name" placeholder="" />
    </n-form-item>
    <n-form-item path="note" label="备注">
      <n-input type="text" v-model:value="model.note" placeholder="" />
    </n-form-item>
  </n-form>
</template>
<style scoped></style>
