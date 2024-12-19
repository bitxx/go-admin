import request from "@/utils/request";
import { ReqPage, ResPage } from "@/utils/request/interface";
export interface GenTableModel {
  id: number;
  tableName: string;
  tableComment: string;
  className: string;
  packageName: string;
  businessName: string;
  moduleName: string;
  functionName: string;
  functionAuthor: string;
  createdAt: Date;
}

export interface DBTableModel {
  tableName: string;
  tableComment: string;
  createdAt: Date;
}

export interface PreviewCode {
  name: string;
  path: string;
  content: string;
}

export interface GenTableDetailModel {
  id?: number;
  tableName?: string;
  tableComment?: string;
  className?: string;
  packageName?: string;
  businessName?: string;
  moduleName?: string;
  functionName?: string;
  functionAuthor?: string;
  remark?: string;
  createBy?: number;
  updateBy?: number;
  createdAt?: Date;
  updatedAt?: Date;
  sysGenColumns?: GenTableColumnModel[];
}

export interface GenTableColumnModel {
  id: number;
  tableId: number;
  columnName: string;
  columnComment: string;
  columnType: string;
  goType: string;
  goField: string;
  jsonField: string;
  isPk: string;
  isRequired: string;
  isQuery: string;
  isList: string;
  queryType: string;
  htmlType: string;
  dictType: string;
  sort: number;
  remark: string;
  createBy: number;
  updateBy: number;
  createdAt: Date;
  updatedAt: Date;
}

export const getGenTablePageApi = (params: ReqPage) => {
  return request.get<ResPage<GenTableModel>>(`/admin-api/v1/admin/sys/sys-table`, {
    ...params,
    pageIndex: params?.current
  });
};

export const delGenTableApi = (params: number[]) => {
  return request.delete<object>(`/admin-api/v1/admin/sys/sys-table`, { ids: params });
};

export const genMenuApi = (id: number) => {
  return request.get<object>(`/admin-api/v1/admin/sys/sys-table/gen/db/` + id);
};

export const genCodeApi = (id: number) => {
  return request.get<object>(`/admin-api/v1/admin/sys/sys-table/gen/` + id);
};

export const getGenTableDetailApi = (id: number) => {
  return request.get<GenTableDetailModel>(`/admin-api/v1/admin/sys/sys-table/` + id);
};

export const updateGenTableApi = (id: number, data: GenTableDetailModel) => {
  return request.put<object>("/admin-api/v1/admin/sys/sys-table/" + id, data);
};

export const getDBTablePageApi = (params: ReqPage) => {
  return request.get<ResPage<DBTableModel>>(`/admin-api/v1/admin/sys/sys-table/db-tables`, {
    ...params,
    pageIndex: params?.current
  });
};

export const importDBTableApi = (params: string[]) => {
  return request.post(`/admin-api/v1/admin/sys/sys-table`, { dbTableNames: params });
};

export const downloadCodeApi = (id: number) => {
  return request.download(`/admin-api/v1/admin/sys/sys-table/gen/download/` + id);
};

export const previewTableApi = (id: number) => {
  return request.get<PreviewCode[]>(`/admin-api/v1/admin/sys/sys-table/preview/` + id);
};
