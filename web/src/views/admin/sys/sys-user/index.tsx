import { DeptModel, getDeptTreeApi } from "@/api/admin/sys/sys-dept";
import { getDictOptions, getDictsApi } from "@/api/admin/sys/sys-dictdata";
import { changeUserStatusApi, delUserApi, getUserPageApi, UserModel } from "@/api/admin/sys/sys-user";
import HocAuth from "@/components/HocAuth";
import LoadingButton from "@/components/LoadingButton";
import SwitchLoading from "@/components/LoadingSwitch";
import { pagination } from "@/config/proTable";
import { ResultEnum } from "@/enums/httpEnum";
import { message, modal } from "@/hooks/useMessage";
import { formatDataForProTable } from "@/utils";
import { DeleteOutlined, EditOutlined, ExclamationCircleOutlined, PlusCircleOutlined } from "@ant-design/icons";
import type { ActionType, ProColumns, ProFormInstance } from "@ant-design/pro-components";
import { ProTable } from "@ant-design/pro-components";
import { Layout, Space, Tree } from "antd";
import { Content } from "antd/es/layout/layout";
import Sider from "antd/es/layout/Sider";
import React, { Key, useEffect, useRef, useState } from "react";
import ChangePwdFormModal, { ChangePwdFormModalRef } from "./components/ChangePwdFormModal";
import FormModal, { FormModalRef } from "./components/FormModal";

