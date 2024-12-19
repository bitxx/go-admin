import { ApiModel, getApiApi, updateApiApi } from "@/api/admin/sys/sys-api";
import { getDictOptions, getDictsApi } from "@/api/admin/sys/sys-dictdata";
import LoadingButton from "@/components/LoadingButton";
import { ResultEnum } from "@/enums/httpEnum";
import { message } from "@/hooks/useMessage";
import { Form, Input, Modal, Select } from "antd";
import { forwardRef, useEffect, useImperativeHandle, useState } from "react";

export interface FormModalRef {
	showEditFormModal: (id: number) => void;
}

interface ModalProps {
	onConfirm: () => void;
}

const FormModal = forwardRef<FormModalRef, ModalProps>(({ onConfirm }, ref) => {
	const [form] = Form.useForm();
	const [isModalOpen, setIsModalOpen] = useState(false);
	const [model, setModel] = useState<ApiModel>({});
	const [methodOptions, setMethodOptions] = useState<Map<string, string>>(new Map());
	const [apiTypeOptions, setApiTypeOptions] = useState<Map<string, string>>(new Map());

	useImperativeHandle(ref, () => ({
		async showEditFormModal(id: number) {
			const { data, msg, code } = await getApiApi(id);
			if (code !== ResultEnum.SUCCESS) {
				message.error(msg);
				return;
			}
			setModel(data);
			form.setFieldsValue(data);
			setIsModalOpen(true);
		}
	}));
	useEffect(() => {
		const initData = async () => {
			const { data: methodData, msg: methodMsg, code: methodCode } = await getDictsApi("admin_sys_api_method");
			if (methodCode !== ResultEnum.SUCCESS) {
				message.error(methodMsg);
				return;
			}
			setMethodOptions(getDictOptions(methodData));
			const { data: apiTypeData, msg: apiTypeMsg, code: apiTypeCode } = await getDictsApi("admin_sys_config_type");
			if (apiTypeCode !== ResultEnum.SUCCESS) {
				message.error(apiTypeMsg);
				return;
			}
			setApiTypeOptions(getDictOptions(apiTypeData));
		};
		initData();
	}, []);

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
					const { msg, code } = await updateApiApi(model.id!, values);
					if (code !== ResultEnum.SUCCESS) {
						message.error(msg);
						return;
					}
					message.success(msg);

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
			title={"编辑"}
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
				<Form.Item name="path" label="接口地址">
					<Input placeholder="请输入接口地址" disabled />
				</Form.Item>
				<Form.Item name="method" label="接口请求方法">
					<Select placeholder="请选择" disabled>
						{Array.from(methodOptions).map(([dictValue, dictLabel]) => (
							<Select.Option key={dictValue} value={dictValue}>
								{dictLabel}
							</Select.Option>
						))}
					</Select>
				</Form.Item>
				<Form.Item name="apiType" label="接口类型" rules={[{ required: true, message: "请输入接口类型" }]}>
					<Select placeholder="请选择">
						{Array.from(apiTypeOptions).map(([dictValue, dictLabel]) => (
							<Select.Option key={dictValue} value={dictValue}>
								{dictLabel}
							</Select.Option>
						))}
					</Select>
				</Form.Item>
				<Form.Item name="description" label="功能描述" rules={[{ required: true, message: "请输入功能描述" }]}>
					<Input placeholder="请输入功能描述" />
				</Form.Item>
				<Form.Item name="remark" label="备注">
					<Input.TextArea placeholder="请输入备注" />
				</Form.Item>
			</Form>
		</Modal>
	);
});

export default FormModal;
