import API from "@/api";
declare module "vue" {
  interface ComponentCustomProperties {
    $api: typeof API;
  }
}

export interface ResponseType<T = any> {
  message: string;
  data: T;
}

type API = typeof API;
