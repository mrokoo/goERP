import warehouse from "@/api/warehouse";
import account from "@/api/account";
import supplier from "@/api/supplier";
import client from "@/api/client";
import budget from "@/api/budget";
import product from "@/api/product";
import unit from "@/api/unit";
import category from "@/api/category";
import user from "@/api/user";
import purchaseOrder from "./purchaseOrder";
import inventory from "./inventory";

const api = {
  warehouse,
  account,
  supplier,
  client,
  budget,
  product,
  unit,
  category,
  user,
  purchaseOrder,
  inventory,
};

// 导出使用
export default api;
