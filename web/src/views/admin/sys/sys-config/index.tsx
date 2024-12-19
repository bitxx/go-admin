import { ConfigModel, delConfigApi, exportConfigApi, getConfigPageApi } from "@/api/admin/sys/sys-config";
import { getDictOptions, getDictsApi } from "@/api/admin/sys/sys-dictdata";
import HocAuth from "@/components/HocAuth";
import LoadingButton from "@/components/LoadingButton";
import { pagination } from "@/config/proTable";
import { ResultEnum } from "@/enums/httpEnum";
import { message, modal } from "@/hooks/useMessage";
import { formatDataForProTable, saveExcelBlob } from "@/utils";
import {
	CloudDownloadOutlined,
	DeleteOutlined,
	EditOutlined,
	ExclamationCircleOutlined,
	PlusCircleOutlined
} from "@ant-design/icons";
import type { ActionType, ProColumns, ProFormInstance } from "@ant-design/pro-components";
import { ProTable } from "@ant-design/pro-components";
import { Space } from "antd";
import React, { useEffect, useRef, useState } from "react";
import FormModal, { FormModalRef } from "./components/FormModal";

const Config: React.FC = () => {
	const actionRef = React.useRef<ActionType>();
	const tableFormRef = React.useRef<ProFormInstance>();
	const formModalRef = useRef<FormModalRef>(null);
	const [configTypeOptions, setConfigTypeOptions] = useState<Map<string, string>>(new Map());
	const [isFrontendOptions, setIsFrontendOptions] = useState<Map<string, string>>(new Map());

	// 定义列
	const columns: ProColumns<ConfigModel>[] = [
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
			title: "配置编号",
			dataIndex: "id",
			hideInSearch: true,
			width: 80,
			align: "left"
		},
		{
			title: "配置名称",
			dataIndex: "configName",
			width: 260,
			align: "left"
		},
		{
			title: "键名",
			dataIndex: "configKey",
			width: 350,
			align: "left"
		},
		{
			title: "键值",
			dataIndex: "configValue",
			hideInSearch: true,
			width: 300,
			align: "left"
		},
		{
			title: "是否前台展示",
			dataIndex: "isFrontend",
			hideInSearch: true,
			valueType: "select",
			valueEnum: isFrontendOptions,
			width: 120,
			align: "left"
		},
		{
			title: "配置类型",
			dataIndex: "configType",
			valueType: "select",
			valueEnum: configTypeOptions,
			width: 120,
			align: "left"
		},
		{
			title: "备注信息",
			dataIndex: "remark",
			hideInSearch: true,
			width: 500,
			align: "left"
		},
		{
			title: "更新时间",
			dataIndex: "updatedAt",
			hideInSearch: true,
			valueType: "dateTime",
			width: 180,
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
				<Space>
					<HocAuth permission={["admin:sys-config:edit"]}>
						<LoadingButton
							key="edit"
							type="link"
							size="small"
							icon={<EditOutlined />}
							onClick={done => handleShowEditFormModal(data.id!, done)}
						>
							编辑
						</LoadingButton>
					</HocAuth>
					<HocAuth permission={["admin:sys-config:del"]}>
						<LoadingButton
							key="delete"
							type="link"
							size="small"
							danger
							icon={<DeleteOutlined />}
							onClick={done => handleDelete(data.id!, done)}
						>
							删除
						</LoadingButton>
					</HocAuth>
				</Space>
			)
		}
	];
	useEffect(() => {
		const initData = async () => {
			const { data: configTypeData, msg: configTypeMsg, code: configTypeCode } = await getDictsApi("admin_sys_config_type");
			if (configTypeCode !== ResultEnum.SUCCESS) {
				message.error(configTypeMsg);
				return;
			}
			setConfigTypeOptions(getDictOptions(configTypeData));
			const {
				data: isFrontendData,
				msg: isFrontendMsg,
				code: isFrontendCode
			} = await getDictsApi("admin_sys_config_is_frontend");
			if (isFrontendCode !== ResultEnum.SUCCESS) {
				message.error(isFrontendMsg);
				return;
			}
			setIsFrontendOptions(getDictOptions(isFrontendData));
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
					saveExcelBlob("Config", await exportConfigApi(tableFormRef.current?.getFieldsValue()));
				} catch (err) {
					message.error("下载失败，请检查网络");
				} finally {
					done();
				}
			}
		});
	};

	const handleDelete = (id: number, done: () => void) => {
		modal.confirm({
			title: "提示",
			icon: <ExclamationCircleOutlined />,
			content: "是否确认删除编号为 " + id + " 的数据项?",
			okText: "确认",
			cancelText: "取消",
			maskClosable: true,
			onCancel: () => {
				done();
			},
			onOk: async () => {
				try {
					const { code, msg } = await delConfigApi([id!]);
					if (code !== ResultEnum.SUCCESS) {
						message.error(msg);
						return;
					}
					actionRef.current?.reload(true);
					message.success(msg);
				} finally {
					done();
				}
			}
		});
	};

	const toolBarRender = () => [
		<HocAuth permission={["admin:sys-config:add"]}>
			<LoadingButton type="primary" key="addTable" icon={<PlusCircleOutlined />} onClick={done => handleShowAddFormModal(done)}>
				新增
			</LoadingButton>
		</HocAuth>,
		<HocAuth permission={["admin:sys-config:export"]}>
			<LoadingButton type="primary" key="importTable" icon={<CloudDownloadOutlined />} onClick={done => handleExport(done)}>
				Excel导出
			</LoadingButton>
		</HocAuth>
	];

	return (
		<>
			<ProTable<ConfigModel>
				className="ant-pro-table-scroll"
				columns={columns}
				actionRef={actionRef}
				formRef={tableFormRef}
				bordered
				cardBordered
				defaultSize="small"
				scroll={{ x: "2000", y: "100%" }}
				request={async params => {
					const { data } = await getConfigPageApi(params);
					return formatDataForProTable<ConfigModel>(data);
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
				headerTitle="参数管理"
				toolBarRender={toolBarRender}
			/>
			<FormModal ref={formModalRef} onConfirm={handleFormModalConfirm} />
		</>
	);
};

export default Config;
