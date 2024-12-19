import { getDictOptions, getDictsApi } from "@/api/admin/sys/sys-dictdata";
import { addPostApi, getPostApi, PostModel, updatePostApi } from "@/api/admin/sys/sys-post";
import LoadingButton from "@/components/LoadingButton";
import { ResultEnum } from "@/enums/httpEnum";
import { message } from "@/hooks/useMessage";
import { Form, Input, InputNumber, Modal, Radio } from "antd";
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
  const [model, setModel] = useState<PostModel>({});
  const [statusOptions, setStatusOptions] = useState<Map<string, string>>(new Map());

  useImperativeHandle(ref, () => ({
    showAddFormModal() {
      reset();
      setIsModalOpen(true);
    },
    async showEditFormModal(id: number) {
      const { data, msg, code } = await getPostApi(id);
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
      const { data: statusData, msg: statusMsg, code: statusCode } = await getDictsApi("admin_sys_status");
      if (statusCode !== ResultEnum.SUCCESS) {
        message.error(statusMsg);
        return;
      }
      setStatusOptions(getDictOptions(statusData));
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
            const { msg, code } = await addPostApi(values);
            if (code !== ResultEnum.SUCCESS) {
              message.error(msg);
              return;
            }
            message.success(msg);
          } else {
            const { msg, code } = await updatePostApi(model.id!, values);
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
        <Form.Item name="postCode" label="岗位编码" rules={[{ required: true, message: "请输入岗位编码" }]}>
          <Input placeholder="请输入岗位编码" />
        </Form.Item>
        <Form.Item name="postName" label="岗位名称" rules={[{ required: true, message: "请输入岗位名称" }]}>
          <Input placeholder="请输入岗位名称" />
        </Form.Item>
        <Form.Item name="remark" label="备注" rules={[{ required: true, message: "请输入备注" }]}>
          <Input.TextArea placeholder="请输入备注" />
        </Form.Item>
        <Form.Item name="sort" label="岗位排序" rules={[{ required: true, message: "请输入岗位排序" }]}>
          <InputNumber placeholder="请输入岗位排序" style={{ width: "100%" }} min={0} />
        </Form.Item>
        <Form.Item name="status" label="状态" rules={[{ required: true, message: "请输入状态" }]}>
          <Radio.Group>
            {Array.from(statusOptions).map(([dictValue, dictLabel]) => (
              <Radio key={dictValue} value={dictValue}>
                {dictLabel}
              </Radio>
            ))}
          </Radio.Group>
        </Form.Item>
      </Form>
    </Modal>
  );
});

export default FormModal;
