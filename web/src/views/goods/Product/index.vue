<script setup lang="ts">
import { type Product } from "@/@types/goods";
import { useGoods } from "@/stores/useGoods";
import { EditOutlined } from "@vicons/antd";
import { TrashOutline, Add } from "@vicons/ionicons5";
import ProductForm from "./components/ProductForm.vue";
import EditForm from "./components/EditForm.vue";
import { ref, provide, Component, h, reactive } from "vue";
import Container from "@/components/Container.vue";
import {
  NButton,
  NIcon,
  NModal,
  NTag,
  NDataTable,
  FormInst,
  useMessage,
  DataTableColumns,
} from "naive-ui";

type ProductModel = Partial<Product>;
const goods = useGoods();
const message = useMessage();

const showModal = ref(false);
const formRef = ref<FormInst | null>(null);
const model = ref<Partial<ProductModel>>({
  openStock: [],
});
provide("form", {
  model,
  formRef,
});

const handler = () => {
  formRef.value?.validate((errors) => {
    if (!errors) {
      goods.createProduct(model.value);
      showModal.value = false;
      model.value = {};
    } else {
      message.error("验证失败");
    }
  });
};

const editModal = ref(false);
const eFormRef = ref<FormInst | null>(null);
const eModel = ref<ProductModel>({
  openStock: [],
});
provide("eform", {
  model: eModel,
  formRef: eFormRef,
});
const editHandler = () => {
  eFormRef.value?.validate((errors) => {
    if (!errors) {
      goods.updateProduct(eModel.value);
      editModal.value = false;
      eModel.value = {};
    } else {
      message.error("验证失败");
    }
  });
};

function deleteHandler(id: string) {
  goods.deleteProduct(id);
}

function renderIcon(icon: Component) {
  return h(NIcon, null, { default: () => h(icon) });
}

type ProductColumns = {
  index: string;
  id: string;
  name: string;
  category_id: string;
  state: string;
};

const columns: DataTableColumns<ProductColumns> = [
  {
    title: "序号",
    key: "index",
    render(row, index) {
      return h("span", {}, index + 1);
    },
  },
  {
    title: "产品编号",
    key: "id",
  },

  {
    title: "产品名称",
    key: "name",
  },
  {
    title: "分类",
    key: "category_id",
    render(rowData, rowIndex) {
      return goods.categoryOptions.find(
        (item) => item.value == rowData.category_id
      )?.label;
    },
  },
  {
    title: "采购价",
    key: "purchase",
  },
  {
    title: "零售价",
    key: "retail",
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
              const index = goods.product.findIndex(
                (item) => item.id == row.id
              );
              const item = goods.product[index];
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
  <container title="产品">
    <div class="container" style="margin-bottom: 10px">
      <n-button type="primary" @click="showModal = true">
        新增产品
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
      title="新增产品"
      preset="card"
      size="huge"
      :segmented="true"
    >
      <ProductForm />
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
      title="编辑产品"
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
      :data="goods.product"
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
