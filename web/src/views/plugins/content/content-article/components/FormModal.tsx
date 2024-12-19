import { getDictOptions, getDictsApi } from "@/api/admin/sys/sys-dictdata";
import {
  addContentArticleApi,
  ContentArticleModel,
  getContentArticleApi,
  updateContentArticleApi
} from "@/api/plugins/content/content-article";
import { ContentCategoryModel } from "@/api/plugins/content/content-category";
import LoadingButton from "@/components/LoadingButton";
import RichTextEditor from "@/components/RichTextEditor";
import { ResultEnum } from "@/enums/httpEnum";
import { message } from "@/hooks/useMessage";
import "@wangeditor/editor/dist/css/style.css";
import { Col, Form, Input, InputNumber, Modal, Row, Select } from "antd";
import { forwardRef, useEffect, useImperativeHandle, useRef, useState } from "react";
import CategorySelectModal, { CategorySelectModalRef } from "../../content-category/components/CategorySelectModal";

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
  const [model, setModel] = useState<ContentArticleModel>({});
  const [statusOptions, setStatusOptions] = useState<Map<string, string>>(new Map());
  const categorySelectModalRef = useRef<CategorySelectModalRef>(null);

  useImperativeHandle(ref, () => ({
    showAddFormModal() {
      reset();
      setIsModalOpen(true);
    },
    async showEditFormModal(id: number) {
      const { data, msg, code } = await getContentArticleApi(id);
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

  const handleShowCategorySelectModal = () => {
    categorySelectModalRef.current?.showCategorySelectModal();
  };

  const handleCategoryConfirm = async (selectedRows: ContentCategoryModel) => {
    setModel({ ...model, cateId: selectedRows.id });
    form.setFieldsValue({ cateId: selectedRows.id });
  };

  const handleConfirm = (done: () => void) => {
    form
      .validateFields()
      .then(async values => {
        try {
          if (model.id! <= 0) {
            const { msg, code } = await addContentArticleApi(values);
            if (code !== ResultEnum.SUCCESS) {
              message.error(msg);
              return;
            }
            message.success(msg);
          } else {
            const { msg, code } = await updateContentArticleApi(model.id!, values);
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
    <>
      <Modal
        title={model.id! > 0 ? "编辑" : "新增"}
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
          <Row gutter={24}>
            <Col span={12}>
              <Form.Item name="cateId" label="分类编号" rules={[{ required: true, message: "请输入分类编号" }]}>
                <InputNumber style={{ width: "100%" }} min={0} onClick={handleShowCategorySelectModal} />
              </Form.Item>
            </Col>
            <Col span={12}>
              <Form.Item name="status" label="状态" rules={[{ required: true, message: "请输入状态" }]}>
                <Select placeholder="请选择">
                  {Array.from(statusOptions).map(([dictValue, dictLabel]) => (
                    <Select.Option key={dictValue} value={dictValue}>
                      {dictLabel}
                    </Select.Option>
                  ))}
                </Select>
              </Form.Item>
            </Col>
            <Col span={24}>
              <Form.Item name="name" label="标题" rules={[{ required: true, message: "请输入标题" }]}>
                <Input placeholder="请输入标题" />
              </Form.Item>
            </Col>
            <Col span={24}>
              <Form.Item name="content" label="内容" rules={[{ required: true, message: "请输入内容" }]}>
                <RichTextEditor placeholder="请输入内容" />
              </Form.Item>
            </Col>
            <Col span={24}>
              <Form.Item name="remark" label="备注信息" rules={[{ required: true, message: "请输入备注信息" }]}>
                <Input.TextArea placeholder="请输入备注信息" />
              </Form.Item>
            </Col>
          </Row>
        </Form>
      </Modal>
      <CategorySelectModal ref={categorySelectModalRef} onConfirm={handleCategoryConfirm} />
    </>
  );
});

export default FormModal;
