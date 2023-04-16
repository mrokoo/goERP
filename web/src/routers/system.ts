export default [
  {
    path: "/role",
    name: "role",
    component: () => import("@/views/Role.vue"),
  },
  {
    path: "/user",
    name: "user",
    component: () => import("@/views/User.vue"),
  },
  {
    path: "/config",
    name: "config",
    component: () => import("@/views/Config.vue"),
  },
];
