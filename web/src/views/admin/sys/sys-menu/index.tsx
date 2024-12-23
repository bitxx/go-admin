import { getDictOptions, getDictsApi } from "@/api/admin/sys/sys-dictdata";
import { delMenuApi, getMenuListApi, MenuModel } from "@/api/admin/sys/sys-menu";
import HocAuth from "@/components/HocAuth";
import { Icon } from "@/components/Icon";
import LoadingButton from "@/components/LoadingButton";
import { pagination } from "@/config/proTable";
import { ResultEnum } from "@/enums/httpEnum";
import { message, modal } from "@/hooks/useMessage";
import { DeleteOutlined, EditOutlined, ExclamationCircleOutlined, PlusCircleOutlined } from "@ant-design/icons";
import type { ActionType, ProColumns, ProFormInstance } from "@ant-design/pro-components";
import { ProTable } from "@ant-design/pro-components";
import { Space } from "antd";
import React, { useEffect, useRef, useState } from "react";
import FormDrawer, { FormDrawerRef } from "./components/FormDrawer";

const Menu: React.FC = () => {
  const actionRef = React.useRef<ActionType>();
  const tableFormRef = React.useRef<ProFormInstance>();
  const formDrawerRef = useRef<FormDrawerRef>(null);
  const [isHiddenOptions, setIsHiddenOptions] = useState<Map<string, string>>(new Map());
  const [menuTypeOptions, setMenuTypeOptions] = useState<Map<string, string>>(new Map());

  // 定义列
  const columns: ProColumns<MenuModel>[] = [
    {
      title: "菜单名称",
      dataIndex: "title",
      width: 300,
      align: "left",
      render: (_, record) => (
        <Space>
          <Icon name={record.icon!} /> {record.title}
        </Space>
      )
    },
    {
      title: "菜单编号",
      dataIndex: "id",
      width: 80,
      align: "center"
    },
    {
      title: "菜单类型",
      dataIndex: "menuType",
      hideInSearch: true,
      valueType: "select",
      valueEnum: menuTypeOptions,
      width: 80,
      align: "left"
    },
    {
      title: "权限标识",
      dataIndex: "permission",
      hideInSearch: true,
      width: 350,
      align: "left"
    },
    {
      title: "是否隐藏",
      dataIndex: "isHidden",
      valueType: "select",
      valueEnum: isHiddenOptions,
      width: 80,
      align: "left"
    },
    {
      title: "路由地址",
      dataIndex: "path",
      hideInSearch: true,
      width: 300,
      align: "left"
    },
    {
      title: "组件路径",
      dataIndex: "element",
      hideInSearch: true,
      width: 350,
      align: "left"
    },
    {
      title: "排序",
      dataIndex: "sort",
      hideInSearch: true,
      width: 50,
      align: "left"
    },
    {
      title: "创建时间",
      dataIndex: "createdAt",
      hideInSearch: true,
      valueType: "dateTime",
      width: 180,
      align: "left"
    },

    {
      title: "操作",
      valueType: "option",
      align: "center",
      fixed: "right",
      width: 230,
      render: (_, data) =>
        data.id !== undefined &&
        data.id > 0 && (
          <Space>
            <HocAuth permission={["admin:sys-menu:add"]}>
              <LoadingButton
                key="addItem"
                type="link"
                size="small"
                icon={<PlusCircleOutlined />}
                onClick={done => handleShowAddFormModal(data.id!, done)}
              >
                新增
              </LoadingButton>
            </HocAuth>
            <HocAuth permission={["admin:sys-menu:edit"]}>
              <LoadingButton
                key="edit"
                type="link"
                size="small"
                icon={<EditOutlined />}
                onClick={done => handleShowEditFormModal(data.id!, done)}
              >
                编辑
              </LoadingButton>
            </HocAuth>
            <HocAuth permission={["admin:sys-menu:del"]}>
              <LoadingButton
                key="delete"
                type="link"
                size="small"
                danger
                icon={<DeleteOutlined />}
                onClick={done => handleDelete(data.id!, done)}
              >
                删除
              </LoadingButton>
            </HocAuth>
          </Space>
        )
    }
  ];

  useEffect(() => {
    const initData = async () => {
      const { data: isHiddenData, msg: isHiddenMsg, code: isHiddenCode } = await getDictsApi("admin_sys_menu_show_hide");
      if (isHiddenCode !== ResultEnum.SUCCESS) {
        message.error(isHiddenMsg);
        return;
      }
      setIsHiddenOptions(getDictOptions(isHiddenData));
      const { data: menuTypeData, msg: menuTypeMsg, code: menuTypeCode } = await getDictsApi("admin_sys_menu_type");
      if (menuTypeCode !== ResultEnum.SUCCESS) {
        message.error(menuTypeMsg);
        return;
      }
      setMenuTypeOptions(getDictOptions(menuTypeData));
    };
    initData();
  }, []);

  const handleShowAddFormModal = (id: number, done: () => void) => {
    formDrawerRef.current?.showAddFormDrawer(id);
    setTimeout(() => done(), 1000);
  };

  const handleShowEditFormModal = (id: number, done: () => void) => {
    formDrawerRef.current?.showEditFormDrawer(id);
    setTimeout(() => done(), 1000);
  };

  const handleFormDrawerConfirm = () => {
    actionRef.current?.reload(true);
  };

  const handleDelete = (id: number, done: () => void) => {
    modal.confirm({
      title: "提示",
      icon: <ExclamationCircleOutlined />,
      content: "是否确认删除编号为 " + id + " 的数据项?",
      okText: "确认",
      cancelText: "取消",
      maskClosable: true,
      onCancel: () => {
        done();
      },
      onOk: async () => {
        try {
          const { code, msg } = await delMenuApi([id!]);
          if (code !== ResultEnum.SUCCESS) {
            message.error(msg);
            return;
          }
          actionRef.current?.reload(true);
          message.success(msg);
        } finally {
          done();
        }
      }
    });
  };

  const toolBarRender = () => [
    <HocAuth permission={["admin:sys-menu:add"]}>
      <LoadingButton
        type="primary"
        key="addTable"
        icon={<PlusCircleOutlined />}
        onClick={done => handleShowAddFormModal(0, done)}
      >
        新增
      </LoadingButton>
    </HocAuth>
  ];

  return (
    <>
      <ProTable<MenuModel>
        className="ant-pro-table-scroll"
        columns={columns}
        actionRef={actionRef}
        formRef={tableFormRef}
        bordered
        cardBordered
        defaultSize="small"
        scroll={{ x: "2000", y: "100%" }}
        request={async params => {
          const { data } = await getMenuListApi(params);
          return {
            success: true,
            data: data,
            total: data.length
          };
        }}
        columnsState={{
          persistenceKey: "use-pro-table-key",
          persistenceType: "localStorage"
        }}
        options={{
          reload: true,
          density: true,
          fullScreen: true
        }}
        rowKey="id"
        search={{ labelWidth: "auto", showHiddenNum: true }}
        pagination={pagination}
        dateFormatter="string"
        headerTitle="菜单管理"
        expandable={{
          // 支持展开/折叠
          defaultExpandedRowKeys: [0],
          childrenColumnName: "children", // 指定子项字段
          defaultExpandAllRows: true // 默认展开所有行
        }}
        toolBarRender={toolBarRender}
      />
      <FormDrawer ref={formDrawerRef} onConfirm={handleFormDrawerConfirm} />
    </>
  );
};

export default Menu;
