import { delOperLogApi, exportOperLogApi, getOperLogPageApi, OperLogModel } from "@/api/admin/sys/sys-operlog";
import HocAuth from "@/components/HocAuth";
import LoadingButton from "@/components/LoadingButton";
import { pagination } from "@/config/proTable";
import { ResultEnum } from "@/enums/httpEnum";
import { message, modal } from "@/hooks/useMessage";
import { formatDataForProTable, saveExcelBlob } from "@/utils";
import { CloudDownloadOutlined, DeleteOutlined, EditOutlined, ExclamationCircleOutlined } from "@ant-design/icons";
import type { ActionType, ProColumns, ProFormInstance } from "@ant-design/pro-components";
import { ProTable } from "@ant-design/pro-components";
import { Space } from "antd";
import React, { useRef } from "react";
import FormModal, { FormModalRef } from "./components/FormModal";

const OperLog: React.FC = () => {
  const actionRef = React.useRef<ActionType>();
  const tableFormRef = React.useRef<ProFormInstance>();
  const formModalRef = useRef<FormModalRef>(null);

  // 定义列
  const columns: ProColumns<OperLogModel>[] = [
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
      title: "日志编号",
      dataIndex: "id",
      hideInSearch: true,
      width: 80,
      align: "left"
    },
    {
      title: "用户编号",
      dataIndex: "userId",
      width: 80,
      align: "left"
    },
    {
      title: "请求方法",
      dataIndex: "requestMethod",
      hideInSearch: true,
      width: 80,
      align: "left"
    },
    {
      title: "请求地址",
      dataIndex: "operUrl",
      hideInSearch: true,
      width: 500,
      align: "left"
    },
    {
      title: "请求ip",
      dataIndex: "operIp",
      hideInSearch: true,
      width: 120,
      align: "left"
    },
    {
      title: "访问位置",
      dataIndex: "operLocation",
      hideInSearch: true,
      width: 100,
      align: "left"
    },
    {
      title: "返回码",
      dataIndex: "status",
      hideInSearch: true,
      width: 80,
      align: "left"
    },
    {
      title: "耗时",
      dataIndex: "latencyTime",
      hideInSearch: true,
      width: 120,
      align: "left"
    },
    {
      title: "操作时间",
      dataIndex: "operTime",
      hideInSearch: true,
      valueType: "dateTime",
      width: 180,
      align: "left"
    },
    {
      title: "用户代理",
      dataIndex: "userAgent",
      hideInSearch: true,
      width: 500,
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
      search: { transform: value => ({ beginCreatedAt: value[0], endCreatedAt: value[1] }) }
    },
    {
      title: "操作",
      valueType: "option",
      align: "center",
      fixed: "right",
      width: 150,
      render: (_, data) => (
        <Space>
          <LoadingButton
            key="edit"
            type="link"
            size="small"
            icon={<EditOutlined />}
            onClick={done => handleShowEditFormModal(data.id!, done)}
          >
            详情
          </LoadingButton>
          <HocAuth permission={["admin:sys-oper-log:query"]}>
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

  const handleShowEditFormModal = (id: number, done: () => void) => {
    formModalRef.current?.showEditFormModal(id);
    setTimeout(() => done(), 1000);
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
          saveExcelBlob("操作日志", await exportOperLogApi(tableFormRef.current?.getFieldsValue()));
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
      content: "是否确认删除编号为 " + id + " 的数据项?",
      okText: "确认",
      cancelText: "取消",
      maskClosable: true,
      onCancel: () => {
        done();
      },
      onOk: async () => {
        try {
          const { code, msg } = await delOperLogApi([id!]);
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
    <LoadingButton type="primary" key="importTable" icon={<CloudDownloadOutlined />} onClick={done => handleExport(done)}>
      Excel导出
    </LoadingButton>
  ];

  return (
    <>
      <ProTable<OperLogModel>
        className="ant-pro-table-scroll"
        columns={columns}
        actionRef={actionRef}
        formRef={tableFormRef}
        bordered
        cardBordered
        defaultSize="small"
        scroll={{ x: "2000", y: "100%" }}
        request={async params => {
          const { data } = await getOperLogPageApi(params);
          return formatDataForProTable<OperLogModel>(data);
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
        headerTitle="操作日志"
        toolBarRender={toolBarRender}
      />
      <FormModal ref={formModalRef} />
    </>
  );
};

export default OperLog;
