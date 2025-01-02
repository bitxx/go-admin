import type { IDomEditor } from "@wangeditor/editor"; // 导入类型
import { Editor, Toolbar } from "@wangeditor/editor-for-react";
import "@wangeditor/editor/dist/css/style.css"; // 引入样式
import React, { useEffect, useState } from "react";
import LoadingButton from "../LoadingButton";
import "./index.less"; // 自定义样式

interface RichTextEditorProps {
	value?: string; // 初始内容
	onChange?: (html: string) => void; // 内容变化回调
	placeholder?: string; // 占位提示
	toolbarConfig?: Record<string, unknown>; // 工具栏配置
	editorStyle?: React.CSSProperties; // 编辑器样式
}

const RichTextEditor: React.FC<RichTextEditorProps> = ({
	value = "",
	onChange,
	placeholder = "请输入内容...",
	toolbarConfig = {},
	editorStyle = {}
}) => {
	const [editor, setEditor] = useState<IDomEditor | null>(null); // 编辑器实例
	const [html, setHtml] = useState<string>(value); // 当前内容
	const [isToolbarVisible, setIsToolbarVisible] = useState(false); // 工具栏是否可见

	// 在组件卸载时销毁编辑器
	useEffect(() => {
		return () => {
			if (editor) {
				editor.destroy();
				setEditor(null); // 清除 editor 状态
			}
		};
	}, []); // 空数组确保仅在组件卸载时触发清理

	const handleEditorChange = (currentEditor: IDomEditor) => {
		const content = currentEditor.getHtml(); // 获取HTML内容
		setHtml(content); // 更新本地内容
		if (onChange) {
			onChange(content); // 调用外部回调
		}
	};

	return (
		<div className="rich-text-editor">
			{/* 折叠按钮 */}
			<div className="rich-text-editor-toolbar-toggle">
				<LoadingButton
					type="text"
					onClick={done => {
						setIsToolbarVisible(prev => !prev);
						setTimeout(() => done(), 500);
					}}
					style={{ fontSize: 14 }}
				>
					{isToolbarVisible ? "隐藏工具栏 ▲" : "显示工具栏 ▼"}
				</LoadingButton>
			</div>
			{/* 工具栏 */}
			{isToolbarVisible && (
				<div className="rich-text-editor-toolbar">
					<Toolbar
						editor={editor}
						defaultConfig={{
							excludeKeys: ["fullScreen"], // 默认移除全屏按钮
							...toolbarConfig
						}}
					/>
				</div>
			)}
			{/* 编辑器 */}
			<div className="rich-text-editor-content">
				<Editor
					defaultConfig={{
						placeholder,
						onChange: handleEditorChange
					}}
					value={html}
					onCreated={currentEditor => setEditor(currentEditor)} // 编辑器创建时保存实例
					style={{
						height: "325px", // 不得小于325px这高度，否则控制台会警告
						border: "1px solid #ccc", // 边框，增强可见性
						padding: "10px", // 增加内边距
						...editorStyle
					}}
				/>
			</div>
		</div>
	);
};

export default RichTextEditor;
