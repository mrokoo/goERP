<script setup lang="ts">
import { type Budget } from "@/@types/basic";
import { useBasic } from "@/stores/useBasic";
import { EditOutlined } from "@vicons/antd";
import { TrashOutline, Add } from "@vicons/ionicons5";
import BudgetForm from "./components/BudgetForm.vue";
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
} from "naive-ui";

type BudgetModel = Partial<Budget>;
const basic = useBasic();
const message = useMessage();

const showModal = ref(false);
const formRef = ref<FormInst | null>(null);
const model = ref<Partial<BudgetModel>>({});
provide("form", {
  model,
  formRef,
});

const handler = () => {
  formRef.value?.validate((errors) => {
    if (!errors) {
      basic.createBudget(model.value);
      showModal.value = false;
      model.value = {};
    } else {
      message.error("验证失败");
    }
  });
};

const editModal = ref(false);
const eFormRef = ref<FormInst | null>(null);
const eModel = ref<BudgetModel>({});
provide("eform", {
  model: eModel,
  formRef: eFormRef,
});
const editHandler = () => {
  eFormRef.value?.validate((errors) => {
    if (!errors) {
      basic.updateBudget(eModel.value);
      editModal.value = false;
      eModel.value = {};
    } else {
      message.error("验证失败");
    }
  });
};

function deleteHandler(id: string) {
  basic.deleteBudget(id);
}

function renderIcon(icon: Component) {
  return h(NIcon, null, { default: () => h(icon) });
}

type BudgetColumns = {
  id: string;
  name: string;
  type: string;
  note?: string;
};

const columns: DataTableColumns<BudgetColumns> = [
  {
    title: "序号",
    key: "index",
    render(row, index) {
      return h("span", {}, index + 1);
    },
  },
  {
    title: "收支项目",
    key: "name",
  },
  {
    title: "类型",
    key: "type",
    render(row, rowIndex) {
      return row.type === "in" ? "收入" : "支出";
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
              const index = basic.budget.findIndex((item) => item.id == row.id);
              const item = basic.budget[index];
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
  <container title="收支项目">
    <div class="container" style="margin-bottom: 10px">
      <n-button type="primary" @click="showModal = true">
        新增收支项目
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
      title="新增收支项目"
      preset="card"
      size="huge"
      :segmented="true"
    >
      <BudgetForm />
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
      title="编辑收支项目"
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
      :data="basic.budget"
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
