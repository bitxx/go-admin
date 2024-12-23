import { getDictOptions, getDictsApi } from "@/api/admin/sys/sys-dictdata";
import { getAllDictTypesApi } from "@/api/admin/sys/sys-dicttype";
import {
  GenTableColumnModel,
  GenTableDetailModel,
  getGenTableDetailApi,
  updateGenTableApi
} from "@/api/admin/sys/sys-tools/sys-gen";
import LoadingButton from "@/components/LoadingButton";
import { ResultEnum } from "@/enums/httpEnum";
import { message } from "@/hooks/useMessage";
import { useDispatch } from "@/redux";
import { removeTab } from "@/redux/modules/tabs";
import type { ProColumns } from "@ant-design/pro-components";
import { EditableProTable } from "@ant-design/pro-components";
import { AutoComplete, Card, Col, Form, Input, Row, Space, Tabs } from "antd";
import { debounce } from "lodash";
import React, { useEffect, useState } from "react";
import { useLocation } from "react-router-dom";

const GenTable: React.FC = () => {
  const dispatch = useDispatch();
  const { pathname, search, state } = useLocation();
  const path = pathname + search;
  const [baseInfoForm] = Form.useForm();
  const [genInfoForm] = Form.useForm();

  const [editableKeys, setEditableRowKeys] = useState<React.Key[]>([]);
  const [dataSource, setDataSource] = useState<GenTableDetailModel>();

  const [sysGenGoTypeOptoins, setSysGenGoTypeOptoins] = useState<Map<string, string>>(new Map());
  const [sysGenQueryTypeOptoins, setSysGenQueryTypeOptoins] = useState<Map<string, string>>(new Map());
  const [sysGenHtmlTypeOptoins, setSysGenHtmlTypeOptoins] = useState<Map<string, string>>(new Map());
  const [sysYesNoOptoins, setSysYesNoOptoins] = useState<Map<string, string>>(new Map());
  const [allDictTypeOptions, setAllDictTypeOptions] = useState<Map<string, string>>(new Map());
  const [autoQueryDictTypeOptions, setAutoQueryDictTypeOptions] = useState<Map<string, string>>(new Map());

  const defaultActiveTab = "2";
  const [activeTab, setActiveTab] = useState(defaultActiveTab);
  const tabItems = [
    {
      key: "1",
      label: "基本信息",
      children: null
    },
    {
      key: "2",
      label: "字段信息",
      children: null
    },
    {
      key: "3",
      label: "生成信息",
      children: null
    }
  ];

  const columns: ProColumns<GenTableColumnModel>[] = [
    {
      title: "序号",
      dataIndex: "index",
      valueType: "index",
      width: 50,
      align: "center",
      fixed: "left",
      className: "gray-cell",
      editable: false
    },
    {
      title: "字段列名",
      dataIndex: "columnName",
      tooltip: "只读，数据库表名",
      align: "center",
      fixed: "left",
      width: 180,
      editable: false
    },
    {
      title: "字段描述",
      dataIndex: "columnComment",
      tooltip: "可编辑，表描述",
      fixed: "left",
      width: 260
    },
    {
      title: "物理类型",
      dataIndex: "columnType",
      tooltip: "只读，字段对应在数据库中的类型",
      width: 120,
      editable: false
    },
    {
      title: "go类型",
      dataIndex: "goType",
      tooltip: "可编辑，在golang中对应的数据类型",
      width: 120,
      valueType: "select",
      valueEnum: sysGenGoTypeOptoins
    },
    {
      title: "go属性",
      dataIndex: "goField",
      tooltip: "可编辑，在golang中的字段名",
      width: 180
    },
    {
      title: "json属性",
      dataIndex: "jsonField",
      tooltip: "可编辑，对应json的字段名",
      width: 180
    },
    {
      title: "查询方式",
      dataIndex: "queryType",
      tooltip: "数据库查询方式",
      width: 120,
      valueType: "select",
      valueEnum: sysGenQueryTypeOptoins
    },
    {
      title: "显示类型",
      dataIndex: "htmlType",
      tooltip: "根据字段类型选择要使用的组件，比如int类型，则选择数字文本框",
      width: 140,
      valueType: "select",
      valueEnum: sysGenHtmlTypeOptoins
    },
    {
      title: "字典类型",
      dataIndex: "dictType",
      tooltip: "注意，只有选择下拉框、单选框等支持单选的组件时，该选项才有效，否则会造成代码混乱",
      width: 300,
      valueType: "text",
      renderFormItem: (text, { record, type, defaultRender }) => (
        <AutoComplete
          options={(() => {
            const options: { value: string; label: string }[] = [];
            autoQueryDictTypeOptions.forEach((key, value) => {
              options.push({ value: value, label: key + " => " + value });
            });
            return options;
          })()}
          dropdownStyle={{ width: 500 }} // 设置下拉框宽度
          onSearch={handleSearch}
          style={{ width: "100%" }}
          defaultValue={text}
        >
          <Input />
        </AutoComplete>
      )
    },
    {
      title: "编辑",
      dataIndex: "isRequired",
      tooltip: "是否可编辑",
      width: 120,
      valueType: "select",
      valueEnum: sysYesNoOptoins
    },
    {
      title: "列表",
      dataIndex: "isList",
      tooltip: "是否在列表展示",
      width: 120,
      valueType: "select",
      valueEnum: sysYesNoOptoins
    },
    {
      title: "查询",
      dataIndex: "isQuery",
      tooltip: "是否查询",
      width: 120,
      valueType: "select",
      valueEnum: sysYesNoOptoins
    }
  ];

  // 本地模糊查询函数
  const handleSearch = debounce((value: string) => {
    if (!value || value === "" || value === undefined) {
      setAutoQueryDictTypeOptions(allDictTypeOptions);
    } else {
      const filter = new Map<string, string>(
        allDictTypeOptions.entries().filter(([key]) => key.toLowerCase().includes(value.toLowerCase()))
      );
      setAutoQueryDictTypeOptions(filter);
    }
  }, 300);

  const onSubmit = async (done: () => void) => {
    try {
      baseInfoForm.validateFields();
      genInfoForm.validateFields();
      const { msg, code } = await updateGenTableApi(dataSource?.id!, dataSource!);
      if (code !== ResultEnum.SUCCESS) {
        message.error(msg);
        return;
      }
      message.success(msg);
      dispatch(removeTab({ path, isCurrent: true }));
    } catch (error) {
      console.error("validate error：", error);
      message.error("表单校验失败");
    } finally {
      done();
    }
  };

  useEffect(() => {
    const getGenTableDetail = async () => {
      const { data: genTableDetailData, msg: genTableDetailMsg, code: genTableDetailCode } = await getGenTableDetailApi(state.id);
      if (genTableDetailCode !== ResultEnum.SUCCESS) {
        message.error(genTableDetailMsg);
        return;
      }

      const { data: sysGenGoTypeData, msg: sysGenGoTypeMsg, code: sysGenGoTypeCode } = await getDictsApi("admin_sys_gen_go_type");
      if (sysGenGoTypeCode !== ResultEnum.SUCCESS) {
        message.error(sysGenGoTypeMsg);
        return;
      }

      const {
        data: sysGenQueryTypeData,
        msg: sysGenQueryTypeMsg,
        code: sysGenQueryTypeCode
      } = await getDictsApi("admin_sys_gen_query_type");
      if (sysGenQueryTypeCode !== ResultEnum.SUCCESS) {
        message.error(sysGenQueryTypeMsg);
        return;
      }

      const {
        data: sysGenHtmlTypeData,
        msg: sysGenHtmlTypeMsg,
        code: sysGenHtmlTypeCode
      } = await getDictsApi("admin_sys_gen_html_type");
      if (sysGenHtmlTypeCode !== ResultEnum.SUCCESS) {
        message.error(sysGenHtmlTypeMsg);
        return;
      }

      const { data: sysYesNoData, msg: sysYesNoMsg, code: sysYesNoCode } = await getDictsApi("admin_sys_yes_no");
      if (sysYesNoCode !== ResultEnum.SUCCESS) {
        message.error(sysYesNoMsg);
        return;
      }

      const { data: allDictTypesData, msg: allDictTypesMsg, code: allDictTypesCode } = await getAllDictTypesApi();
      if (allDictTypesCode !== ResultEnum.SUCCESS) {
        message.error(allDictTypesMsg);
        return;
      }

      setSysGenGoTypeOptoins(getDictOptions(sysGenGoTypeData));
      setSysGenQueryTypeOptoins(getDictOptions(sysGenQueryTypeData));
      setSysGenHtmlTypeOptoins(getDictOptions(sysGenHtmlTypeData));
      setSysYesNoOptoins(getDictOptions(sysYesNoData));

      setAllDictTypeOptions(new Map(allDictTypesData.map(({ dictType, dictName }) => [dictType || "", dictName || ""])));
      setAutoQueryDictTypeOptions(new Map(allDictTypesData.map(({ dictType, dictName }) => [dictType || "", dictName || ""])));

      setDataSource(genTableDetailData);
      setEditableRowKeys(genTableDetailData.sysGenColumns!.map(item => item.id));
    };
    getGenTableDetail();
  }, []);

  return (
    <Card>
      <Tabs activeKey={activeTab} onChange={key => setActiveTab(key)} centered items={tabItems} />
      {activeTab === "1" && (
        <Form
          form={baseInfoForm}
          layout="vertical"
          initialValues={dataSource}
          onValuesChange={(param, params) => {
            setDataSource({ ...dataSource, ...params });
          }}
        >
          <Row gutter={24}>
            <Col span={12}>
              <Form.Item name="tableName" label="数据表名称">
                <Input readOnly disabled />
              </Form.Item>
            </Col>
            <Col span={12}>
              <Form.Item name="tableComment" label="菜单名称" rules={[{ required: true, message: "请输入菜单名称" }]}>
                <Input placeholder="请输入菜单名称" />
              </Form.Item>
            </Col>
            <Col span={12}>
              <Form.Item name="className" label="结构体模型名称" rules={[{ required: true, message: "请输入结构体模型名称" }]}>
                <Input placeholder="请输入结构体模型名称" />
              </Form.Item>
            </Col>
            <Col span={12}>
              <Form.Item name="functionAuthor" label="作者名称" rules={[{ required: true, message: "请输入作者名称" }]}>
                <Input placeholder="请输入作者名称" />
              </Form.Item>
            </Col>
            <Col span={24}>
              <Form.Item name="remark" label="备注">
                <Input.TextArea placeholder="请输入内容" />
              </Form.Item>
            </Col>
          </Row>
        </Form>
      )}

      {activeTab === "2" && (
        <EditableProTable<GenTableColumnModel>
          rowKey="id"
          className="ant-pro-table-scroll"
          scroll={{ x: 1000, y: "100%" }}
          bordered
          cardBordered
          pagination={false}
          recordCreatorProps={false}
          loading={false}
          columns={columns}
          value={dataSource?.sysGenColumns}
          editable={{
            type: "multiple",
            editableKeys,
            onChange: setEditableRowKeys,
            onValuesChange: (record, recordList) => {
              setDataSource({ ...dataSource, sysGenColumns: recordList });
            }
          }}
        />
      )}
      {activeTab === "3" && (
        <Form
          form={genInfoForm}
          layout="vertical"
          initialValues={dataSource}
          onValuesChange={(param, params) => {
            setDataSource({ ...dataSource, ...params });
          }}
        >
          <Row gutter={24}>
            <Col span={12}>
              <Form.Item name="packageName" label="应用名" rules={[{ required: true, message: "请输入应用名" }]}>
                <Input placeholder="请输入应用名" />
              </Form.Item>
            </Col>
            <Col span={12}>
              <Form.Item name="businessName" label="业务名" rules={[{ required: true, message: "请输入业务名" }]}>
                <Input placeholder="请输入业务名" />
              </Form.Item>
            </Col>
            <Col span={12}>
              <Form.Item name="functionName" label="功能描述" rules={[{ required: true, message: "请输入功能描述" }]}>
                <Input placeholder="请输入功能描述" />
              </Form.Item>
            </Col>
            <Col span={12}>
              <Form.Item name="moduleName" label="接口路径">
                <Input addonBefore="api/{version}/" placeholder="请输入接口路径" />
              </Form.Item>
            </Col>
          </Row>
        </Form>
      )}

      <Space align="center" style={{ width: "100%", display: "flex", justifyContent: "center", padding: "10px 16px" }}>
        <LoadingButton
          key="confirm"
          type="primary"
          style={{
            padding: "6px 16px",
            fontSize: "14px"
          }}
          onClick={done => onSubmit(done)}
        >
          确定
        </LoadingButton>
        <LoadingButton
          key="close"
          type="default"
          style={{
            padding: "6px 16px",
            fontSize: "14px"
          }}
          onClick={done => {
            dispatch(removeTab({ path, isCurrent: true }));
            done();
          }}
        >
          关闭
        </LoadingButton>
      </Space>
    </Card>
  );
};

export default GenTable;
