import axios from "@/utils/request";

interface UserAPI {
  getUserList: () => Promise<any>;
  createUser: (data: User) => Promise<any>;
  updateUser: (data: User) => Promise<any>;
  deleteUser: (id: string) => Promise<any>;
}

const userAPI: UserAPI = {
  getUserList: () => axios.get("/users"),
  createUser: (data: User) => {
    return axios.post("/users", data);
  },

  updateUser: (data: User) => {
    return axios.put(`/users/${data.id}`, data);
  },

  deleteUser: (id: string) => {
    return axios.delete(`/users/${id}`);
  },
};

export default userAPI;
