import { getDictOptions, getDictsApi } from "@/api/admin/sys/sys-dictdata";
import { getUserConfPageApi, UserConfModel } from "@/api/app/user/user-conf";
import HocAuth from "@/components/HocAuth";
import LoadingButton from "@/components/LoadingButton";
import { pagination } from "@/config/proTable";
import { ResultEnum } from "@/enums/httpEnum";
import { message } from "@/hooks/useMessage";
import { formatDataForProTable } from "@/utils";
import { EditOutlined } from "@ant-design/icons";
import type { ActionType, ProColumns, ProFormInstance } from "@ant-design/pro-components";
import { ProTable } from "@ant-design/pro-components";
import { Space } from "antd";
import React, { useEffect, useRef, useState } from "react";
import FormModal, { FormModalRef } from "./components/FormModal";

const UserConf: React.FC = () => {
  const actionRef = React.useRef<ActionType>();
  const tableFormRef = React.useRef<ProFormInstance>();
  const formModalRef = useRef<FormModalRef>(null);
  const [canLoginOptions, setCanLoginOptions] = useState<Map<string, string>>(new Map());

  // 定义列
  const columns: ProColumns<UserConfModel>[] = [
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
      title: "配置编号",
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
      title: "是否允许登录",
      dataIndex: "canLogin",
      valueType: "select",
      valueEnum: canLoginOptions,
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
    },
    {
      title: "操作",
      valueType: "option",
      align: "center",
      fixed: "right",
      width: 150,
      render: (_, data) => (
        <Space>
          <HocAuth permission={["app:user-conf:edit"]}>
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
        </Space>
      )
    }
  ];
  useEffect(() => {
    const initData = async () => {
      const { data: canLoginData, msg: canLoginMsg, code: canLoginCode } = await getDictsApi("admin_sys_yes_no");
      if (canLoginCode !== ResultEnum.SUCCESS) {
        message.error(canLoginMsg);
        return;
      }
      setCanLoginOptions(getDictOptions(canLoginData));
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

  return (
    <>
      <ProTable<UserConfModel>
        className="ant-pro-table-scroll"
        columns={columns}
        actionRef={actionRef}
        formRef={tableFormRef}
        bordered
        cardBordered
        defaultSize="small"
        scroll={{ x: "2000", y: "100%" }}
        request={async params => {
          const { data } = await getUserConfPageApi(params);
          return formatDataForProTable<UserConfModel>(data);
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
        headerTitle="用户配置"
      />
      <FormModal ref={formModalRef} onConfirm={handleFormModalConfirm} />
    </>
  );
};

export default UserConf;
