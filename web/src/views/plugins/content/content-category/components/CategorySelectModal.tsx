import { ContentCategoryModel, getContentCategoryPageApi } from "@/api/plugins/content/content-category";
import LoadingButton from "@/components/LoadingButton";
import { pagination } from "@/config/proTable";
import { formatDataForProTable } from "@/utils";
import type { ProColumns } from "@ant-design/pro-components";
import { ProTable } from "@ant-design/pro-components";
import { Modal } from "antd";
import { forwardRef, useImperativeHandle, useState } from "react";

export interface CategorySelectModalRef {
  showCategorySelectModal: () => void;
}

interface ModalProps {
  onConfirm: (selectedRows: ContentCategoryModel) => void;
}

const CategorySelectModal = forwardRef<CategorySelectModalRef, ModalProps>(({ onConfirm }, ref) => {
  const [isModalOpen, setIsModalOpen] = useState(false);

  // 暴露方法给父组件
  useImperativeHandle(ref, () => ({
    showCategorySelectModal() {
      setIsModalOpen(true);
    }
  }));

  // 表格列定义
  const columns: ProColumns<ContentCategoryModel>[] = [
    {
      title: "序号",
      dataIndex: "index",
      valueType: "index",
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
      title: "分类编号",
      dataIndex: "id",
      width: 80,
      align: "left",
      hideInSearch: true
    },
    {
      title: "分类名称",
      dataIndex: "name",
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
    }
  ];

  // 处理确认
  const handleRowConfirm = (record: ContentCategoryModel) => {
    onConfirm(record);
    setIsModalOpen(false);
  };

  return (
    <Modal
      title="内容分类选择"
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
        </LoadingButton>
      ]}
      width={800}
    >
      <ProTable<ContentCategoryModel>
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
        onRow={record => ({
          onClick: () => handleRowConfirm(record) // 选中行时触发
        })}
        request={async params => {
          const { data } = await getContentCategoryPageApi(params);
          return formatDataForProTable<ContentCategoryModel>(data);
        }}
      />
    </Modal>
  );
});

export default CategorySelectModal;
