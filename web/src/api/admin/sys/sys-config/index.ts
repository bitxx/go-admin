import request from "@/utils/request";
import { ReqPage, ResPage } from "@/utils/request/interface";

export interface ConfigModel {
  configKey?: string;
  configName?: string;
  configType?: string;
  configValue?: string;
  createBy?: number;
  createdAt?: Date;
  id?: number;
  isFrontend?: string;
  remark?: string;
  updateBy?: number;
  updatedAt?: Date;
}

export const getConfigPageApi = (params: ReqPage) => {
  return request.get<ResPage<ConfigModel>>(`/admin-api/v1/admin/sys/sys-config`, { ...params, pageIndex: params?.current });
};

export const getConfigApi = (id: number) => {
  return request.get<ConfigModel>(`/admin-api/v1/admin/sys/sys-config/` + id);
};

export const addConfigApi = (data: object) => {
  return request.post<object>(`/admin-api/v1/admin/sys/sys-config`, data);
};

export const updateConfigApi = (id: number, data: object) => {
  return request.put<object>("/admin-api/v1/admin/sys/sys-config/" + id, data);
};

export const delConfigApi = (params: number[]) => {
  return request.delete<object>(`/admin-api/v1/admin/sys/sys-config`, { ids: params });
};

export const exportConfigApi = (query: object) => {
  return request.download(`/admin-api/v1/admin/sys/sys-config/export`, query);
};

export const getConfigByKey = (key: string) => {
  return request.get<ConfigModel>(`/admin-api/v1/admin/sys/sys-config/` + key);
};
