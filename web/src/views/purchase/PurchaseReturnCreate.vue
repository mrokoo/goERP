<script setup lang="ts">
import { TrashOutline } from "@vicons/ionicons5";
import { reactive, ref, h, Component, computed } from "vue";
import {
  NForm,
  NFormItem,
  NInput,
  NSelect,
  NDivider,
  NButton,
  NModal,
  NDataTable,
  NIcon,
  NInputNumber,
  NSpace,
  useMessage,
  type DataTableColumns,
  type FormRules,
  type FormInst,
} from "naive-ui";
import Container from "@/components/Container.vue";
import { useBasic } from "@/stores/useBasic";
import { useSystem } from "@/stores/useSystem";
import { useGoods } from "@/stores/useGoods";
import { usePurchaseOrder } from "@/stores/usePurchaseOrder";
const basic = useBasic();
const system = useSystem();
const goods = useGoods();
const message = useMessage();
const purchase = usePurchaseOrder();
type RowData = {
  product_id: string;
  quantity: number;
  price: number;
};

function renderIcon(icon: Component) {
  return h(NIcon, null, { default: () => h(icon) });
}

function getProductName(id: string): string {
  const p = goods.product.find((item) => {
    if (item.id === id) {
      return true;
    }
    return false;
  });
  return p?.name ?? "";
}

function getCategoryName(id: string): string {
  const p = goods.product.find((item) => {
    if (item.id === id) {
      return true;
    }
    return false;
  });
  const c = goods.category.find((item) => {
    if (item.id === p?.category_id) {
      return true;
    }
    return false;
  });
  return c?.name ?? "";
}

function getUnitName(id: string): string {
  const p = goods.product.find((item) => {
    if (item.id === id) {
      return true;
    }
    return false;
  });
  const u = goods.unit.find((item) => {
    if (item.id === p?.unit_id) {
      return true;
    }
    return false;
  });
  return u?.name ?? "";
}

const columns: DataTableColumns<RowData> = [
  {
    title: "序号",
    key: "index",
    align: "center",
    render(rowData, rowIndex) {
      return rowIndex + 1;
    },
  },
  {
    title: "名称",
    key: "name",
    align: "center",
    render(rowData, rowIndex) {
      const name = getProductName(rowData.product_id);
      return h("div", name);
    },
  },
  {
    title: "编号",
    key: "product_id",
    align: "center",
  },
  {
    title: "规格",
    key: "category",
    align: "center",
    render(rowData, rowIndex) {
      return h("div", getCategoryName(rowData.product_id));
    },
  },
  {
    title: "单位",
    key: "unit",
    align: "center",
    render(rowData, rowIndex) {
      return h("div", getUnitName(rowData.product_id));
    },
  },
  {
    title: "退货价",
    key: "price",
    align: "center",
    width: "125px",
    render(row, index) {
      return h(NInputNumber, {
        value: model.items![index].price,
        onUpdateValue(v) {
          model.items![index].price = v || 0;
        },
        min: 0,
        style: {
          width: "125px",
        },
      });
    },
  },
  {
    title: "退货数量",
    key: "count",
    align: "center",
    width: "125px",
    render(row, index) {
      return h(NInputNumber, {
        value: model.items![index].quantity,
        onUpdateValue(v) {
          model.items![index].quantity = v || 0;
        },
        min: 1,
        style: {
          width: "125px",
        },
      });
    },
  },
  {
    title: "金额",
    key: "amount",
    align: "center",
    render(rowData, rowIndex) {
      return h("div", rowData.price * rowData.quantity);
    },
  },
  {
    title: "操作",
    key: "action",
    align: "center",
    render(row, index) {
      return h(
        NButton,
        {
          size: "small",
          type: "error",
          onClick: () => {
            model.items!.splice(index, 1);
          },
        },
        {
          default: () => "移除",
          icon: () => renderIcon(TrashOutline),
        }
      );
    },
  },
];

const order = () => {
  return {
    id: "",
    warehouse_id: "",
    basic: "",
    supplier_id: "",
    user_id: "",
    account_id: "",
    other_cost: 0,
    totalCost: 0,
    actal_payment: 0,
    debt: 0,
    created_at: "",
    is_validated: false,
    kind: "ReturnOrder",
    items: [],
  };
};

type PurchaseOrderOptiontal = Partial<PurchaseOrder>;
const model = reactive<PurchaseOrderOptiontal>(order());
const showModal = ref(false);

const id = ref("");
const totalCost = computed(() => {
  const other = model.other_cost ?? 0;
  const sum = model.items?.reduce((prevValue, row) => {
    return prevValue + row.price * row.quantity;
  }, 0);
  return (sum ?? 0) + other;
});
const debt = computed(() => {
  const other = model.other_cost ?? 0;
  const sum = model.items?.reduce((prevValue, row) => {
    return prevValue + row.price * row.quantity;
  }, 0);
  const t = (sum ?? 0) + other;
  return t - (model.actal_payment ?? 0);
});

function addItem(productID: string) {
  if (model.items?.findIndex((item) => item.product_id === productID) !== -1) {
    message.error("该商品已存在");
    return;
  }
  let price = 0;
  goods.product.find((item) => {
    if (item.id === productID) {
      price = item.purchase;
    }
  });
  model.items ?? (model.items = []);
  model.items.push({
    product_id: productID,
    quantity: 0,
    price: price,
  });
  id.value = "";
  showModal.value = false;
}

