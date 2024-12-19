import { RequestData } from "@ant-design/pro-components";
import { ResPage } from "./request/interface";

export function saveExcelBlob(fileName: string, blobData: BlobPart) {
	var now = new Date();
	var year = now.getFullYear(); // 年
	var month = now.getMonth() + 1; // 月
	var day = now.getDate(); // 日
	var hh = now.getHours(); // 时
	var mm = now.getMinutes(); // 分
	var ss = now.getSeconds(); // 秒
	fileName = fileName + "_" + year + "_" + month + "_" + day + "_" + hh + "_" + mm + "_" + ss;

	const blob = new Blob([blobData], { type: "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet" });
	// 创建一个 URL 对象来下载 Blob 数据
	const url = window.URL.createObjectURL(blob);

	// 创建一个虚拟的 <a> 标签
	const link = document.createElement("a");
	link.href = url;

	// 设置下载的文件名
	link.download = fileName + ".xlsx";

	// 模拟点击该链接，触发下载
	link.click();

	// 释放 URL 对象
	window.URL.revokeObjectURL(url);
}

export function saveZipBlob(fileName: string, blobData: BlobPart) {
	const blob = new Blob([blobData]);
	// 创建一个 URL 对象来下载 Blob 数据
	const url = window.URL.createObjectURL(blob);

	// 创建一个虚拟的 <a> 标签
	const link = document.createElement("a");
	link.href = url;

	// 设置下载的文件名
	link.download = fileName + ".zip";

	// 模拟点击该链接，触发下载
	link.click();

	// 释放 URL 对象
	window.URL.revokeObjectURL(url);
}

export function formatDataForProTable<T>(data: ResPage<T>): Partial<RequestData<T>> {
	let list: T[] = [];
	let total = 0;
	if (data && data.list) {
		list = data.list;
		total = data.count;
	}
	return {
		success: true,
		data: list,
		total: total
	};
}

export function formatDataListForProTable<T>(data: T[]): Partial<RequestData<T>> {
	if (!data) {
		data = [];
	}
	let total = data.length;

	return {
		success: true,
		data: data,
		total: total
	};
}
