export interface SystemInfo {
  localIp: string
  cpuUsage: number
  memoryUsage: number
  totalMemory: number
  uptime: number
}

export enum SystemStatus {
  NORMAL = 'NORMAL',
  ERROR = 'ERROR',
}

export interface SystemStatusResponse {
  status: SystemStatus;
}
