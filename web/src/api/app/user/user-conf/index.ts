import request from "@/utils/request";
import { ReqPage, ResPage } from "@/utils/request/interface";
import { UserModel } from "../user";

export interface UserConfModel {
  id?: number;
  userId?: number;
  canLogin?: string;
  remark?: string;
  status?: string;
  createBy?: number;
  updateBy?: number;
  createdAt?: Date;
  updatedAt?: Date;
  user?: UserModel;
}

export const getUserConfPageApi = (params: ReqPage) => {
  return request.get<ResPage<UserConfModel>>(`/admin-api/v1/app/user/user-conf`, { ...params, pageIndex: params?.current });
};

export const getUserConfApi = (id: number) => {
  return request.get<UserConfModel>(`/admin-api/v1/app/user/user-conf/` + id);
};

export const addUserConfApi = (data: object) => {
  return request.post<object>(`/admin-api/v1/app/user/user-conf`, data);
};

export const updateUserConfApi = (id: number, data: object) => {
  return request.put<object>("/admin-api/v1/app/user/user-conf/" + id, data);
};

export const delUserConfApi = (params: number[]) => {
  return request.delete<object>(`/admin-api/v1/app/user/user-conf`, { ids: params });
};

export const exportUserConfApi = (query: object) => {
  return request.download(`/admin-api/v1/app/user/user-conf/export`, query);
};
