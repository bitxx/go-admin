import { getDictOptions, getDictsApi } from "@/api/admin/sys/sys-dictdata";
import { changeRoleStatusApi, delRoleApi, getRolePageApi, RoleModel } from "@/api/admin/sys/sys-role";
import HocAuth from "@/components/HocAuth";
import LoadingButton from "@/components/LoadingButton";
import SwitchLoading from "@/components/LoadingSwitch";
import { pagination } from "@/config/proTable";
import { ResultEnum } from "@/enums/httpEnum";
import { message, modal } from "@/hooks/useMessage";
import { formatDataForProTable } from "@/utils";
import { DeleteOutlined, EditOutlined, ExclamationCircleOutlined, PlusCircleOutlined } from "@ant-design/icons";
import type { ActionType, ProColumns, ProFormInstance } from "@ant-design/pro-components";
import { ProTable } from "@ant-design/pro-components";
import { Space } from "antd";
import React, { useEffect, useRef, useState } from "react";
import DataScopeFormModal, { DataScopeFormModalRef } from "./components/DataScopeModa";
import FormModal, { FormModalRef } from "./components/FormModal";

const Role: React.FC = () => {
  const STATUS_YES = "1";
  const STATUS_NO = "2";
  const actionRef = React.useRef<ActionType>();
  const tableFormRef = React.useRef<ProFormInstance>();
  const formModalRef = useRef<FormModalRef>(null);
  const dataScopeFormModalRef = useRef<DataScopeFormModalRef>(null);
  const [dataScopeOptions, setDataScopeOptions] = useState<Map<string, string>>(new Map());
  const [statusOptions, setStatusOptions] = useState<Map<string, string>>(new Map());

  // 定义列
  const columns: ProColumns<RoleModel>[] = [
    {
      title: "序号",
      dataIndex: "index",
      valueType: "index",
      width: 50,
      align: "center",
      className: "gray-cell",
      render: (_, __, index, action) => {
        // 根据分页计算实际序号
        const currentPage = action?.pageInfo?.current || 1;
        const pageSize = action?.pageInfo?.pageSize || 10;
        return (currentPage - 1) * pageSize + index + 1;
      }
    },
    {
      title: "角色编号",
      dataIndex: "id",
      width: 80,
      hideInSearch: true,
      align: "left"
    },
    {
      title: "角色名称",
      dataIndex: "roleName",
      width: 120,
      align: "left"
    },
    {
      title: "角色键",
      dataIndex: "roleKey",
      width: 80,
      align: "left"
    },
    {
      title: "排序",
      dataIndex: "roleSort",
      hideInSearch: true,
      width: 80,
      align: "left"
    },
    {
      title: "角色状态",
      dataIndex: "status",
      valueType: "select",
      valueEnum: statusOptions,
      width: 120,
      align: "left",
      render: (text, record, index, action) => (
        <>
          <SwitchLoading
            checked={record.status === STATUS_YES}
            checkedChildren="开启"
            unCheckedChildren="关闭"
            onChange={checked => handleStatusChange(checked, record, action)}
          />
        </>
      )
    },
    {
      title: "数据范围",
      dataIndex: "dataScope",
      hideInSearch: true,
      valueType: "select",
      valueEnum: dataScopeOptions,
      width: 180,
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
      width: 300,
      render: (_, data) =>
        data.roleKey !== "admin" && (
          <Space>
            <HocAuth permission={["admin:sys-role:edit"]}>
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
            <HocAuth permission={["admin:sys-role:datascope"]}>
              <LoadingButton
                key="dataScope"
                type="link"
                size="small"
                icon={<EditOutlined />}
                onClick={done => handleShowDataScopeFormModal(data.id!, done)}
              >
                分配数据权限
              </LoadingButton>
            </HocAuth>
            <HocAuth permission={["admin:sys-role:del"]}>
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
      const { data: dataScopeData, msg: dataScopeMsg, code: dataScopeCode } = await getDictsApi("admin_sys_role_data_scope");
      if (dataScopeCode !== ResultEnum.SUCCESS) {
        message.error(dataScopeMsg);
        return;
      }
      setDataScopeOptions(getDictOptions(dataScopeData));
      const { data: statusData, msg: statusMsg, code: statusCode } = await getDictsApi("admin_sys_status");
      if (statusCode !== ResultEnum.SUCCESS) {
        message.error(statusMsg);
        return;
      }
      setStatusOptions(getDictOptions(statusData));
    };
    initData();
  }, []);

  const handleShowAddFormModal = (done: () => void) => {
    formModalRef.current?.showAddFormModal();
    setTimeout(() => done(), 1000);
  };

  const handleShowEditFormModal = (id: number, done: () => void) => {
    formModalRef.current?.showEditFormModal(id);
    setTimeout(() => done(), 1000);
  };

  const handleShowDataScopeFormModal = (id: number, done: () => void) => {
    dataScopeFormModalRef.current?.showDataScopeFormModal(id);
    setTimeout(() => done(), 1000);
  };

  const handleFormModalConfirm = () => {
    actionRef.current?.reload(false);
  };

  const handleDataScopeFormModalConfirm = () => {
    actionRef.current?.reload(false);
  };

  const handleStatusChange = async (checked: boolean, record: RoleModel, action: any) => {
    const newStatus = checked ? STATUS_YES : STATUS_NO;
    const { code, msg } = await changeRoleStatusApi(record.id!, newStatus);
    if (code !== ResultEnum.SUCCESS) {
      message.error(msg);
      return;
    }
    action.reload();
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
          const { code, msg } = await delRoleApi([id!]);
          if (code !== ResultEnum.SUCCESS) {
            message.error(msg);
            return;
          }
          actionRef.current?.reload(false);
          message.success(msg);
        } finally {
          done();
        }
      }
    });
  };

  const toolBarRender = () => [
    <HocAuth permission={["admin:sys-role:add"]}>
      <LoadingButton
        type="primary"
        key="addTable"
        icon={<PlusCircleOutlined />}
        onClick={async done => handleShowAddFormModal(done)}
      >
        新增
      </LoadingButton>
    </HocAuth>
  ];

  return (
    <>
      <ProTable<RoleModel>
        className="ant-pro-table-scroll"
        columns={columns}
        actionRef={actionRef}
        formRef={tableFormRef}
        bordered
        cardBordered
        defaultSize="small"
        scroll={{ x: "2000", y: "100%" }}
        request={async params => {
          const { data } = await getRolePageApi(params);
          return formatDataForProTable<RoleModel>(data);
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
        headerTitle="角色管理"
        toolBarRender={toolBarRender}
      />
      <FormModal ref={formModalRef} onConfirm={handleFormModalConfirm} />
      <DataScopeFormModal ref={dataScopeFormModalRef} onConfirm={handleDataScopeFormModalConfirm} />
    </>
  );
};

export default Role;
