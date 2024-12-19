import request from "@/utils/request";
import { ReqPage, ResPage } from "@/utils/request/interface";

export interface UserLevelModel {
  id?: number;
  name?: string;
  levelType?: string;
  level?: number;
  status?: string;
  remark?: string;
  createBy?: number;
  updateBy?: number;
  createdAt?: Date;
  updatedAt?: Date;
}

export const getUserLevelPageApi = (params: ReqPage) => {
  return request.get<ResPage<UserLevelModel>>(`/admin-api/v1/app/user/user-level`, { ...params, pageIndex: params?.current });
};

export const getUserLevelApi = (id: number) => {
  return request.get<UserLevelModel>(`/admin-api/v1/app/user/user-level/` + id);
};

export const addUserLevelApi = (data: object) => {
  return request.post<object>(`/admin-api/v1/app/user/user-level`, data);
};

export const updateUserLevelApi = (id: number, data: object) => {
  return request.put<object>("/admin-api/v1/app/user/user-level/" + id, data);
};

export const delUserLevelApi = (params: number[]) => {
  return request.delete<object>(`/admin-api/v1/app/user/user-level`, { ids: params });
};

export const exportUserLevelApi = (query: object) => {
  return request.download(`/admin-api/v1/app/user/user-level/export`, query);
};
