<template>
  <div v-if="showWarning" class="absolute top-0 left-0 z-50 m-2 ml-2 sm:ml-4 lg:ml-4">
    <div class="flex items-center gap-2 px-2 py-1.5 sm:px-3 sm:py-2 bg-destructive text-destructive-foreground rounded-md shadow-lg">
      <svg
        xmlns="http://www.w3.org/2000/svg"
        width="14"
        height="14"
        viewBox="0 0 24 24"
        fill="none"
        stroke="currentColor"
        stroke-width="2"
        stroke-linecap="round"
        stroke-linejoin="round"
        class="flex-shrink-0 sm:w-4 sm:h-4"
      >
        <path d="m21.73 18-8-14a2 2 0 0 0-3.48 0l-8 14A2 2 0 0 0 4 21h16a2 2 0 0 0 1.73-3Z" />
        <path d="M12 9v4" />
        <path d="m12 17 .01 0" />
      </svg>
      <span class="text-xs font-medium whitespace-nowrap">
        Error System
      </span>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useSystemStatusQuery } from '@/composables/use-system'
import { SystemStatus } from '@/types/system-info'

// Sử dụng options tối ưu hóa để giảm network requests và ẩn progress bar
const { data: systemStatus} = useSystemStatusQuery({
  refetchInterval: 3000,  //3s
  axiosOpts: {
    doNotShowLoading: true, // Ẩn progress bar
  },
})

const showWarning = computed(() => {
  return systemStatus.value?.status === SystemStatus.ERROR
})
</script>

<style scoped>
@media (max-width: 639px) {
  .absolute {
    left: 50px !important; 
  }
}

@media (min-width: 640px) and (max-width: 1023px) {
  .absolute {
    left: 16px !important;
  }
}

@media (min-width: 1024px) {
  .absolute {
    left: 16px !important;
  }
}
</style>
