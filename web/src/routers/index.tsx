import { RouteObjectType } from "@/api/admin/sys/sys-menu";
import { SysMenuType, SysStatus } from "@/enums/base";
import useMessage from "@/hooks/useMessage";
import { store } from "@/redux";
import { parseFlatMenuList } from "@/utils/util";
import Login from "@/views/admin/sys/login/index";
import React, { lazy, useEffect, useState } from "react";
import { Navigate, useRoutes } from "react-router-dom";
import { LayoutIndex } from "./constant";
import lazyLoad from "./utils/lazyLoad";

export const rootRouter: RouteObjectType[] = [
	{
		path: "/",
		element: <Navigate to="/login" />
	},
	{
		path: "/login",
		element: <Login />,
		title: "登录页"
	},
	{
		path: "/403",
		element: lazyLoad(React.lazy(() => import("@/components/ErrorMessage/403"))),
		title: "403页面"
	},
	{
		path: "/404",
		element: lazyLoad(React.lazy(() => import("@/components/ErrorMessage/404"))),
		title: "404页面"
	},
	{
		path: "/500",
		element: lazyLoad(React.lazy(() => import("@/components/ErrorMessage/500"))),
		title: "500页面"
	}
];

const Router = () => {
	const [routerList, setRouterList] = useState<RouteObjectType[]>(rootRouter);
	const rList: RouteObjectType[] = store.getState().global.routeList;
	useMessage();
	// const token: string = store.getState().global.token;
	// const uInfo: LoginUserInfo = store.getState().global.userinfo;
	useEffect(() => {
		if (rList && rList.length > 0) {
			setRouterList([...rootRouter, ...dynamicRouter(rList)]);
		} else {
			setRouterList([...rootRouter]);
		}
	}, [rList]);

	const routes = useRoutes(routerList);
	return routes;
};

const modules = import.meta.glob("@/views/**/*.tsx") as Record<string, Parameters<typeof lazy>[number]>;
// console.log(modules["../views/admin/sys/home/index.tsx"]);

export const dynamicRouter = (mList: RouteObjectType[]) => {
	const list = parseFlatMenuList(mList);
	const handleMenuList = list.map(item => {
		item.children && delete item.children;
		if (item.redirect) item.element = <Navigate to={item.redirect} />;
		if (item.element && typeof item.element === "string") {
			// let ip: string = "../views" + String(item.element);
			item.element = lazyLoad(lazy(modules["/src/views" + item.element + ".tsx"]));
		}
		return item;
	});

	const dynamicRouter: RouteObjectType[] = [{ element: <LayoutIndex />, children: [] }];
	handleMenuList.forEach(item => {
		if (item.isFrame == SysStatus.FALSE && item.menuType == SysMenuType.MENU) dynamicRouter.push(item);
		else if (item.menuType == SysMenuType.MENU || item.menuType == SysMenuType.DIRECT) dynamicRouter[0].children?.push(item);
	});
	return dynamicRouter;
};

export default Router;
