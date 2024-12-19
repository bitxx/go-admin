import { RouteObject } from "@/api/admin/sys/sys-menu";
import { LoginUserInfo } from "@/api/admin/sys/sys-user";
import { ThemeConfigProp } from "@/redux/interface/index";
import * as types from "@/redux/mutation-types";

// * setToken
export const setToken = (token: string) => ({
	type: types.SET_TOKEN,
	token
});

// * setAssemblySize
export const setAssemblySize = (assemblySize: string) => ({
	type: types.SET_ASSEMBLY_SIZE,
	assemblySize
});

export const setUserInfo = (userInfo: LoginUserInfo) => {
	if (!userInfo.avatar?.startsWith("http")) {
		userInfo.avatar = import.meta.env.VITE_API_URL + userInfo.avatar;
	}

	return {
		type: types.SET_USER_INFO,
		userInfo
	};
};

export const setRouteList = (routeList: RouteObject[]) => ({
	type: types.SET_ROUTE_LIST,
	routeList
});

// * setLanguage
export const setLanguage = (language: string) => ({
	type: types.SET_LANGUAGE,
	language
});

// * setThemeConfig
export const setThemeConfig = (themeConfig: ThemeConfigProp) => ({
	type: types.SET_THEME_CONFIG,
	themeConfig
});
