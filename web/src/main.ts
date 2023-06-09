import { createApp } from "vue";
import "vfonts/Lato.css";
import "vfonts/FiraCode.css";
import "./style.css";
import App from "./App.vue";
import router from "./routers";
import { createPinia } from "pinia";
import API from "@/api";

const app = createApp(App);
app.config.globalProperties.$api = API;

app.provide("$api", API);

app.use(router).use(createPinia()).mount("#app");
