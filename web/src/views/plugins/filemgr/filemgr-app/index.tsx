import { getDictOptions, getDictsApi } from "@/api/admin/sys/sys-dictdata";
import { delFilemgrAppApi, exportFilemgrAppApi, FilemgrAppModel, getFilemgrAppPageApi } from "@/api/plugins/filemgr/filemgr-app";
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

const FilemgrApp: React.FC = () => {
  const actionRef = React.useRef<ActionType>();
  const tableFormRef = React.useRef<ProFormInstance>();
  const formModalRef = useRef<FormModalRef>(null);
  const [platformOptions, setPlatformOptions] = useState<Map<string, string>>(new Map());
  const [appTypeOptions, setAppTypeOptions] = useState<Map<string, string>>(new Map());
  const [downloadTypeOptions, setDownloadTypeOptions] = useState<Map<string, string>>(new Map());
  const [statusOptions, setStatusOptions] = useState<Map<string, string>>(new Map());

  // 定义列
  const columns: ProColumns<FilemgrAppModel>[] = [
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
      title: "App编号",
      dataIndex: "id",
      hideInSearch: true,
      width: 80,
      align: "left"
    },
    {
      title: "版本号",
      dataIndex: "version",
      width: 80,
      align: "left"
    },
    {
      title: "系统平台",
      dataIndex: "platform",
      valueType: "select",
      valueEnum: platformOptions,
      width: 80,
      align: "left"
    },
    {
      title: "版本类型",
      dataIndex: "appType",
      valueType: "select",
      valueEnum: appTypeOptions,
      width: 80,
      align: "left"
    },
    {
      title: "本地地址",
      dataIndex: "localAddress",
      hideInSearch: true,
      width: 300,
      align: "left"
    },
    {
      title: "下载方式",
      dataIndex: "downloadType",
      valueType: "select",
      valueEnum: downloadTypeOptions,
      width: 80,
      align: "left"
    },
    {
      title: "下载地址",
      dataIndex: "downloadUrl",
      hideInSearch: true,
      width: 300,
      align: "left"
    },
    {
      title: "备注信息",
      dataIndex: "remark",
      hideInSearch: true,
      width: 80,
      align: "left"
    },
    {
      title: "发布状态",
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
    },
    {
      title: "操作",
      valueType: "option",
      align: "center",
      fixed: "right",
      width: 150,
      render: (_, data) => (
        <Space>
          <HocAuth permission={["plugins:filemgr-app:edit"]}>
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
          <HocAuth permission={["plugins:filemgr-app:del"]}>
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
      const { data: platformData, msg: platformMsg, code: platformCode } = await getDictsApi("plugin_filemgr_app_platform");
      if (platformCode !== ResultEnum.SUCCESS) {
        message.error(platformMsg);
        return;
      }
      setPlatformOptions(getDictOptions(platformData));
      const { data: appTypeData, msg: appTypeMsg, code: appTypeCode } = await getDictsApi("plugin_filemgr_app_type");
      if (appTypeCode !== ResultEnum.SUCCESS) {
        message.error(appTypeMsg);
        return;
      }
      setAppTypeOptions(getDictOptions(appTypeData));
      const {
        data: downloadTypeData,
        msg: downloadTypeMsg,
        code: downloadTypeCode
      } = await getDictsApi("plugin_filemgr_app_download_type");
      if (downloadTypeCode !== ResultEnum.SUCCESS) {
        message.error(downloadTypeMsg);
        return;
      }
      setDownloadTypeOptions(getDictOptions(downloadTypeData));
      const { data: statusData, msg: statusMsg, code: statusCode } = await getDictsApi("plugin_filemgr_publish_status");
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

  const handleFormModalConfirm = () => {
    actionRef.current?.reload(true);
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
          saveExcelBlob("App管理", await exportFilemgrAppApi(tableFormRef.current?.getFieldsValue()));
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
          const { code, msg } = await delFilemgrAppApi([id!]);
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
    <HocAuth permission={["plugins:filemgr-app:add"]}>
      <LoadingButton type="primary" key="addTable" icon={<PlusCircleOutlined />} onClick={done => handleShowAddFormModal(done)}>
        新增
      </LoadingButton>
    </HocAuth>,
    <HocAuth permission={["plugins:filemgr-app:export"]}>
      <LoadingButton type="primary" key="importTable" icon={<CloudDownloadOutlined />} onClick={done => handleExport(done)}>
        Excel导出
      </LoadingButton>
    </HocAuth>
  ];

  return (
    <>
      <ProTable<FilemgrAppModel>
        className="ant-pro-table-scroll"
        columns={columns}
        actionRef={actionRef}
        formRef={tableFormRef}
        bordered
        cardBordered
        defaultSize="small"
        scroll={{ x: "2000", y: "100%" }}
        request={async params => {
          const { data } = await getFilemgrAppPageApi(params);
          return formatDataForProTable<FilemgrAppModel>(data);
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
        headerTitle="App管理"
        toolBarRender={toolBarRender}
      />
      <FormModal ref={formModalRef} onConfirm={handleFormModalConfirm} />
    </>
  );
};

export default FilemgrApp;
