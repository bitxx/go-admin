import { LoginUserInfo } from "@/api/admin/sys/sys-user";
import { store } from "@/redux";
import { setUserInfo as setUInfo } from "@/redux/modules/global/action";

export const setUserInfo = (uInfo: LoginUserInfo) => {
	const dispatch = store.dispatch;
	dispatch(setUInfo(uInfo));
};
