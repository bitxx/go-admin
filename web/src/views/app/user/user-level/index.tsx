import { getDictOptions, getDictsApi } from "@/api/admin/sys/sys-dictdata";
import { delUserLevelApi, exportUserLevelApi, getUserLevelPageApi, UserLevelModel } from "@/api/app/user/user-level";
import HocAuth from "@/components/HocAuth";
import LoadingButton from "@/components/LoadingButton";
import { pagination } from "@/config/proTable";
import { ResultEnum } from "@/enums/httpEnum";
import { message, modal } from "@/hooks/useMessage";
import { formatDataForProTable, saveExcelBlob } from "@/utils";
import {
  CloudDownloadOutlined,
  DeleteOutlined,
  EditOutlined,
  ExclamationCircleOutlined,
  PlusCircleOutlined
} from "@ant-design/icons";
import type { ActionType, ProColumns, ProFormInstance } from "@ant-design/pro-components";
import { ProTable } from "@ant-design/pro-components";
import { Space } from "antd";
import React, { useEffect, useRef, useState } from "react";
import FormModal, { FormModalRef } from "./components/FormModal";

const UserLevel: React.FC = () => {
  const actionRef = React.useRef<ActionType>();
  const tableFormRef = React.useRef<ProFormInstance>();
  const formModalRef = useRef<FormModalRef>(null);
  const [levelTypeOptions, setLevelTypeOptions] = useState<Map<string, string>>(new Map());

  // 定义列
  const columns: ProColumns<UserLevelModel>[] = [
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
      title: "等级编号",
      dataIndex: "id",
      hideInSearch: true,
      width: 80,
      align: "left"
    },
    {
      title: "等级名称",
      dataIndex: "name",
      width: 80,
      align: "left"
    },
    {
      title: "等级类型",
      dataIndex: "levelType",
      valueType: "select",
      valueEnum: levelTypeOptions,
      fieldProps: {
        getPopupContainer: (triggerNode: { parentElement: any }) => triggerNode.parentElement || document.body
      },
      width: 120,
      align: "left"
    },
    {
      title: "等级",
      dataIndex: "level",
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
      title: "操作",
      valueType: "option",
      align: "center",
      fixed: "right",
      width: 150,
      render: (_, data) => (
        <Space>
          <HocAuth permission={["app:user-level:edit"]}>
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
          <HocAuth permission={["app:user-level:del"]}>
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
      const { data: levelTypeData, msg: levelTypeMsg, code: levelTypeCode } = await getDictsApi("app_user_level_type");
      if (levelTypeCode !== ResultEnum.SUCCESS) {
        message.error(levelTypeMsg);
        return;
      }
      setLevelTypeOptions(getDictOptions(levelTypeData));
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
          saveExcelBlob("用户等级", await exportUserLevelApi(tableFormRef.current?.getFieldsValue()));
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
          const { code, msg } = await delUserLevelApi([id!]);
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
    <HocAuth permission={["app:user-level:add"]}>
      <LoadingButton type="primary" key="addTable" icon={<PlusCircleOutlined />} onClick={done => handleShowAddFormModal(done)}>
        新增
      </LoadingButton>
    </HocAuth>,
    <HocAuth permission={["app:user-level:export"]}>
      <LoadingButton type="primary" key="importTable" icon={<CloudDownloadOutlined />} onClick={done => handleExport(done)}>
        Excel导出
      </LoadingButton>
    </HocAuth>
  ];

  return (
    <>
      <ProTable<UserLevelModel>
        className="ant-pro-table-scroll"
        columns={columns}
        actionRef={actionRef}
        formRef={tableFormRef}
        bordered
        cardBordered
        defaultSize="small"
        scroll={{ x: "2000", y: "100%" }}
        request={async params => {
          const { data } = await getUserLevelPageApi(params);
          return formatDataForProTable<UserLevelModel>(data);
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
        headerTitle="用户等级"
        toolBarRender={toolBarRender}
      />
      <FormModal ref={formModalRef} onConfirm={handleFormModalConfirm} />
    </>
  );
};

export default UserLevel;
