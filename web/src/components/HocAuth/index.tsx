import { LoginUserInfo } from "@/api/admin/sys/sys-user";
import { store } from "@/redux";
import React from "react";

type AuthButtonProps = {
	permission: string[];
	children: React.ReactNode;
};

const AuthButton: React.FC<AuthButtonProps> = ({ permission, children }) => {
	const uInfo: LoginUserInfo = store.getState().global.userInfo;
	let isAuth = false;
	if (uInfo && permission instanceof Array && permission.length) {
		const hasPermission = permission.every(item => uInfo.permissions?.includes(item));
		hasPermission && (isAuth = true);
	}
	if (uInfo.permissions?.length == 1 && uInfo.permissions?.includes("*:*:*")) {
		isAuth = true;
	}

	return <React.Fragment>{isAuth && children}</React.Fragment>;
};

export default React.memo(AuthButton);