const User: React.FC = () => {
  const STATUS_YES = "1";
  const STATUS_NO = "2";
  const actionRef = React.useRef<ActionType>();
  const tableFormRef = React.useRef<ProFormInstance>();
  const formModalRef = useRef<FormModalRef>(null);
  const changePwdFormModalRef = useRef<ChangePwdFormModalRef>(null);
  //const [sexOptions, setSexOptions] = useState<Map<string, string>>(new Map());
  const [statusOptions, setStatusOptions] = useState<Map<string, string>>(new Map());
  const [deptList, setDeptList] = useState<DeptModel[]>();
  const [selectDeptKey, setSelectDeptKey] = useState<Key>(0);

  // 定义列
  const columns: ProColumns<UserModel>[] = [
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
      title: "用户",
      dataIndex: "id",
      hideInSearch: true,
      width: 50,
      align: "left"
    },
    {
      title: "用户名",
      dataIndex: "username",
      width: 80,
      align: "left"
    },
    // {
    //   title: "昵称",
    //   dataIndex: "nickName",
    //   hideInSearch: true,
    //   width: 80,
    //   align: "left"
    // },
    // {
    //   title: "性别",
    //   dataIndex: "sex",
    //   hideInSearch: true,
    //   valueType: "select",
    //   valueEnum: sexOptions,
    //   width: 60,
    //   align: "left"
    // },
    {
      title: "角色",
      dataIndex: "roleId",
      hideInSearch: true,
      width: 100,
      render: (text, record) => record.role?.roleName,
      align: "left"
    },
    {
      title: "部门",
      dataIndex: "deptId",
      hideInSearch: true,
      render: (text, record) => record.dept?.deptName,
      width: 100,
      align: "left"
    },
    {
      title: "岗位",
      dataIndex: "postId",
      hideInSearch: true,
      render: (text, record) => record.post?.postName,
      width: 120,
      align: "left"
    },
    {
      title: "手机号",
      dataIndex: "phone",
      width: 120,
      align: "left"
    },
    {
      title: "邮箱",
      dataIndex: "email",
      hideInSearch: true,
      width: 180,
      align: "left"
    },
    {
      title: "状态",
      dataIndex: "status",
      valueType: "select",
      valueEnum: statusOptions,
      width: 80,
      align: "left",
      render: (text, record, index, action) => (
        <>
          {record.username !== "admin" && (
            <SwitchLoading
              checked={record.status === STATUS_YES}
              checkedChildren="开启"
              unCheckedChildren="关闭"
              onChange={checked => handleStatusChange(checked, record, action)}
            />
          )}
        </>
      )
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
      title: "操作",
      valueType: "option",
      align: "center",
      fixed: "right",
      width: 250,
      render: (_, data) => (
        <Space>
          <HocAuth permission={["admin:sys-user:edit"]}>
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
          <HocAuth permission={["admin:sys-user:edit-pwd"]}>
            <LoadingButton
              key="changePwd"
              type="link"
              size="small"
              icon={<EditOutlined />}
              onClick={done => handleShowChangePwdFormModal(data.id!, done)}
            >
              修改密码
            </LoadingButton>
          </HocAuth>
          {data.username !== "admin" && (
            <HocAuth permission={["admin:sys-user:del"]}>
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
          )}
        </Space>
      )
    }
  ];
  useEffect(() => {
    const initData = async () => {
      // const { data: sexData, msg: sexMsg, code: sexCode } = await getDictsApi("admin_sys_user_sex");
      // if (sexCode !== ResultEnum.SUCCESS) {
      //   message.error(sexMsg);
      //   return;
      // }
      // setSexOptions(getDictOptions(sexData));
      const { data: statusData, msg: statusMsg, code: statusCode } = await getDictsApi("admin_sys_status");
      if (statusCode !== ResultEnum.SUCCESS) {
        message.error(statusMsg);
        return;
      }
      setStatusOptions(getDictOptions(statusData));

      const { data: deptListData, msg: deptListMsg, code: deptListCode } = await getDeptTreeApi({});
      if (deptListCode !== ResultEnum.SUCCESS) {
        message.error(deptListMsg);
        return;
      }
      // setTimeout(() => setDeptList(deptListData), 1000);

      setDeptList(deptListData);
      setSelectDeptKey("");
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

  const handleShowChangePwdFormModal = (id: number, done: () => void) => {
    changePwdFormModalRef.current?.showChangePwdFormModal(id);
    setTimeout(() => done(), 1000);
  };

  const handleFormModalConfirm = () => {
    actionRef.current?.reload(true);
  };
  const handleChangePwdFormModalConfirm = () => {
    actionRef.current?.reload(true);
  };

  const handleStatusChange = async (checked: boolean, record: UserModel, action: any) => {
    const newStatus = checked ? STATUS_YES : STATUS_NO;
    const { code, msg } = await changeUserStatusApi(record.id!, newStatus);
    if (code !== ResultEnum.SUCCESS) {
      message.error(msg);
      return;
    }
    action.reload();
  };

  useEffect(() => {
    if (actionRef.current) {
      actionRef.current.reload(); // 每次 selectedKey 变化时刷新表格
    }
  }, [selectDeptKey]);

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
          const { code, msg } = await delUserApi([id!]);
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
    <HocAuth permission={["admin:sys-user:add"]}>
      <LoadingButton type="primary" key="addTable" icon={<PlusCircleOutlined />} onClick={done => handleShowAddFormModal(done)}>
        新增
      </LoadingButton>
    </HocAuth>
  ];

  return (
    <>
      <Layout>
        <Sider style={{ paddingTop: 20 }}>
          {deptList && (
            <>
              <LoadingButton
                type="link"
                onClick={done => {
                  setSelectDeptKey("");
                  done();
                }}
              >
                清空选中
              </LoadingButton>

              <Tree
                defaultExpandAll
                fieldNames={{ title: "deptName", key: "id", children: "children" }}
                treeData={deptList as any[]}
                selectedKeys={[selectDeptKey]}
                onSelect={selectedKeys => {
                  if (selectedKeys.length > 0) {
                    setSelectDeptKey(selectedKeys[0]);
                  }
                }}
              />
            </>
          )}
        </Sider>
        <Layout>
          <Content
            style={{
              padding: 0,
              margin: 0,
              display: "flex",
              flexDirection: "column",
              justifyContent: "space-between"
            }}
          >
            <ProTable<UserModel>
              className="ant-pro-table-scroll"
              columns={columns}
              actionRef={actionRef}
              formRef={tableFormRef}
              bordered
              cardBordered
              defaultSize="small"
              scroll={{ x: "100%", y: "100%" }}
              request={async params => {
                let reqData = { ...params, deptId: selectDeptKey as number };
                const { data } = await getUserPageApi(reqData);
                return formatDataForProTable<UserModel>(data);
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
              headerTitle="用户管理"
              toolBarRender={toolBarRender}
            />
          </Content>
        </Layout>
      </Layout>
      <FormModal ref={formModalRef} onConfirm={handleFormModalConfirm} />
      <ChangePwdFormModal ref={changePwdFormModalRef} onConfirm={handleChangePwdFormModalConfirm} />
    </>
  );
};

export default User;
