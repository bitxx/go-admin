import request from "@/utils/request";
import { ReqPage, ResPage } from "@/utils/request/interface";

export interface ContentArticleModel {
  id?: number;
  cateId?: number;
  name?: string;
  content?: string;
  remark?: string;
  status?: string;
  createBy?: number;
  updateBy?: number;
  updatedAt?: Date;
  createdAt?: Date;
}

export const getContentArticlePageApi = (params: ReqPage) => {
  return request.get<ResPage<ContentArticleModel>>(`/admin-api/v1/plugins/content/content-article`, { ...params, pageIndex: params?.current });
};

export const getContentArticleApi = (id: number) => {
  return request.get<ContentArticleModel>(`/admin-api/v1/plugins/content/content-article/` + id);
};

export const addContentArticleApi = (data: object) => {
  return request.post<object>(`/admin-api/v1/plugins/content/content-article`, data);
};

export const updateContentArticleApi = (id: number, data: object) => {
  return request.put<object>("/admin-api/v1/plugins/content/content-article/" + id, data);
};

export const delContentArticleApi = (params: number[]) => {
  return request.delete<object>(`/admin-api/v1/plugins/content/content-article`, { ids: params });
};

export const exportContentArticleApi = (query: object) => {
  return request.download(`/admin-api/v1/plugins/content/content-article/export`, query);
};
