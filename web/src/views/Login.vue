<script lang="ts" setup>
import { NButton, NInput, NIcon, NSpace, NImage } from "naive-ui";
import { UserOutlined } from "@vicons/antd";
import { LockClosedOutline } from "@vicons/ionicons5";
import LoginLayout from "@/layouts/Login.vue";
const props = defineProps(["currentLayout"]);
const emits = defineEmits(["update:currentLayout"]);
emits("update:currentLayout", LoginLayout);
</script>
<template>
  <div class="container">
    <div class="center">
      <NImage src="../src/assets/logo.png" preview-disabled width="100" />
      <div style="font-size: 36px">GoInvoic</div>
    </div>
    <n-space vertical>
      <n-input placeholder="请输入用户名" round size="large">
        <template #prefix>
          <n-icon :component="UserOutlined" :depth="1" />
        </template>
      </n-input>
      <n-input
        type="password"
        show-password-on="mousedown"
        placeholder="请输入密码"
        style="margin-top: 15px"
        round
        size="large"
        :maxlength="256"
        >\
        <template #prefix>
          <n-icon :component="LockClosedOutline" :depth="1" /> </template
      ></n-input>
      <n-button
        type="primary"
        style="text-align: center; width: 100%; margin-top: 15px"
      >
        登录
      </n-button>
    </n-space>
  </div>
</template>

<style scoped>
.container {
  background-color: #fff;
  border-radius: 10px;
  box-shadow: 0px 0px 5px rgba(0, 0, 0, 0.035), 0px 0px 40px rgba(0, 0, 0, 0.07);
  width: 400px;
  padding: 2rem;
}

.center {
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
  margin: 40px 0;
}
</style>
