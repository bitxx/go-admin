import request from "@/utils/request";
import { ReqPage, ResPage } from "@/utils/request/interface";
import { UserModel } from "../user";

export interface UserAccountLogModel {
  id?: number;
  userId?: number;
  changeMoney?: string;
  beforeMoney?: string;
  afterMoney?: string;
  moneyType?: string;
  changeType?: string;
  status?: string;
  createBy?: number;
  createdAt?: Date;
  updateBy?: number;
  updatedAt?: Date;
  remarks?: string;
  user?: UserModel;
}

export const getUserAccountLogPageApi = (params: ReqPage) => {
  return request.get<ResPage<UserAccountLogModel>>(`/admin-api/v1/app/user/user-account-log`, {
    ...params,
    pageIndex: params?.current
  });
};

export const getUserAccountLogApi = (id: number) => {
  return request.get<UserAccountLogModel>(`/admin-api/v1/app/user/user-account-log/` + id);
};

export const addUserAccountLogApi = (data: object) => {
  return request.post<object>(`/admin-api/v1/app/user/user-account-log`, data);
};

export const updateUserAccountLogApi = (id: number, data: object) => {
  return request.put<object>("/admin-api/v1/app/user/user-account-log/" + id, data);
};

export const delUserAccountLogApi = (params: number[]) => {
  return request.delete<object>(`/admin-api/v1/app/user/user-account-log`, { ids: params });
};

export const exportUserAccountLogApi = (query: object) => {
  return request.download(`/admin-api/v1/app/user/user-account-log/export`, query);
};
