import request from "@/utils/request";
import { ReqPage, ResPage } from "@/utils/request/interface";

export interface PostModel {
  createBy?: number;
  createdAt?: Date;
  id?: number;
  postCode?: string;
  postName?: string;
  remark?: string;
  sort?: number;
  status?: string;
  updateBy?: number;
  updatedAt?: Date;
}

export const getPostTotalListApi = (params: ReqPage) => {
  return request.get<PostModel[]>(`/admin-api/v1/admin/sys/sys-post/list`, params);
};

export const getPostPageApi = (params: ReqPage) => {
  return request.get<ResPage<PostModel>>(`/admin-api/v1/admin/sys/sys-post`, { ...params, pageIndex: params?.current });
};

export const getPostApi = (id: number) => {
  return request.get<PostModel>(`/admin-api/v1/admin/sys/sys-post/` + id);
};

export const addPostApi = (data: object) => {
  return request.post<object>(`/admin-api/v1/admin/sys/sys-post`, data);
};

export const updatePostApi = (id: number, data: object) => {
  return request.put<object>("/admin-api/v1/admin/sys/sys-post/" + id, data);
};

export const delPostApi = (params: number[]) => {
  return request.delete<object>(`/admin-api/v1/admin/sys/sys-post`, { ids: params });
};

export const exportPostApi = (query: object) => {
  return request.download(`/admin-api/v1/admin/sys/sys-post/export`, query);
};
