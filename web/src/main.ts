import { createApp } from "vue";
import "vfonts/Lato.css";
import "vfonts/FiraCode.css";
import "./style.css";
import App from "./App.vue";
import router from "./routers";
import { createPinia } from "pinia";

createApp(App)
  .use(router)
  .use(createPinia())
  .mount("#app");
