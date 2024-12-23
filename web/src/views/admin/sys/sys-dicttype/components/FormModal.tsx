import { addDictTypeApi, DictTypeModel, getDictTypeApi, updateDictTypeApi } from "@/api/admin/sys/sys-dicttype";
import LoadingButton from "@/components/LoadingButton";
import { ResultEnum } from "@/enums/httpEnum";
import { message } from "@/hooks/useMessage";
import { Form, Input, Modal } from "antd";
import { forwardRef, useImperativeHandle, useState } from "react";

export interface FormModalRef {
  showAddFormModal: () => void;
  showEditFormModal: (id: number) => void;
}

interface ModalProps {
  onConfirm: () => void;
}

const FormModal = forwardRef<FormModalRef, ModalProps>(({ onConfirm }, ref) => {
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [model, setModel] = useState<DictTypeModel>({});
  const [form] = Form.useForm();

  useImperativeHandle(ref, () => ({
    showAddFormModal() {
      reset();
      setIsModalOpen(true);
    },
    async showEditFormModal(id: number) {
      const { data, msg, code } = await getDictTypeApi(id);
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
            const { msg, code } = await addDictTypeApi(values);
            if (code !== ResultEnum.SUCCESS) {
              message.error(msg);
              return;
            }
            message.success(msg);
          } else {
            const { msg, code } = await updateDictTypeApi(model.id!, values);
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
        <Form.Item name="dictName" label="字典名称" rules={[{ required: true, message: "请输入字典名称" }]}>
          <Input placeholder="请输入字典名称" />
        </Form.Item>
        <Form.Item name="dictType" label="字典类型" rules={[{ required: true, message: "请输入字典类型" }]}>
          <Input placeholder="请输入字典类型" />
        </Form.Item>
        <Form.Item name="remark" label="备注">
          <Input.TextArea placeholder="请输入内容" />
        </Form.Item>
      </Form>
    </Modal>
  );
});

export default FormModal;
