import API from "@/api";
declare module "vue" {
  interface ComponentCustomProperties {
    $api: typeof API;
  }
}

type API = typeof API;
