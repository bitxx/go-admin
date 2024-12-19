import { getDictOptions, getDictsApi } from "@/api/admin/sys/sys-dictdata";
import { getUserLevelPageApi, UserLevelModel } from "@/api/app/user/user-level";
import LoadingButton from "@/components/LoadingButton";
import { pagination } from "@/config/proTable";
import { ResultEnum } from "@/enums/httpEnum";
import { message } from "@/hooks/useMessage";
import { formatDataForProTable } from "@/utils";
import type { ProColumns } from "@ant-design/pro-components";
import { ProTable } from "@ant-design/pro-components";
import { Modal } from "antd";
import { forwardRef, useEffect, useImperativeHandle, useState } from "react";

export interface UserLevelSelectModalRef {
  showUserLevelSelectModal: () => void;
}

interface ModalProps {
  onConfirm: (selectedRows: UserLevelModel) => void;
}

const UserLevelSelectModal = forwardRef<UserLevelSelectModalRef, ModalProps>(({ onConfirm }, ref) => {
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [levelTypeOptions, setLevelTypeOptions] = useState<Map<string, string>>(new Map());

  // 暴露方法给父组件
  useImperativeHandle(ref, () => ({
    showUserLevelSelectModal() {
      setIsModalOpen(true);
    }
  }));

  // 表格列定义
  const columns: ProColumns<UserLevelModel>[] = [
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
      width: 120,
      align: "left"
    },
    {
      title: "等级",
      dataIndex: "level",
      width: 80,
      align: "left"
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

  // 处理确认
  const handleRowConfirm = (record: UserLevelModel) => {
    onConfirm(record);
    setIsModalOpen(false);
  };

  return (
    <Modal
      title="等级选择"
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
      <ProTable<UserLevelModel>
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
          const { data } = await getUserLevelPageApi(params);
          return formatDataForProTable<UserLevelModel>(data);
        }}
      />
    </Modal>
  );
});

export default UserLevelSelectModal;
