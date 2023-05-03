import { defineStore } from "pinia";
import api from "@/api";
import { createDiscreteApi } from "naive-ui";
const { message } = createDiscreteApi(["message"]);

export const useInventory = defineStore("inventory", {
  state: () => ({
    task: [] as Task[],
  }),
  getters: {
    inTask(state) {
      return state.task.filter((item) => {
        return item.kind.includes("in");
      });
    },
    taskRecord(state) {
      const record = [] as TaskRecord[];
      state.task.forEach((item) => {
        record.push(...item.records);
      });
      return record;
    },
  },
  actions: {
    async getTaskList() {
      api.inventory
        .getTaskList()
        .then((res: any) => {
          this.task.push(...res.data.data);
        })
        .catch((err: any) => {
          throw err;
        });
    },
  },
});
