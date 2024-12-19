import request from "@/utils/request";
import { ReqPage, ResPage } from "@/utils/request/interface";
import md5 from "js-md5";
import { DeptModel } from "../sys-dept";
import { PostModel } from "../sys-post";
import { RoleModel } from "../sys-role";

export interface ReqLogin {
	username?: string;
	password?: string;
	uuid?: string;
	code?: string;
}

export interface RespLogin {
	token: string;
	username: string;
}

export interface Captcha {
	id: string;
	data: string;
}

export interface LoginUserInfo {
	id?: number;
	username?: string;
	avatar?: string;
	phone?: string;
	sex?: string;
	email?: string;
	deptName?: string;
	roleName?: string;
	permissions?: string[];
	roleKyes?: string[];
	createdAt?: string;
}

export interface PassWdChange {
	oldPassword?: string;
	newPassword?: string;
}

export interface UserInfoChange {
	id?: number;
	username?: string;
	phone?: string;
	email?: string;
	sex?: string;
}

export interface UserModel {
	avatar?: string;
	createBy?: number;
	createdAt?: Date;
	deptId?: number;
	email?: string;
	id?: number;
	nickName?: string;
	password?: string;
	phone?: string;
	postId?: number;
	remark?: string;
	roleId?: number;
	salt?: string;
	sex?: string;
	status?: string;
	updateBy?: number;
	updatedAt?: Date;
	username?: string;
	dept?: DeptModel;
	role?: RoleModel;
	post?: PostModel;
}

// login
export const loginApi = (params: ReqLogin) => {
	return request.post<RespLogin>(`/admin-api/v1/login`, { ...params, password: md5(params.password!) });
};

// get captcha
export const getCaptchaApi = () => {
	return request.get<Captcha>(`/admin-api/v1/captcha`);
};

// User logout
export const logoutApi = () => {
	return request.get(`/admin-api/v1/admin/sys/sys-user/logout`, {}, { loading: true });
};

// get user info
export const getUserProfileApi = () => {
	return request.get<LoginUserInfo>(`/admin-api/v1/admin/sys/sys-user/profile`);
};

// update password
export const updateProfilePwdApi = (passwdChange: PassWdChange) => {
	return request.put(`/admin-api/v1/admin/sys/sys-user/profile/pwd`, {
		...passwdChange,
		oldPassword: md5(passwdChange.oldPassword!),
		newPassword: md5(passwdChange.newPassword!)
	});
};

// update user info
export const updateProfileInfoApi = (userInfoChange: UserInfoChange) => {
	return request.put(`/admin-api/v1/admin/sys/sys-user/profile`, userInfoChange);
};

// update avatar
export const updateProfileAvatar = (params: FormData) => {
	return request.post<string>(`/admin-api/v1/admin/sys/sys-user/profile/avatar`, params);
};

export const getUserPageApi = (params: ReqPage) => {
	return request.get<ResPage<UserModel>>(`/admin-api/v1/admin/sys/sys-user`, { ...params, pageIndex: params?.current });
};

export const getUserApi = (id: number) => {
	return request.get<UserModel>(`/admin-api/v1/admin/sys/sys-user/` + id);
};

export const addUserApi = (data: object) => {
	return request.post<object>(`/admin-api/v1/admin/sys/sys-user`, data);
};

export const updateUserApi = (id: number, data: object) => {
	return request.put<object>("/admin-api/v1/admin/sys/sys-user/" + id, data);
};

export const delUserApi = (params: number[]) => {
	return request.delete<object>(`/admin-api/v1/admin/sys/sys-user`, { ids: params });
};

export const exportUserApi = (query: object) => {
	return request.download(`/admin-api/v1/admin/sys/sys-user/export`, query);
};

export const changeUserPwdApi = (userId: number, password: string) => {
	return request.put<object>("/admin-api/v1/admin/sys/sys-user/pwd-reset", { userId, password: md5(password!) });
};

export const changeUserStatusApi = (roleId: number, status: string) => {
	return request.put<object>("/admin-api/v1/admin/sys/sys-user/update-status", { userId: roleId, status: status });
};
