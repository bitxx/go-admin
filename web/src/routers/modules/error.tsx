import { RouteObjectType } from "@/api/admin/sys/sys-menu";
import lazyLoad from "@/routers/utils/lazyLoad";
import React from "react";

// 错误页面模块
const errorRouter: Array<RouteObjectType> = [
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

export default errorRouter;
