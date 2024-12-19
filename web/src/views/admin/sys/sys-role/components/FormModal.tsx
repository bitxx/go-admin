import { getDictOptions, getDictsApi } from "@/api/admin/sys/sys-dictdata";
import { MenuTreeRole, roleMenuTreeselectApi } from "@/api/admin/sys/sys-menu";
import { addRoleApi, getRoleApi, RoleModel, updateRoleApi } from "@/api/admin/sys/sys-role";
import LoadingButton from "@/components/LoadingButton";
import { ResultEnum } from "@/enums/httpEnum";
import { message } from "@/hooks/useMessage";
import { Col, Form, Input, InputNumber, Modal, Radio, Row, Tree } from "antd";
import { forwardRef, Key, useEffect, useImperativeHandle, useState } from "react";

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
  const [model, setModel] = useState<RoleModel>({});
  const [statusOptions, setStatusOptions] = useState<Map<string, string>>(new Map());
  const [menuSelect, setMenuSelect] = useState<Key[]>([]);
  const [menuTreeRole, setMenuTreeRole] = useState<MenuTreeRole>({});

  useImperativeHandle(ref, () => ({
    showAddFormModal() {
      reset();
      roleMenuTreeselect(0);
      setIsModalOpen(true);
    },
    async showEditFormModal(id: number) {
      const { data, msg, code } = await getRoleApi(id);
      if (code !== ResultEnum.SUCCESS) {
        message.error(msg);
        return;
      }

      setTimeout(() => {
        setMenuSelect(data.menuIds || []);
      }, 500);

      setModel(data);
      roleMenuTreeselect(id);
      form.setFieldsValue(data);
      setIsModalOpen(true);
    }
  }));

  const roleMenuTreeselect = async (id: number) => {
    const { data } = await roleMenuTreeselectApi(id);
    setMenuTreeRole(data);
  };

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
    setMenuSelect([]);
    setMenuTreeRole({});

    setTimeout(() => form.resetFields(), 100);
  };

  const handleConfirm = async (done: () => void) => {
    form
      .validateFields()
      .then(async values => {
        try {
          if (model.id! <= 0) {
            const { msg, code } = await addRoleApi({ ...values, menuIds: menuSelect });
            if (code !== ResultEnum.SUCCESS) {
              message.error(msg);
              return;
            }
            message.success(msg);
          } else {
            const { msg, code } = await updateRoleApi(model.id!, { ...values, menuIds: menuSelect });
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
          onClick={async done => {
            reset();
            setIsModalOpen(false);
            done();
          }}
        >
          取消
        </LoadingButton>,
        <LoadingButton key="confirm" type="primary" onClick={handleConfirm}>
          确定
        </LoadingButton>
      ]}
    >
      <Form form={form} layout="vertical" initialValues={model}>
        <Row gutter={24}>
          <Col span={12}>
            <Form.Item name="roleName" label="角色名称" rules={[{ required: true, message: "请输入角色名称" }]}>
              <Input placeholder="请输入角色名称" />
            </Form.Item>
          </Col>
          <Col span={12}>
            <Form.Item name="roleKey" label="角色键" rules={[{ required: true, message: "请输入角色键" }]}>
              <Input placeholder="请输入角色键" />
            </Form.Item>
          </Col>
          <Col span={12}>
            <Form.Item name="roleSort" label="排序" rules={[{ required: true, message: "请输入排序" }]}>
              <InputNumber placeholder="请输入排序" style={{ width: "100%" }} min={0} />
            </Form.Item>
          </Col>
          <Col span={12}>
            <Form.Item name="status" label="角色状态" rules={[{ required: true, message: "请输入角色状态" }]}>
              <Radio.Group>
                {Array.from(statusOptions).map(([dictValue, dictLabel]) => (
                  <Radio key={dictValue} value={dictValue}>
                    {dictLabel}
                  </Radio>
                ))}
              </Radio.Group>
            </Form.Item>
          </Col>
          <Col span={24}>
            <Form.Item name="menuOptions" label="菜单权限">
              <Tree
                checkable
                checkedKeys={menuSelect}
                onCheck={(checked, halfChecked) => {
                  const keys = Array.isArray(checked) ? checked : checked.checked;
                  setMenuSelect(keys);
                }}
                defaultExpandedKeys={[0]}
                fieldNames={{ title: "title", key: "id", children: "children" }}
                treeData={menuTreeRole.menus as any[]}
                style={{
                  width: "100%"
                }}
              />
            </Form.Item>
          </Col>
        </Row>

        <Form.Item name="remark" label="备注">
          <Input.TextArea placeholder="请输入备注" />
        </Form.Item>
      </Form>
    </Modal>
  );
});

export default FormModal;
