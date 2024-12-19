import request from "@/utils/request";

export interface DeptModel {
  createBy?: number;
  createdAt?: Date;
  deptName?: string;
  deptPath?: string;
  email?: string;
  id?: number;
  leader?: string;
  parentId?: number;
  phone?: string;
  sort?: number;
  status?: number;
  updateBy?: number;
  updatedAt?: Date;
  children?: DeptModel[];
}

export interface DeptTreeRole {
  depts?: DeptModel[];
  checkedKeys?: number[];
}

export const getDeptTreeApi = (params: object) => {
  return request.get<DeptModel[]>(`/admin-api/v1/admin/sys/sys-dept/dept-tree`, params);
};

export const roleDeptTreeselectApi = (roleId: number) => {
  return request.get<DeptTreeRole>(`/admin-api/v1/admin/sys/sys-dept/role-dept-tree-select/` + roleId);
};

export const getDeptApi = (id: number) => {
  return request.get<DeptModel>(`/admin-api/v1/admin/sys/sys-dept/` + id);
};

export const addDeptApi = (data: object) => {
  return request.post<object>(`/admin-api/v1/admin/sys/sys-dept`, data);
};

export const updateDeptApi = (id: number, data: object) => {
  return request.put<object>("/admin-api/v1/admin/sys/sys-dept/" + id, data);
};

export const delDeptApi = (params: number[]) => {
  return request.delete<object>(`/admin-api/v1/admin/sys/sys-dept`, { ids: params });
};

export const exportDeptApi = (query: object) => {
  return request.download(`/admin-api/v1/admin/sys/sys-dept`, query);
};
