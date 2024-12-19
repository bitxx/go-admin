import { getMonitorApi, Monitor } from "@/api/admin/sys/sys-tools/sys-monitor";
import { ResultEnum } from "@/enums/httpEnum";
import { Card, Divider, Layout, Progress, Table } from "antd";
import type { ColumnsType } from "antd/es/table";
import React, { useEffect, useState } from "react";
import "./index.less";

interface DiskStatus {
	key: string;
	path: string;
	fs: string;
	total: string;
	available: string;
	used: string;
	percentage: string;
}

const App: React.FC = () => {
	const [monitor, setMonitor] = useState<Monitor>();
	const [diskStatus, setDiskStatus] = useState<DiskStatus[]>();
	useEffect(() => {
		const getServerMonitor = async () => {
			const { data, code } = await getMonitorApi();
			if (code === ResultEnum.SUCCESS) {
				// const cpuName = data?.cpu.cpuInfo[0].modelName.split("@ ");
				// if (cpuName.length > 1) {
				//   data.cpu.cpuInfo[0].modelName = cpuName[0];
				// }
				const dk: DiskStatus[] = [];

				data.diskList.forEach((disk, index) => {
					let tmp = {
						key: String(index),
						path: disk.path,
						fs: disk.fstype,
						total: disk.total + "M",
						available: disk.free + "M",
						used: disk.used + "M",
						percentage: disk.usedPercent + "%"
					};
					dk.push(tmp);
				});

				setDiskStatus(dk);

				setMonitor(data);
			}
		};
		getServerMonitor();
	}, []);

	// 表格列定义
	const columns: ColumnsType<any> = [
		{ title: "盘符路径", dataIndex: "path", key: "path" },
		{ title: "文件系统", dataIndex: "fs", key: "fs" },
		{ title: "总大小", dataIndex: "total", key: "total" },
		{ title: "可用大小", dataIndex: "available", key: "available" },
		{ title: "已用大小", dataIndex: "used", key: "used" },
		{ title: "已用百分比", dataIndex: "percentage", key: "percentage" }
	];

	return (
		<Layout>
			{/* 性能监控 */}
			{monitor?.cpu && (
				<div className="monitor-item-cpu">
					<Card title="CPU使用率" style={{ flex: 1 }}>
						<div style={{ textAlign: "center" }}>
							<Progress type="circle" percent={monitor?.cpu.Percent} format={percent => `${percent}%`} />
							<Divider />
							<div style={{ textAlign: "left" }}>
								<div className="monitor-item">
									<span>CPU主频:</span>
									<span>{monitor?.cpu.cpuInfo[0].modelName}</span>
								</div>
								<Divider />
								<div className="monitor-item">
									<span>核心数:</span>
									<span>{monitor?.cpu.cpuNum}</span>
								</div>
							</div>
						</div>
					</Card>
					<Card title="内存使用率" style={{ flex: 1 }}>
						<div style={{ textAlign: "center" }}>
							<Progress type="circle" percent={monitor?.mem.usage} format={percent => `${percent}%`} />
							<Divider />
							<div style={{ textAlign: "left" }}>
								<div className="monitor-item">
									<span>总内存:</span>
									<span>{monitor?.mem.total}G</span>
								</div>
								<Divider />
								<div className="monitor-item">
									<span>已用内存:</span>
									<span>{monitor?.mem.used}G</span>
								</div>
							</div>
						</div>
					</Card>
					<Card title="磁盘信息" style={{ flex: 1 }}>
						<div style={{ textAlign: "center" }}>
							<Progress
								type="circle"
								percent={
									monitor?.disk.total && monitor?.disk.total > 0
										? Number((((monitor?.disk.total - monitor?.disk.free) / monitor?.disk.total) * 100).toFixed(2))
										: 0
								}
								format={percent => `${percent}%`}
							/>
							<Divider />
							<div style={{ textAlign: "left" }}>
								<div className="monitor-item">
									<span>总磁盘:</span>
									<span>{monitor?.disk.total}G</span>
								</div>
								<Divider />
								<div className="monitor-item">
									<span>已用磁盘:</span>
									{monitor?.disk.total && monitor?.disk.free ? `${(monitor?.disk.total - monitor?.disk.free).toFixed(2)}G` : "0G"}
								</div>
							</div>
						</div>
					</Card>
				</div>
			)}
			{/* Go运行环境 */}
			{monitor?.os && (
				<Card title="Go运行环境" style={{ marginBottom: 10 }}>
					<div>
						<div className="monitor-item-go-version">
							<span>GO 版本:</span>
							<span>{monitor?.os.version}</span>
						</div>
						<Divider />
						<div className="monitor-item-go-version">
							<span>Goroutine:</span>
							<span>{monitor?.os.numGoroutine}</span>
						</div>
						<Divider />
						<div className="monitor-item-go-version">
							<span>项目地址:</span>
							<span>{monitor?.os.projectDir}</span>
						</div>
					</div>
				</Card>
			)}
			{/* 系统信息 */}
			{monitor?.os && (
				<Card title="服务器信息" style={{ marginBottom: 10 }}>
					<div>
						<div className="monitor-item">
							<span>主机名称:</span>
							<span>{monitor?.os.hostName}</span>
						</div>
						<Divider />
						<div className="monitor-item">
							<span>操作系统:</span>
							<span>{monitor?.os.goOs}</span>
						</div>
						<Divider />
						<div className="monitor-item">
							<span>服务器IP:</span>
							<span>{monitor?.os.ip}</span>
						</div>
						<Divider />
						<div className="monitor-item">
							<span>系统架构:</span>
							<span>{monitor?.os.arch}</span>
						</div>
						<Divider />
						{monitor?.cpu && monitor?.cpu.cpuInfo.length > 0 && (
							<>
								<div className="monitor-item">
									<span>CPU:</span>
									<span>{monitor?.cpu.cpuInfo[0].modelName}</span>
								</div>
								<Divider />
							</>
						)}

						<div className="monitor-item">
							<span>当前时间:</span>
							<span>{monitor?.os.time}</span>
						</div>
					</div>
				</Card>
			)}
			{/* 磁盘状态 */}
			<Card title="磁盘状态">
				<Table columns={columns} dataSource={diskStatus} pagination={false} />
			</Card>
		</Layout>
	);
};

export default App;
