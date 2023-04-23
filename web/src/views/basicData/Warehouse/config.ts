import type { DataTableColumns } from "naive-ui";
import { h, Component, reactive, inject } from "vue";
import { NButton, NIcon, NTag } from "naive-ui";
import { EditOutlined } from "@vicons/antd";
import { TrashOutline } from "@vicons/ionicons5";


// const { model, editModal } = inject("emodel") as any;
function renderIcon(icon: Component) {
  return h(NIcon, null, { default: () => h(icon) });
}
export type WarehousetColumns = {
  id: string;
  name: string;
  state: string;
};

export const columns: DataTableColumns<WarehousetColumns> = [
  {
    title: "序号",
    key: "index",
    render(row, index) {
      return h("span", {}, index + 1);
    },
  },
  {
    title: "仓库编号",
    key: "id",
  },
  {
    title: "仓库名称",
    key: "name",
  },
  {
    title: "状态",
    key: "state",
    render(row) {
      let t: "success" | "error" = "success"; // tag类型
      if (row.state != "active") {
        t = "error";
      }
      return h(
        NTag,
        {
          type: t,
          round: true,
        },
        {
          default: () => (row.state == "active" ? "激活" : "冻结"),
        }
      );
    },
  },
  {
    title: "操作",
    key: "action",
    width: 180,
    render(row) {
      return h("div", {}, [
        h(
          NButton,
          {
            size: "small",
            // text: "编辑",
            onClick: () => {
              // const index = basic.warehouse.findIndex(
              //   (item) => item.id == row.id
              // );
              // const item = basic.warehouse[index];
              // model.value = item;
              // editModal.value = true;
            },
          },
          {
            default: () => "编辑",
            icon: () => renderIcon(EditOutlined),
          }
        ),
        h(
          NButton,
          {
            size: "small",
            strong: true,
            secondary: true,
            type: "error",
            onClick: () => {
              console.log(row);
            },
          },
          {
            default: () => "删除",
            icon: () => renderIcon(TrashOutline),
          }
        ),
      ]);
    },
  },
];

export const pagination = reactive({
  page: 1,
  pageSize: 12,
  onChange: (page: number) => {
    pagination.page = page;
  },
});
