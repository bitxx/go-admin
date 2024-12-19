import request from "@/utils/request";
import { ReqPage, ResPage } from "@/utils/request/interface";

export interface DictDataModel {
  id?: number;
  dictSort?: number;
  dictLabel?: string;
  dictValue?: string;
  dictType?: string;
  cssClass?: string;
  listClass?: string;
  isDefault?: string;
  status?: string;
  default?: string;
  remark?: string;
  createBy?: number;
  updateBy?: number;
  createdAt?: Date;
  updatedAt?: Date;
}

export const getDictDataPageApi = (params: ReqPage, dictType: string) => {
  return request.get<ResPage<DictDataModel>>(`/admin-api/v1/admin/sys/sys-dict/data?dictType=` + dictType, {
    ...params,
    pageIndex: params?.current
  });
};

export const getDictDataApi = (id: number) => {
  return request.get<DictDataModel>(`/admin-api/v1/admin/sys/sys-dict/data/` + id);
};

export const addDictDataApi = (data: object) => {
  return request.post<object>(`/admin-api/v1/admin/sys/sys-dict/data`, data);
};

export const updateDictDataApi = (id: number, data: DictDataModel) => {
  return request.put<object>("/admin-api/v1/admin/sys/sys-dict/data/" + id, data);
};

export const delDictDataApi = (params: number[]) => {
  return request.delete<object>(`/admin-api/v1/admin/sys/sys-dict/data`, { ids: params });
};

export const getDictsApi = (dictType: string) => {
  return request.get<DictDataModel[]>(`/admin-api/v1/admin/sys/sys-dict/data/select?dictType=` + dictType);
};

export const getDictOptions = (datas: DictDataModel[]) => {
  return new Map(datas.map(({ dictValue, dictLabel }) => [dictValue || "", dictLabel || ""]));
};
