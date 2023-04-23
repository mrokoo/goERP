<script lang="ts" setup>
import { h, Component } from "vue";
import { NIcon, NMenu, NDivider, NScrollbar } from "naive-ui";
import type { MenuOption } from "naive-ui";
import { RouterLink } from "vue-router";
import {
  HomeOutline as HomeIcon,
  PieChartOutline as ChartIcon,
  CubeOutline as BasicIcon,
  FileTrayOutline as FileIcon,
  CartOutline as CartIcon,
  BagHandleOutline as BagIcon,
  LayersOutline as WarehouseIcon,
  SettingsOutline as SettingIcon,
} from "@vicons/ionicons5";

import { DollarOutlined as FinanceIcon } from "@vicons/antd";

function renderIcon(icon: Component) {
  return () => h(NIcon, null, { default: () => h(icon) });
}

const menuOptions: MenuOption[] = [
  {
    label: () =>
      h(
        RouterLink,
        {
          to: {
            name: "home",
          },
        },
        { default: () => "首页" }
      ),
    key: "home",
    icon: renderIcon(HomeIcon),
  },
  {
    label: "报表统计",
    key: "report-statistics",
    icon: renderIcon(ChartIcon),
    children: [
      {
        label: () =>
          h(
            RouterLink,
            {
              to: {
                name: "sale_report",
              },
            },
            { default: () => "销售报表" }
          ),
        key: "sale-report",
      },
      {
        label: () =>
          h(
            RouterLink,
            {
              to: {
                name: "purchase_report",
              },
            },
            { default: () => "采购报表" }
          ),
        key: "purchase-report",
      },
      {
        label: () =>
          h(
            RouterLink,
            {
              to: {
                name: "stock_report",
              },
            },
            { default: () => "库存报表" }
          ),
        key: "stock-report",
      },
      {
        label: () =>
          h(
            RouterLink,
            {
              to: {
                name: "income_report",
              },
            },
            { default: () => "收支统计" }
          ),
        key: "income",
      },
    ],
  },
  {
    label: "基础数据",
    key: "basic-data",
    icon: renderIcon(BasicIcon),
    children: [
      {
        label: () =>
          h(
            RouterLink,
            {
              to: {
                name: "client",
              },
            },
            { default: () => "客户管理" }
          ),
        key: "cilent",
      },
      {
        label: () =>
          h(
            RouterLink,
            {
              to: {
                name: "supplier",
              },
            },
            { default: () => "供应商管理" }
          ),
        key: "supplier",
      },
      {
        label: () =>
          h(
            RouterLink,
            {
              to: {
                name: "warehousem",
              },
            },
            { default: () => "仓库管理" }
          ),
        key: "warehouse",
      },
      {
        label: () =>
          h(
            RouterLink,
            {
              to: {
                name: "account",
              },
            },
            { default: () => "结算账户" }
          ),
        key: "account",
      },
      {
        label: () =>
          h(
            RouterLink,
            {
              to: {
                name: "budget",
              },
            },
            { default: () => "收支项目" }
          ),
        key: "budget",
      },
    ],
  },
  {
    label: "产品管理",
    key: "goods",
    icon: renderIcon(FileIcon),
    children: [
      {
        label: () =>
          h(
            RouterLink,
            {
              to: {
                name: "category",
              },
            },
            { default: () => "产品分类" }
          ),
        key: "category",
      },
      {
        label: () =>
          h(
            RouterLink,
            {
              to: {
                name: "unit",
              },
            },
            { default: () => "产品单位" }
          ),
        key: "unit",
      },
      {
        label: () =>
          h(
            RouterLink,
            {
              to: {
                name: "information",
              },
            },
            { default: () => "产品信息" }
          ),
        key: "information",
      },
    ],
  },
  {
    label: "采购管理",
    key: "purchase",
    icon: renderIcon(CartIcon),
    children: [
      {
        label: () =>
          h(
            RouterLink,
            {
              to: {
                name: "purchase_create",
              },
            },
            { default: () => "采购开单" }
          ),
        key: "purchase_create",
      },
      {
        label: () =>
          h(
            RouterLink,
            {
              to: {
                name: "purchase_record",
              },
            },
            { default: () => "采购记录" }
          ),
        key: "purchase_record",
      },
      {
        label: () =>
          h(
            RouterLink,
            {
              to: {
                name: "purchase_return_create",
              },
            },
            { default: () => "采购退货" }
          ),
        key: "purchase_return_create",
      },
      {
        label: () =>
          h(
            RouterLink,
            {
              to: {
                name: "purchase_return_record",
              },
            },
            { default: () => "退货记录" }
          ),
        key: "purchase_return_record",
      },
    ],
  },
  {
    label: "销售管理",
    key: "sale",
    icon: renderIcon(BagIcon),
    children: [
      {
        label: () =>
          h(
            RouterLink,
            {
              to: {
                name: "sale_create",
              },
            },
            { default: () => "销售开单" }
          ),
        key: "sale_create",
      },
      {
        label: () =>
          h(
            RouterLink,
            {
              to: {
                name: "sale_record",
              },
            },
            { default: () => "销售记录" }
          ),
        key: "sale_record",
      },
      {
        label: () =>
          h(
            RouterLink,
            {
              to: {
                name: "sale_return_create",
              },
            },
            { default: () => "销售退货" }
          ),
        key: "sale_return_create",
      },
      {
        label: () =>
          h(
            RouterLink,
            {
              to: {
                name: "sale_return_record",
              },
            },
            { default: () => "退货记录" }
          ),
        key: "sale_return_record",
      },
    ],
  },
  {
    label: "库存管理",
    key: "warehouses",
    icon: renderIcon(WarehouseIcon),
    children: [
      // please help me write InStock OutStock
      {
        label: () =>
          h(
            RouterLink,
            {
              to: {
                name: "inStock",
              },
            },
            { default: () => "入库任务" }
          ),
        key: "inStock",
      },
      {
        label: () =>
          h(
            RouterLink,
            {
              to: {
                name: "outStock",
              },
            },
            { default: () => "出库任务" }
          ),
        key: "outStock",
      },
      {
        label: () =>
          h(
            RouterLink,
            {
              to: {
                name: "inventory",
              },
            },
            { default: () => "盘点" }
          ),
        key: "inventory",
      },
      {
        label: () =>
          h(
            RouterLink,
            {
              to: {
                name: "allocation",
              },
            },
            { default: () => "调拨" }
          ),
        key: "allocation",
      },
      {
        label: () =>
          h(
            RouterLink,
            {
              to: {
                name: "flow",
              },
            },
            { default: () => "库存流水" }
          ),
        key: "flow",
      },
    ],
  },
  {
    label: "财务管理",
    key: "finance",
    icon: renderIcon(FinanceIcon),
    children: [
      // please help me write InStock OutStock
      {
        label: () =>
          h(
            RouterLink,
            {
              to: {
                name: "arrears_payable",
              },
            },
            { default: () => "应付欠款" }
          ),
        key: "arrears_payable",
      },
      {
        label: () =>
          h(
            RouterLink,
            {
              to: {
                name: "payment",
              },
            },
            { default: () => "付款" }
          ),
        key: "payment",
      },
      {
        label: () =>
          h(
            RouterLink,
            {
              to: {
                name: "arrears_receivable",
              },
            },
            { default: () => "应收欠款" }
          ),
        key: "arrears_receivable",
      },
      {
        label: () =>
          h(
            RouterLink,
            {
              to: {
                name: "collection",
              },
            },
            { default: () => "收款" }
          ),
        key: "collection",
      },
      {
        label: () =>
          h(
            RouterLink,
            {
              to: {
                name: "account_transfer",
              },
            },
            { default: () => "账户转账" }
          ),
        key: "account_transfer",
      },
      {
        label: () =>
          h(
            RouterLink,
            {
              to: {
                name: "income_and_pay",
              },
            },
            { default: () => "日常收支" }
          ),
        key: "income_and_pay",
      },
      {
        label: () =>
          h(
            RouterLink,
            {
              to: {
                name: "flow",
              },
            },
            { default: () => "资金流水" }
          ),
        key: "flow",
      },
    ],
  },
  {
    label: "系统管理",
    key: "system",
    icon: renderIcon(SettingIcon),
    children: [
      // please help me write InStock OutStock
      {
        label: () =>
          h(
            RouterLink,
            {
              to: {
                name: "role",
              },
            },
            { default: () => "角色管理" }
          ),
        key: "role",
      },
      {
        label: () =>
          h(
            RouterLink,
            {
              to: {
                name: "user",
              },
            },
            { default: () => "用户管理" }
          ),
        key: "user",
      },
      {
        label: () =>
          h(
            RouterLink,
            {
              to: {
                name: "config",
              },
            },
            { default: () => "系统配置" }
          ),
        key: "config",
      },
    ],
  },
];
</script>

<template>
  <div class="logoc">
    <img src="../assets/logo.png" alt="logo" class="logo" />
    <span class="logo-text">GoInvic</span>
  </div>
  <NDivider style="margin: 0" />
  <n-menu :options="menuOptions" :indent="32" />
</template>

<style scoped>
.logoc {
  height: 40px;
  font-size: 24px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.logo {
  height: 35px;
  width: 48px;
  margin-bottom: 8px;
}

.logo-text {
  font-weight: bold;
  margin-left: 10px;
  text-align: center;
}
</style>
