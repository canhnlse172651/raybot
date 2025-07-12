import type { AxiosRequestConfig } from 'axios'
import { useMutation, useQuery } from '@tanstack/vue-query'
import systemAPI from '@/api/system'

export const SYSTEM_INFO_QUERY_KEY = 'system-info'
export const SYSTEM_STATUS_QUERY_KEY = 'system-status'

export function useSystemRebootMutation() {
  return useMutation({
    mutationFn: systemAPI.reboot,
  })
}

export function useSystemStopEmergencyMutation() {
  return useMutation({
    mutationFn: systemAPI.stopEmergency,
  })
}

export function useSystemGetInfoQuery(opts?: {
  axiosOpts?: Partial<AxiosRequestConfig>
  refetchInterval?: number
}) {
  return useQuery({
    queryKey: [SYSTEM_INFO_QUERY_KEY],
    queryFn: () => systemAPI.getInfo(opts?.axiosOpts),
    refetchInterval: opts?.refetchInterval,
  })
}

export function useSystemStatusQuery(opts?: {
  axiosOpts?: Partial<AxiosRequestConfig>
  refetchInterval?: number
}) {
  return useQuery({
    queryKey: [SYSTEM_STATUS_QUERY_KEY],
    queryFn: () => systemAPI.getStatus({
      ...opts?.axiosOpts,
      doNotShowLoading: true, // Ẩn progress bar cho system status
    }),
    refetchInterval: opts?.refetchInterval || 10000, // Tăng lên 10 giây
    refetchOnWindowFocus: false, // Không refetch khi focus window
    refetchOnMount: true, // Chỉ fetch khi mount
    staleTime: 5000, // Data được coi là fresh trong 5 giây
  })
}
