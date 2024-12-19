import { addDeptApi, DeptModel, getDeptApi, getDeptTreeApi, updateDeptApi } from "@/api/admin/sys/sys-dept";
import LoadingButton from "@/components/LoadingButton";
import { ResultEnum } from "@/enums/httpEnum";
import { message } from "@/hooks/useMessage";
import { Col, Form, Input, InputNumber, Modal, Row, TreeSelect } from "antd";
import { forwardRef, useImperativeHandle, useState } from "react";

export interface FormModalRef {
	showAddFormModal: () => void;
	showEditFormModal: (id: number) => void;
}

interface ModalProps {
	onConfirm: () => void;
}

const FormModal = forwardRef<FormModalRef, ModalProps>(({ onConfirm }, ref) => {
	const [form] = Form.useForm();
	const [isModalOpen, setIsModalOpen] = useState(false);
	const [model, setModel] = useState<DeptModel>({});
	const [parentDept, setParentDept] = useState<DeptModel>();
	const [deptList, setDeptList] = useState<DeptModel[]>();

	useImperativeHandle(ref, () => ({
		async showAddFormModal() {
			reset();

			const { data: deptListData, msg: deptListMsg, code: deptListCode } = await getDeptTreeApi({});
			if (deptListCode !== ResultEnum.SUCCESS) {
				message.error(deptListMsg);
				return;
			}
			setDeptList(deptListData);

			setIsModalOpen(true);
		},
		async showEditFormModal(id: number) {
			const { data, msg, code } = await getDeptApi(id);
			if (code !== ResultEnum.SUCCESS) {
				message.error(msg);
				return;
			}
			const { data: deptListData, msg: deptListMsg, code: deptListCode } = await getDeptTreeApi({});
			if (deptListCode !== ResultEnum.SUCCESS) {
				message.error(deptListMsg);
				return;
			}
			// let totalDeptListData = [{ id: 0, deptName: "主类目", parentId: 0, children: deptListData }];

			const parentNode = getParentDeptModel(data, deptListData);
			setParentDept(parentNode);
			setModel(data);
			form.setFieldsValue(data);
			setDeptList(deptListData);
			setIsModalOpen(true);
		}
	}));

	const reset = () => {
		if (model.id! > 0) {
			setModel({});
		} else {
			setModel({ id: 0 });
		}
		setTimeout(() => form.resetFields(), 100);
	};

	const handleConfirm = (done: () => void) => {
		form
			.validateFields()
			.then(async values => {
				try {
					if (model.id! <= 0) {
						const { msg, code } = await addDeptApi(values);
						if (code !== ResultEnum.SUCCESS) {
							message.error(msg);
							return;
						}
						message.success(msg);
					} else {
						const { msg, code } = await updateDeptApi(model.id!, values);
						if (code !== ResultEnum.SUCCESS) {
							message.error(msg);
							return;
						}
						message.success(msg);
					}
					reset();
					setIsModalOpen(false);
					onConfirm();
				} finally {
					done();
				}
			})
			.catch(error => {
				console.error("validate error：", error);
				message.error("表单校验失败");
				done();
			});
	};

	return (
		<Modal
			title={model.id! > 0 ? "编辑" : "新增"}
			getContainer={false}
			width={500}
			open={isModalOpen}
			maskClosable={false}
			keyboard={false}
			onCancel={() => {
				reset();
				setIsModalOpen(false);
			}}
			destroyOnClose
			footer={[
				<LoadingButton
					key="cancel"
					onClick={done => {
						reset();
						setIsModalOpen(false);
						done();
					}}
				>
					取消
				</LoadingButton>,
				<LoadingButton key="confirm" type="primary" onClick={done => handleConfirm(done)}>
					确定
				</LoadingButton>
			]}
		>
			<Form form={form} layout="vertical" initialValues={model}>
				<Form.Item name="parentId" label="上级部门" rules={[{ required: true, message: "选择上级部门" }]}>
					<TreeSelect
						showSearch
						style={{ width: "100%" }}
						value={parentDept}
						treeNodeFilterProp="deptName"
						fieldNames={{ label: "deptName", value: "id", children: "children" }}
						dropdownStyle={{ maxHeight: 400, overflow: "auto" }}
						placeholder="选择上级部门"
						allowClear
						treeDefaultExpandAll
						onChange={(newDept: DeptModel) => {
							setParentDept(newDept);
						}}
						treeData={deptList}
					/>
				</Form.Item>
				<Row gutter={24}>
					<Col span={12}>
						<Form.Item name="deptName" label="部门名称" rules={[{ required: true, message: "请输入部门名称" }]}>
							<Input placeholder="请输入部门名称" />
						</Form.Item>
					</Col>
					<Col span={12}>
						<Form.Item name="leader" label="负责人" rules={[{ required: true, message: "请输入负责人" }]}>
							<Input placeholder="请输入负责人" />
						</Form.Item>
					</Col>
					<Col span={12}>
						<Form.Item name="email" label="邮箱">
							<Input placeholder="请输入邮箱" />
						</Form.Item>
					</Col>
					<Col span={12}>
						<Form.Item name="phone" label="电话">
							<Input placeholder="请输入电话" />
						</Form.Item>
					</Col>
					<Col span={12}>
						<Form.Item name="sort" label="排序">
							<InputNumber placeholder="请输入排序" style={{ width: "100%" }} min={0} />
						</Form.Item>
					</Col>
				</Row>
			</Form>
		</Modal>
	);
});

export default FormModal;

// 获取当前选中项的父级路径
const getParentDeptModel = (currentNode: DeptModel, treeList: DeptModel[]) => {
	if (!currentNode) {
		return;
	}
	for (let node of treeList) {
		if (currentNode.parentId === node.id || currentNode.id == node.id) {
			return node; // 如果找到匹配项，返回当前路径
		}
		if (node.children) {
			getParentDeptModel(currentNode, node.children);
		}
	}
};
