import { getMenuRoleApi } from "@/api/admin/sys/sys-menu";
import { getCaptchaApi, getUserProfileApi, loginApi } from "@/api/admin/sys/sys-user";
import LoadingButton from "@/components/LoadingButton";
import { HOME_URL } from "@/config";
import { ResultEnum } from "@/enums/httpEnum";
import { setRouteList, setToken, setUserInfo } from "@/redux/modules/global/action";
import { setTabsList } from "@/redux/modules/tabs/action";
import { CloseCircleOutlined, LockOutlined, UserOutlined } from "@ant-design/icons";
import { Form, Input, message } from "antd";
import { useEffect, useState } from "react";
import { useTranslation } from "react-i18next";
import { connect } from "react-redux";
import { useNavigate } from "react-router-dom";

const LoginForm = (props: any) => {
	const { t } = useTranslation();
	const { setToken, setRouteList, setUserInfo, setTabsList } = props;
	const [form] = Form.useForm();
	const [captchaInfo, setCaptchaInfo] = useState("");
	const [captchaId, setCaptchaId] = useState("");
	const navigate = useNavigate();

	const onCaptcha = async () => {
		const { data, msg, code } = await getCaptchaApi();
		if (code !== ResultEnum.SUCCESS) {
			message.error(msg);
			return;
		}
		setCaptchaInfo(data.data);
		setCaptchaId(data.id);
	};

	const onLogin = (done: () => void) => {
		form
			.validateFields()
			.then(async values => {
				try {
					setTabsList([]);
					message.open({ key: "loading", type: "loading", content: "登录中..." });
					// values = { ...values, password: md5(values.password), uuid: captchaId };
					values = { ...values, uuid: captchaId };
					const { data, msg, code } = await loginApi(values);
					if (code !== ResultEnum.SUCCESS) {
						message.error(msg);
						return;
					}
					setToken(data.token);
					const { data: userInfo, code: userCode, msg: userMsg } = await getUserProfileApi();
					if (userCode !== ResultEnum.SUCCESS) {
						message.error(userMsg);
						return;
					}
					const { data: routeList, code: routeCode, msg: routeMsg } = await getMenuRoleApi();
					if (routeCode !== ResultEnum.SUCCESS) {
						message.error(routeMsg);
						return;
					}

					setUserInfo(userInfo);
					setRouteList(routeList);

					message.success("登录成功！");
					navigate(HOME_URL);
				} catch (error) {
					onCaptcha();
				} finally {
					done();
					message.destroy("loading");
				}
			})
			.catch(() => {
				onCaptcha();
				message.error("表单校验失败");
				done();
			});
	};

	useEffect(() => {
		onCaptcha();
	}, []);

	return (
		<div className="login-form-content">
			<Form form={form} name="basic" initialValues={{ username: "admin", password: "123456" }} size="large" autoComplete="off">
				<Form.Item name="username" rules={[{ required: true, message: "请输入用户名!" }]}>
					<Input prefix={<UserOutlined />} placeholder="用户名: admin / test" />
				</Form.Item>
				<Form.Item name="password" rules={[{ required: true, message: "请输入密码!" }]}>
					<Input.Password prefix={<LockOutlined />} placeholder="密码：123456" />
				</Form.Item>
				<Form.Item name="code" rules={[{ required: true, message: "请输入验证码!" }]}>
					<div className="login-form-captcha-input-group">
						<Input placeholder="验证码:" />
						<img src={captchaInfo} className="login-form-captcha-img" onClick={onCaptcha} />
					</div>
				</Form.Item>
				<Form.Item className="login-btn">
					<LoadingButton
						onClick={done => {
							form.resetFields();
							setTimeout(() => done(), 1000);
						}}
						icon={<CloseCircleOutlined />}
					>
						{t("login.reset")}
					</LoadingButton>
					<LoadingButton type="primary" onClick={done => onLogin(done)} htmlType="submit" icon={<UserOutlined />}>
						{t("login.confirm")}
					</LoadingButton>
				</Form.Item>
			</Form>
		</div>
	);
};

const mapDispatchToProps = { setToken, setUserInfo, setRouteList, setTabsList };
export default connect(null, mapDispatchToProps)(LoginForm);
