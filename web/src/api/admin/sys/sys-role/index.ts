import request from "@/utils/request";
import { ReqPage, ResPage } from "@/utils/request/interface";

export interface RoleModel {
  createBy?: number;
  createdAt?: Date;
  dataScope?: string;
  id?: number;
  remark?: string;
  roleKey?: string;
  roleName?: string;
  roleSort?: number;
  status?: string;
  updateBy?: number;
  updatedAt?: Date;
  menuIds?: number[];
  deptIds?: number[];
}

export const getRoleTotalListApi = (params: ReqPage) => {
  return request.get<RoleModel[]>(`/admin-api/v1/admin/sys/sys-role/list`, params);
};

export const getRolePageApi = (params: ReqPage) => {
  return request.get<ResPage<RoleModel>>(`/admin-api/v1/admin/sys/sys-role`, { ...params, pageIndex: params?.current });
};

export const getRoleApi = (id: number) => {
  return request.get<RoleModel>(`/admin-api/v1/admin/sys/sys-role/` + id);
};

export const addRoleApi = (data: object) => {
  return request.post<object>(`/admin-api/v1/admin/sys/sys-role`, data);
};

export const updateRoleApi = (id: number, data: object) => {
  return request.put<object>("/admin-api/v1/admin/sys/sys-role/" + id, data);
};

export const delRoleApi = (params: number[]) => {
  return request.delete<object>(`/admin-api/v1/admin/sys/sys-role`, { ids: params });
};

export const dataScopeApi = (data: object) => {
  return request.put<object>("/admin-api/v1/admin/sys/sys-role/role-data-scope", data);
};

export const changeRoleStatusApi = (roleId: number, status: string) => {
  return request.put<object>("/admin-api/v1/admin/sys/sys-role/role-status", { roleId: roleId, status: status });
};
