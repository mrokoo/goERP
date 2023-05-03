import axios from "@/utils/request";

interface PurchaseOrderAPI {
  getPurchaseOrderList: () => Promise<any>;
  addPurchaseOrder: (data: any) => Promise<any>;
  invalidatePurchaseOrder: (id: string) => Promise<any>;
  addPurchaseReturnOrder: (data: any) => Promise<any>;
  invalidatePurchaseReturnOrder: (id: string) => Promise<any>;
}

const purchaseOrderAPI: PurchaseOrderAPI = {
  getPurchaseOrderList: () => axios.get("/purchaseOrders"),
  addPurchaseOrder: (data: any) => {
    return axios.post("/purchaseOrders", data);
  },
  invalidatePurchaseOrder: (id: string) => {
    return axios.put(`/purchaseOrders/${id}`);
  },
  addPurchaseReturnOrder: (data: any) => {
    return axios.post("/purchaseReturnOrders", data);
  },
  invalidatePurchaseReturnOrder: (id: string) => {
    return axios.put(`/purchaseReturnOrders/${id}`);
  },
};

export default purchaseOrderAPI;
