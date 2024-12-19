import { addDictDataApi, DictDataModel, getDictDataApi, updateDictDataApi } from "@/api/admin/sys/sys-dictdata";
import LoadingButton from "@/components/LoadingButton";
import { ResultEnum } from "@/enums/httpEnum";
import { message } from "@/hooks/useMessage";
import { Form, Input, InputNumber, Modal } from "antd";
import { forwardRef, useImperativeHandle, useState } from "react";

export interface FormModalRef {
  showAddFormModal: (dictType: string) => void;
  showEditFormModal: (id: number) => void;
}

interface ModalProps {
  onConfirm: () => void;
}

const FormModal = forwardRef<FormModalRef, ModalProps>(({ onConfirm }, ref) => {
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [model, setModel] = useState<DictDataModel>({});
  const [form] = Form.useForm();

  useImperativeHandle(ref, () => ({
    showAddFormModal(dictType: string) {
      reset(dictType);
      setIsModalOpen(true);
    },
    async showEditFormModal(id: number) {
      const { data, msg, code } = await getDictDataApi(id);
      if (code !== ResultEnum.SUCCESS) {
        message.error(msg);
        return;
      }
      setModel(data);
      form.setFieldsValue(data);
      setIsModalOpen(true);
    }
  }));

  const reset = (dictType: string) => {
    if (model.id! > 0) {
      setModel({});
    } else {
      setModel({ id: 0, dictSort: 0, dictType });
    }
    setTimeout(() => form.resetFields(), 100);
  };

  const handleConfirm = (done: () => void) => {
    form
      .validateFields()
      .then(async values => {
        try {
          if (model.id! <= 0) {
            const { msg, code } = await addDictDataApi(values);
            if (code !== ResultEnum.SUCCESS) {
              message.error(msg);
              return;
            }
            message.success(msg);
          } else {
            const { msg, code } = await updateDictDataApi(model.id!, values);
            if (code !== ResultEnum.SUCCESS) {
              message.error(msg);
              return;
            }
            message.success(msg);
          }
          reset("");
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
        reset("");
        setIsModalOpen(false);
      }}
      destroyOnClose
      footer={[
        <LoadingButton
          key="cancel"
          onClick={done => {
            reset("");
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
        <Form.Item name="dictType" label="字典类型">
          <Input readOnly disabled />
        </Form.Item>
        <Form.Item name="dictLabel" label="数据标签" rules={[{ required: true, message: "请输入数据标签" }]}>
          <Input placeholder="请输入数据标签" />
        </Form.Item>
        <Form.Item name="dictValue" label="数据键值" rules={[{ required: true, message: "请输入数据键值" }]}>
          <Input placeholder="请输入数据键值" />
        </Form.Item>
        <Form.Item name="dictSort" label="显示排序" rules={[{ required: true, message: "数据顺序不能为空" }]}>
          <InputNumber style={{ width: "100%" }} min={0} />
        </Form.Item>
        <Form.Item name="remark" label="备注">
          <Input.TextArea placeholder="请输入内容" />
        </Form.Item>
      </Form>
    </Modal>
  );
});

export default FormModal;
