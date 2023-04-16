import * as VueRouter from "vue-router";
import Home from "../views/Home.vue";
import saleRoutes from "@/routers/sale";
import warehouseRoutes from "@/routers/warehouse";
import financeRoutes from "@/routers/finance";
import systemRoutes from "@/routers/system";
const routes = [
  {
    path: "/home",
    name: "home",
    component: Home,
  },
  {
    path: "/report",
    children: [
      {
        path: "sale",
        name: "sale_report",
        component: () => import("@/views/reports/SaleReport.vue"),
      },
      {
        path: "purchase",
        name: "purchase_report",
        component: () => import("@/views/reports/PurchaseReport.vue"),
      },
      {
        path: "stock",
        name: "stock_report",
        component: () => import("@/views/reports/StockReport.vue"),
      },
      {
        path: "income",
        name: "income_report",
        component: () => import("@/views/reports/IncomeReport.vue"),
      },
    ],
  },
  {
    path: "/basicData",
    children: [
      {
        path: "client",
        name: "client",
        component: () => import("@/views/basicData/Client.vue"),
      },
      // 帮我写supplier、account、budget、warehouse的路
      {
        path: "supplier",
        name: "supplier",
        component: () => import("@/views/basicData/Supplier.vue"),
      },
      {
        path: "account",
        name: "account",
        component: () => import("@/views/basicData/Account.vue"),
      },
      {
        path: "budget",
        name: "budget",
        component: () => import("@/views/basicData/Budget.vue"),
      },
      {
        path: "warehouse",
        name: "warehouse",
        component: () => import("@/views/basicData/Warehouse.vue"),
      },
    ],
  },
  {
    path: "/goods",
    children: [
      {
        path: "category",
        name: "category",
        component: () => import("@/views/goods/Category.vue"),
      },
      {
        path: "unit",
        name: "unit",
        component: () => import("@/views/goods/Unit.vue"),
      },
      {
        path: "information",
        name: "information",
        component: () => import("@/views/goods/Information.vue"),
      },
    ],
  },
  {
    path: "/purchase",
    children: [
      {
        path: "purchase_create",
        name: "purchase_create",
        component: () => import("@/views/purchase/PurchaseCreate.vue"),
      },
      {
        path: "purchase_record",
        name: "purchase_record",
        component: () => import("@/views/purchase/PurchaseRecord.vue"),
      },
      {
        path: "purchase_return_create",
        name: "purchase_return_create",
        component: () => import("@/views/purchase/PurchaseReturnCreate.vue"),
      },
      {
        path: "purchase_return_record",
        name: "purchase_return_record",
        component: () => import("@/views/purchase/PurchaseReturnRecord.vue"),
      },
    ],
  },
  saleRoutes,
  warehouseRoutes,
  financeRoutes,
  ...systemRoutes,
  {
    path: "/:pathMatch(.*)*",
    name: "NotFound",
    component: () => import("@/views/NotFound.vue"),
  },
];

const router = VueRouter.createRouter({
  history: VueRouter.createWebHistory(import.meta.env.BASE_URL),
  routes,
});

export default router;
