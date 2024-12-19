import request from "@/utils/request";
import { ReqPage, ResPage } from "@/utils/request/interface";

export interface DictTypeModel {
  id?: number;
  dictName?: string;
  dictType?: string;
  status?: string;
  remark?: string;
  createBy?: number;
  updateBy?: number;
  createdAt?: string;
  updatedAt?: string;
}

export const getDictTypePageApi = (params: ReqPage) => {
  return request.get<ResPage<DictTypeModel>>(`/admin-api/v1/admin/sys/sys-dict/type`, { ...params, pageIndex: params?.current });
};

export const getDictTypeApi = (id: number) => {
  return request.get<DictTypeModel>(`/admin-api/v1/admin/sys/sys-dict/type/` + id);
};

export const getAllDictTypesApi = () => {
  return request.get<DictTypeModel[]>(`/admin-api/v1/admin/sys/sys-dict/type/option-select`);
};

export const addDictTypeApi = (data: object) => {
  return request.post<object>(`/admin-api/v1/admin/sys/sys-dict/type`, data);
};

export const updateDictTypeApi = (id: number, data: object) => {
  return request.put<object>("/admin-api/v1/admin/sys/sys-dict/type/" + id, data);
};

export const delDictTypeApi = (params: number[]) => {
  return request.delete<object>(`/admin-api/v1/admin/sys/sys-dict/type`, { ids: params });
};

export const exportDictTypeApi = (query: object) => {
  return request.download(`/admin-api/v1/admin/sys/sys-dict/type/export`, query);
};
