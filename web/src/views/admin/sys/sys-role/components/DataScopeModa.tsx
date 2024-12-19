import { DeptTreeRole, roleDeptTreeselectApi } from "@/api/admin/sys/sys-dept";
import { getDictOptions, getDictsApi } from "@/api/admin/sys/sys-dictdata";
import { dataScopeApi, getRoleApi, RoleModel } from "@/api/admin/sys/sys-role";
import LoadingButton from "@/components/LoadingButton";
import { ResultEnum } from "@/enums/httpEnum";
import { message } from "@/hooks/useMessage";
import { Col, Form, Input, Modal, Row, Select, Tree } from "antd";
import { forwardRef, Key, useEffect, useImperativeHandle, useState } from "react";

const dataScopeSelf = "2";

export interface DataScopeFormModalRef {
  showDataScopeFormModal: (id: number) => void;
}

interface ModalProps {
  onConfirm: () => void;
}

const DataScopeFormModal = forwardRef<DataScopeFormModalRef, ModalProps>(({ onConfirm }, ref) => {
  const [form] = Form.useForm();
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [model, setModel] = useState<RoleModel>({});
  const [dataScopeOptions, setDataScopeOptions] = useState<Map<string, string>>(new Map());

  const [deptSelect, setDeptSelect] = useState<Key[]>([]);
  const [deptTreeRole, setDeptTreeRole] = useState<DeptTreeRole>({});
  const [currentDataScope, setCurrentDataScope] = useState<string | undefined>("");

  useImperativeHandle(ref, () => ({
    async showDataScopeFormModal(id: number) {
      const { data: roleData, msg: roleMsg, code: roleCode } = await getRoleApi(id);
      if (roleCode !== ResultEnum.SUCCESS) {
        message.error(roleMsg);
        return;
      }

      const { data: deptData, msg: deptMsg, code: roleDeptCode } = await roleDeptTreeselectApi(id);
      if (roleDeptCode !== ResultEnum.SUCCESS) {
        message.error(deptMsg);
        return;
      }

      setTimeout(() => {
        setDeptSelect(roleData.deptIds || []);
      }, 500);
      setTimeout(() => {
        setDeptTreeRole(deptData);
      }, 500);

      setModel(roleData);
      setCurrentDataScope(roleData.dataScope);
      form.setFieldsValue(roleData);
      setIsModalOpen(true);
    }
  }));

  // 监听表单值变化
  const onDataScopeChange = (changedValues: Partial<RoleModel>, allValues: RoleModel) => {
    setCurrentDataScope(allValues.dataScope!);
  };

  useEffect(() => {
    const initData = async () => {
      const { data: dataScopeData, msg: dataScopeMsg, code: dataScopeCode } = await getDictsApi("admin_sys_role_data_scope");
      if (dataScopeCode !== ResultEnum.SUCCESS) {
        message.error(dataScopeMsg);
        return;
      }
      setDataScopeOptions(getDictOptions(dataScopeData));
    };
    initData();
  }, []);

  const reset = () => {
    if (model.id! > 0) {
      setModel({});
    } else {
      setModel({ id: 0 });
    }
    setDeptSelect([]);
    setDeptTreeRole({});

    setTimeout(() => form.resetFields(), 100);
  };

  const handleConfirm = async (done: () => void) => {
    form
      .validateFields()
      .then(async values => {
        try {
          let deptIds = deptSelect;
          if (currentDataScope !== dataScopeSelf) {
            deptIds = [];
          }
          const { msg, code } = await dataScopeApi({ ...values, id: model.id!, deptIds: deptIds });
          if (code !== ResultEnum.SUCCESS) {
            message.error(msg);
            return;
          }
          message.success(msg);
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
      title={"数据权限"}
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
      <Form form={form} layout="vertical" onValuesChange={onDataScopeChange} initialValues={model}>
        <Row gutter={24}>
          <Col span={12}>
            <Form.Item name="roleName" label="角色名称">
              <Input disabled />
            </Form.Item>
          </Col>
          <Col span={12}>
            <Form.Item name="roleKey" label="角色键">
              <Input disabled />
            </Form.Item>
          </Col>
          <Col span={24}>
            <Form.Item name="dataScope" label="数据范围" rules={[{ required: true, message: "请输入数据范围" }]}>
              <Select placeholder="请选择">
                {Array.from(dataScopeOptions).map(([dictValue, dictLabel]) => (
                  <Select.Option key={dictValue} value={dictValue}>
                    {dictLabel}
                  </Select.Option>
                ))}
              </Select>
            </Form.Item>
          </Col>

          {currentDataScope == dataScopeSelf && (
            <Col span={24}>
              <Form.Item name="deptOptions" label="权限范围">
                <Tree
                  checkable
                  defaultExpandAll
                  checkedKeys={deptSelect}
                  onCheck={(checked, halfChecked) => {
                    const keys = Array.isArray(checked) ? checked : checked.checked;
                    setDeptSelect(keys);
                  }}
                  fieldNames={{ title: "deptName", key: "id", children: "children" }}
                  treeData={deptTreeRole.depts as any[]}
                  style={{
                    width: "100%"
                  }}
                />
              </Form.Item>
            </Col>
          )}
        </Row>
      </Form>
    </Modal>
  );
});

export default DataScopeFormModal;
