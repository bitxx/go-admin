import { getDictOptions, getDictsApi } from "@/api/admin/sys/sys-dictdata";
import { exportUserApi, getUserPageApi, UserModel } from "@/api/app/user/user";
import HocAuth from "@/components/HocAuth";
import LoadingButton from "@/components/LoadingButton";
import { pagination } from "@/config/proTable";
import { ResultEnum } from "@/enums/httpEnum";
import { message, modal } from "@/hooks/useMessage";
import { formatDataForProTable, saveExcelBlob } from "@/utils";
import { CloudDownloadOutlined, EditOutlined, ExclamationCircleOutlined, PlusCircleOutlined } from "@ant-design/icons";
import type { ActionType, ProColumns, ProFormInstance } from "@ant-design/pro-components";
import { ProTable } from "@ant-design/pro-components";
import { Space } from "antd";
import React, { useEffect, useRef, useState } from "react";
import FormModal, { FormModalRef } from "./components/FormModal";

const User: React.FC = () => {
	const actionRef = React.useRef<ActionType>();
	const tableFormRef = React.useRef<ProFormInstance>();
	const formModalRef = useRef<FormModalRef>(null);
	const [statusOptions, setStatusOptions] = useState<Map<string, string>>(new Map());
	const [levelTypeOptions, setLevelTypeOptions] = useState<Map<string, string>>(new Map());

	// 定义列
	const columns: ProColumns<UserModel>[] = [
		{
			title: "序号",
			dataIndex: "index",
			valueType: "index",
			width: 50,
			align: "center",
			className: "gray-cell",
			render: (_, __, index, action) => {
				// 根据分页计算实际序号
				const currentPage = action?.pageInfo?.current || 1;
				const pageSize = action?.pageInfo?.pageSize || 10;
				return (currentPage - 1) * pageSize + index + 1;
			}
		},
		{
			title: "用户编号",
			dataIndex: "id",
			width: 80,
			align: "left"
		},
		{
			title: "等级编号",
			dataIndex: "levelId",
			width: 80,
			align: "left",
			hideInSearch: true
		},
		{
			title: "用户名",
			dataIndex: "userName",
			width: 80,
			align: "left"
		},
		{
			title: "真实姓名",
			dataIndex: "trueName",
			width: 80,
			align: "left"
		},
		{
			title: "账户余额",
			dataIndex: "money",
			hideInSearch: true,
			width: 80,
			align: "left"
		},
		{
			title: "电子邮箱",
			dataIndex: "email",
			width: 180,
			align: "left"
		},
		{
			title: "国家区号",
			dataIndex: "mobileTitle",
			hideInSearch: true,
			width: 80,
			align: "left"
		},
		{
			title: "手机号码",
			dataIndex: "mobile",
			width: 120,
			align: "left"
		},
		{
			title: "等级类型",
			dataIndex: "levelType",
			valueType: "select", // 保证筛选框显示为下拉框
			width: 80,
			align: "left",
			hideInSearch: true,
			valueEnum: Object.fromEntries(levelTypeOptions), // 将 Map 转为 Object，提供给 ProTable 搜索框
			filters: Object.entries(levelTypeOptions).map(([value, label]) => ({
				text: label,
				value
			})),
			onFilter: (value, record) => record.userLevel?.levelType === value,
			render: (_, record) => {
				const levelType = record.userLevel?.levelType;
				return levelType ? levelTypeOptions.get(levelType) || levelType : "-";
			}
		},
		{
			title: "等级",
			dataIndex: "level",
			width: 80,
			align: "left",
			hideInSearch: true,
			render: (text, record) => record.userLevel?.level
		},
		{
			title: "当前用户邀请码",
			dataIndex: "refCode",
			width: 120,
			align: "left"
		},
		{
			title: "上级用户邀请码",
			dataIndex: "parentRefCode",
			width: 120,
			align: "left"
		},
		{
			title: "上级用户编号",
			dataIndex: "parentId",
			width: 120,
			align: "left",
			hideInSearch: true
		},
		{
			title: "状态",
			dataIndex: "status",
			valueType: "select",
			valueEnum: statusOptions,
			width: 80,
			align: "left"
		},
		{
			title: "创建时间",
			dataIndex: "createdAt",
			hideInSearch: true,
			valueType: "dateTime",
			width: 180,
			align: "left"
		},
		{
			title: "创建时间",
			dataIndex: "createdAt",
			valueType: "dateTimeRange",
			hideInTable: true,
			search: { transform: value => ({ beginCreatedAt: value[0], endCreatedAt: value[1] }) }
		},
		{
			title: "操作",
			valueType: "option",
			align: "center",
			fixed: "right",
			width: 150,
			render: (_, data) => (
				<HocAuth permission={["app:user:edit"]}>
					<Space>
						<LoadingButton
							key="edit"
							type="link"
							size="small"
							icon={<EditOutlined />}
							onClick={done => handleShowEditFormModal(data.id!, done)}
						>
							编辑
						</LoadingButton>
					</Space>
				</HocAuth>
			)
		}
	];

	useEffect(() => {
		const initData = async () => {
			const { data: statusData, msg: statusMsg, code: statusCode } = await getDictsApi("admin_sys_status");
			if (statusCode !== ResultEnum.SUCCESS) {
				message.error(statusMsg);
				return;
			}
			setStatusOptions(getDictOptions(statusData));
			const { data: levelTypeData, msg: levelTypeMsg, code: levelTypeCode } = await getDictsApi("app_user_level_type");
			if (levelTypeCode !== ResultEnum.SUCCESS) {
				message.error(levelTypeMsg);
				return;
			}
			setLevelTypeOptions(getDictOptions(levelTypeData));
		};
		initData();
	}, []);

	const handleShowAddFormModal = (done: () => void) => {
		formModalRef.current?.showAddFormModal();
		setTimeout(() => done(), 1000);
	};

	const handleShowEditFormModal = (id: number, done: () => void) => {
		formModalRef.current?.showEditFormModal(id);
		setTimeout(() => done(), 1000);
	};

	const handleFormModalConfirm = () => {
		actionRef.current?.reload(true);
	};

	const handleExport = (done: () => void) => {
		modal.confirm({
			title: "提示",
			icon: <ExclamationCircleOutlined />,
			content: "是否确认导出所选数据？",
			okText: "确认",
			cancelText: "取消",
			maskClosable: true,
			onCancel: () => {
				done();
			},
			onOk: async () => {
				try {
					saveExcelBlob("用户管理", await exportUserApi(tableFormRef.current?.getFieldsValue()));
				} catch (err) {
					message.error("下载失败，请检查网络");
				} finally {
					done();
				}
			}
		});
	};

	const toolBarRender = () => [
		<HocAuth permission={["app:user:add"]}>
			<LoadingButton type="primary" key="addTable" icon={<PlusCircleOutlined />} onClick={done => handleShowAddFormModal(done)}>
				新增
			</LoadingButton>
		</HocAuth>,
		<HocAuth permission={["app:user:export"]}>
			<LoadingButton type="primary" key="importTable" icon={<CloudDownloadOutlined />} onClick={done => handleExport(done)}>
				Excel导出
			</LoadingButton>
		</HocAuth>
	];

	return (
		<>
			<ProTable<UserModel>
				className="ant-pro-table-scroll"
				columns={columns}
				actionRef={actionRef}
				formRef={tableFormRef}
				bordered
				cardBordered
				defaultSize="small"
				scroll={{ x: "2000", y: "100%" }}
				request={async params => {
					const { data } = await getUserPageApi(params);
					return formatDataForProTable<UserModel>(data);
				}}
				columnsState={{
					persistenceKey: "use-pro-table-key",
					persistenceType: "localStorage"
				}}
				options={{
					reload: true,
					density: true,
					fullScreen: true
				}}
				rowKey="id"
				search={{ labelWidth: "auto", showHiddenNum: true }}
				pagination={pagination}
				dateFormatter="string"
				headerTitle="用户管理"
				toolBarRender={toolBarRender}
			/>
			<FormModal ref={formModalRef} onConfirm={handleFormModalConfirm} />
		</>
	);
};

export default User;
