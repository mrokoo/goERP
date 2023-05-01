import { defineStore } from "pinia";
import api from "@/api";
import { createDiscreteApi } from "naive-ui";
const { message } = createDiscreteApi(["message"]);

export const useSystem = defineStore("system", {
  state: () => ({
    user: [] as User[],
  }),
  getters: {
    userOptions: (state) => {
      const option: any[] = [];
      state.user.map((item) => {
        option.push({
          label: item.name,
          value: item.id,
        });
      });
      return option;
    },
  },
  actions: {
    async getUser() {
      api.user
        .getUserList()
        .then((res: any) => {
          this.user.push(...res.data.data);
        })
        .catch((err: any) => {
          throw err;
        });
    },
    async createUser(data: any) {
      try {
        const res = await api.user.createUser(data);
        this.user.push(res.data.data);
        message.success("创建成功");
      } catch (error) {
        message.error("编号验证无效");
      }
    },
    async updateUser(data: any) {
      try {
        const res = await api.user.updateUser(data);
        this.user.push(res.data.data);
        message.success("更新成功");
      } catch (error) {
        message.error("编号验证无效");
      }
    },
    async deleteUser(data: any) {
      try {
        const res = await api.user.deleteUser(data);
        this.user.push(res.data.data);
        message.success("删除成功");
      } catch (error) {
        message.error("编号验证无效");
      }
    },
  },
});
