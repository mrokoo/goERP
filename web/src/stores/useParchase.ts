import { defineStore } from "pinia";
import api from "@/api";
import { createDiscreteApi } from "naive-ui";
const { message } = createDiscreteApi(["message"]);

export const useParchase = defineStore("parchase", {
  state: () => ({}),
  getters: {},
});
