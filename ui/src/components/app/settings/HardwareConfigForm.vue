<script setup lang="ts">
import type { HardwareConfig } from '@/types/config'
import { useQueryClient } from '@tanstack/vue-query'
import { toTypedSchema } from '@vee-validate/zod'
import { Loader } from 'lucide-vue-next'
import { useForm } from 'vee-validate'
import { z } from 'zod'
import { Button } from '@/components/ui/button'
import { FormControl, FormField, FormItem, FormLabel, FormMessage } from '@/components/ui/form'
import { Input } from '@/components/ui/input'
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '@/components/ui/select'
import { Switch } from '@/components/ui/switch'
import { HARDWARE_CONFIG_QUERY_KEY, useHardwareConfigMutation } from '@/composables/use-config'
import { useListAvailableSerialPortsQuery } from '@/composables/use-peripheral'

interface Props {
  initialValues: HardwareConfig
}

const props = defineProps<Props>()
const BAUD_RATES = [9600, 19200, 38400, 57600, 115200]

const serialConfigSchema = z.object({
  port: z.string().min(1, 'Port is required'),
  baudRate: z.number().int().positive('Baud rate must be positive'),
  parity: z.enum(['NONE', 'EVEN', 'ODD']).default('NONE'),
  dataBits: z.union([z.literal(5), z.literal(6), z.literal(7), z.literal(8)]).default(8),
  stopBits: z.union([z.literal(1), z.literal(1.5), z.literal(2)]).default(1),
  readTimeout: z.number().int().nonnegative('Read timeout must be non-negative'),
})

const ledConfigSchema = z.object({
  pin: z.string().min(1, 'Pin is required'),
})

const hardwareConfigSchema = z.object({
  esp: z.object({
    serial: serialConfigSchema,
    enableAck: z.boolean().default(false),
    commandAckTimeout: z.number().int().nonnegative('Command ack timeout must be non-negative'),
  }),
  pic: z.object({
    serial: serialConfigSchema,
    enableAck: z.boolean().default(false),
    commandAckTimeout: z.number().int().nonnegative('Command ack timeout must be non-negative'),
  }),
  leds: z.object({
    system: ledConfigSchema,
    alert: ledConfigSchema,
  }),
    system: ledConfigSchema.default({ pin: '' }),
    alert: ledConfigSchema.default({ pin: '' }),
  }).default({ system: { pin: '' }, alert: { pin: '' } }),
}).superRefine((data, ctx) => {
  if (data.esp.serial.port === data.pic.serial.port) {
    ctx.addIssue({
      code: z.ZodIssueCode.custom,
      message: 'ESP and PIC cannot use the same port',
      path: ['esp.serial.port'],
    })
    ctx.addIssue({
      code: z.ZodIssueCode.custom,
      message: 'ESP and PIC cannot use the same port',
      path: ['pic.serial.port'],
    })
  }

  if (data.leds.system.pin === data.leds.alert.pin) {
    ctx.addIssue({
      code: z.ZodIssueCode.custom,
      message: 'System and alert LEDs cannot use the same pin',
      path: ['leds.system.pin'],
    })
    ctx.addIssue({
      code: z.ZodIssueCode.custom,
      message: 'System and alert LEDs cannot use the same pin',
      path: ['leds.alert.pin'],
    })
  }
})

const queryClient = useQueryClient()
const { mutate, isPending } = useHardwareConfigMutation()
const { data: ports, refetch: refetchPorts } = useListAvailableSerialPortsQuery({ doNotShowLoading: true })
const form = useForm({
  validationSchema: toTypedSchema(hardwareConfigSchema),
  initialValues: props.initialValues,
})


const onSubmit = form.handleSubmit((values) => {
  mutate(values, {
    onSuccess: () => {
      queryClient.setQueryData([HARDWARE_CONFIG_QUERY_KEY], values)
      notification.success('Hardware configuration updated successfully!')
    },
    onError: () => {
      notification.error('Failed to update hardware configuration')
    },
  })
})

function fetchPorts(newValue: boolean) {
  if (newValue) {
    refetchPorts()
  }
}
</script>

