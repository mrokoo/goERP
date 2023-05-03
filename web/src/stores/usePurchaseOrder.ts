import { defineStore } from "pinia";
import api from "@/api";
import { createDiscreteApi } from "naive-ui";
const { message } = createDiscreteApi(["message"]);

export const usePurchaseOrder = defineStore("purchaseOrder", {
  state: () => ({
    purchaseOrder: [] as PurchaseOrder[],
  }),
  getters: {
    purchaseOrderOptions: (state) => {
      const option: any[] = [];
      state.purchaseOrder.map((item) => {
        if (item.kind == "Order") {
          option.push({
            label: item.id,
            value: item.id,
          });
        }
      });
      return option;
    },
    order: (state) => {
      return state.purchaseOrder.filter((item) => {
        return item.kind === "Order";
      });
    },
    returnOrder: (state) => {
      return state.purchaseOrder.filter((item) => {
        return item.kind === "ReturnOrder";
      });
    },
  },
  actions: {
    async getPurchaseOrder() {
      try {
        const res = await api.purchaseOrder.getPurchaseOrderList();
        console.log(res.data.data);
        this.purchaseOrder.push(...res.data.data);
      } catch (error) {
        throw error;
      }
    },
    async addPurchaseOrder(data: any) {
      try {
        const res = await api.purchaseOrder.addPurchaseOrder(data);
        this.purchaseOrder.push(res.data.data);
        message.success("创建成功");
      } catch (error) {
        message.error("创建订单失败");
      }
    },

    async invalidatePurchaseOrder(id: string) {
      try {
        const res = await api.purchaseOrder.invalidatePurchaseOrder(id);
        message.success("作废成功");
        return true;
      } catch (error) {
        message.error("作废失败");
        return false;
      }
    },

    async addPurchaseReturnOrder(data: any) {
      try {
        const res = await api.purchaseOrder.addPurchaseReturnOrder(data);
        this.purchaseOrder.push(res.data.data);
        message.success("创建成功");
      } catch (error) {
        message.error("创建订单失败");
      }
    },

    async invalidatePurchaseReturnOrder(id: string) {
      try {
        const res = await api.purchaseOrder.invalidatePurchaseReturnOrder(id);
        message.success("作废成功");
        return true;
      } catch (error) {
        message.error("作废失败");
        return false;
      }
    },
  },
});
