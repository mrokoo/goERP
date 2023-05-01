import { defineStore } from "pinia";
import api from "@/api";
import { createDiscreteApi } from "naive-ui";
const { message } = createDiscreteApi(["message"]);
import { type Product, type Unit, type Category } from "@/@types/goods";

export const useGoods = defineStore("goods", {
  state: () => ({
    product: [] as Product[],
    unit: [] as Unit[],
    category: [] as Category[],
  }),
  getters: {
    categoryOptions(state) {
      return state.category.map((item) => {
        return {
          label: item.name,
          value: item.id,
        };
      });
    },
    unitOptions(state) {
      return state.unit.map((item) => {
        return {
          label: item.name,
          value: item.id,
        };
      });
    },
  },

  actions: {
    async getProduct() {
      api.product
        .getProductList()
        .then((res: any) => {
          this.product.push(...res.data.data);
        })
        .catch((err: any) => {
          throw err;
        });
    },
    async createProduct(data: any) {
      try {
        const res = await api.product.createProduct(data);
        this.product.push(res.data.data);
        message.success("创建成功");
      } catch (error) {
        message.error("编号验证无效");
      }
    },
    async updateProduct(data: any) {
      try {
        const res = await api.product.updateProduct(data);
        const index = this.product.findIndex((item) => item.id === data.id);
        this.product[index] = data;
        message.success("更新成功");
      } catch (error) {
        message.error("编号验证无效");
      }
    },
    async deleteProduct(id: string) {
      try {
        const res = await api.product.deleteProduct(id);
        const index = this.product.findIndex((item) => item.id === id);
        this.product.splice(index, 1);
        message.success("删除成功");
      } catch (error) {
        message.error("编号验证无效");
      }
    },

    async getUnit() {
      api.unit
        .getUnitList()
        .then((res: any) => {
          this.unit.push(...res.data.data);
        })
        .catch((err: any) => {
          throw err;
        });
    },

    async createUnit(data: any) {
      try {
        const res = await api.unit.createUnit(data);
        this.unit.push(res.data.data);
        message.success("创建成功");
      } catch (error) {
        message.error("编号验证无效");
      }
    },

    async updateUnit(data: any) {
      try {
        const res = await api.unit.updateUnit(data);
        const index = this.unit.findIndex((item) => item.id === data.id);
        this.unit[index] = data;
        message.success("更新成功");
      } catch (error) {
        message.error("编号验证无效");
      }
    },

    async deleteUnit(id: string) {
      try {
        const res = await api.unit.deleteUnit(id);
        const index = this.unit.findIndex((item) => item.id === id);
        this.unit.splice(index, 1);
        message.success("删除成功");
      } catch (error) {
        message.error("编号验证无效");
      }
    },

    async getCategory() {
      api.category
        .getCategoryList()
        .then((res: any) => {
          this.category.push(...res.data.data);
        })
        .catch((err: any) => {
          throw err;
        });
    },

    async createCategory(data: any) {
      try {
        const res = await api.category.createCategory(data);
        this.category.push(res.data.data);
        message.success("创建成功");
      } catch (error) {
        message.error("编号验证无效");
      }
    },

    async updateCategory(data: any) {
      try {
        const res = await api.category.updateCategory(data);
        const index = this.category.findIndex((item) => item.id === data.id);
        this.category[index] = data;
        message.success("更新成功");
      } catch (error) {
        message.error("编号验证无效");
      }
    },

    async deleteCategory(id: string) {
      try {
        const res = await api.category.deleteCategory(id);
        const index = this.category.findIndex((item) => item.id === id);
        this.category.splice(index, 1);
        message.success("删除成功");
      } catch (error) {
        message.error("编号验证无效");
      }
    },
  },
});
