import { addUserApi, getUserApi, updateUserApi, UserModel } from "@/api/app/user/user";
import { UserLevelModel } from "@/api/app/user/user-level";
import LoadingButton from "@/components/LoadingButton";
import { ResultEnum } from "@/enums/httpEnum";
import { message } from "@/hooks/useMessage";
import { Col, Form, Input, InputNumber, Modal, Row } from "antd";
import { forwardRef, useImperativeHandle, useRef, useState } from "react";
import UserLevelSelectModal, { UserLevelSelectModalRef } from "../../user-level/components/LevelSelectModal";

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
  const [model, setModel] = useState<UserModel>({});
  const userLevelSelectModalRef = useRef<UserLevelSelectModalRef>(null);

  useImperativeHandle(ref, () => ({
    showAddFormModal() {
      reset();
      setIsModalOpen(true);
    },
    async showEditFormModal(id: number) {
      const { data, msg, code } = await getUserApi(id);
      if (code !== ResultEnum.SUCCESS) {
        message.error(msg);
        return;
      }
      setModel(data);
      form.setFieldsValue(data);
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
            const { msg, code } = await addUserApi(values);
            if (code !== ResultEnum.SUCCESS) {
              message.error(msg);
              return;
            }
            message.success(msg);
          } else {
            const { msg, code } = await updateUserApi(model.id!, values);
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

  const handleShowUserLevelSelectModal = () => {
    userLevelSelectModalRef.current?.showUserLevelSelectModal();
  };

  const handleUserLevelConfirm = async (selectedRows: UserLevelModel) => {
    setModel({ ...model, levelId: selectedRows.id });
    form.setFieldsValue({ levelId: selectedRows.id });
  };

  return (
    <>
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
        {model.id! <= 0 && (
          <Form form={form} layout="vertical" initialValues={model}>
            <Form.Item name="mobileTitle" label="国家区号">
              <Input placeholder="请输入国家区号" />
            </Form.Item>
            <Form.Item name="mobiles" label="手机号码集合">
              <Input.TextArea placeholder="请输入手机号(多个手机号之间使用半角逗号分割)" />
            </Form.Item>
            <Form.Item name="emails" label="邮箱集合">
              <Input.TextArea placeholder="请输入邮箱(多个邮箱之间使用半角逗号分割)" />
            </Form.Item>
            <Form.Item name="refCode" label="邀请码">
              <Input placeholder="请输入邀请码" />
            </Form.Item>
          </Form>
        )}
        {model.id! > 0 && (
          <Form form={form} layout="vertical" initialValues={model}>
            <Row gutter={24}>
              <Col span={12}>
                <Form.Item name="id" label="用户编号">
                  <InputNumber style={{ width: "100%" }} min={0} disabled />
                </Form.Item>
              </Col>
              <Col span={12}>
                <Form.Item name="levelId" label="用户等级编号" rules={[{ required: true, message: "请输入用户等级编号" }]}>
                  <InputNumber style={{ width: "100%" }} min={0} onClick={handleShowUserLevelSelectModal} />
                </Form.Item>
              </Col>

              <Col span={24}>
                <Form.Item name="userName" label="用户名" rules={[{ required: true, message: "请输入用户名" }]}>
                  <Input placeholder="请输入用户名" />
                </Form.Item>
              </Col>

              <Col span={12}>
                <Form.Item name="trueName" label="真实姓名" rules={[{ required: true, message: "请输入真实姓名" }]}>
                  <Input placeholder="请输入真实姓名" />
                </Form.Item>
              </Col>
              <Col span={12}>
                <Form.Item name="mobileTitle" label="国家区号" rules={[{ required: true, message: "请输入国家区号" }]}>
                  <Input placeholder="请输入国家区号" />
                </Form.Item>
              </Col>
              <Col span={12}>
                <Form.Item name="mobile" label="手机号码" rules={[{ required: true, message: "请输入手机号码" }]}>
                  <Input placeholder="请输入手机号码" />
                </Form.Item>
              </Col>
              <Col span={12}>
                <Form.Item name="email" label="电子邮箱" rules={[{ required: true, message: "请输入电子邮箱" }]}>
                  <Input placeholder="请输入电子邮箱" />
                </Form.Item>
              </Col>
            </Row>
          </Form>
        )}
      </Modal>
      <UserLevelSelectModal ref={userLevelSelectModalRef} onConfirm={handleUserLevelConfirm} />
    </>
  );
});

export default FormModal;
