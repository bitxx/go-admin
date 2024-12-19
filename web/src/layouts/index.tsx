import { updateCollapse } from "@/redux/modules/menu/action";
import { Layout } from "antd";
import { useEffect } from "react";
import KeepAlive from "react-activation";
import { connect } from "react-redux";
import { Outlet, useLocation } from "react-router-dom";
import LayoutFooter from "./components/Footer";
import LayoutHeader from "./components/Header";
import LayoutMenu from "./components/Menu";
import LayoutTabs from "./components/Tabs";
import "./index.less";

const LayoutIndex = (props: any) => {
	const { Sider, Content } = Layout;
	const { isCollapse, updateCollapse } = props;
	const { pathname } = useLocation();

	// 获取按钮权限列表
	// const getAuthButtonsList = async () => {
	// 	const { data } = await getAuthorButtons();
	// 	setAuthButtons(data);
	// };

	// 监听窗口大小变化
	const listeningWindow = () => {
		window.onresize = () => {
			return (() => {
				let screenWidth = document.body.clientWidth;
				if (!isCollapse && screenWidth < 1200) updateCollapse(true);
				if (!isCollapse && screenWidth > 1200) updateCollapse(false);
			})();
		};
	};

	useEffect(() => {
		listeningWindow();
		// getAuthButtonsList();
	}, []);

	return (
		// 这里不用 Layout 组件原因是切换页面时样式会先错乱然后在正常显示，造成页面闪屏效果
		<section className="container">
			<Sider trigger={null} collapsed={props.isCollapse} width={220} theme="dark">
				<LayoutMenu></LayoutMenu>
			</Sider>
			<Layout>
				<LayoutHeader></LayoutHeader>
				<LayoutTabs></LayoutTabs>
				<KeepAlive name={pathname} when={false}>
					<Content>
						<Outlet></Outlet>
					</Content>
				</KeepAlive>
				<LayoutFooter></LayoutFooter>
			</Layout>
		</section>
	);
};

const mapStateToProps = (state: any) => state.menu;
const mapDispatchToProps = { updateCollapse };
export default connect(mapStateToProps, mapDispatchToProps)(LayoutIndex);
