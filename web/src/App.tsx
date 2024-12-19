import { setLanguage } from "@/redux/modules/global/action";
import Router from "@/routers/index";
import AuthRouter from "@/routers/utils/authRouter";
import { getBrowserLang } from "@/utils/util";
import { App as AppLayout, ConfigProvider } from "antd";
import enUS from "antd/lib/locale/en_US";
import zhCN from "antd/lib/locale/zh_CN";
import i18n from "i18next";
import "moment/dist/locale/zh-cn";
import { useEffect, useState } from "react";
import { AliveScope } from "react-activation";
import { connect } from "react-redux";
import { HashRouter } from "react-router-dom";

const App = (props: any) => {
	const { language, assemblySize, setLanguage } = props;
	const [i18nLocale, setI18nLocale] = useState(zhCN);

	// 设置 antd 语言国际化
	const setAntdLanguage = () => {
		// 如果 redux 中有默认语言就设置成 redux 的默认语言，没有默认语言就设置成浏览器默认语言
		if (language && language == "zh") return setI18nLocale(zhCN);
		if (language && language == "en") return setI18nLocale(enUS);
		if (getBrowserLang() == "zh") return setI18nLocale(zhCN);
		if (getBrowserLang() == "en") return setI18nLocale(enUS);
	};

	useEffect(() => {
		// 全局使用国际化
		i18n.changeLanguage(language || getBrowserLang());
		setLanguage(language || getBrowserLang());
		setAntdLanguage();
	}, [language]);

	return (
		<ConfigProvider
			locale={i18nLocale}
			componentSize={"middle"}
			button={{
				autoInsertSpace: true
			}}
			// theme={{
			// 	algorithm: theme.defaultAlgorithm
			// }}
		>
			<AppLayout style={{ height: "100vh", display: "flex", flexDirection: "column" }}>
				<AliveScope>
					<HashRouter>
						<AuthRouter>
							<Router />
						</AuthRouter>
					</HashRouter>
				</AliveScope>
			</AppLayout>
		</ConfigProvider>
	);
};

const mapStateToProps = (state: any) => state.global;
const mapDispatchToProps = { setLanguage };
export default connect(mapStateToProps, mapDispatchToProps)(App);
