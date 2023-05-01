<script setup lang="ts">
import { ref, provide, h, Component, reactive } from "vue";
import { Add } from "@vicons/ionicons5";
import type { DataTableColumns } from "naive-ui";
import { EditOutlined } from "@vicons/antd";
import { TrashOutline } from "@vicons/ionicons5";
import {
  NDataTable,
  NButton,
  NModal,
  NIcon,
  NTag,
  FormInst,
  useMessage,
} from "naive-ui";
import WarehouseForm from "./components/WarehouseForm.vue";
import Container from "@/components/Container.vue";
import EditForm from "./components/EditForm.vue";
import { useBasic } from "@/stores/useBasic";

const basic = useBasic();
const message = useMessage();

interface WarehouseModel {
  id?: string;
  name?: string;
  address?: string;
  phone?: string;
  admin?: string;
  note?: string;
  state?: string;
}

const showModal = ref(false);
const formRef = ref<FormInst | null>(null);
const model = ref<WarehouseModel>({});
provide("form", {
  model: model,
  formRef: formRef,
});
const handler = () => {
  formRef.value?.validate((errors) => {
    if (!errors) {
      basic.createWarehouse(model.value);
      showModal.value = false;
      model.value = {};
    } else {
      message.error("验证失败");
    }
  });
};

const editModal = ref(false);
const eFormRef = ref<FormInst | null>(null);
const eModel = ref<WarehouseModel>({});
provide("eform", {
  model: eModel,
  formRef: eFormRef,
});
const editHandler = () => {
  eFormRef.value?.validate((errors) => {
    if (!errors) {
      basic.updateWarehouse(eModel.value);
      editModal.value = false;
      eModel.value = {};
    } else {
      message.error("验证失败");
    }
  });
};
function deleteHandler(id: string) {
  basic.deleteWarehouse(id);
}

function renderIcon(icon: Component) {
  return h(NIcon, null, { default: () => h(icon) });
}
type WarehousetColumns = {
  id: string;
  name: string;
  state: string;
};
const columns: DataTableColumns<WarehousetColumns> = [
  {
    title: "序号",
    key: "index",
    render(row, index) {
      return h("span", {}, index + 1);
    },
  },
  {
    title: "仓库编号",
    key: "id",
  },
  {
    title: "仓库名称",
    key: "name",
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
              const index = basic.warehouse.findIndex(
                (item) => item.id == row.id
              );
              const item = basic.warehouse[index];
              eModel.value = item;
              editModal.value = true;
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
  <container title="仓库">
    <div class="container" style="margin-bottom: 10px">
      <n-button type="primary" @click="showModal = true">
        新增仓库
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
      title="新增仓库"
      preset="card"
      size="huge"
      :segmented="true"
    >
      <WarehouseForm />
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
      title="编辑仓库"
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
      :data="basic.warehouse"
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
