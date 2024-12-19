import { ApiModel, getApiListApi } from "@/api/admin/sys/sys-api";
import { getDictOptions, getDictsApi } from "@/api/admin/sys/sys-dictdata";
import { addMenuApi, getMenuApi, getMenuListApi, MenuModel, updateMenuApi } from "@/api/admin/sys/sys-menu";
import IconSelect from "@/components/IconSelect";
import LoadingButton from "@/components/LoadingButton";
import { ResultEnum } from "@/enums/httpEnum";
import { message } from "@/hooks/useMessage";
import { Col, Drawer, Form, Input, InputNumber, Radio, Row, Select, Space, Transfer, TreeSelect } from "antd";
import { TransferProps } from "antd/lib";
import { forwardRef, useEffect, useImperativeHandle, useState } from "react";

export interface FormDrawerRef {
  showAddFormDrawer: (id: number) => void;
  showEditFormDrawer: (id: number) => void;
}

interface DrawerProps {
  onConfirm: () => void;
}

const MENU_TYPE_DIRECT = "1";
const MENU_TYPE_MENU = "2";
const MENU_TYPE_BUTTON = "3";

const FormDrawer = forwardRef<FormDrawerRef, DrawerProps>(({ onConfirm }, ref) => {
  const [form] = Form.useForm();
  const [isDrawerOpen, setIsDrawerOpen] = useState(false);
  const [drawer, setDrawer] = useState<MenuModel>({});
  const [isFrameOptions, setIsFrameOptions] = useState<Map<string, string>>(new Map());
  const [isAffixOptions, setIsAffixOptions] = useState<Map<string, string>>(new Map());
  const [isHiddenOptions, setIsHiddenOptions] = useState<Map<string, string>>(new Map());
  const [isKeepAliveOptions, setIsKeepAliveOptions] = useState<Map<string, string>>(new Map());
  const [menuTypeOptions, setMenuTypeOptions] = useState<Map<string, string>>(new Map());
  const [parentMenu, setParentMenu] = useState<MenuModel>();
  const [menuList, setMenuList] = useState<MenuModel[]>();
  const [menuType, setMenuType] = useState<string>("");
  const [apiList, setApiList] = useState<ApiModel[]>([]);
  const [apiSelectKeys, setApiSelectKeys] = useState<TransferProps["targetKeys"]>([]);

  // const [selectedIcon, setSelectedIcon] = useState<JSX.Element | null>(null);
  // const [searchIconText, setSearchIconText] = useState<string>("");

  useImperativeHandle(ref, () => ({
    async showAddFormDrawer(id: number) {
      reset();

      const { data: menuListData, msg: menuListMsg, code: menuListCode } = await getMenuListApi({});
      if (menuListCode !== ResultEnum.SUCCESS) {
        message.error(menuListMsg);
        return;
      }
      setMenuList(menuListData);

      if (id > 0) {
        const { data, msg, code } = await getMenuApi(id);
        if (code !== ResultEnum.SUCCESS) {
          message.error(msg);
          return;
        }
        setParentMenu(data);
        form.setFieldValue("parentId", data.id);
      }

      setIsDrawerOpen(true);
    },
    async showEditFormDrawer(id: number) {
      const { data, msg, code } = await getMenuApi(id);
      if (code !== ResultEnum.SUCCESS) {
        message.error(msg);
        return;
      }
      setApiSelectKeys(data.apis);

      const { data: menuListData, msg: menuListMsg, code: menuListCode } = await getMenuListApi({});
      if (menuListCode !== ResultEnum.SUCCESS) {
        message.error(menuListMsg);
        return;
      }
      const parentNode = getParentMenuModel(data, menuListData);
      setParentMenu(parentNode);
      setDrawer(data);
      setMenuType(data.menuType!);
      form.setFieldsValue(data);
      setMenuList(menuListData);
      setIsDrawerOpen(true);
    }
  }));
  useEffect(() => {
    const getDict = async () => {
      const { data: isFrameData, msg: isFrameMsg, code: isFrameCode } = await getDictsApi("admin_sys_yes_no");
      if (isFrameCode !== ResultEnum.SUCCESS) {
        message.error(isFrameMsg);
        return;
      }
      setIsFrameOptions(getDictOptions(isFrameData));

      const { data: isAffixData, msg: isAffixMsg, code: isAffixCode } = await getDictsApi("admin_sys_yes_no");
      if (isAffixCode !== ResultEnum.SUCCESS) {
        message.error(isAffixMsg);
        return;
      }
      setIsAffixOptions(getDictOptions(isAffixData));

      const { data: isHiddenData, msg: isHiddenMsg, code: isHiddenCode } = await getDictsApi("admin_sys_menu_show_hide");
      if (isHiddenCode !== ResultEnum.SUCCESS) {
        message.error(isHiddenMsg);
        return;
      }
      setIsHiddenOptions(getDictOptions(isHiddenData));
      const { data: isKeepAliveData, msg: isKeepAliveMsg, code: isKeepAliveCode } = await getDictsApi("admin_sys_yes_no");
      if (isKeepAliveCode !== ResultEnum.SUCCESS) {
        message.error(isKeepAliveMsg);
        return;
      }
      setIsKeepAliveOptions(getDictOptions(isKeepAliveData));
      const { data: menuTypeData, msg: menuTypeMsg, code: menuTypeCode } = await getDictsApi("admin_sys_menu_type");
      if (menuTypeCode !== ResultEnum.SUCCESS) {
        message.error(menuTypeMsg);
        return;
      }
      setMenuTypeOptions(getDictOptions(menuTypeData));

      const { data: apiData, msg: apiMsg, code: apiCode } = await getApiListApi({});
      if (apiCode !== ResultEnum.SUCCESS) {
        message.error(apiMsg);
        return;
      }
      setApiList(apiData);
    };
    getDict();
  }, []);

  const reset = () => {
    if (drawer.id! > 0) {
      setDrawer({});
    } else {
      setDrawer({ id: 0 });
    }
    // setMenuType("");
    // setApiList([]);
    setApiSelectKeys([]);
    // setMenuList([]);
    setMenuType("");
    setTimeout(() => form.resetFields(), 100);
  };

  const handleConfirm = (done: () => void) => {
    form
      .validateFields()
      .then(async values => {
        try {
          if (drawer.id! <= 0) {
            const { msg, code } = await addMenuApi(values);
            if (code !== ResultEnum.SUCCESS) {
              message.error(msg);
              return;
            }
            message.success(msg);
          } else {
            const { msg, code } = await updateMenuApi(drawer.id!, values);
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
        setIsDrawerOpen(false);
        onConfirm();
      })
      .catch(error => {
        console.error("validate error：", error);
        message.error("表单校验失败");
        done();
      });
  };

  const onMenuTypeChange = (changedValues: Partial<MenuModel>, allValues: MenuModel) => {
    setMenuType(allValues.menuType!);
  };

  return (
    <Drawer
      title={drawer.id! > 0 ? "编辑" : "新增"}
      getContainer={false}
      width={800}
      open={isDrawerOpen}
      maskClosable={false}
      keyboard={false}
      onClose={() => {
        reset();
        setIsDrawerOpen(false);
      }}
      destroyOnClose
      extra={[
        <Space key="space">
          <LoadingButton
            key="cancel1"
            onClick={done => {
              reset();
              setIsDrawerOpen(false);
              done();
            }}
          >
            取消
          </LoadingButton>
          <LoadingButton key="confirm2" type="primary" onClick={done => handleConfirm(done)}>
            确定
          </LoadingButton>
        </Space>
      ]}
    >
      <Form form={form} layout="vertical" initialValues={drawer} onValuesChange={onMenuTypeChange}>
        <Row gutter={24}>
          <Col span={24}>
            <Form.Item name="parentId" label="上级菜单" rules={[{ required: true, message: "选择上级菜单" }]}>
              <TreeSelect
                showSearch
                style={{ width: "100%" }}
                value={parentMenu}
                treeNodeFilterProp="title"
                fieldNames={{ label: "title", value: "id", children: "children" }}
                dropdownStyle={{ maxHeight: 400, overflow: "auto" }}
                placeholder="选择上级菜单"
                allowClear
                treeDefaultExpandAll
                onChange={(newMenu: MenuModel) => {
                  setParentMenu(newMenu);
                }}
                treeData={menuList}
              />
            </Form.Item>
          </Col>
          <Col span={18}>
            <Form.Item name="title" label="菜单标题" rules={[{ required: true, message: "请输入路由名称" }]}>
              <Input placeholder="请输入菜单标题" />
            </Form.Item>
          </Col>
          <Col span={6}>
            <Form.Item name="sort" label="排序" rules={[{ required: true, message: "请输入排序" }]}>
              <InputNumber style={{ width: "100%" }} min={0} />
            </Form.Item>
          </Col>
          <Col span={24}>
            <Form.Item name="icon" label="图标">
              <IconSelect></IconSelect>
            </Form.Item>
          </Col>
          <Col span={24}>
            <Form.Item name="menuType" label="菜单类型" rules={[{ required: true, message: "请输入菜单类型" }]}>
              <Radio.Group>
                {Array.from(menuTypeOptions).map(([dictValue, dictLabel]) => (
                  <Radio key={dictValue} value={dictValue}>
                    {dictLabel}
                  </Radio>
                ))}
              </Radio.Group>
            </Form.Item>
          </Col>

          {(menuType === MENU_TYPE_DIRECT || menuType === MENU_TYPE_MENU) && (
            <>
              <Col span={12}>
                <Form.Item name="path" label="路由地址" rules={[{ required: true, message: "请输入路由地址" }]}>
                  <Input placeholder="请输入路由地址" />
                </Form.Item>
              </Col>
              <Col span={12}>
                <Form.Item name="isHidden" label="是否隐藏" rules={[{ required: true, message: "请输入是否隐藏" }]}>
                  <Select placeholder="请选择">
                    {Array.from(isHiddenOptions).map(([dictValue, dictLabel]) => (
                      <Select.Option key={dictValue} value={dictValue}>
                        {dictLabel}
                      </Select.Option>
                    ))}
                  </Select>
                </Form.Item>
              </Col>
              {menuType === MENU_TYPE_DIRECT && (
                <Col span={12}>
                  <Form.Item name="redirect" label="跳转路由" rules={[{ required: true, message: "请输入跳转路由" }]}>
                    <Input placeholder="请输入跳转路由" />
                  </Form.Item>
                </Col>
              )}
              {menuType === MENU_TYPE_MENU && (
                <>
                  <Col span={12}>
                    <Form.Item name="element" label="组件路径" rules={[{ required: true, message: "请输入组件路径" }]}>
                      <Input placeholder="请输入组件路径" />
                    </Form.Item>
                  </Col>
                  <Col span={12}>
                    <Form.Item name="isAffix" label="是否固定" rules={[{ required: true, message: "请输入是否固定" }]}>
                      <Select placeholder="请选择">
                        {Array.from(isAffixOptions).map(([dictValue, dictLabel]) => (
                          <Select.Option key={dictValue} value={dictValue}>
                            {dictLabel}
                          </Select.Option>
                        ))}
                      </Select>
                    </Form.Item>
                  </Col>
                  <Col span={12}>
                    <Form.Item name="isFrame" label="是否内嵌" rules={[{ required: true, message: "请输入是否内嵌" }]}>
                      <Select placeholder="请选择">
                        {Array.from(isFrameOptions).map(([dictValue, dictLabel]) => (
                          <Select.Option key={dictValue} value={dictValue}>
                            {dictLabel}
                          </Select.Option>
                        ))}
                      </Select>
                    </Form.Item>
                  </Col>
                  <Col span={12}>
                    <Form.Item name="isKeepAlive" label="是否缓存" rules={[{ required: true, message: "请输入是否缓存" }]}>
                      <Select placeholder="请选择">
                        {Array.from(isKeepAliveOptions).map(([dictValue, dictLabel]) => (
                          <Select.Option key={dictValue} value={dictValue}>
                            {dictLabel}
                          </Select.Option>
                        ))}
                      </Select>
                    </Form.Item>
                  </Col>
                </>
              )}
            </>
          )}
          {menuType === MENU_TYPE_BUTTON && (
            <Col span={12}>
              <Form.Item name="permission" label="权限标识" rules={[{ required: true, message: "请输入权限标识" }]}>
                <Input placeholder="请输入权限标识" />
              </Form.Item>
            </Col>
          )}

          {(menuType === MENU_TYPE_MENU || menuType === MENU_TYPE_BUTTON) && (
            <Col span={24}>
              <Form.Item name="apis" label="授权接口">
                <Transfer
                  dataSource={apiList}
                  showSearch
                  filterOption={(inputValue: string, option: ApiModel) => {
                    return option.description!.indexOf(inputValue) > -1;
                  }}
                  listStyle={{
                    width: 500,
                    height: 500
                  }}
                  rowKey={record => record.id!}
                  targetKeys={apiSelectKeys}
                  onChange={newTargetKeys => {
                    setApiSelectKeys(newTargetKeys);
                  }}
                  onSearch={(dir, value) => {
                    //console.log("search:", dir, value);
                  }}
                  render={item => item.description || "未知API"}
                />
              </Form.Item>
            </Col>
          )}
        </Row>
      </Form>
    </Drawer>
  );
});

export default FormDrawer;

// 获取当前选中项的父级路径
const getParentMenuModel = (currentNode: MenuModel, treeList: MenuModel[]) => {
  if (!currentNode) {
    return;
  }
  for (let node of treeList) {
    if (currentNode.parentId === node.id || currentNode.id == node.id) {
      return node; // 如果找到匹配项，返回当前路径
    }
    if (node.children) {
      getParentMenuModel(currentNode, node.children);
    }
  }
};
