import { getDictOptions, getDictsApi } from "@/api/admin/sys/sys-dictdata";
import {
  addFilemgrAppApi,
  exportUploadFileAppApi,
  FilemgrAppModel,
  getFilemgrAppApi,
  updateFilemgrAppApi
} from "@/api/plugins/filemgr/filemgr-app";
import LoadingButton from "@/components/LoadingButton";
import { ResultEnum } from "@/enums/httpEnum";
import { message } from "@/hooks/useMessage";
import { InboxOutlined } from "@ant-design/icons";
import { Col, Form, Input, Modal, Row, Select, Upload, UploadProps } from "antd";
import { forwardRef, useEffect, useImperativeHandle, useState } from "react";
const { Dragger } = Upload;

const appDownloadTypeLocal = "1"; //local
const appDownloadTypeUrl = "2"; //outlet
const appDownloadTypeOss = "3"; //OSS

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
  const [model, setModel] = useState<FilemgrAppModel>({});
  const [platformOptions, setPlatformOptions] = useState<Map<string, string>>(new Map());
  const [appTypeOptions, setAppTypeOptions] = useState<Map<string, string>>(new Map());
  const [downloadTypeOptions, setDownloadTypeOptions] = useState<Map<string, string>>(new Map());
  const [statusOptions, setStatusOptions] = useState<Map<string, string>>(new Map());
  const [downloadType, setDownloadType] = useState<string>("");

  const props: UploadProps = {
    name: "file",
    multiple: false,
    accept: ".apk, .ipa",
    maxCount: 1,
    customRequest: exportUploadFileAppApi,
    // action: import.meta.env.VITE_API_URL + "/admin-api/v1/plugins/filemgr/filemgr-app/upload",
    // headers: { Authorization: "Bearer " + store.getState().user.token },
    onChange(info) {
      const { status } = info.file;
      // if (status !== "uploading") {
      //   console.log(info.file, info.fileList);
      // }
      if (status === "done") {
        form.setFieldValue("localAddress", info.file.response);
        message.success(`${info.file.name} 上传成功`);
      } else if (status === "error") {
        message.error(`${info.file.name} 上传失败`);
      }
    },
    onDrop(e) {
      //console.log("Dropped files", e.dataTransfer.files);
    }
  };

  useImperativeHandle(ref, () => ({
    showAddFormModal() {
      reset();
      setIsModalOpen(true);
    },
    async showEditFormModal(id: number) {
      const { data, msg, code } = await getFilemgrAppApi(id);
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
      const { data: platformData, msg: platformMsg, code: platformCode } = await getDictsApi("plugin_filemgr_app_platform");
      if (platformCode !== ResultEnum.SUCCESS) {
        message.error(platformMsg);
        return;
      }
      setPlatformOptions(getDictOptions(platformData));
      const { data: appTypeData, msg: appTypeMsg, code: appTypeCode } = await getDictsApi("plugin_filemgr_app_type");
      if (appTypeCode !== ResultEnum.SUCCESS) {
        message.error(appTypeMsg);
        return;
      }
      setAppTypeOptions(getDictOptions(appTypeData));
      const {
        data: downloadTypeData,
        msg: downloadTypeMsg,
        code: downloadTypeCode
      } = await getDictsApi("plugin_filemgr_app_download_type");
      if (downloadTypeCode !== ResultEnum.SUCCESS) {
        message.error(downloadTypeMsg);
        return;
      }
      setDownloadTypeOptions(getDictOptions(downloadTypeData));
      const { data: statusData, msg: statusMsg, code: statusCode } = await getDictsApi("plugin_filemgr_publish_status");
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
    setDownloadType("");
    setTimeout(() => form.resetFields(), 100);
  };

  const handleConfirm = (done: () => void) => {
    form
      .validateFields()
      .then(async values => {
        try {
          if (model.id! <= 0) {
            const { msg, code } = await addFilemgrAppApi(values);
            if (code !== ResultEnum.SUCCESS) {
              message.error(msg);
              return;
            }
            message.success(msg);
          } else {
            const { msg, code } = await updateFilemgrAppApi(model.id!, values);
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

  // 监听表单值变化
  const onDownloadTypeChange = (changedValues: Partial<FilemgrAppModel>, allValues: FilemgrAppModel) => {
    // 获取 field1 的最新值
    setDownloadType(allValues.downloadType!);
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
      <Form form={form} layout="vertical" onValuesChange={onDownloadTypeChange} initialValues={model}>
        {model.id! <= 0 && (
          <>
            <Row gutter={24}>
              <Col span={12}>
                <Form.Item name="version" label="版本号" rules={[{ required: true, message: "请输入版本号" }]}>
                  <Input placeholder="请输入版本号" />
                </Form.Item>
              </Col>
              <Col span={12}>
                <Form.Item name="platform" label="系统平台" rules={[{ required: true, message: "请输入系统平台" }]}>
                  <Select placeholder="请选择">
                    {Array.from(platformOptions).map(([dictValue, dictLabel]) => (
                      <Select.Option key={dictValue} value={dictValue}>
                        {dictLabel}
                      </Select.Option>
                    ))}
                  </Select>
                </Form.Item>
              </Col>

              <Col span={12}>
                <Form.Item name="downloadType" label="下载方式" rules={[{ required: true, message: "请输入下载方式" }]}>
                  <Select placeholder="请选择">
                    {Array.from(downloadTypeOptions).map(([dictValue, dictLabel]) => (
                      <Select.Option key={dictValue} value={dictValue}>
                        {dictLabel}
                      </Select.Option>
                    ))}
                  </Select>
                </Form.Item>
              </Col>
              <Col span={12}>
                <Form.Item name="appType" label="版本类型" rules={[{ required: true, message: "请输入版本类型" }]}>
                  <Select placeholder="请选择">
                    {Array.from(appTypeOptions).map(([dictValue, dictLabel]) => (
                      <Select.Option key={dictValue} value={dictValue}>
                        {dictLabel}
                      </Select.Option>
                    ))}
                  </Select>
                </Form.Item>
              </Col>
            </Row>
            {downloadType === appDownloadTypeUrl && (
              <Form.Item name="downloadUrl" label="下载地址" rules={[{ required: true, message: "请输入下载地址" }]}>
                <Input placeholder="请输入下载地址" />
              </Form.Item>
            )}
            {downloadType === appDownloadTypeLocal && (
              <>
                <Form.Item name="localRootUrl" label="本地根地址" rules={[{ required: true, message: "请输入本地根地址" }]}>
                  <Input placeholder="请输入本地根地址" />
                </Form.Item>
              </>
            )}
            {(downloadType === appDownloadTypeLocal || downloadType === appDownloadTypeOss) && (
              <>
                <Dragger {...props}>
                  <p className="ant-upload-drag-icon">
                    <InboxOutlined />
                  </p>
                  <p className="ant-upload-text">
                    将文件拖到此处，或<em>点击上传</em>
                  </p>
                  <p className="ant-upload-hint">提示：仅允许导入“apk”或“ipa”格式文件！</p>
                </Dragger>
                <Form.Item name="localAddress" style={{ display: "none" }}>
                  <Input />
                </Form.Item>
              </>
            )}
            <Form.Item name="remark" label="更新内容" rules={[{ required: true, message: "请输入更新内容" }]}>
              <Input.TextArea placeholder="请输入更新内容" />
            </Form.Item>
          </>
        )}
        {model.id! > 0 && (
          <Form.Item name="status" label="发布状态" rules={[{ required: true, message: "请输入发布状态" }]}>
            <Select placeholder="请选择">
              {Array.from(statusOptions).map(([dictValue, dictLabel]) => (
                <Select.Option key={dictValue} value={dictValue}>
                  {dictLabel}
                </Select.Option>
              ))}
            </Select>
          </Form.Item>
        )}
      </Form>
    </Modal>
  );
});

export default FormModal;
