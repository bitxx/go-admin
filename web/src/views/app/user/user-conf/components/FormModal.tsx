import { getDictOptions, getDictsApi } from "@/api/admin/sys/sys-dictdata";
import { getUserConfApi, updateUserConfApi, UserConfModel } from "@/api/app/user/user-conf";
import LoadingButton from "@/components/LoadingButton";
import { ResultEnum } from "@/enums/httpEnum";
import { message } from "@/hooks/useMessage";
import { Form, InputNumber, Modal, Select } from "antd";
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
  const [model, setModel] = useState<UserConfModel>({});
  const [canLoginOptions, setCanLoginOptions] = useState<Map<string, string>>(new Map());

  useImperativeHandle(ref, () => ({
    async showEditFormModal(id: number) {
      const { data, msg, code } = await getUserConfApi(id);
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
      const { data: canLoginData, msg: canLoginMsg, code: canLoginCode } = await getDictsApi("admin_sys_yes_no");
      if (canLoginCode !== ResultEnum.SUCCESS) {
        message.error(canLoginMsg);
        return;
      }
      setCanLoginOptions(getDictOptions(canLoginData));
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
          const { msg, code } = await updateUserConfApi(model.id!, values);
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
        <Form.Item name="userId" label="用户编号" rules={[{ required: true, message: "请输入用户编号" }]}>
          <InputNumber style={{ width: "100%" }} disabled min={0} />
        </Form.Item>
        <Form.Item name="canLogin" label="是否允许登录" rules={[{ required: true, message: "请输入是否允许登录" }]}>
          <Select placeholder="请选择">
            {Array.from(canLoginOptions).map(([dictValue, dictLabel]) => (
              <Select.Option key={dictValue} value={dictValue}>
                {dictLabel}
              </Select.Option>
            ))}
          </Select>
        </Form.Item>
      </Form>
    </Modal>
  );
});

export default FormModal;
