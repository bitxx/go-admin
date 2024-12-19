import { LoginUserInfo } from "@/api/admin/sys/sys-user";
import { store } from "@/redux";
import { Layout } from "antd";
import AvatarIcon from "./components/AvatarIcon";
import BreadcrumbNav from "./components/BreadcrumbNav";
import CollapseIcon from "./components/CollapseIcon";
import Fullscreen from "./components/Fullscreen";
import Language from "./components/Language";
import "./index.less";

const LayoutHeader = () => {
	const { Header } = Layout;
	const uInfo: LoginUserInfo = store.getState().global.userInfo;

	return (
		<Header style={{ background: "#fff", padding: 12 }}>
			<div className="header-lf">
				<CollapseIcon />
				<BreadcrumbNav />
			</div>
			<div className="header-ri">
				<Language />
				<Fullscreen />
				<span className="username">{uInfo.username}</span>
				<AvatarIcon />
			</div>
		</Header>
	);
};

export default LayoutHeader;
