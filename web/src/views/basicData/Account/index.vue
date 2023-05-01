<script setup lang="ts">
import { type Account } from "@/@types/basic";
import { useBasic } from "@/stores/useBasic";
import { EditOutlined } from "@vicons/antd";
import { TrashOutline, Add } from "@vicons/ionicons5";
import AccountForm from "./components/AccountForm.vue";
import EditForm from "./components/EditForm.vue";
import { ref, provide, Component, h, reactive } from "vue";
import Container from "@/components/Container.vue";
import {
  NButton,
  NIcon,
  NModal,
  NDataTable,
  FormInst,
  useMessage,
  DataTableColumns,
  NTag,
} from "naive-ui";

type AccountModel = Partial<Account>;
const basic = useBasic();
const message = useMessage();

const showModal = ref(false);
const formRef = ref<FormInst | null>(null);
const model = ref<Partial<AccountModel>>({});
provide("form", {
  model,
  formRef,
});

const handler = () => {
  formRef.value?.validate((errors) => {
    if (!errors) {
      basic.createAccount(model.value);
      showModal.value = false;
      model.value = {};
    } else {
      message.error("验证失败");
    }
  });
};

const editModal = ref(false);
const eFormRef = ref<FormInst | null>(null);
const eModel = ref<AccountModel>({});
provide("eform", {
  model: eModel,
  formRef: eFormRef,
});
const editHandler = () => {
  eFormRef.value?.validate((errors) => {
    if (!errors) {
      basic.updateAccount(eModel.value);
      editModal.value = false;
      eModel.value = {};
    } else {
      message.error("验证失败");
    }
  });
};

function deleteHandler(id: string) {
  basic.deleteAccount(id);
}

function renderIcon(icon: Component) {
  return h(NIcon, null, { default: () => h(icon) });
}

type AccountColumns = {
  id: string;
  name: string;
  type: string;
  state: string;
  balance: number;
};

const columns: DataTableColumns<AccountColumns> = [
  {
    title: "序号",
    key: "index",
    render(row, index) {
      return h("span", {}, index + 1);
    },
  },
  {
    title: "账户编号",
    key: "id",
  },
  {
    title: "账户名称",
    key: "name",
  },
  {
    title: "类型",
    key: "type",
    render(row) {
      let t: string = "";
      switch (row.type) {
        case "cash":
          t = "现金";
          break;
        case "weipay":
          t = "微信钱包";
          break;
        case "alipay":
          t = "支付宝";
          break;
        case "other":
          t = "其他";
          break;
      }
      return t;
    },
  },
  {
    title: "状态",
    key: "state",
    render(row) {
      let t: "success" | "error" = "success"; // tag类型
      if (row.state != "active") {
        t = "error";
      }
      return h(
        NTag,
        {
          type: t,
          round: true,
        },
        {
          default: () => (row.state == "active" ? "激活" : "冻结"),
        }
      );
    },
  },
  {
    title: "余额",
    key: "balance",
  },
  {
    title: "操作",
    key: "action",
    width: 180,
    render(row) {
      return h("div", {}, [
        h(
          NButton,
          {
            size: "small",
            // text: "编辑",
            onClick: () => {
              const index = basic.account.findIndex(
                (item) => item.id == row.id
              );
              const item = basic.account[index];
              eModel.value = item;
              editModal.value = true;
              // to do
            },
          },
          {
            default: () => "编辑",
            icon: () => renderIcon(EditOutlined),
          }
        ),
        h(
          NButton,
          {
            size: "small",
            strong: true,
            secondary: true,
            type: "error",
            onClick: () => {
              deleteHandler(row.id);
            },
          },
          {
            default: () => "删除",
            icon: () => renderIcon(TrashOutline),
          }
        ),
      ]);
    },
  },
];

const pagination = reactive({
  page: 1,
  pageSize: 12,
  onChange: (page: number) => {
    pagination.page = page;
  },
});
</script>
<template>
  <container title="账户">
    <div class="container" style="margin-bottom: 10px">
      <n-button type="primary" @click="showModal = true">
        新增账户
        <template #icon>
          <n-icon>
            <Add />
          </n-icon>
        </template>
      </n-button>
    </div>
    <n-modal
      v-model:show="showModal"
      :mask-closable="false"
      style="width: 550px"
      title="新增账户"
      preset="card"
      size="huge"
      :segmented="true"
    >
      <AccountForm />
      <template #footer>
        <div class="container">
          <n-button @click="showModal = false" style="margin-right: 10px">
            取消
          </n-button>
          <n-button type="primary" @click="handler()"> 确定 </n-button>
        </div>
      </template>
    </n-modal>
    <n-modal
      v-model:show="editModal"
      :mask-closable="false"
      style="width: 550px"
      title="编辑账户"
      preset="card"
      size="huge"
    >
      <EditForm />
      <template #footer>
        <div class="container">
          <n-button @click="editModal = false" style="margin-right: 10px">
            取消
          </n-button>
          <n-button type="primary" @click="editHandler()"> 确定 </n-button>
        </div>
      </template>
    </n-modal>
    <n-data-table
      :columns="columns"
      :data="basic.account"
      :pagination="pagination"
    />
  </container>
</template>
<style scoped>
.container {
  display: flex;
  justify-content: flex-end;
  align-items: center;
}
</style>
