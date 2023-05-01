export default {
  path: "/finance",
  name: "finance",
  children: [
    {
      path: "arrears_payable",
      name: "arrears_payable",
      component: () => import("@/views/finance/ArrearsPayable.vue"),
    },
    {
      path: "payment",
      name: "payment",
      component: () => import("@/views/finance/Payment.vue"),
    },
    {
      path: "arrears_receivable",
      name: "arrears_receivable",
      component: () => import("@/views/finance/ArrearsReceivable.vue"),
    },
    {
      path: "collection",
      name: "collection",
      component: () => import("@/views/finance/Collection.vue"),
    },
    {
      path: "account_transfer",
      name: "account_transfer",
      component: () => import("@/views/finance/AccountTransfer.vue"),
    },
    {
      path: "income_and_pay",
      name: "income_and_pay",
      component: () => import("@/views/finance/IncomeAndPay.vue"),
    },
    {
      path: "flow",
      name: "flow",
      component: () => import("@/views/finance/Flow.vue"),
    },
  ],
};
