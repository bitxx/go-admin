import { RouteObjectType } from "@/api/admin/sys/sys-menu";
import { SysStatus } from "@/enums/base";
import { store } from "@/redux";
import { setAuthRouter } from "@/redux/modules/auth/action";
import { setBreadcrumbList } from "@/redux/modules/breadcrumb/action";
import { setMenuList } from "@/redux/modules/menu/action";
import { findAllBreadcrumb, getOpenKeys, handleRouter, searchRoute } from "@/utils/util";
import * as Icons from "@ant-design/icons";
import type { MenuProps } from "antd";
import { Menu, Spin } from "antd";
import React, { useEffect, useState } from "react";
import { connect } from "react-redux";
import { useLocation, useNavigate } from "react-router-dom";
import Logo from "./components/Logo";
import "./index.less";

const LayoutMenu = (props: any) => {
	const { pathname } = useLocation();
	const { isCollapse, setBreadcrumbList, setAuthRouter, setMenuList: setMenuListAction } = props;
	const [selectedKeys, setSelectedKeys] = useState<string[]>([pathname]);
	const [openKeys, setOpenKeys] = useState<string[]>([]);
	const rList: RouteObjectType[] = store.getState().global.routeList;

	// 刷新页面菜单保持高亮
	useEffect(() => {
		setSelectedKeys([pathname]);
		isCollapse ? null : setOpenKeys(getOpenKeys(pathname));
	}, [pathname, isCollapse]);

	// 设置当前展开的 subMenu
	const onOpenChange = (openKeys: string[]) => {
		if (openKeys.length === 0 || openKeys.length === 1) return setOpenKeys(openKeys);
		const latestOpenKey = openKeys[openKeys.length - 1];
		if (latestOpenKey.includes(openKeys[0])) return setOpenKeys(openKeys);
		setOpenKeys([latestOpenKey]);
	};

	// 定义 menu 类型
	type MenuItem = Required<MenuProps>["items"][number];
	const getItem = (
		label: React.ReactNode,
		key?: React.Key | null,
		icon?: React.ReactNode,
		children?: MenuItem[],
		type?: "group"
	): MenuItem => {
		return {
			key,
			icon,
			children,
			label,
			type
		} as MenuItem;
	};

	// 动态渲染 Icon 图标
	const customIcons: { [key: string]: any } = Icons;
	const addIcon = (name: string) => {
		return React.createElement(customIcons[name]);
	};

	// 处理后台返回菜单 key 值为 antd 菜单需要的 key 值
	const deepLoopFloat = (menuList: RouteObjectType[], newArr: MenuItem[] = []) => {
		menuList.forEach((item: RouteObjectType) => {
			if (item.isHidden === SysStatus.FALSE) {
				if (!item?.children?.length) return newArr.push(getItem(item.title, item.path, addIcon(item.icon!)));
				newArr.push(getItem(item.title, item.path, addIcon(item.icon!), deepLoopFloat(item.children)));
			}
			// 下面判断代码解释 *** !item?.children?.length   ==>   (!item.children || item.children.length === 0)
		});
		return newArr;
	};

	// 获取菜单列表并处理成 antd menu 需要的格式
	const [menuList, setMenuList] = useState<MenuItem[]>([]);
	const [loading, setLoading] = useState(false);
	const getMenuData = async () => {
		setLoading(true);
		try {
			setMenuList(deepLoopFloat(rList));
			// 存储处理过后的所有面包屑导航栏到 redux 中
			setBreadcrumbList(findAllBreadcrumb(rList));
			// 把路由菜单处理成一维数组，存储到 redux 中，做菜单权限判断
			const dynamicRouter = handleRouter(rList);
			setAuthRouter(dynamicRouter);
			setMenuListAction(rList);
		} finally {
			setLoading(false);
		}
	};
	useEffect(() => {
		getMenuData();
	}, []);

	// 点击当前菜单跳转页面
	const navigate = useNavigate();
	const clickMenu: MenuProps["onClick"] = ({ key }: { key: string }) => {
		const route = searchRoute(key, props.menuList);
		if (route.isFrame === SysStatus.FALSE) window.open(route.isFrame, "_blank");
		navigate(key);
	};

	return (
		<div className="menu">
			<Spin spinning={loading} tip="Loading...">
				<Logo></Logo>
				<Menu
					theme="dark"
					mode="inline"
					triggerSubMenuAction="click"
					openKeys={openKeys}
					selectedKeys={selectedKeys}
					items={menuList}
					onClick={clickMenu}
					onOpenChange={onOpenChange}
				></Menu>
			</Spin>
		</div>
	);
};

const mapStateToProps = (state: any) => state.menu;
const mapDispatchToProps = { setMenuList, setBreadcrumbList, setAuthRouter };
export default connect(mapStateToProps, mapDispatchToProps)(LayoutMenu);