<template>
  <form class="flex flex-col w-full space-y-6" @submit="onSubmit">
    <div class="grid grid-cols-1 gap-8">
      <!-- ESP Controller Section -->
      <div class="space-y-3">
        <h4 class="text-lg font-medium tracking-tight">
          ESP Controller
        </h4>

        <div class="px-4 space-y-6">
          <div class="space-y-6">
            <h4 class="pb-2 text-lg font-medium border-b">
              Serial Configuration
            </h4>

            <!-- Port and Timeout in same row -->
            <div class="grid grid-cols-1 gap-4 md:grid-cols-2">
              <FormField v-slot="{ componentField, value }" name="esp.serial.port">
                <FormItem>
                  <FormLabel>Port</FormLabel>
                  <Select v-bind="componentField" required @update:open="fetchPorts">
                    <FormControl>
                      <SelectTrigger :disabled="isPending">
                        <SelectValue :placeholder="value" />
                      </SelectTrigger>
                    </FormControl>
                    <SelectContent>
                      <SelectItem value="NONEESP">
                        None
                      </SelectItem>
                      <template v-if="ports">
                        <SelectItem v-for="port in ports" :key="port.port" :value="port.port">
                          {{ port.port }}
                        </SelectItem>
                      </template>
                    </SelectContent>
                  </Select>
                  <FormMessage />
                </FormItem>
              </FormField>

              <FormField v-slot="{ componentField }" name="esp.serial.readTimeout">
                <FormItem>
                  <FormLabel>Read Timeout (s)</FormLabel>
                  <FormControl>
                    <Input v-bind="componentField" type="number" :disabled="isPending" placeholder="e.g. 1" />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              </FormField>
            </div>

            <!-- Other settings in a 4-column grid -->
            <div class="grid grid-cols-1 gap-4 md:grid-cols-2 lg:grid-cols-4">
              <FormField v-slot="{ componentField }" name="esp.serial.baudRate">
                <FormItem>
                  <FormLabel>Baud Rate</FormLabel>
                  <Select v-bind="componentField">
                    <FormControl>
                      <SelectTrigger :disabled="isPending">
                        <SelectValue placeholder="Select baud rate" />
                      </SelectTrigger>
                    </FormControl>
                    <SelectContent>
                      <SelectItem v-for="baudRate in BAUD_RATES" :key="baudRate" :value="baudRate">
                        {{ baudRate }}
                      </SelectItem>
                    </SelectContent>
                  </Select>
                  <FormMessage />
                </FormItem>
              </FormField>

              <FormField v-slot="{ componentField }" name="esp.serial.parity">
                <FormItem>
                  <FormLabel>Parity</FormLabel>
                  <Select v-bind="componentField">
                    <FormControl>
                      <SelectTrigger :disabled="isPending">
                        <SelectValue placeholder="Select parity" />
                      </SelectTrigger>
                    </FormControl>
                    <SelectContent>
                      <SelectItem value="NONE">
                        None
                      </SelectItem>
                      <SelectItem value="EVEN">
                        Even
                      </SelectItem>
                      <SelectItem value="ODD">
                        Odd
                      </SelectItem>
                    </SelectContent>
                  </Select>
                  <FormMessage />
                </FormItem>
              </FormField>

              <FormField v-slot="{ componentField }" name="esp.serial.dataBits">
                <FormItem>
                  <FormLabel>Data Bits</FormLabel>
                  <Select v-bind="componentField">
                    <FormControl>
                      <SelectTrigger :disabled="isPending">
                        <SelectValue placeholder="Select data bits" />
                      </SelectTrigger>
                    </FormControl>
                    <SelectContent>
                      <SelectItem :value="5">
                        5
                      </SelectItem>
                      <SelectItem :value="6">
                        6
                      </SelectItem>
                      <SelectItem :value="7">
                        7
                      </SelectItem>
                      <SelectItem :value="8">
                        8
                      </SelectItem>
                    </SelectContent>
                  </Select>
                  <FormMessage />
                </FormItem>
              </FormField>

              <FormField v-slot="{ componentField }" name="esp.serial.stopBits">
                <FormItem>
                  <FormLabel>Stop Bits</FormLabel>
                  <Select v-bind="componentField">
                    <FormControl>
                      <SelectTrigger :disabled="isPending">
                        <SelectValue placeholder="Select stop bits" />
                      </SelectTrigger>
                    </FormControl>
                    <SelectContent>
                      <SelectItem :value="1">
                        1
                      </SelectItem>
                      <SelectItem :value="1.5">
                        1.5
                      </SelectItem>
                      <SelectItem :value="2">
                        2
                      </SelectItem>
                    </SelectContent>
                  </Select>
                  <FormMessage />
                </FormItem>
              </FormField>
            </div>
          </div>

          <div class="space-y-6">
            <h4 class="pb-2 text-lg font-medium border-b">
              Command Acknowledgment
            </h4>

            <!-- Port and Timeout in same row -->
            <div class="grid grid-cols-1 gap-4 md:grid-cols-2">
              <div class="space-y-6">
                <FormField v-slot="{ value, handleChange }" name="esp.enableAck">
                  <FormItem class="flex flex-row items-center justify-between p-4 border rounded-lg">
                    <FormLabel>Enable Ack</FormLabel>
                    <FormControl>
                      <Switch
                        :model-value="value"
                        :disabled="isPending"
                        aria-readonly
                        @update:model-value="handleChange"
                      />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                </FormField>
                <FormField v-slot="{ componentField }" name="esp.commandAckTimeout">
                  <FormItem>
                    <FormLabel>Ack Timeout (ms)</FormLabel>
                    <FormControl>
                      <Input v-bind="componentField" type="number" :disabled="isPending" placeholder="e.g. 1" />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                </FormField>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- PIC Controller Section -->
      <div class="space-y-3">
        <h4 class="text-lg font-medium tracking-tight">
          PIC Controller
        </h4>
        <div class="px-4 space-y-6">
          <div class="space-y-6">
            <h4 class="pb-2 text-lg font-medium border-b">
              Serial Configuration
            </h4>

            <!-- Port and Timeout in same row -->
            <div class="grid grid-cols-1 gap-4 md:grid-cols-2">
              <FormField v-slot="{ componentField, value }" name="pic.serial.port">
                <FormItem>
                  <FormLabel>Port</FormLabel>
                  <Select v-bind="componentField" required @update:open="fetchPorts">
                    <FormControl>
                      <SelectTrigger :disabled="isPending">
                        <SelectValue :placeholder="value" />
                      </SelectTrigger>
                    </FormControl>
                    <SelectContent>
                      <SelectItem value="NONEPIC">
                        None
                      </SelectItem>
                      <template v-if="ports">
                        <SelectItem v-for="port in ports" :key="port.port" :value="port.port">
                          {{ port.port }}
                        </SelectItem>
                      </template>
                    </SelectContent>
                  </Select>
                  <FormMessage />
                </FormItem>
              </FormField>

              <FormField v-slot="{ componentField }" name="pic.serial.readTimeout">
                <FormItem>
                  <FormLabel>Read Timeout (s)</FormLabel>
                  <FormControl>
                    <Input v-bind="componentField" type="number" :disabled="isPending" placeholder="e.g. 1" />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              </FormField>
            </div>

            <!-- Other settings in a 4-column grid -->
            <div class="grid grid-cols-1 gap-4 md:grid-cols-2 lg:grid-cols-4">
              <FormField v-slot="{ componentField }" name="pic.serial.baudRate">
                <FormItem>
                  <FormLabel>Baud Rate</FormLabel>
                  <Select v-bind="componentField">
                    <FormControl>
                      <SelectTrigger :disabled="isPending">
                        <SelectValue placeholder="Select baud rate" />
                      </SelectTrigger>
                    </FormControl>
                    <SelectContent>
                      <SelectItem v-for="baudRate in BAUD_RATES" :key="baudRate" :value="baudRate">
                        {{ baudRate }}
                      </SelectItem>
                    </SelectContent>
                  </Select>
                  <FormMessage />
                </FormItem>
              </FormField>

              <FormField v-slot="{ componentField }" name="pic.serial.parity">
                <FormItem>
                  <FormLabel>Parity</FormLabel>
                  <Select v-bind="componentField">
                    <FormControl>
                      <SelectTrigger :disabled="isPending">
                        <SelectValue placeholder="Select parity" />
                      </SelectTrigger>
                    </FormControl>
                    <SelectContent>
                      <SelectItem value="NONE">
                        None
                      </SelectItem>
                      <SelectItem value="EVEN">
                        Even
                      </SelectItem>
                      <SelectItem value="ODD">
                        Odd
                      </SelectItem>
                    </SelectContent>
                  </Select>
                  <FormMessage />
                </FormItem>
              </FormField>

              <FormField v-slot="{ componentField }" name="pic.serial.dataBits">
                <FormItem>
                  <FormLabel>Data Bits</FormLabel>
                  <Select v-bind="componentField">
                    <FormControl>
                      <SelectTrigger :disabled="isPending">
                        <SelectValue placeholder="Select data bits" />
                      </SelectTrigger>
                    </FormControl>
                    <SelectContent>
                      <SelectItem :value="5">
                        5
                      </SelectItem>
                      <SelectItem :value="6">
                        6
                      </SelectItem>
                      <SelectItem :value="7">
                        7
                      </SelectItem>
                      <SelectItem :value="8">
                        8
                      </SelectItem>
                    </SelectContent>
                  </Select>
                  <FormMessage />
                </FormItem>
              </FormField>

              <FormField v-slot="{ componentField }" name="pic.serial.stopBits">
                <FormItem>
                  <FormLabel>Stop Bits</FormLabel>
                  <Select v-bind="componentField">
                    <FormControl>
                      <SelectTrigger :disabled="isPending">
                        <SelectValue placeholder="Select stop bits" />
                      </SelectTrigger>
                    </FormControl>
                    <SelectContent>
                      <SelectItem :value="1">
                        1
                      </SelectItem>
                      <SelectItem :value="1.5">
                        1.5
                      </SelectItem>
                      <SelectItem :value="2">
                        2
                      </SelectItem>
                    </SelectContent>
                  </Select>
                  <FormMessage />
                </FormItem>
              </FormField>
            </div>
          </div>

          <div class="space-y-6">
            <h3 class="pb-2 text-lg font-medium border-b">
              Command Acknowledgment
            </h3>

            <div class="grid grid-cols-1 gap-4 md:grid-cols-2">
              <div class="space-y-6">
                <FormField v-slot="{ value, handleChange }" name="pic.enableAck">
                  <FormItem class="flex flex-row items-center justify-between p-4 border rounded-lg">
                    <FormLabel>Enable Ack</FormLabel>
                    <FormControl>
                      <Switch
                        :model-value="value"
                        :disabled="isPending"
                        aria-readonly
                        @update:model-value="handleChange"
                      />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                </FormField>
                <FormField v-slot="{ componentField }" name="pic.commandAckTimeout">
                  <FormItem>
                    <FormLabel>Ack Timeout (ms)</FormLabel>
                    <FormControl>
                      <Input v-bind="componentField" type="number" :disabled="isPending" placeholder="e.g. 1" />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                </FormField>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- LEDs Configuration Section -->
      <div class="space-y-3">
        <h4 class="text-lg font-medium tracking-tight">
          LEDs Configuration
        </h4>
        <div class="px-4 space-y-6">
          <div class="space-y-6">
            <h4 class="pb-2 text-lg font-medium border-b">
              LED Pins
            </h4>

            <div class="grid grid-cols-1 gap-4 md:grid-cols-2">
              <FormField v-slot="{ componentField }" name="leds.system.pin">
                <FormItem>
                  <FormLabel>System LED Pin</FormLabel>
                  <FormControl>
                    <Input v-bind="componentField" :disabled="isPending" placeholder="e.g. GPIO2" />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              </FormField>

              <FormField v-slot="{ componentField }" name="leds.alert.pin">
                <FormItem>
                  <FormLabel>Alert LED Pin</FormLabel>
                  <FormControl>
                    <Input v-bind="componentField" :disabled="isPending" placeholder="e.g. GPIO2" />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              </FormField>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- LEDs Section -->
    <div class="space-y-3">
      <h4 class="text-lg font-medium tracking-tight">
        LED Configuration
      </h4>
      <div class="px-4 space-y-6">
        <div class="grid grid-cols-1 gap-4 md:grid-cols-2">
          <FormField v-slot="{ componentField }" name="leds.system.pin">
            <FormItem>
              <FormLabel>System LED Pin</FormLabel>
              <FormControl>
                <Input v-bind="componentField" :disabled="isPending" placeholder="e.g. GPIO2" />
              </FormControl>
              <FormMessage />
            </FormItem>
          </FormField>

          <FormField v-slot="{ componentField }" name="leds.alert.pin">
            <FormItem>
              <FormLabel>Alert LED Pin</FormLabel>
              <FormControl>
                <Input v-bind="componentField" :disabled="isPending" placeholder="e.g. GPIO2" />
              </FormControl>
              <FormMessage />
            </FormItem>
          </FormField>
        </div>
      </div>
    </div>

    <div>
      <Button type="submit" :disabled="isPending">
        <Loader v-if="isPending" class="w-4 h-4 mr-2 animate-spin" />
        Save
      </Button>
    </div>
  </form>
</template>
