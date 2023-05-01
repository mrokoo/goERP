<script setup lang="ts">
import { type Category } from "@/@types/goods";
import { useGoods } from "@/stores/useGoods";
import { EditOutlined } from "@vicons/antd";
import { TrashOutline, Add } from "@vicons/ionicons5";
import CategoryForm from "./components/CategoryForm.vue";
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

type CategoryModel = Partial<Category>;
const goods = useGoods();
const message = useMessage();

const showModal = ref(false);
const formRef = ref<FormInst | null>(null);
const model = ref<Partial<CategoryModel>>({});
provide("form", {
  model,
  formRef,
});

const handler = () => {
  formRef.value?.validate((errors) => {
    if (!errors) {
      goods.createCategory(model.value);
      showModal.value = false;
      model.value = {};
    } else {
      message.error("验证失败");
    }
  });
};

const editModal = ref(false);
const eFormRef = ref<FormInst | null>(null);
const eModel = ref<CategoryModel>({});
provide("eform", {
  model: eModel,
  formRef: eFormRef,
});
const editHandler = () => {
  eFormRef.value?.validate((errors) => {
    if (!errors) {
      goods.updateCategory(eModel.value);
      editModal.value = false;
      eModel.value = {};
    } else {
      message.error("验证失败");
    }
  });
};

function deleteHandler(id: string) {
  goods.deleteCategory(id);
}

function renderIcon(icon: Component) {
  return h(NIcon, null, { default: () => h(icon) });
}

type CategoryColumns = {
  id: string;
  name: string;
  type: string;
  state: string;
  balance: number;
};

const columns: DataTableColumns<CategoryColumns> = [
  {
    title: "序号",
    key: "index",
    render(row, index) {
      return h("span", {}, index + 1);
    },
  },

  {
    title: "分类名称",
    key: "name",
  },
  {
    title: "备注",
    key: "note",
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
              const index = goods.category.findIndex((item) => item.id == row.id);
              const item = goods.category[index];
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
  <container title="分类">
    <div class="container" style="margin-bottom: 10px">
      <n-button type="primary" @click="showModal = true">
        新增分类
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
      title="新增分类"
      preset="card"
      size="huge"
      :segmented="true"
    >
      <CategoryForm />
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
      title="编辑分类"
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
      :data="goods.category"
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
