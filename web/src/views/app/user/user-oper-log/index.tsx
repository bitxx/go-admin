import { getDictOptions, getDictsApi } from "@/api/admin/sys/sys-dictdata";
import { exportUserOperLogApi, getUserOperLogPageApi, UserOperLogModel } from "@/api/app/user/user-oper-log";
import HocAuth from "@/components/HocAuth";
import LoadingButton from "@/components/LoadingButton";
import { pagination } from "@/config/proTable";
import { ResultEnum } from "@/enums/httpEnum";
import { modal } from "@/hooks/useMessage";
import { formatDataForProTable, saveExcelBlob } from "@/utils";
import { CloudDownloadOutlined, ExclamationCircleOutlined } from "@ant-design/icons";
import type { ActionType, ProColumns, ProFormInstance } from "@ant-design/pro-components";
import { ProTable } from "@ant-design/pro-components";
import { message } from "antd";
import React, { useEffect, useState } from "react";

const UserOperLog: React.FC = () => {
  const actionRef = React.useRef<ActionType>();
  const tableFormRef = React.useRef<ProFormInstance>();
  const [actionTypeOptions, setActionTypeOptions] = useState<Map<string, string>>(new Map());
  const [byTypeOptions, setByTypeOptions] = useState<Map<string, string>>(new Map());

  // 定义列
  const columns: ProColumns<UserOperLogModel>[] = [
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
      title: "日志编码",
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
      title: "用户名",
      dataIndex: "userName",
      width: 80,
      align: "left",
      render: (text, record) => record.user?.userName
    },
    {
      title: "用户邮箱",
      dataIndex: "email",
      width: 160,
      align: "left",
      render: (text, record) => record.user?.email
    },
    {
      title: "用户手机号",
      dataIndex: "mobile",
      width: 120,
      align: "left",
      render: (text, record) => record.user?.mobile
    },
    {
      title: "用户行为类型",
      dataIndex: "actionType",
      valueType: "select",
      valueEnum: actionTypeOptions,
      fieldProps: {
        getPopupContainer: (triggerNode: { parentElement: any }) => triggerNode.parentElement || document.body
      },
      width: 120,
      align: "left"
    },
    {
      title: "更新用户类型",
      dataIndex: "byType",
      valueType: "select",
      valueEnum: byTypeOptions,
      fieldProps: {
        getPopupContainer: (triggerNode: { parentElement: any }) => triggerNode.parentElement || document.body
      },
      width: 120,
      align: "left"
    },
    {
      title: "创建者",
      dataIndex: "createBy",
      hideInSearch: true,
      width: 80,
      align: "left"
    },
    {
      title: "更新者",
      dataIndex: "updateBy",
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
      search: { transform: value => ({ beginCreatedAt: value[0], endCreatedAt: value[1] }) },
      fieldProps: {
        getPopupContainer: (triggerNode: { parentElement: any }) => triggerNode.parentElement || document.body // 确保弹出框在合适的容器中
      }
    },
    {
      title: "更新时间",
      dataIndex: "updatedAt",
      hideInSearch: true,
      valueType: "dateTime",
      width: 180,
      align: "left"
    }
  ];
  useEffect(() => {
    const initData = async () => {
      const { data: actionTypeData, msg: actionTypeMsg, code: actionTypeCode } = await getDictsApi("app_user_action_type");
      if (actionTypeCode !== ResultEnum.SUCCESS) {
        message.error(actionTypeMsg);
        return;
      }
      setActionTypeOptions(getDictOptions(actionTypeData));
      const { data: byTypeData, msg: byTypeMsg, code: byTypeCode } = await getDictsApi("app_user_by_type");
      if (byTypeCode !== ResultEnum.SUCCESS) {
        message.error(byTypeMsg);
        return;
      }
      setByTypeOptions(getDictOptions(byTypeData));
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
          saveExcelBlob("用户关键行为日志表", await exportUserOperLogApi(tableFormRef.current?.getFieldsValue()));
        } catch (err) {
          message.error("下载失败，请检查网络");
        } finally {
          done();
        }
      }
    });
  };

  const toolBarRender = () => [
    <HocAuth permission={["app:user-oper-log:export"]}>
      <LoadingButton type="primary" key="importTable" icon={<CloudDownloadOutlined />} onClick={done => handleExport(done)}>
        Excel导出
      </LoadingButton>
    </HocAuth>
  ];

  return (
    <ProTable<UserOperLogModel>
      className="ant-pro-table-scroll"
      columns={columns}
      actionRef={actionRef}
      formRef={tableFormRef}
      bordered
      cardBordered
      defaultSize="small"
      scroll={{ x: "2000", y: "100%" }}
      request={async params => {
        const { data } = await getUserOperLogPageApi(params);
        return formatDataForProTable<UserOperLogModel>(data);
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
      headerTitle="用户关键行为日志表"
      toolBarRender={toolBarRender}
    />
  );
};

export default UserOperLog;
