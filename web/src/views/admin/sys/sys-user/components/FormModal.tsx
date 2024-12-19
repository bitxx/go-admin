import { DeptModel, getDeptTreeApi } from "@/api/admin/sys/sys-dept";
import { getDictOptions, getDictsApi } from "@/api/admin/sys/sys-dictdata";
import { getPostTotalListApi } from "@/api/admin/sys/sys-post";
import { getRoleTotalListApi } from "@/api/admin/sys/sys-role";
import { addUserApi, getUserApi, updateUserApi, UserModel } from "@/api/admin/sys/sys-user";
import LoadingButton from "@/components/LoadingButton";
import { ResultEnum } from "@/enums/httpEnum";
import { message } from "@/hooks/useMessage";
import { Col, Form, Input, Modal, Radio, Row, Select, TreeSelect } from "antd";
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
  const [model, setModel] = useState<UserModel>({});
  const [sexOptions, setSexOptions] = useState<Map<string, string>>(new Map());
  const [statusOptions, setStatusOptions] = useState<Map<string, string>>(new Map());
  const [postOptions, setPostOptions] = useState<Map<number, string>>(new Map());
  const [roleOptions, setRoleOptions] = useState<Map<number, string>>(new Map());
  const [selectDept, setSelectDept] = useState<DeptModel>();
  const [deptList, setDeptList] = useState<DeptModel[]>();

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
      setSelectDept(data.dept);

      setIsModalOpen(true);
    }
  }));
  useEffect(() => {
    const initData = async () => {
      const { data: sexData, msg: sexMsg, code: sexCode } = await getDictsApi("admin_sys_user_sex");
      if (sexCode !== ResultEnum.SUCCESS) {
        message.error(sexMsg);
        return;
      }
      setSexOptions(getDictOptions(sexData));
      const { data: statusData, msg: statusMsg, code: statusCode } = await getDictsApi("admin_sys_status");
      if (statusCode !== ResultEnum.SUCCESS) {
        message.error(statusMsg);
        return;
      }
      setStatusOptions(getDictOptions(statusData));
      const { data: postData, msg: postMsg, code: postCode } = await getPostTotalListApi({});
      if (postCode !== ResultEnum.SUCCESS) {
        message.error(postMsg);
        return;
      }
      setPostOptions(new Map(postData.map(item => [item.id!, item.postName!])));

      const { data: roleData, msg: roleMsg, code: roleCode } = await getRoleTotalListApi({});
      if (roleCode !== ResultEnum.SUCCESS) {
        message.error(roleMsg);
        return;
      }
      setRoleOptions(new Map(roleData.map(item => [item.id!, item.roleName!])));

      const { data: deptListData, msg: deptListMsg, code: deptListCode } = await getDeptTreeApi({});
      if (deptListCode !== ResultEnum.SUCCESS) {
        message.error(deptListMsg);
        return;
      }
      setDeptList(deptListData);
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
        <Row gutter={24}>
          <Col span={12}>
            <Form.Item name="username" label="用户名" rules={[{ required: true, message: "请输入用户名" }]}>
              <Input placeholder="请输入用户名" />
            </Form.Item>
          </Col>
          <Col span={12}>
            <Form.Item name="nickName" label="昵称" rules={[{ required: true, message: "请输入昵称" }]}>
              <Input placeholder="请输入昵称" />
            </Form.Item>
          </Col>
          <Col span={12}>
            <Form.Item name="sex" label="性别" rules={[{ required: true, message: "请输入性别" }]}>
              <Select placeholder="请选择">
                {Array.from(sexOptions).map(([dictValue, dictLabel]) => (
                  <Select.Option key={dictValue} value={dictValue}>
                    {dictLabel}
                  </Select.Option>
                ))}
              </Select>
            </Form.Item>
          </Col>
          {model.id! <= 0 && (
            <Col span={12}>
              <Form.Item name="password" label="初始密码" rules={[{ required: true, message: "请输入初始密码" }]}>
                <Input placeholder="请输入初始密码" />
              </Form.Item>
            </Col>
          )}
          <Col span={12}>
            <Form.Item name="phone" label="手机号" rules={[{ required: true, message: "请输入手机号" }]}>
              <Input placeholder="请输入手机号" />
            </Form.Item>
          </Col>
          <Col span={12}>
            <Form.Item name="email" label="邮箱" rules={[{ required: true, message: "请输入邮箱" }]}>
              <Input placeholder="请输入邮箱" />
            </Form.Item>
          </Col>
          <Col span={12}>
            <Form.Item name="status" label="状态" rules={[{ required: true, message: "请输入状态" }]}>
              <Radio.Group>
                {Array.from(statusOptions).map(([dictValue, dictLabel]) => (
                  <Radio key={dictValue} value={dictValue}>
                    {dictLabel}
                  </Radio>
                ))}
              </Radio.Group>
            </Form.Item>
          </Col>
          <Col span={12}>
            <Form.Item name="deptId" label="部门" rules={[{ required: true, message: "请选择部门" }]}>
              <TreeSelect
                showSearch
                style={{ width: "100%" }}
                value={selectDept}
                treeNodeFilterProp="deptName"
                fieldNames={{ label: "deptName", value: "id", children: "children" }}
                dropdownStyle={{ maxHeight: 400, overflow: "auto" }}
                placeholder="请选择部门"
                allowClear
                treeDefaultExpandAll
                onChange={(newDept: DeptModel) => {
                  setSelectDept(newDept);
                }}
                treeData={deptList}
                //onPopupScroll={onPopupScroll => {}}
              />
            </Form.Item>
          </Col>
          <Col span={12}>
            <Form.Item name="postId" label="岗位" rules={[{ required: true, message: "请选择岗位" }]}>
              <Select placeholder="请选择">
                {Array.from(postOptions).map(([id, postName]) => (
                  <Select.Option key={id} value={id}>
                    {postName}
                  </Select.Option>
                ))}
              </Select>
            </Form.Item>
          </Col>
          <Col span={12}>
            <Form.Item name="roleId" label="角色" rules={[{ required: true, message: "请选择角色" }]}>
              <Select placeholder="请选择">
                {Array.from(roleOptions).map(([id, roleName]) => (
                  <Select.Option key={id} value={id}>
                    {roleName}
                  </Select.Option>
                ))}
              </Select>
            </Form.Item>
          </Col>
          <Col span={24}>
            <Form.Item name="remark" label="备注">
              <Input.TextArea placeholder="请输入备注" />
            </Form.Item>
          </Col>
        </Row>
      </Form>
    </Modal>
  );
});

export default FormModal;