const rule1: FormRules = {
  id: [
    {
      required: true,
      message: "请输入采购编号",
      trigger: "blur",
    },
  ],
  warehouse_id: [
    {
      required: true,
      message: "请选择仓库",
      trigger: "blur",
    },
  ],
  supplier_id: [
    {
      required: true,
      message: "请选择供应商",
      trigger: "blur",
    },
  ],
  user_id: [
    {
      required: true,
      message: "请选择经手人",
      trigger: "blur",
    },
  ],
};
const form1 = ref<FormInst | null>(null);
function addProductItem() {
  form1.value?.validate(async (errors) => {
    if (errors) {
      message.error("请填写完整信息");
    } else {
      showModal.value = true;
    }
  });
}

const rule2: FormRules = {
  account_id: [
    {
      required: true,
      message: "请选择账户",
      trigger: "blur",
    },
  ],
};
const form2 = ref<FormInst | null>(null);
function submitForm() {
  form2.value?.validate(async (errors) => {
    if (errors) {
      message.error("请填写完整信息");
    } else {
      purchase.addPurchaseReturnOrder(model);
      Object.assign(model, order());
    }
  });
}

function loadItem() {
  const target = purchase.purchaseOrder.find((item) => {
    return item.id === model.basic;
  });
  model.items!.push(...(target?.items ?? []));
  model.warehouse_id = target?.warehouse_id;
  model.supplier_id = target?.supplier_id;
  model.user_id = target?.user_id;
}
</script>
<template>
  <container title="采购退货单">
    <n-form
      inline
      ref="form1"
      :model="model"
      :label-placement="'left'"
      :label-width="100"
      :rules="rule1"
    >
      <n-form-item label="退货单编号:" path="id">
        <NInput v-model:value="model.id" style="width: 200px" placeholder="" />
      </n-form-item>
      <n-form-item label="采购单据:" path="basic">
        <NSelect
          :options="purchase.purchaseOrderOptions"
          v-model:value="model.basic"
          @update:value="loadItem"
          style="width: 200px"
          placeholder=""
        />
      </n-form-item>
      <n-form-item label="仓库:" path="warehouse_id">
        <NSelect
          v-model:value="model.warehouse_id"
          :options="basic.warehouseOptions"
          style="width: 200px"
          placeholder=""
        />
      </n-form-item>
      <n-form-item label="供应商:" path="supplier_id">
        <NSelect
          v-model:value="model.supplier_id"
          :options="basic.supplierOptions"
          style="width: 200px"
          placeholder=""
        />
      </n-form-item>
      <n-form-item label="经手人:" path="user_id">
        <NSelect
          v-model:value="model.user_id"
          :options="system.userOptions"
          style="width: 200px"
          placeholder=""
        />
      </n-form-item>
    </n-form>
    <n-divider title-placement="left">产品信息</n-divider>
    <n-modal
      v-model:show="showModal"
      :mask-closable="false"
      style="width: 550px"
      title="新增产品项"
      preset="card"
      size="huge"
      :segmented="true"
    >
      <n-form>
        <n-form-item label="产品选择" path="name" label-placement="left">
          <NSelect
            :options="goods.productOptions"
            placeholder=""
            v-model:value="id"
          />
        </n-form-item>
      </n-form>
      <template #footer>
        <div class="container">
          <n-button @click="showModal = false" style="margin-right: 10px">
            取消
          </n-button>
          <n-button type="primary" @click="addItem(id)"> 添加 </n-button>
        </div>
      </template>
    </n-modal>
    <n-button
      type="primary"
      @click="addProductItem"
      style="margin-bottom: 15px"
    >
      添加产品
    </n-button>
    <NDataTable :columns="columns" :data="model.items" />

    <n-divider title-placement="left">账单信息</n-divider>
    <n-form
      inline
      ref="form2"
      :rules="rule2"
      :model="model"
      :label-placement="'left'"
      :label-width="100"
    >
      <n-space>
        <n-form-item label="其他费用" path="other_cost">
          <NInputNumber v-model:value="model.other_cost" :default-value="0" />
        </n-form-item>
        <n-form-item label="总计费用" path="total_cost">
          <NInputNumber
            :value="totalCost"
            placeholder=""
            disabled
            :default-value="0"
          />
        </n-form-item>
        <br />
        <n-form-item label="结算账户" path="account_id">
          <NSelect
            :options="basic.accountOptions"
            v-model:value="model.account_id"
            placeholder=""
            style="width: 212px"
          />
        </n-form-item>
        <n-form-item label="实收金额" path="actal_payment">
          <NInputNumber
            v-model:value="model.actal_payment"
            placeholder=""
            :default-value="0"
          />
        </n-form-item>
        <n-form-item label="本单欠款" path="debt">
          <NInputNumber
            :value="debt"
            placeholder=""
            disabled
            :default-value="0"
          />
        </n-form-item>
      </n-space>
    </n-form>
    <n-button type="primary" @click="submitForm"> 创建退货单 </n-button>
  </container>
</template>
<style scoped>
.container {
  display: flex;
  justify-content: flex-end;
  align-items: center;
}
</style>
