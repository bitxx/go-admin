import { modal } from "@/hooks/useMessage";

export function removeTab({ path, isCurrent }: { path: string; isCurrent: boolean }) {
	//因业务相关页面代码不能修改，这里方法仅用作兼容
	// console.log("请从标签关闭当前页面");
	// message.warning("因业务相关页面代码原则上不能修改，该按钮");

	modal.confirm({
		title: "提示",
		content:
			"因该框架兼容性问题，而业务相关页面代码需保持和我内部版一致。导致该按钮无法完成页面关闭功能。若需要关闭该页面，请从页面标签处点击'x'关闭",
		okText: "确认",
		// cancelText: "取消",
		maskClosable: true
	});
}
