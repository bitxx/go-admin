import { getDictOptions, getDictsApi } from "@/api/admin/sys/sys-dictdata";
import { addUserLevelApi, getUserLevelApi, updateUserLevelApi, UserLevelModel } from "@/api/app/user/user-level";
import LoadingButton from "@/components/LoadingButton";
import { ResultEnum } from "@/enums/httpEnum";
import { message } from "@/hooks/useMessage";
import { Form, Input, InputNumber, Modal, Select } from "antd";
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
  const [model, setModel] = useState<UserLevelModel>({});
  const [levelTypeOptions, setLevelTypeOptions] = useState<Map<string, string>>(new Map());

  useImperativeHandle(ref, () => ({
    showAddFormModal() {
      reset();
      setIsModalOpen(true);
    },
    async showEditFormModal(id: number) {
      const { data, msg, code } = await getUserLevelApi(id);
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
      const { data: levelTypeData, msg: levelTypeMsg, code: levelTypeCode } = await getDictsApi("app_user_level_type");
      if (levelTypeCode !== ResultEnum.SUCCESS) {
        message.error(levelTypeMsg);
        return;
      }
      setLevelTypeOptions(getDictOptions(levelTypeData));
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
            const { msg, code } = await addUserLevelApi(values);
            if (code !== ResultEnum.SUCCESS) {
              message.error(msg);
              return;
            }
            message.success(msg);
          } else {
            const { msg, code } = await updateUserLevelApi(model.id!, values);
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
        <Form.Item name="name" label="等级名称" rules={[{ required: true, message: "请输入等级名称" }]}>
          <Input placeholder="请输入等级名称" />
        </Form.Item>
        <Form.Item name="levelType" label="等级类型" rules={[{ required: true, message: "请输入等级类型" }]}>
          <Select placeholder="请选择">
            {Array.from(levelTypeOptions).map(([dictValue, dictLabel]) => (
              <Select.Option key={dictValue} value={dictValue}>
                {dictLabel}
              </Select.Option>
            ))}
          </Select>
        </Form.Item>
        <Form.Item name="level" label="等级" rules={[{ required: true, message: "请输入等级" }]}>
          <InputNumber style={{ width: "100%" }} min={0} />
        </Form.Item>
      </Form>
    </Modal>
  );
});

export default FormModal;
