import request from "@/utils/request";
import { ReqPage } from "@/utils/request/interface";

export interface RouteObject {
	id?: number;
	title?: string;
	icon?: string;
	redirect?: string; //针对目录跳转，比如搜索出菜单
	permission?: string;
	parentIds?: string;
	index?: false;
	parentId?: number;
	sort?: number;
	menuType?: string;
	isKeepAlive?: string;
	isAffix?: string;
	isHidden?: string;
	isFrame?: string;
	path?: string;
	caseSensitive?: boolean;
	element?: React.ReactNode;
	children?: RouteObject[];
}

export interface MenuModel {
	createBy?: number;
	createdAt?: Date;
	element?: string;
	icon?: string;
	id?: number;
	isAffix?: string;
	isFrame?: string;
	isHidden?: string;
	isKeepAlive?: string;
	menuType?: string;
	parentId?: number;
	parentIds?: string;
	path?: string;
	permission?: string;
	redirect?: string;
	sort?: number;
	title?: string;
	updateBy?: number;
	updatedAt?: Date;
	apis?: number[];
	children?: MenuModel[];
}

export interface MenuTreeRole {
	menus?: MenuModel[];
	checkedKeys?: number[];
}

export const getMenuRoleApi = () => {
	return request.get<RouteObject[]>(`/admin-api/v1/admin/sys/sys-menu/menu-role`);
};

export const roleMenuTreeselectApi = (id: number) => {
	return request.get<MenuTreeRole>(`/admin-api/v1/admin/sys/sys-menu/role-menu-tree-select/` + id);
};

export const getMenuListApi = (params: ReqPage) => {
	return request.get<MenuModel[]>(`/admin-api/v1/admin/sys/sys-menu`, params);
};

export const getMenuApi = (id: number) => {
	return request.get<MenuModel>(`/admin-api/v1/admin/sys/sys-menu/` + id);
};

export const addMenuApi = (data: object) => {
	return request.post<object>(`/admin-api/v1/admin/sys/sys-menu`, data);
};

export const updateMenuApi = (id: number, data: object) => {
	return request.put<object>("/admin-api/v1/admin/sys/sys-menu/" + id, data);
};

export const delMenuApi = (params: number[]) => {
	return request.delete<object>(`/admin-api/v1/admin/sys/sys-menu`, { ids: params });
};

export const exportMenuApi = (query: object) => {
	return request.download(`/admin-api/v1/admin/sys/sys-menu`, query);
};
