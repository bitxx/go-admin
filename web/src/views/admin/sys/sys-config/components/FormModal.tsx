import { addConfigApi, ConfigModel, getConfigApi, updateConfigApi } from "@/api/admin/sys/sys-config";
import { getDictOptions, getDictsApi } from "@/api/admin/sys/sys-dictdata";
import LoadingButton from "@/components/LoadingButton";
import { ResultEnum } from "@/enums/httpEnum";
import { message } from "@/hooks/useMessage";
import { Form, Input, Modal, Select } from "antd";
import { forwardRef, useEffect, useImperativeHandle, useState } from "react";

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
	const [model, setModel] = useState<ConfigModel>({});
	const [configTypeOptions, setConfigTypeOptions] = useState<Map<string, string>>(new Map());
	const [isFrontendOptions, setIsFrontendOptions] = useState<Map<string, string>>(new Map());

	useImperativeHandle(ref, () => ({
		showAddFormModal() {
			reset();
			setIsModalOpen(true);
		},
		async showEditFormModal(id: number) {
			const { data, msg, code } = await getConfigApi(id);
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
						const { msg, code } = await addConfigApi(values);
						if (code !== ResultEnum.SUCCESS) {
							message.error(msg);
							return;
						}
						message.success(msg);
					} else {
						const { msg, code } = await updateConfigApi(model.id!, values);
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
				<Form.Item name="configKey" label="键名" rules={[{ required: true, message: "请输入键名" }]}>
					<Input placeholder="请输入键名" />
				</Form.Item>
				<Form.Item name="configName" label="配置名称" rules={[{ required: true, message: "请输入配置名称" }]}>
					<Input placeholder="请输入配置名称" />
				</Form.Item>
				<Form.Item name="configType" label="配置类型" rules={[{ required: true, message: "请输入配置类型" }]}>
					<Select placeholder="请选择">
						{Array.from(configTypeOptions).map(([dictValue, dictLabel]) => (
							<Select.Option key={dictValue} value={dictValue}>
								{dictLabel}
							</Select.Option>
						))}
					</Select>
				</Form.Item>
				<Form.Item name="configValue" label="键值" rules={[{ required: true, message: "请输入键值" }]}>
					<Input placeholder="请输入键值" />
				</Form.Item>
				<Form.Item name="isFrontend" label="是否前台展示" rules={[{ required: true, message: "请输入是否前台展示" }]}>
					<Select placeholder="请选择">
						{Array.from(isFrontendOptions).map(([dictValue, dictLabel]) => (
							<Select.Option key={dictValue} value={dictValue}>
								{dictLabel}
							</Select.Option>
						))}
					</Select>
				</Form.Item>
				<Form.Item name="remark" label="备注信息">
					<Input.TextArea placeholder="请输入备注信息" />
				</Form.Item>
			</Form>
		</Modal>
	);
});

export default FormModal;
