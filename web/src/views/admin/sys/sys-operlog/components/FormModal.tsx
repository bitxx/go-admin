import { getOperLogApi, OperLogModel } from "@/api/admin/sys/sys-operlog";
import LoadingButton from "@/components/LoadingButton";
import { ResultEnum } from "@/enums/httpEnum";
import { message } from "@/hooks/useMessage";
import { Col, Form, Input, Modal, Row } from "antd";
import { forwardRef, useImperativeHandle, useState } from "react";

export interface FormModalRef {
  showEditFormModal: (id: number) => void;
}

interface ModalProps {}

const FormModal = forwardRef<FormModalRef, ModalProps>(({}, ref) => {
  const [form] = Form.useForm();
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [model, setModel] = useState<OperLogModel>({});

  useImperativeHandle(ref, () => ({
    async showEditFormModal(id: number) {
      const { data, msg, code } = await getOperLogApi(id);
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

  return (
    <Modal
      title={"详情"}
      getContainer={false}
      width={800}
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
          key="confirm"
          type="primary"
          onClick={done => {
            reset();
            setIsModalOpen(false);
            done();
          }}
        >
          确定
        </LoadingButton>
      ]}
    >
      <Form form={form} layout="vertical" initialValues={model}>
        <Row gutter={24}>
          <Col span={12}>
            <Form.Item name="id" label="日志编号">
              <Input disabled />
            </Form.Item>
          </Col>
          <Col span={12}>
            <Form.Item name="operUrl" label="请求地址">
              <Input disabled />
            </Form.Item>
          </Col>
          <Col span={12}>
            <Form.Item name="operIp" label="请求ip">
              <Input disabled />
            </Form.Item>
          </Col>
          <Col span={12}>
            <Form.Item name="requestMethod" label="请求方法">
              <Input disabled />
            </Form.Item>
          </Col>
          <Col span={12}>
            <Form.Item name="latencyTime" label="耗时">
              <Input disabled />
            </Form.Item>
          </Col>
          <Col span={12}>
            <Form.Item name="operLocation" label="访问位置">
              <Input disabled />
            </Form.Item>
          </Col>
          <Col span={24}>
            <Form.Item name="jsonResult" label="返回参数">
              <Input.TextArea autoSize={{ minRows: 2 }} disabled />
            </Form.Item>
          </Col>
          <Col span={12}>
            <Form.Item name="status" label="返回码">
              <Input disabled />
            </Form.Item>
          </Col>
          <Col span={12}>
            <Form.Item name="operTime" label="操作时间">
              <Input disabled />
            </Form.Item>
          </Col>
        </Row>
      </Form>
    </Modal>
  );
});

export default FormModal;
