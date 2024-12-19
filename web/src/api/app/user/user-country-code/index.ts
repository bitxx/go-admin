import request from "@/utils/request";
import { ReqPage, ResPage } from "@/utils/request/interface";

export interface UserCountryCodeModel {
  id?: number;
  country?: string;
  code?: string;
  status?: string;
  remark?: string;
  createBy?: number;
  updateBy?: number;
  createdAt?: Date;
  updatedAt?: Date;
}

export const getUserCountryCodePageApi = (params: ReqPage) => {
  return request.get<ResPage<UserCountryCodeModel>>(`/admin-api/v1/app/user/user-country-code`, { ...params, pageIndex: params?.current });
};

export const getUserCountryCodeApi = (id: number) => {
  return request.get<UserCountryCodeModel>(`/admin-api/v1/app/user/user-country-code/` + id);
};

export const addUserCountryCodeApi = (data: object) => {
  return request.post<object>(`/admin-api/v1/app/user/user-country-code`, data);
};

export const updateUserCountryCodeApi = (id: number, data: object) => {
  return request.put<object>("/admin-api/v1/app/user/user-country-code/" + id, data);
};

export const delUserCountryCodeApi = (params: number[]) => {
  return request.delete<object>(`/admin-api/v1/app/user/user-country-code`, { ids: params });
};

export const exportUserCountryCodeApi = (query: object) => {
  return request.download(`/admin-api/v1/app/user/user-country-code/export`, query);
};
