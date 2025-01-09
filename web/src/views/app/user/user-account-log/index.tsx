import { getDictOptions, getDictsApi } from "@/api/admin/sys/sys-dictdata";
import { exportUserAccountLogApi, getUserAccountLogPageApi, UserAccountLogModel } from "@/api/app/user/user-account-log";
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

const UserAccountLog: React.FC = () => {
  const actionRef = React.useRef<ActionType>();
  const tableFormRef = React.useRef<ProFormInstance>();
  const [moneyTypeOptions, setMoneyTypeOptions] = useState<Map<string, string>>(new Map());
  const [changeTypeOptions, setChangeTypeOptions] = useState<Map<string, string>>(new Map());

  const columns: ProColumns<UserAccountLogModel>[] = [
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
      title: "账变编号",
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
      width: 180,
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
      title: "账变金额",
      dataIndex: "changeMoney",
      hideInSearch: true,
      width: 150,
      align: "left"
    },
    {
      title: "账变前金额",
      dataIndex: "beforeMoney",
      hideInSearch: true,
      width: 150,
      align: "left"
    },
    {
      title: "账变后金额",
      dataIndex: "afterMoney",
      hideInSearch: true,
      width: 150,
      align: "left"
    },
    {
      title: "金额类型",
      dataIndex: "moneyType",
      valueType: "select",
      valueEnum: moneyTypeOptions,
      width: 120,
      align: "left"
    },
    {
      title: "帐变类型",
      dataIndex: "changeType",
      valueType: "select",
      valueEnum: changeTypeOptions,
      width: 120,
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
    }
  ];

  useEffect(() => {
    const initData = async () => {
      const { data: moneyTypeData, msg: moneyTypeMsg, code: moneyTypeCode } = await getDictsApi("app_money_type");
      if (moneyTypeCode !== ResultEnum.SUCCESS) {
        message.error(moneyTypeMsg);
        return;
      }
      setMoneyTypeOptions(getDictOptions(moneyTypeData));
      const { data: changeTypeData, msg: changeTypeMsg, code: changeTypeCode } = await getDictsApi("app_account_change_type");
      if (changeTypeCode !== ResultEnum.SUCCESS) {
        message.error(changeTypeMsg);
        return;
      }
      setChangeTypeOptions(getDictOptions(changeTypeData));
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
          saveExcelBlob("账变记录", await exportUserAccountLogApi(tableFormRef.current?.getFieldsValue()));
        } catch (err) {
          message.error("下载失败，请检查网络");
        } finally {
          done();
        }
      }
    });
  };

  const toolBarRender = () => [
    <HocAuth permission={["app:user-account-log:export"]}>
      <LoadingButton type="primary" key="importTable" icon={<CloudDownloadOutlined />} onClick={done => handleExport(done)}>
        Excel导出
      </LoadingButton>
    </HocAuth>
  ];

  return (
    <>
      <ProTable<UserAccountLogModel>
        className="ant-pro-table-scroll"
        columns={columns}
        actionRef={actionRef}
        formRef={tableFormRef}
        bordered
        cardBordered
        defaultSize="small"
        scroll={{ x: "2000", y: "100%" }}
        request={async params => {
          const { data } = await getUserAccountLogPageApi(params);
          return formatDataForProTable<UserAccountLogModel>(data);
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
        headerTitle="账变记录"
        toolBarRender={toolBarRender}
      />
    </>
  );
};

export default UserAccountLog;
