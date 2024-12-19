import { getDictOptions, getDictsApi } from "@/api/admin/sys/sys-dictdata";
import { exportMsgCodeApi, getMsgCodePageApi, MsgCodeModel } from "@/api/plugins/msg/msg-code";
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

const MsgCode: React.FC = () => {
  const actionRef = React.useRef<ActionType>();
  const tableFormRef = React.useRef<ProFormInstance>();
  const [codeTypeOptions, setCodeTypeOptions] = useState<Map<string, string>>(new Map());
  const [statusOptions, setStatusOptions] = useState<Map<string, string>>(new Map());

  // 定义列
  const columns: ProColumns<MsgCodeModel>[] = [
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
      title: "验证码编号",
      dataIndex: "id",
      width: 80,
      align: "left",
      hideInSearch: true
    },
    {
      title: "用户编号",
      dataIndex: "userId",
      width: 80,
      align: "left"
    },
    {
      title: "验证码",
      dataIndex: "code",
      hideInSearch: true,
      width: 80,
      align: "left"
    },
    {
      title: "验证码类型",
      dataIndex: "codeType",
      valueType: "select",
      valueEnum: codeTypeOptions,
      width: 120,
      align: "left"
    },
    {
      title: "备注异常",
      dataIndex: "remark",
      hideInSearch: true,
      width: 80,
      align: "left"
    },
    {
      title: "验证码状态",
      dataIndex: "status",
      valueType: "select",
      valueEnum: statusOptions,
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
      const { data: codeTypeData, msg: codeTypeMsg, code: codeTypeCode } = await getDictsApi("plugin_msg_code_type");
      if (codeTypeCode !== ResultEnum.SUCCESS) {
        message.error(codeTypeMsg);
        return;
      }
      setCodeTypeOptions(getDictOptions(codeTypeData));
      const { data: statusData, msg: statusMsg, code: statusCode } = await getDictsApi("plugin_msg_sendstatus");
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
          saveExcelBlob("验证码记录", await exportMsgCodeApi(tableFormRef.current?.getFieldsValue()));
        } catch (err) {
          message.error("下载失败，请检查网络");
        } finally {
          done();
        }
      }
    });
  };

  const toolBarRender = () => [
    <HocAuth permission={["plugins:msg-code:export"]}>
      <LoadingButton type="primary" key="importTable" icon={<CloudDownloadOutlined />} onClick={done => handleExport(done)}>
        Excel导出
      </LoadingButton>
    </HocAuth>
  ];

  return (
    <>
      <ProTable<MsgCodeModel>
        className="ant-pro-table-scroll"
        columns={columns}
        actionRef={actionRef}
        formRef={tableFormRef}
        bordered
        cardBordered
        defaultSize="small"
        scroll={{ x: "2000", y: "100%" }}
        request={async params => {
          const { data } = await getMsgCodePageApi(params);
          return formatDataForProTable<MsgCodeModel>(data);
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
        headerTitle="验证码记录"
        toolBarRender={toolBarRender}
      />
    </>
  );
};

export default MsgCode;
