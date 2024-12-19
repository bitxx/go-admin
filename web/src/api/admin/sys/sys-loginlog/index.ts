import request from "@/utils/request";
import { ReqPage, ResPage } from "@/utils/request/interface";

export interface LoginLogModel {
  agent?: string;
  browser?: string;
  createBy?: number;
  createdAt?: Date;
  id?: number;
  ipaddr?: string;
  loginLocation?: string;
  loginTime?: Date;
  os?: string;
  platform?: string;
  remark?: string;
  status?: string;
  updateBy?: number;
  updatedAt?: Date;
  userId?: number;
}

export const getLoginLogPageApi = (params: ReqPage) => {
  return request.get<ResPage<LoginLogModel>>(`/admin-api/v1/admin/sys/sys-login-log`, { ...params, pageIndex: params?.current });
};

export const getLoginLogApi = (id: number) => {
  return request.get<LoginLogModel>(`/admin-api/v1/admin/sys/sys-login-log/` + id);
};

export const addLoginLogApi = (data: object) => {
  return request.post<object>(`/admin-api/v1/admin/sys/sys-login-log`, data);
};

export const updateLoginLogApi = (id: number, data: object) => {
  return request.put<object>("/admin-api/v1/admin/sys/sys-login-log/" + id, data);
};

export const delLoginLogApi = (params: number[]) => {
  return request.delete<object>(`/admin-api/v1/admin/sys/sys-login-log`, { ids: params });
};

export const exportLoginLogApi = (query: object) => {
  return request.download(`/admin-api/v1/admin/sys/sys-login-log/export`, query);
};
