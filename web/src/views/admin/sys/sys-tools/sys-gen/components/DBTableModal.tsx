import { DBTableModel, getDBTablePageApi } from "@/api/admin/sys/sys-tools/sys-gen";
import LoadingButton from "@/components/LoadingButton";
import { pagination } from "@/config/proTable";
import { formatDataForProTable } from "@/utils";
import type { ProColumns } from "@ant-design/pro-components";
import { ProTable } from "@ant-design/pro-components";
import { Modal } from "antd";
import { forwardRef, useImperativeHandle, useState } from "react";

export interface DBTableModalRef {
  showDBTableModal: () => void; // 显示模态框
}

interface ModalProps {
  onConfirm: (selectedRows: DBTableModel[]) => void; // 确认时回传选中的数据
}

const DBTableModal = forwardRef<DBTableModalRef, ModalProps>(({ onConfirm }, ref) => {
  const [isModalOpen, setIsModalOpen] = useState(false); // 控制模态框显示
  const [selectedRows, setSelectedRows] = useState<DBTableModel[]>([]); // 存储选中的行

  // 暴露方法给父组件
  useImperativeHandle(ref, () => ({
    showDBTableModal() {
      setIsModalOpen(true);
    }
  }));

  // 表格列定义
  const columns: ProColumns<DBTableModel>[] = [
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
      title: "表名称",
      dataIndex: "tableName",
      align: "center",
      width: 200
    },
    {
      title: "表描述",
      dataIndex: "tableComment",
      align: "center",
      width: 150
    },
    {
      title: "创建时间",
      key: "createdAt",
      dataIndex: "createdAt",
      valueType: "dateTime",
      hideInSearch: true
    }
  ];

  // 处理确认
  const handleConfirm = async (done: () => void) => {
    await onConfirm(selectedRows);
    setIsModalOpen(false);
    done();
  };

  return (
    <Modal
      title="导入表"
      open={isModalOpen}
      onCancel={() => setIsModalOpen(false)}
      destroyOnClose
      footer={[
        <LoadingButton
          key="cancel"
          onClick={done => {
            setIsModalOpen(false);
            done();
          }}
        >
          取消
        </LoadingButton>,
        <LoadingButton
          key="confirm"
          type="primary"
          onClick={done => handleConfirm(done)}
          disabled={selectedRows.length === 0} // 没有选择时禁用按钮
        >
          确定
        </LoadingButton>
      ]}
      width={800}
    >
      <ProTable<DBTableModel>
        style={{ maxHeight: "50vh", overflowY: "auto" }}
        columns={columns}
        bordered
        cardBordered
        defaultSize="small"
        rowKey="tableName"
        toolBarRender={false} // 隐藏工具栏
        dateFormatter="string"
        search={{ labelWidth: "auto" }}
        pagination={{ ...pagination, defaultPageSize: 5 }}
        columnsState={{
          persistenceKey: "use-pro-table-key",
          persistenceType: "localStorage"
        }}
        rowSelection={{
          type: "checkbox",
          onChange: (_, selectedRows) => setSelectedRows(selectedRows)
        }}
        request={async params => {
          const { data } = await getDBTablePageApi(params);
          return formatDataForProTable<DBTableModel>(data);
        }}
      />
    </Modal>
  );
});

export default DBTableModal;
