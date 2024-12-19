import { delDeptApi, DeptModel, getDeptTreeApi } from "@/api/admin/sys/sys-dept";
import HocAuth from "@/components/HocAuth";
import LoadingButton from "@/components/LoadingButton";
import { pagination } from "@/config/proTable";
import { ResultEnum } from "@/enums/httpEnum";
import { message, modal } from "@/hooks/useMessage";
import { formatDataListForProTable } from "@/utils";
import { DeleteOutlined, EditOutlined, ExclamationCircleOutlined, PlusCircleOutlined } from "@ant-design/icons";
import type { ActionType, Key, ProColumns, ProFormInstance } from "@ant-design/pro-components";
import { ProTable } from "@ant-design/pro-components";
import { Space } from "antd";
import React, { useRef, useState } from "react";
import FormModal, { FormModalRef } from "./components/FormModal";

const Dept: React.FC = () => {
	const actionRef = React.useRef<ActionType>();
	const tableFormRef = React.useRef<ProFormInstance>();
	const formModalRef = useRef<FormModalRef>(null);
	const [expandedRowKeys, setExpandedRowKeys] = useState<string[]>([]);

	// 定义列
	const columns: ProColumns<DeptModel>[] = [
		{
			title: "部门名称",
			dataIndex: "deptName",
			width: 150,
			align: "left"
		},
		{
			title: "部门编号",
			dataIndex: "id",
			width: 80,
			hideInSearch: true,
			align: "left"
		},
		{
			title: "负责人",
			dataIndex: "leader",
			width: 80,
			align: "left"
		},
		{
			title: "电话",
			dataIndex: "phone",
			hideInSearch: true,
			width: 120,
			align: "left"
		},
		{
			title: "邮箱",
			dataIndex: "email",
			hideInSearch: true,
			width: 120,
			align: "left"
		},
		{
			title: "排序",
			dataIndex: "sort",
			hideInSearch: true,
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
			title: "操作",
			valueType: "option",
			align: "center",
			fixed: "right",
			width: 150,
			render: (_, data) =>
				data.id !== undefined &&
				data.id > 0 && (
					<Space>
						<HocAuth permission={["admin:sys-dept:edit"]}>
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
						<HocAuth permission={["admin:sys-dept:del"]}>
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

	const getAllRowKeys = (data: any[], keys: string[] = []) => {
		data.forEach(item => {
			keys.push(item.id);
			if (item.children) {
				getAllRowKeys(item.children, keys);
			}
		});
		return keys;
	};

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
					const { code, msg } = await delDeptApi([id!]);
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
		<HocAuth permission={["admin:sys-dept:add"]}>
			<LoadingButton type="primary" key="addTable" icon={<PlusCircleOutlined />} onClick={done => handleShowAddFormModal(done)}>
				新增
			</LoadingButton>
		</HocAuth>
	];

	return (
		<>
			<ProTable<DeptModel>
				className="ant-pro-table-scroll"
				columns={columns}
				actionRef={actionRef}
				formRef={tableFormRef}
				bordered
				cardBordered
				defaultSize="small"
				scroll={{ x: "2000", y: "100%" }}
				expandable={{
					expandedRowKeys: expandedRowKeys,
					onExpandedRowsChange: (expandedKeys: readonly Key[]) => {
						setExpandedRowKeys(expandedKeys as string[]);
					}
				}}
				request={async params => {
					const { data } = await getDeptTreeApi(params);
					if (data) {
						const allRowKeys = getAllRowKeys(data);
						setExpandedRowKeys(allRowKeys);
					}
					return formatDataListForProTable(data);
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
				headerTitle="部门管理"
				toolBarRender={toolBarRender}
			/>
			<FormModal ref={formModalRef} onConfirm={handleFormModalConfirm} />
		</>
	);
};

export default Dept;
