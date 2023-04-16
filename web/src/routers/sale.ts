export default {
  path: "/sale",
  name: "sale",
  children: [
    // sale_create sale_record sale_return_create sale_return_record
    {
      path: "sale_create",
      name: "sale_create",
      component: () => import("@/views/sale/SaleCreate.vue"),
    },
    {
      path: "sale_record",
      name: "sale_record",
      component: () => import("@/views/sale/SaleRecord.vue"),
    },
    {
      path: "sale_return_create",
      name: "sale_return_create",
      component: () => import("@/views/sale/SaleReturnCreate.vue"),
    },
    {
      path: "sale_return_record",
      name: "sale_return_record",
      component: () => import("@/views/sale/SaleReturnRecord.vue"),
    },
  ],
};
