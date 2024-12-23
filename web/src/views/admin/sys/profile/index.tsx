import { getUserProfileApi, updateProfileAvatar, updateProfileInfoApi, updateProfilePwdApi } from "@/api/admin/sys/sys-user";
import LoadingButton from "@/components/LoadingButton";
import { ResultEnum } from "@/enums/httpEnum";
import { message } from "@/hooks/useMessage";
import { RootState, useDispatch, useSelector } from "@/redux";
import { removeTab } from "@/redux/modules/tabs";
import { setUserInfo } from "@/redux/modules/user";
import { UserOutlined } from "@ant-design/icons";
import { Avatar, Card, Col, Divider, Form, Input, Radio, Row, Tabs, Upload } from "antd";
import ImgCrop from "antd-img-crop";
import React, { useEffect } from "react";
import { useLocation } from "react-router-dom";
import "./index.less";

const Profile: React.FC = () => {
  const dispatch = useDispatch();
  const userInfo = useSelector((state: RootState) => state.user.userInfo);
  const { pathname, search } = useLocation();
  const path = pathname + search;

  const [userInfoform] = Form.useForm();
  const [passwordForm] = Form.useForm();

  useEffect(() => {
    const getUserInfo = async () => {
      const { data: userInfo, code: userCode } = await getUserProfileApi();
      if (userCode === ResultEnum.SUCCESS) {
        dispatch(setUserInfo(userInfo));
      }
    };
    getUserInfo();
  }, []);

  const onUpdateUserInfo = async (done: () => void) => {
    try {
      userInfoform.validateFields().then(async values => {
        const { code, msg } = await updateProfileInfoApi({ ...values, id: userInfo.id });
        if (code !== ResultEnum.SUCCESS) {
          message.error(msg);
          return;
        }
        dispatch(
          setUserInfo({
            ...userInfo,
            username: values.username,
            phone: values.phone,
            email: values.email,
            sex: values.sex
          })
        );
        message.success(msg);
      });
    } finally {
      done();
    }
  };

  // 处理头像更改
  const handleAvatarChange = async ({ file }: any) => {
    if (file.status === "done") {
      const formData = new FormData();
      formData.append("avatar", file.originFileObj);
      const { data, msg, code } = await updateProfileAvatar(formData);
      if (code != ResultEnum.SUCCESS) {
        message.error(msg);
        return;
      }
      dispatch(setUserInfo({ ...userInfo, avatar: data }));
      message.success(msg);
    }
  };

  const onPasswordChange = async (done: () => void) => {
    try {
      passwordForm.validateFields().then(async values => {
        const { code, msg } = await updateProfilePwdApi(values);
        if (code !== ResultEnum.SUCCESS) {
          message.error(msg);
          return;
        }
        message.success(msg);
        passwordForm.resetFields();
      });
    } finally {
      done();
    }
  };

  return (
    <>
      <Row gutter={24}>
        {/* 左侧个人信息 */}
        <Col span={8}>
          <Card>
            <div className="userinfo-head">
              <ImgCrop>
                <Upload
                  name="avatar"
                  showUploadList={false}
                  customRequest={({ file, onSuccess }) => {
                    onSuccess && onSuccess("ok"); // 模拟上传成功
                  }}
                  onChange={handleAvatarChange}
                  accept="image/*"
                >
                  <Avatar
                    size={100}
                    src={userInfo.avatar || <UserOutlined />}
                    style={{ cursor: "pointer", marginBottom: "16px" }}
                  />
                </Upload>
              </ImgCrop>
              <p>点击头像修改</p>
            </div>
            <Divider />
            <ul className="userinfo">
              <li>
                <strong>用户名：</strong> {userInfo.username}
              </li>
              <Divider />
              <li>
                <strong>手机号：</strong> {userInfo.phone}
              </li>
              <Divider />
              <li>
                <strong>用户邮箱：</strong> {userInfo.email}
              </li>
              <Divider />
              <li>
                <strong>所属部门：</strong> {userInfo.deptName}
              </li>
              <Divider />
              <li>
                <strong>所属角色：</strong> {userInfo.roleName}
              </li>
              <Divider />
              <li>
                <strong>创建日期：</strong> {userInfo.createdAt}
              </li>
            </ul>
          </Card>
        </Col>

        {/* 右侧编辑表单 */}
        <Col span={16}>
          <Card>
            <Tabs defaultActiveKey="1">
              <Tabs.TabPane tab="基本资料" key="1">
                <Form
                  form={userInfoform}
                  layout="vertical"
                  initialValues={{
                    username: userInfo.username,
                    phone: userInfo.phone,
                    email: userInfo.email,
                    sex: userInfo.sex
                  }}
                >
                  {/* 用户名 */}
                  <Form.Item label="用户名" name="username" rules={[{ required: true, message: "请输入用户名！" }]}>
                    <Input />
                  </Form.Item>

                  {/* 手机号 */}
                  <Form.Item
                    label="手机号"
                    name="phone"
                    rules={[
                      { required: true, message: "请输入手机号！" },
                      { pattern: /^1[3-9]\d{9}$/, message: "请输入有效的手机号！" }
                    ]}
                  >
                    <Input />
                  </Form.Item>

                  {/* 邮箱 */}
                  <Form.Item
                    label="邮箱"
                    name="email"
                    rules={[
                      { required: true, message: "请输入邮箱！" },
                      { type: "email", message: "请输入有效的邮箱地址！" }
                    ]}
                  >
                    <Input />
                  </Form.Item>

                  {/* 性别 */}
                  <Form.Item label="性别" name="sex">
                    <Radio.Group>
                      <Radio value="1">男</Radio>
                      <Radio value="2">女</Radio>
                    </Radio.Group>
                  </Form.Item>

                  {/* 按钮组 */}
                  <Form.Item>
                    <div className="profile-submit">
                      <LoadingButton type="primary" htmlType="submit" onClick={done => onUpdateUserInfo(done)}>
                        保存
                      </LoadingButton>
                      <LoadingButton
                        htmlType="button"
                        onClick={done => {
                          dispatch(removeTab({ path, isCurrent: true }));
                          done();
                        }}
                      >
                        关闭
                      </LoadingButton>
                    </div>
                  </Form.Item>
                </Form>
              </Tabs.TabPane>
              <Tabs.TabPane tab="修改密码" key="2">
                <Form form={passwordForm} layout="vertical">
                  {/* Old Password */}
                  <Form.Item label="原密码" name="oldPassword" rules={[{ required: true, message: "请输入原密码！" }]}>
                    <Input.Password />
                  </Form.Item>

                  {/* New Password */}
                  <Form.Item label="新密码" name="newPassword" rules={[{ required: true, message: "请输入新密码！" }]}>
                    <Input.Password />
                  </Form.Item>

                  {/* Confirm New Password */}
                  <Form.Item
                    label="确认新密码"
                    name="confirmPassword"
                    rules={[
                      { required: true, message: "请确认新密码！" },
                      ({ getFieldValue }) => ({
                        validator(_, value) {
                          if (!value || getFieldValue("newPassword") === value) {
                            return Promise.resolve();
                          }
                          return Promise.reject(new Error("新密码和确认密码不匹配！"));
                        }
                      })
                    ]}
                  >
                    <Input.Password />
                  </Form.Item>

                  {/* Submit buttons */}
                  <Form.Item>
                    <div className="profile-submit">
                      <LoadingButton type="primary" htmlType="submit" onClick={done => onPasswordChange(done)}>
                        修改密码
                      </LoadingButton>
                      <LoadingButton
                        htmlType="button"
                        onClick={done => {
                          dispatch(removeTab({ path, isCurrent: true }));
                          done();
                        }}
                      >
                        关闭
                      </LoadingButton>
                    </div>
                  </Form.Item>
                </Form>
              </Tabs.TabPane>
            </Tabs>
          </Card>
        </Col>
      </Row>
    </>
  );
};

export default Profile;
