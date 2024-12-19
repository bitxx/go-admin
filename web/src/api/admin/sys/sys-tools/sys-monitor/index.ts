import request from "@/utils/request";

export interface Monitor {
  cpu: Cpu;
  disk: Disk;
  diskList: DiskList[];
  mem: Mem;
  os: Os;
}

export interface CpuInfo {
  cpu: number;
  vendorId: string;
  family: string;
  model: string;
  stepping: number;
  physicalId: string;
  coreId: string;
  cores: number;
  modelName: string;
  mhz: number;
  cacheSize: number;
  flags?: null;
  microcode: string;
}
export interface Cpu {
  Percent: number;
  cpuInfo: CpuInfo[];
  cpuNum: number;
  cpus: number[];
}
export interface Disk {
  free: number;
  total: number;
}
export interface DiskList {
  path: string;
  fstype: string;
  total: number;
  free: number;
  used: number;
  usedPercent: number;
  inodesTotal: number;
  inodesUsed: number;
  inodesFree: number;
  inodesUsedPercent: number;
}
export interface Mem {
  free: number;
  total: number;
  usage: number;
  used: number;
}
export interface Os {
  arch: string;
  compiler: string;
  goOs: string;
  hostName: string;
  ip: string;
  mem: number;
  numGoroutine: number;
  projectDir: string;
  time: string;
  version: string;
}

export const getMonitorApi = () => {
  return request.get<Monitor>(`/admin-api/v1/admin/sys/sys-monitor`);
};

export const getMonitorPingApi = () => {
  return request.get<object>(`/admin-api/v1/admin/sys/sys-monitor/ping`);
};
