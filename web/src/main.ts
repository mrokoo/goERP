import { createApp } from "vue";
import "vfonts/Lato.css";
import "vfonts/FiraCode.css";
import "./style.css";
import App from "./App.vue";
import router from "./routers";

createApp(App).use(router).mount("#app");
