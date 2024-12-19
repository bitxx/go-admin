import { HOME_URL, LOGIN_URL } from "@/config";
import { store } from "@/redux/index";
import { AxiosCanceler } from "@/utils/request/helper/axiosCancel";
import { useEffect } from "react";
import { useLocation, useNavigate } from "react-router-dom";

const axiosCanceler = new AxiosCanceler();

/**
 * @description 路由守卫组件
 * */
const AuthRouter = (props: { children: JSX.Element }) => {
	const { pathname } = useLocation();
	const navigate = useNavigate();
	const token = store.getState().global.token;

	useEffect(() => {
		axiosCanceler.removeAllPending();
		const update = async () => {
			if (!token) {
				return navigate(LOGIN_URL);
			}
			if (token && pathname === LOGIN_URL) {
				return navigate(HOME_URL);
			}

			if (!token && pathname !== LOGIN_URL) {
				return navigate(LOGIN_URL, { replace: true });
			}
		};
		update();
	}, [token, pathname]);
	// * 当前账号有权限返回 Router，正常访问页面
	return props.children;
};

export default AuthRouter;
