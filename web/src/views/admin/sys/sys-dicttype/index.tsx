import { delDictTypeApi, DictTypeModel, exportDictTypeApi, getDictTypePageApi } from "@/api/admin/sys/sys-dicttype";
import HocAuth from "@/components/HocAuth";
import LoadingButton from "@/components/LoadingButton";
import { DICT_DATA_URL } from "@/config";
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
import { Button, Space, Tooltip } from "antd";
import React, { useRef } from "react";
import { useAliveController } from "react-activation";
import { useNavigate } from "react-router-dom";
import FormModal, { FormModalRef } from "./components/FormModal";

const DictType: React.FC = () => {
  const navigate = useNavigate();
  const actionRef = React.useRef<ActionType>();
  const formRef = React.useRef<ProFormInstance>();
  const formModalRef = useRef<FormModalRef>(null);
  const { drop } = useAliveController();

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

  // 定义列
  const columns: ProColumns<DictTypeModel>[] = [
    {
      title: "序号", // 显示列标题
      dataIndex: "index", // 自定义索引
      valueType: "index", // 自动生成行号
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
      title: "字典编号",
      dataIndex: "id",
      width: 80,
      align: "center",
      hideInSearch: true
    },
    {
      title: "字典名称",
      dataIndex: "dictName",
      width: 220
    },
    {
      title: "字典类型",
      dataIndex: "dictType",
      width: 220,
      render: text => (
        <Tooltip title="点击查看详情">
          <Button type="link" style={{ padding: 0 }}>
            {text}
          </Button>
        </Tooltip>
      ),
      onCell: data => ({
        onClick: () => handleToDataClick(data.dictType!)
      })
    },
    {
      title: "备注",
      dataIndex: "remark",
      ellipsis: true,
      hideInSearch: true,
      width: 280
    },
    {
      title: "创建时间",
      key: "createdAt",
      dataIndex: "createdAt",
      valueType: "dateTime",
      hideInSearch: true,
      width: 180
    },
    {
      title: "创建时间",
      dataIndex: "createTime",
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
          <HocAuth permission={["admin:sys-dict-type:edit"]}>
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
          <HocAuth permission={["admin:sys-dict-type:del"]}>
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

  // 事件处理
  const handleToDataClick = async (dictType: string) => {
    await drop(DICT_DATA_URL);
    navigate(DICT_DATA_URL, { state: { dictType } });
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
          saveExcelBlob("字典列表", await exportDictTypeApi(formRef.current?.getFieldsValue()));
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
          const { code, msg } = await delDictTypeApi([id!]);
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
    <HocAuth permission={["admin:sys-dict-type:add"]}>
      <LoadingButton type="primary" key="addTable" icon={<PlusCircleOutlined />} onClick={done => handleShowAddFormModal(done)}>
        新增
      </LoadingButton>
    </HocAuth>,
    <HocAuth permission={["admin:sys-dict-type:export"]}>
      <LoadingButton type="primary" key="importTable" icon={<CloudDownloadOutlined />} onClick={done => handleExport(done)}>
        Excel导出
      </LoadingButton>
    </HocAuth>
  ];

  return (
    <>
      <ProTable<DictTypeModel>
        className="ant-pro-table-scroll"
        columns={columns}
        actionRef={actionRef}
        formRef={formRef}
        bordered
        cardBordered
        defaultSize="small"
        scroll={{ x: "2000", y: "100%" }}
        request={async params => {
          const { data } = await getDictTypePageApi(params);
          return formatDataForProTable<DictTypeModel>(data);
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
        headerTitle="字典类型列表"
        toolBarRender={toolBarRender}
      />
      <FormModal ref={formModalRef} onConfirm={handleFormModalConfirm} />
    </>
  );
};

export default DictType;
