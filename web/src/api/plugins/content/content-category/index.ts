import request from "@/utils/request";
import { ReqPage, ResPage } from "@/utils/request/interface";

export interface ContentCategoryModel {
  id?: number;
  name?: string;
  status?: string;
  remark?: string;
  createBy?: number;
  updateBy?: number;
  updatedAt?: Date;
  createdAt?: Date;
}

export const getContentCategoryPageApi = (params: ReqPage) => {
  return request.get<ResPage<ContentCategoryModel>>(`/admin-api/v1/plugins/content/content-category`, { ...params, pageIndex: params?.current });
};

export const getContentCategoryApi = (id: number) => {
  return request.get<ContentCategoryModel>(`/admin-api/v1/plugins/content/content-category/` + id);
};

export const addContentCategoryApi = (data: object) => {
  return request.post<object>(`/admin-api/v1/plugins/content/content-category`, data);
};

export const updateContentCategoryApi = (id: number, data: object) => {
  return request.put<object>("/admin-api/v1/plugins/content/content-category/" + id, data);
};

export const delContentCategoryApi = (params: number[]) => {
  return request.delete<object>(`/admin-api/v1/plugins/content/content-category`, { ids: params });
};

export const exportContentCategoryApi = (query: object) => {
  return request.download(`/admin-api/v1/plugins/content/content-category/export`, query);
};
