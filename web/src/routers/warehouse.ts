export default {
  path: "/warehouse",
  children: [
    // convert  (inStock outStock inventory allocation flow ) to {path:...,name:...,component:...}
    {
      path: "inStock",
      name: "inStock",
      component: () => import("@/views/warehouse/InStock.vue"),
    },
    {
      path: "outStock",
      name: "outStock",
      component: () => import("@/views/warehouse/OutStock.vue"),
    },
    {
      path: "inventory",
      name: "inventory",
      component: () => import("@/views/warehouse/Inventory.vue"),
    },
    {
      path: "allocation",
      name: "allocation",
      component: () => import("@/views/warehouse/Allocation.vue"),
    },
    {
      path: "flow",
      name: "flow",
      component: () => import("@/views/warehouse/Flow.vue"),
    },
  ],
};
