import { changeUserPwdApi } from "@/api/admin/sys/sys-user";
import LoadingButton from "@/components/LoadingButton";
import { ResultEnum } from "@/enums/httpEnum";
import { message } from "@/hooks/useMessage";
import { Col, Form, Input, Modal, Row } from "antd";
import { forwardRef, useImperativeHandle, useState } from "react";

export interface ChangePwdFormModalRef {
  showChangePwdFormModal: (id: number) => void;
}

interface ModalProps {
  onConfirm: () => void;
}

const ChangePwdFormModal = forwardRef<ChangePwdFormModalRef, ModalProps>(({ onConfirm }, ref) => {
  const [form] = Form.useForm();
  const [isModalOpen, setIsModalOpen] = useState(false);

  const [userId, setUserId] = useState<number>();

  useImperativeHandle(ref, () => ({
    async showChangePwdFormModal(id: number) {
      setUserId(id);
      setIsModalOpen(true);
    }
  }));

  const reset = () => {
    setUserId(undefined);
    setTimeout(() => form.resetFields(), 100);
  };

  const handleConfirm = (done: () => void) => {
    form
      .validateFields()
      .then(async values => {
        try {
          const { msg, code } = await changeUserPwdApi(userId!, values.password);
          if (code !== ResultEnum.SUCCESS) {
            message.error(msg);
            return;
          }
          message.success(msg);
        } finally {
          done();
        }
        reset();
        setIsModalOpen(false);
        onConfirm();
      })
      .catch(error => {
        console.error("validate error：", error);
        message.error("表单校验失败");
        done();
      });
  };

  return (
    <Modal
      title={"修改密码"}
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
      <Form form={form} layout="vertical">
        <Row gutter={24}>
          <Col span={24}>
            <Form.Item name="password" label="新密码" rules={[{ required: true, message: "请输入新密码" }]}>
              <Input placeholder="请输入新密码" />
            </Form.Item>
          </Col>
        </Row>
      </Form>
    </Modal>
  );
});

export default ChangePwdFormModal;
