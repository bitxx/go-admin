import request from "@/utils/request";
import { ReqPage, ResPage } from "@/utils/request/interface";

export interface OperLogModel {
  createBy?: number;
  createdAt?: Date;
  id?: number;
  jsonResult?: string;
  latencyTime?: string;
  operIp?: string;
  operLocation?: string;
  operTime?: Date;
  operUrl?: string;
  remark?: string;
  requestMethod?: string;
  status?: string;
  updateBy?: number;
  updatedAt?: Date;
  userAgent?: string;
  userId?: number;
}

export const getOperLogPageApi = (params: ReqPage) => {
  return request.get<ResPage<OperLogModel>>(`/admin-api/v1/admin/sys/sys-oper-log`, { ...params, pageIndex: params?.current });
};

export const getOperLogApi = (id: number) => {
  return request.get<OperLogModel>(`/admin-api/v1/admin/sys/sys-oper-log/` + id);
};

export const delOperLogApi = (params: number[]) => {
  return request.delete<object>(`/admin-api/v1/admin/sys/sys-oper-log`, { ids: params });
};

export const exportOperLogApi = (query: object) => {
  return request.download(`/admin-api/v1/admin/sys/sys-oper-log/export`, query);
};
