import { defineStore } from "pinia";
import api from "@/api";
import { createDiscreteApi } from "naive-ui";
const { message } = createDiscreteApi(["message"]);

export const useInventory = defineStore("inventory", {
  state: () => ({
    task: [] as Task[],
    flow: [] as InventoryFlow[],
  }),
  getters: {
    inTask(state) {
      return state.task.filter((item) => {
        return item.kind.includes("in");
      });
    },
    intaskRecord(state) {
      const record = [] as TaskRecord[];
      const inTask = state.task.filter((item) => {
        return item.kind.includes("in");
      });

      inTask.forEach((item) => {
        if (item.records == null) {
          return;
        } else {
          record.push(...item.records);
        }
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

    async getInventoryFlowList() {
      api.inventory
        .getInventoryFlowList()
        .then((res: any) => {
          this.flow.push(...res.data.data);
        })
        .catch((err: any) => {
          throw err;
        });
    },
  },
});
