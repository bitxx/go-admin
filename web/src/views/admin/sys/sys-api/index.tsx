import { ApiModel, delApiApi, exportApiApi, getApiPageApi, syncApiApi } from "@/api/admin/sys/sys-api";
import { getDictOptions, getDictsApi } from "@/api/admin/sys/sys-dictdata";
import HocAuth from "@/components/HocAuth";
import LoadingButton from "@/components/LoadingButton";
import { pagination } from "@/config/proTable";
import { ResultEnum } from "@/enums/httpEnum";
import { message, modal } from "@/hooks/useMessage";
import { formatDataForProTable, saveExcelBlob } from "@/utils";
import { CloudDownloadOutlined, DeleteOutlined, EditOutlined, ExclamationCircleOutlined, SyncOutlined } from "@ant-design/icons";
import type { ActionType, ProColumns, ProFormInstance } from "@ant-design/pro-components";
import { ProTable } from "@ant-design/pro-components";
import { Button, Popover, Space, Table } from "antd";
import React, { useEffect, useRef, useState } from "react";
import FormModal, { FormModalRef } from "./components/FormModal";

const Api: React.FC = () => {
  const actionRef = React.useRef<ActionType>();
  const tableFormRef = React.useRef<ProFormInstance>();
  const formModalRef = useRef<FormModalRef>(null);
  const [methodOptions, setMethodOptions] = useState<Map<string, string>>(new Map());
  const [apiTypeOptions, setApiTypeOptions] = useState<Map<string, string>>(new Map());
  const [menuTypeOptions, setMenuTypeOptions] = useState<Map<string, string>>(new Map());

  const popoverColumns = [
    {
      title: "菜单编号",
      dataIndex: "id",
      key: "id"
    },
    {
      title: "菜单名称",
      dataIndex: "title",
      key: "title"
    },
    {
      title: "菜单类型",
      dataIndex: "menuType",
      key: "menuType",
      render: (menuType: string) => menuTypeOptions.get(menuType) || "未知类型"
    }
  ];

  const columns: ProColumns<ApiModel>[] = [
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
      title: "接口编号",
      dataIndex: "id",
      hideInSearch: true,
      width: 80,
      align: "left"
    },
    {
      title: "功能描述",
      dataIndex: "description",
      width: 300,
      align: "left"
    },
    {
      title: "接口地址",
      dataIndex: "path",
      width: 400,
      align: "left"
    },
    {
      title: "接口请求方法",
      dataIndex: "method",
      valueType: "select",
      valueEnum: methodOptions,
      fieldProps: {
        getPopupContainer: (triggerNode: { parentElement: any }) => triggerNode.parentElement || document.body
      },
      width: 110,
      align: "center"
    },
    {
      title: "绑定菜单/按钮数量",
      dataIndex: "sysMenu",
      width: 150,
      align: "center",
      hideInSearch: true,
      render: (text, record) => (
        <Popover
          content={
            <Table
              columns={popoverColumns}
              dataSource={record.sysMenu?.map(item => ({
                key: item.id,
                ...item
              }))}
              pagination={false}
              size="small"
            />
          }
        >
          <Button type="link" style={{ padding: 0 }}>
            {record.sysMenu?.length}
          </Button>
        </Popover>
      )
    },
    {
      title: "接口类型",
      dataIndex: "apiType",
      valueType: "select",
      valueEnum: apiTypeOptions,
      fieldProps: {
        getPopupContainer: (triggerNode: { parentElement: any }) => triggerNode.parentElement || document.body
      },
      width: 80,
      align: "center"
    },
    {
      title: "备注",
      dataIndex: "remark",
      width: 300,
      align: "left"
    },
    {
      title: "更新时间",
      dataIndex: "updatedAt",
      hideInSearch: true,
      valueType: "dateTime",
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
      title: "创建时间",
      dataIndex: "createdAt",
      valueType: "dateTimeRange",
      hideInTable: true,
      search: { transform: value => ({ beginCreatedAt: value[0], endCreatedAt: value[1] }) },
      fieldProps: {
        getPopupContainer: (triggerNode: { parentElement: any }) => triggerNode.parentElement || document.body // 确保弹出框在合适的容器中
      }
    },
    {
      title: "操作",
      valueType: "option",
      align: "center",
      fixed: "right",
      width: 150,
      render: (_, data) => (
        <Space>
          <HocAuth permission={["admin:sys-api:edit"]}>
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
          <HocAuth permission={["admin:sys-api:del"]}>
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
      const { data: methodData, msg: methodMsg, code: methodCode } = await getDictsApi("admin_sys_api_method");
      if (methodCode !== ResultEnum.SUCCESS) {
        message.error(methodMsg);
        return;
      }
      setMethodOptions(getDictOptions(methodData));
      const { data: apiTypeData, msg: apiTypeMsg, code: apiTypeCode } = await getDictsApi("admin_sys_config_type");
      if (apiTypeCode !== ResultEnum.SUCCESS) {
        message.error(apiTypeMsg);
        return;
      }
      setApiTypeOptions(getDictOptions(apiTypeData));
      const { data: menuTypeData, msg: menuTypeMsg, code: menuTypeCode } = await getDictsApi("admin_sys_menu_type");
      if (menuTypeCode !== ResultEnum.SUCCESS) {
        message.error(menuTypeMsg);
        return;
      }
      setMenuTypeOptions(getDictOptions(menuTypeData));
    };
    initData();
  }, []);

  const handleShowEditFormModal = (id: number, done: () => void) => {
    formModalRef.current?.showEditFormModal(id);
    setTimeout(() => done(), 1000);
  };

  const handleFormModalConfirm = () => {
    actionRef.current?.reload(false);
  };

  const handleExport = (done: () => void) => {
    modal.confirm({
      title: "提示",
      icon: <ExclamationCircleOutlined />,
      content: "是否确认导出所选数据？",
      okText: "确认",
      cancelText: "取消",
      maskClosable: true,
      onCancel: () => {
        done();
      },
      onOk: async () => {
        try {
          saveExcelBlob("Api", await exportApiApi(tableFormRef.current?.getFieldsValue()));
        } catch (err) {
          message.error("下载失败，请检查网络");
        } finally {
          done();
        }
      }
    });
  };

  const handleDelete = (id: number, done: () => void) => {
    modal.confirm({
      title: "提示",
      icon: <ExclamationCircleOutlined />,
      content: "若该接口与菜单有关联，删除接口的同时，将会取消该关联。是否确认删除编号为 " + id + " 的接口?",
      okText: "确认",
      cancelText: "取消",
      maskClosable: true,
      onCancel: () => {
        done();
      },
      onOk: async () => {
        try {
          const { code, msg } = await delApiApi([id!]);
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

  const handleSync = (done: () => void) => {
    modal.confirm({
      title: "提示",
      icon: <ExclamationCircleOutlined />,
      content:
        "开始同步一段时间后，再次回到该页面，即可看到新的数据。本次同步会删除库中的无效路由，同时会新增检测到的新路由。是否开始同步数据？",
      okText: "确认",
      cancelText: "取消",
      maskClosable: true,
      onCancel: () => {
        done();
      },
      onOk: async () => {
        try {
          const { code, msg } = await syncApiApi();
          if (code !== ResultEnum.SUCCESS) {
            message.error(msg);
            return;
          }
          message.success("同步完成！");
          actionRef.current?.reload(false);
        } finally {
          done();
        }
      }
    });
  };

  const toolBarRender = () => [
    <>
      <HocAuth permission={["admin:sys-api:export"]}>
        <LoadingButton type="primary" key="importTable" icon={<CloudDownloadOutlined />} onClick={done => handleExport(done)}>
          Excel导出
        </LoadingButton>
      </HocAuth>
      <HocAuth permission={["admin:sys-api:sync"]}>
        <LoadingButton type="primary" key="syncApi" icon={<SyncOutlined />} onClick={done => handleSync(done)}>
          接口数据同步
        </LoadingButton>
      </HocAuth>
    </>
  ];

  return (
    <>
      <ProTable<ApiModel>
        className="ant-pro-table-scroll"
        columns={columns}
        actionRef={actionRef}
        formRef={tableFormRef}
        bordered
        cardBordered
        defaultSize="small"
        scroll={{ x: "2000", y: "100%" }}
        request={async params => {
          const { data } = await getApiPageApi(params);
          return formatDataForProTable<ApiModel>(data);
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
        headerTitle="接口管理"
        toolBarRender={toolBarRender}
      />
      <FormModal ref={formModalRef} onConfirm={handleFormModalConfirm} />
    </>
  );
};

export default Api;
