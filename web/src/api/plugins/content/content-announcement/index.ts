import request from "@/utils/request";
import { ReqPage, ResPage } from "@/utils/request/interface";

export interface ContentAnnouncementModel {
  id?: number;
  title?: string;
  content?: string;
  num?: number;
  remark?: string;
  status?: string;
  createBy?: number;
  updateBy?: number;
  updatedAt?: Date;
  createdAt?: Date;
}

export const getContentAnnouncementPageApi = (params: ReqPage) => {
  return request.get<ResPage<ContentAnnouncementModel>>(`/admin-api/v1/plugins/content/content-announcement`, { ...params, pageIndex: params?.current });
};

export const getContentAnnouncementApi = (id: number) => {
  return request.get<ContentAnnouncementModel>(`/admin-api/v1/plugins/content/content-announcement/` + id);
};

export const addContentAnnouncementApi = (data: object) => {
  return request.post<object>(`/admin-api/v1/plugins/content/content-announcement`, data);
};

export const updateContentAnnouncementApi = (id: number, data: object) => {
  return request.put<object>("/admin-api/v1/plugins/content/content-announcement/" + id, data);
};

export const delContentAnnouncementApi = (params: number[]) => {
  return request.delete<object>(`/admin-api/v1/plugins/content/content-announcement`, { ids: params });
};

export const exportContentAnnouncementApi = (query: object) => {
  return request.download(`/admin-api/v1/plugins/content/content-announcement/export`, query);
};
