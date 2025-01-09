import {
  DBTableModel,
  delGenTableApi,
  downloadCodeApi,
  genCodeApi,
  genMenuApi,
  GenTableModel,
  getGenTablePageApi,
  importDBTableApi
} from "@/api/admin/sys/sys-tools/sys-gen";
import HocAuth from "@/components/HocAuth";
import LoadingButton from "@/components/LoadingButton";
import { GEN_EDIT_TABLE } from "@/config";
import { pagination } from "@/config/proTable";
import { ResultEnum } from "@/enums/httpEnum";
import { message, modal } from "@/hooks/useMessage";
import { formatDataForProTable, saveZipBlob } from "@/utils";
import { CloudUploadOutlined, DeleteOutlined, EditOutlined, ExclamationCircleOutlined, EyeOutlined } from "@ant-design/icons";
import type { ActionType, ProColumns } from "@ant-design/pro-components";
import { ProTable } from "@ant-design/pro-components";
import { Space } from "antd";
import { useRef } from "react";
import { useAliveController } from "react-activation";
import { useNavigate } from "react-router-dom";
import DBTableModal, { DBTableModalRef } from "./components/DBTableModal";
import PreviewCodeModal, { PreviewCodeModalRef } from "./components/PreviewCodeModal";

const GenTable = () => {
  const navigate = useNavigate();
  const dbTableModalRef = useRef<DBTableModalRef>(null);
  const actionRef = useRef<ActionType>();
  const previewCodeModalRef = useRef<PreviewCodeModalRef>(null);
  const { drop } = useAliveController();

  const columns: ProColumns<GenTableModel>[] = [
    {
      title: "序号", // 显示列标题
      dataIndex: "index", // 自定义索引
      valueType: "index", // 自动生成行号
      width: 80,
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
      title: "表编号",
      dataIndex: "id",
      width: 80,
      hideInSearch: true
    },
    {
      title: "数据库表名",
      dataIndex: "tableName",
      width: 300
    },
    {
      title: "数据库表描述",
      dataIndex: "tableComment",
      width: 150
    },
    {
      title: "类名",
      dataIndex: "className",
      hideInSearch: true,
      width: 180
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
      key: "option",
      fixed: "right",
      width: 530,
      hideInSearch: true,
      render: (_, data) => (
        <Space>
          <HocAuth permission={["admin:sys-gen:edit"]}>
            <LoadingButton
              key="edit"
              type="link"
              size="small"
              icon={<EditOutlined />}
              onClick={done => {
                handleEditTable(data.id, done);
              }}
            >
              编辑
            </LoadingButton>
          </HocAuth>
          <HocAuth permission={["admin:sys-gen:preview"]}>
            <LoadingButton
              key="preview"
              type="link"
              size="small"
              icon={<EyeOutlined />}
              onClick={done => {
                handlePreviewCode(data.id, done);
              }}
            >
              预览
            </LoadingButton>
          </HocAuth>
          <HocAuth permission={["admin:sys-gen:gen-code"]}>
            <LoadingButton
              key="genCode"
              type="link"
              size="small"
              icon={<EyeOutlined />}
              onClick={done => {
                handleGenCode(data.id, done);
              }}
            >
              生成代码
            </LoadingButton>
          </HocAuth>
          <HocAuth permission={["admin:sys-gen:download-code"]}>
            <LoadingButton
              key="downCode"
              type="link"
              size="small"
              icon={<EyeOutlined />}
              onClick={done => {
                handleDownloadCode(data.id, done);
              }}
            >
              下载代码
            </LoadingButton>
          </HocAuth>
          <HocAuth permission={["admin:sys-gen:import-config"]}>
            <LoadingButton
              key="genMenu"
              type="link"
              size="small"
              icon={<EyeOutlined />}
              onClick={done => {
                handleGenMenu(data.id, done);
              }}
            >
              配置导入
            </LoadingButton>
          </HocAuth>
          <HocAuth permission={["admin:sys-gen:del"]}>
            <LoadingButton
              key="del"
              type="link"
              size="small"
              danger
              icon={<DeleteOutlined />}
              onClick={done => handleDelete(data.id, done)}
            >
              删除
            </LoadingButton>
          </HocAuth>
        </Space>
      )
    }
  ];

  const handleShowDBTableModal = (done: () => void) => {
    dbTableModalRef.current?.showDBTableModal();
    setTimeout(() => done(), 1000);
  };

  // confirm the select table
  const handleDBTableConfirm = async (selectedRows: DBTableModel[]) => {
    let tbNames: string[] = [];
    selectedRows.forEach(item => tbNames.push(item.tableName));
    const { msg, code } = await importDBTableApi(tbNames);
    if (code != ResultEnum.SUCCESS) {
      message.error(msg);
      return;
    }
    actionRef.current?.reload(false);
  };

  const handlePreviewCodeModalConfirm = () => {};

  const handleEditTable = async (id: number, done: () => void) => {
    await drop(GEN_EDIT_TABLE);
    navigate(GEN_EDIT_TABLE, { state: { id } });
    setTimeout(() => done(), 1000);
  };

  // view code
  const handlePreviewCode = async (id: number, done: () => void) => {
    previewCodeModalRef.current?.showPreviewCodeModal(id);
    setTimeout(() => done(), 1000);
  };

  const handleGenCode = async (id: number, done: () => void) => {
    modal.confirm({
      title: "提示",
      icon: <ExclamationCircleOutlined />,
      content: "是否为编号 " + id + " 生成代码?",
      okText: "确认",
      cancelText: "取消",
      maskClosable: true,
      onCancel: () => {
        done();
      },
      onOk: async () => {
        try {
          const { msg, code } = await genCodeApi(id);
          if (code !== ResultEnum.SUCCESS) {
            message.error(msg);
            return;
          }
          message.success(msg);
        } finally {
          done();
        }
      }
    });
  };
  const handleGenMenu = async (id: number, done: () => void) => {
    modal.confirm({
      title: "提示",
      icon: <ExclamationCircleOutlined />,
      content: "导入完成后，需要重新手动刷新页面。是否为编号 " + id + " 导入菜单到数据库?",
      okText: "确认",
      cancelText: "取消",
      maskClosable: true,
      onCancel: () => {
        done();
      },
      onOk: async () => {
        try {
          const { msg, code } = await genMenuApi(id);
          if (code !== ResultEnum.SUCCESS) {
            message.error(msg);
            return;
          }
          message.success(msg);
        } finally {
          done();
        }
      }
    });
  };
  const handleDownloadCode = async (id: number, done: () => void) => {
    modal.confirm({
      title: "提示",
      icon: <ExclamationCircleOutlined />,
      content: "是否确认下载编号为 " + id + " 的代码?",
      okText: "确认",
      cancelText: "取消",
      maskClosable: true,
      onCancel: () => {
        done();
      },
      onOk: async () => {
        try {
          saveZipBlob("code", await downloadCodeApi(id));
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
          const { code, msg } = await delGenTableApi([id]);
          if (code !== ResultEnum.SUCCESS) {
            message.error(msg);
            return;
          }
          actionRef.current?.reload(false);
          message.success(msg);
        } finally {
          done;
        }
      }
    });
  };

  const toolBarRender = () => [
    <LoadingButton type="primary" key="importTable" icon={<CloudUploadOutlined />} onClick={done => handleShowDBTableModal(done)}>
      导入数据库表
    </LoadingButton>
  ];

  return (
    <>
      <ProTable<GenTableModel>
        className="ant-pro-table-scroll"
        columns={columns}
        actionRef={actionRef}
        bordered
        cardBordered
        defaultSize="small"
        scroll={{ x: "100%", y: "100%" }}
        request={async params => {
          const { data } = await getGenTablePageApi(params);
          return formatDataForProTable<GenTableModel>(data);
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
        search={{ labelWidth: "auto" }}
        pagination={pagination}
        dateFormatter="string"
        headerTitle="代码列表"
        toolBarRender={toolBarRender}
      />
      <DBTableModal ref={dbTableModalRef} onConfirm={handleDBTableConfirm} />
      <PreviewCodeModal ref={previewCodeModalRef} onConfirm={handlePreviewCodeModalConfirm} />
    </>
  );
};

export default GenTable;
