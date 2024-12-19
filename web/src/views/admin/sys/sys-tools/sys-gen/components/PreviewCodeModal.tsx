import { previewTableApi } from "@/api/admin/sys/sys-tools/sys-gen";
import LoadingButton from "@/components/LoadingButton";
import { ResultEnum } from "@/enums/httpEnum";
import { message } from "@/hooks/useMessage";
import { Modal, Tabs } from "antd";
import { forwardRef, useImperativeHandle, useState } from "react";
import { Light as SyntaxHighlighter } from "react-syntax-highlighter";
import { atomOneDark } from "react-syntax-highlighter/dist/esm/styles/hljs";

export interface PreviewCodeModalRef {
	showPreviewCodeModal: (id: number) => void;
}

interface ModalProps {
	onConfirm: () => void;
}

const PreviewCodeModal = forwardRef<PreviewCodeModalRef, ModalProps>(({ onConfirm }, ref) => {
	const defaultActiveTab = "api.go";
	const [isModalVisible, setIsModalVisible] = useState(false);
	const [activeTab, setActiveTab] = useState(defaultActiveTab);
	const [codes, setCodes] = useState<Map<string, string>>(new Map());

	useImperativeHandle(ref, () => ({
		async showPreviewCodeModal(id: number) {
			const { data, msg, code } = await previewTableApi(id);
			if (code !== ResultEnum.SUCCESS) {
				message.error(msg);
				return;
			}
			setCodes(prevMap => {
				const newMap = new Map(prevMap);
				data.forEach(item => {
					newMap.set(item.name, item.content);
				});
				return newMap; // 返回新的 Map
			});
			setIsModalVisible(true);
		}
	}));

	const reset = (done?: () => void) => {
		setIsModalVisible(false);
		setActiveTab(defaultActiveTab);
		setCodes(new Map());
		if (done) {
			done();
		}
	};

	return (
		<>
			<Modal
				title="代码预览"
				getContainer={false}
				open={isModalVisible}
				onCancel={() => reset()}
				destroyOnClose
				width="90vw"
				styles={{
					body: {
						paddingBottom: "60px", // 预留空间给底部按钮
						height: "70vh", // 设置固定高度
						overflow: "hidden" // 启用滚动条
					}
				}}
				footer={null} // 移除默认的 footer
			>
				<div
					style={{
						display: "flex",
						flexDirection: "column",
						height: "100%"
					}}
				>
					{/* Tab 头部，固定位置 */}
					<div
						style={{
							position: "sticky", // 固定位置
							top: 0,
							zIndex: 1 // 确保在内容之上
						}}
					>
						<Tabs
							activeKey={activeTab}
							onChange={key => setActiveTab(key)}
							centered
							tabBarStyle={{ marginBottom: 0 }}
							items={Array.from(codes.entries()).map(([key, code]) => ({
								key,
								label: key,
								children: null // 移除内容部分，放在外部渲染
							}))}
						/>
					</div>

					{/* Tab 内容区域，允许滚动 */}
					<div
						style={{
							flex: 1,
							overflow: "auto", // 内容区域滚动
							padding: "16px"
						}}
					>
						<SyntaxHighlighter
							style={atomOneDark}
							customStyle={{
								padding: "10px",
								borderRadius: "4px",
								fontSize: "14px"
							}}
						>
							{codes.get(activeTab) || "代码加载失败或为空"}
						</SyntaxHighlighter>
					</div>
				</div>
				{/* 固定底部按钮 */}
				<div
					style={{
						position: "absolute",
						bottom: 0,
						left: 0,
						width: "100%",
						textAlign: "right",
						padding: "20px 16px"
					}}
				>
					<LoadingButton key="confirm" type="primary" onClick={done => reset(done)}>
						确定
					</LoadingButton>
				</div>
			</Modal>
		</>
	);
});

export default PreviewCodeModal;
