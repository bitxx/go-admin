import request from "@/utils/request/index";
import { ReqPage, ResPage } from "@/utils/request/interface";
import { MenuModel } from "../sys-menu";

export interface ApiModel {
	method?: string;
	apiType?: string;
	createBy?: number;
	createdAt?: Date;
	id?: number;
	path?: string;
	description?: string;
	updateBy?: number;
	updatedAt?: Date;
	sysMenu?: MenuModel[];
}

export const getApiPageApi = (params: ReqPage) => {
	return request.get<ResPage<ApiModel>>(`/admin-api/v1/admin/sys/sys-api`, { ...params, pageIndex: params?.current });
};

export const getApiListApi = (params: ReqPage) => {
	return request.get<ApiModel[]>(`/admin-api/v1/admin/sys/sys-api/list`, { ...params });
};

export const getApiApi = (id: number) => {
	return request.get<ApiModel>(`/admin-api/v1/admin/sys/sys-api/` + id);
};

export const addApiApi = (data: object) => {
	return request.post<object>(`/admin-api/v1/admin/sys/sys-api`, data);
};

export const updateApiApi = (id: number, data: object) => {
	return request.put<object>("/admin-api/v1/admin/sys/sys-api/" + id, data);
};

export const delApiApi = (params: number[]) => {
	return request.delete<object>(`/admin-api/v1/admin/sys/sys-api`, { ids: params });
};

export const exportApiApi = (query: object) => {
	return request.download(`/admin-api/v1/admin/sys/sys-api/export`, query);
};

export const syncApiApi = () => {
	return request.get<object>(`/admin-api/v1/admin/sys/sys-api/sync`);
};

export const getSyncStatusApiApi = () => {
	return request.get<string>(`/admin-api/v1/admin/sys/sys-api/sync/status`);
};
