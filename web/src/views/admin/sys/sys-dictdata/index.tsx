import { delDictDataApi, DictDataModel, getDictDataPageApi } from "@/api/admin/sys/sys-dictdata";
import HocAuth from "@/components/HocAuth";
import LoadingButton from "@/components/LoadingButton";
import { pagination } from "@/config/proTable";
import { ResultEnum } from "@/enums/httpEnum";
import { message, modal } from "@/hooks/useMessage";
import { formatDataForProTable } from "@/utils";
import { DeleteOutlined, EditOutlined, ExclamationCircleOutlined, PlusCircleOutlined } from "@ant-design/icons";
import type { ActionType, ProColumns } from "@ant-design/pro-components";
import { ProTable } from "@ant-design/pro-components";
import { Space } from "antd";
import { useRef } from "react";
import { useLocation } from "react-router-dom";
import FormModal, { FormModalRef } from "./components/FormModal";

const DictData: React.FC = () => {
  const actionRef = useRef<ActionType>();
  const formModalRef = useRef<FormModalRef>(null);
  const location = useLocation();
  const state = location.state;

  const handleShowAddFormModal = (dictType: string, done: () => void) => {
    formModalRef.current?.showAddFormModal(dictType);
    setTimeout(() => done(), 1000);
  };
  const handleShowEditFormModal = (dictDataId: number, done: () => void) => {
    formModalRef.current?.showEditFormModal(dictDataId);
    setTimeout(() => done(), 1000);
  };

  const handleFormModalConfirm = () => {
    actionRef.current?.reload(true);
  };

  // 定义列
  const columns: ProColumns<DictDataModel>[] = [
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
      title: "字典编号",
      dataIndex: "id",
      width: 80,
      align: "center",
      hideInSearch: true
    },
    {
      title: "字典标签",
      dataIndex: "dictLabel",
      width: 150
    },
    {
      title: "字典键值",
      dataIndex: "dictValue",
      width: 80,
      align: "center",
      hideInSearch: true
    },
    {
      title: "字典排序",
      dataIndex: "dictSort",
      width: 80,
      align: "center",
      hideInSearch: true
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
          <HocAuth permission={["admin:sys-dict-data:edit"]}>
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
          <HocAuth permission={["admin:sys-dict-data:del"]}>
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
          const { code, msg } = await delDictDataApi([id]);
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
    <HocAuth permission={["admin:sys-dict-data:add"]}>
      <LoadingButton
        type="primary"
        key="addTable"
        icon={<PlusCircleOutlined />}
        onClick={done => handleShowAddFormModal(state.dictType, done)}
      >
        新增
      </LoadingButton>
    </HocAuth>
  ];

  return (
    <>
      <ProTable<DictDataModel>
        className="ant-pro-table-scroll"
        columns={columns}
        actionRef={actionRef}
        bordered
        cardBordered
        defaultSize="small"
        scroll={{ x: "2000", y: "100%" }}
        request={async params => {
          const { data } = await getDictDataPageApi(params, state.dictType);
          return formatDataForProTable<DictDataModel>(data);
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
        headerTitle="字典数据列表"
        toolBarRender={toolBarRender}
      />
      <FormModal ref={formModalRef} onConfirm={handleFormModalConfirm} />
    </>
  );
};

export default DictData;
