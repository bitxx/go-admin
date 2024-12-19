import request from "@/utils/request";
import { ReqPage, ResPage } from "@/utils/request/interface";
import { UserModel } from "../user";

export interface UserOperLogModel {
  actionType?: string;
  byType?: string;
  createBy?: number;
  createdAt?: Date;
  id?: number;
  remark?: string;
  status?: string;
  updateBy?: number;
  updatedAt?: Date;
  userId?: number;
  user?: UserModel;
}

export const getUserOperLogPageApi = (params: ReqPage) => {
  return request.get<ResPage<UserOperLogModel>>(`/admin-api/v1/app/user/user-oper-log`, {
    ...params,
    pageIndex: params?.current
  });
};

export const getUserOperLogApi = (id: number) => {
  return request.get<UserOperLogModel>(`/admin-api/v1/app/user/user-oper-log/` + id);
};

export const addUserOperLogApi = (data: object) => {
  return request.post<object>(`/admin-api/v1/app/user/user-oper-log`, data);
};

export const updateUserOperLogApi = (id: number, data: object) => {
  return request.put<object>("/admin-api/v1/app/user/user-oper-log/" + id, data);
};

export const delUserOperLogApi = (params: number[]) => {
  return request.delete<object>(`/admin-api/v1/app/user/user-oper-log`, { ids: params });
};

export const exportUserOperLogApi = (query: object) => {
  return request.download(`/admin-api/v1/app/user/user-oper-log/export`, query);
};
