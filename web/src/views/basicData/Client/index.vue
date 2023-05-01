<script setup lang="ts">
import { type Client } from "@/@types/basic";
import { useBasic } from "@/stores/useBasic";
import { EditOutlined } from "@vicons/antd";
import { TrashOutline, Add } from "@vicons/ionicons5";
import ClientForm from "./components/ClientForm.vue";
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

type ClientModel = Partial<Client>;
const basic = useBasic();
const message = useMessage();

const showModal = ref(false);
const formRef = ref<FormInst | null>(null);
const model = ref<Partial<ClientModel>>({});
provide("form", {
  model,
  formRef,
});

const handler = () => {
  formRef.value?.validate((errors) => {
    if (!errors) {
      basic.createClient(model.value);
      showModal.value = false;
      model.value = {};
    } else {
      message.error("验证失败");
    }
  });
};

const editModal = ref(false);
const eFormRef = ref<FormInst | null>(null);
const eModel = ref<ClientModel>({});
provide("eform", {
  model: eModel,
  formRef: eFormRef,
});
const editHandler = () => {
  eFormRef.value?.validate((errors) => {
    if (!errors) {
      basic.updateClient(eModel.value);
      editModal.value = false;
      eModel.value = {};
    } else {
      message.error("验证失败");
    }
  });
};

function deleteHandler(id: string) {
  basic.deleteClient(id);
}

function renderIcon(icon: Component) {
  return h(NIcon, null, { default: () => h(icon) });
}

type ClientColumns = {
  id: string;
  name: string;
  contact: string;
  phone: string;
  state: string;
};

const columns: DataTableColumns<ClientColumns> = [
  {
    title: "序号",
    key: "index",
    render(row, index) {
      return h("span", {}, index + 1);
    },
  },
  {
    title: "客户编号",
    key: "id",
  },
  {
    title: "客户名称",
    key: "name",
  },
  {
    title: "联系人",
    key: "contact",
  },
  {
    title: "联系电话",
    key: "phone",
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
              const index = basic.client.findIndex((item) => item.id == row.id);
              const item = basic.client[index];
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
  <container title="客户">
    <div class="container" style="margin-bottom: 10px">
      <n-button type="primary" @click="showModal = true">
        新增客户
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
      title="新增客户"
      preset="card"
      size="huge"
      :segmented="true"
    >
      <ClientForm />
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
      title="编辑客户"
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
      :data="basic.client"
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
