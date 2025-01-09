import { getDictOptions, getDictsApi } from "@/api/admin/sys/sys-dictdata";
import { delLoginLogApi, exportLoginLogApi, getLoginLogPageApi, LoginLogModel } from "@/api/admin/sys/sys-loginlog";
import HocAuth from "@/components/HocAuth";
import LoadingButton from "@/components/LoadingButton";
import { pagination } from "@/config/proTable";
import { ResultEnum } from "@/enums/httpEnum";
import { message, modal } from "@/hooks/useMessage";
import { formatDataForProTable, saveExcelBlob } from "@/utils";
import { CloudDownloadOutlined, DeleteOutlined, ExclamationCircleOutlined } from "@ant-design/icons";
import type { ActionType, ProColumns, ProFormInstance } from "@ant-design/pro-components";
import { ProTable } from "@ant-design/pro-components";
import { Space } from "antd";
import React, { useEffect, useState } from "react";

const LoginLog: React.FC = () => {
  const actionRef = React.useRef<ActionType>();
  const tableFormRef = React.useRef<ProFormInstance>();
  const [statusOptions, setStatusOptions] = useState<Map<string, string>>(new Map());

  // 定义列
  const columns: ProColumns<LoginLogModel>[] = [
    {
      title: "序号",
      dataIndex: "index",
      valueType: "index",
      width: 50,
      align: "center",
      className: "gray-cell",
      render: (_, __, index, action) => {
        const currentPage = action?.pageInfo?.current || 1;
        const pageSize = action?.pageInfo?.pageSize || 10;
        return (currentPage - 1) * pageSize + index + 1;
      }
    },
    {
      title: "登录编号",
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
      title: "登录状态",
      dataIndex: "status",
      valueType: "select",
      valueEnum: statusOptions,
      width: 120,
      align: "left"
    },
    {
      title: "ip地址",
      dataIndex: "ipaddr",
      width: 80,
      align: "left"
    },
    {
      title: "归属地",
      dataIndex: "loginLocation",
      hideInSearch: true,
      width: 80,
      align: "left"
    },
    {
      title: "浏览器",
      dataIndex: "browser",
      hideInSearch: true,
      width: 200,
      align: "left"
    },
    {
      title: "代理",
      dataIndex: "agent",
      hideInSearch: true,
      width: 500,
      align: "left"
    },
    {
      title: "系统",
      dataIndex: "os",
      hideInSearch: true,
      width: 200,
      align: "left"
    },
    {
      title: "固件",
      dataIndex: "platform",
      hideInSearch: true,
      width: 120,
      align: "left"
    },
    {
      title: "登录时间",
      dataIndex: "loginTime",
      hideInSearch: true,
      valueType: "dateTime",
      width: 180,
      align: "left"
    },
    {
      title: "备注",
      dataIndex: "remark",
      hideInSearch: true,
      width: 80,
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
          <HocAuth permission={["admin:sys-login-log:del"]}>
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
      const { data: statusData, msg: statusMsg, code: statusCode } = await getDictsApi("admin_sys_loginlog_status");
      if (statusCode !== ResultEnum.SUCCESS) {
        message.error(statusMsg);
        return;
      }
      setStatusOptions(getDictOptions(statusData));
    };
    initData();
  }, []);

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
          saveExcelBlob("LoginLog", await exportLoginLogApi(tableFormRef.current?.getFieldsValue()));
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
          const { code, msg } = await delLoginLogApi([id!]);
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
    <HocAuth permission={["admin:sys-login-log:export"]}>
      <LoadingButton type="primary" key="importTable" icon={<CloudDownloadOutlined />} onClick={done => handleExport(done)}>
        Excel导出
      </LoadingButton>
    </HocAuth>
  ];

  return (
    <>
      <ProTable<LoginLogModel>
        className="ant-pro-table-scroll"
        columns={columns}
        actionRef={actionRef}
        formRef={tableFormRef}
        bordered
        cardBordered
        defaultSize="small"
        scroll={{ x: "2000", y: "100%" }}
        request={async params => {
          const { data } = await getLoginLogPageApi(params);
          return formatDataForProTable<LoginLogModel>(data);
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
        headerTitle="登录日志"
        toolBarRender={toolBarRender}
      />
    </>
  );
};

export default LoginLog;
