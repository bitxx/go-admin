import request from "@/utils/request";
import { ReqPage, ResPage } from "@/utils/request/interface";

export interface MsgCodeModel {
  id?: number;
  userId?: number;
  code?: string;
  codeType?: string;
  remark?: string;
  status?: string;
  createBy?: number;
  updateBy?: number;
  createdAt?: Date;
  updatedAt?: Date;
}

export const getMsgCodePageApi = (params: ReqPage) => {
  return request.get<ResPage<MsgCodeModel>>(`/admin-api/v1/plugins/msg/msg-code`, { ...params, pageIndex: params?.current });
};

export const getMsgCodeApi = (id: number) => {
  return request.get<MsgCodeModel>(`/admin-api/v1/plugins/msg/msg-code/` + id);
};

export const addMsgCodeApi = (data: object) => {
  return request.post<object>(`/admin-api/v1/plugins/msg/msg-code`, data);
};

export const updateMsgCodeApi = (id: number, data: object) => {
  return request.put<object>("/admin-api/v1/plugins/msg/msg-code/" + id, data);
};

export const delMsgCodeApi = (params: number[]) => {
  return request.delete<object>(`/admin-api/v1/plugins/msg/msg-code`, { ids: params });
};

export const exportMsgCodeApi = (query: object) => {
  return request.download(`/admin-api/v1/plugins/msg/msg-code/export`, query);
};
