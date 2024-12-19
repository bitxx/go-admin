import request from "@/utils/request";
import { ReqPage, ResPage } from "@/utils/request/interface";
import { UserLevelModel } from "../user-level";

export interface UserModel {
  id?: number;
  levelId?: number;
  userName?: string;
  trueName?: string;
  money?: string;
  email?: string;
  mobileTitle?: string;
  mobile?: string;
  avatar?: string;
  payPwd?: string;
  pwd?: string;
  refCode?: string;
  parentId?: number;
  parentIds?: string;
  treeSort?: string;
  treeSorts?: string;
  treeLeaf?: string;
  treeLevel?: number;
  status?: string;
  remark?: string;
  createBy?: number;
  updateBy?: number;
  createdAt?: Date;
  updatedAt?: Date;
  userLevel?: UserLevelModel;
}

export const getUserPageApi = (params: ReqPage) => {
  return request.get<ResPage<UserModel>>(`/admin-api/v1/app/user/user`, { ...params, pageIndex: params?.current });
};

export const getUserApi = (id: number) => {
  return request.get<UserModel>(`/admin-api/v1/app/user/user/` + id);
};

export const addUserApi = (data: object) => {
  return request.post<object>(`/admin-api/v1/app/user/user`, data);
};

export const updateUserApi = (id: number, data: object) => {
  return request.put<object>("/admin-api/v1/app/user/user/" + id, data);
};

export const delUserApi = (params: number[]) => {
  return request.delete<object>(`/admin-api/v1/app/user/user`, { ids: params });
};

export const exportUserApi = (query: object) => {
  return request.download(`/admin-api/v1/app/user/user/export`, query);
};
