<!--
    @auther: bypanghu<bypanghu@163.com>
    @date: 2024/5/8
!-->

<template>
  <div class="">
    <div class="flex items-center justify-between mb-2">
      <div v-if="title" class="font-bold">
        {{ title }}
      </div>
      <slot v-else name="title" />
    </div>
    <div class="w-full relative">
      <div v-if="type !== 4">
        <div class="mt-4 text-gray-600 text-3xl font-mono">
          <!-- <el-statistic v-if="type === 1" :value="1000" /> -->
          <el-statistic :value="t" />
        </div>
        <div class="mt-2 text-green-600 text-sm font-bold font-mono">
          {{ p }} <el-icon><TopRight /></el-icon>
        </div>
      </div>
      <div class="absolute top-0 right-2 w-[50%] h-20">
        <!-- <charts-people-number v-if="type === 1" :data="data[0]" height="100%" />
        <charts-people-number v-if="type === 2" :data="data[1]" height="100%" />
        <charts-people-number v-if="type === 3" :data="data[2]" height="100%" /> -->
        <charts-people-number :data="data[type - 1]" height="100%" />
      </div>
      <!-- <charts-content-number v-if="type === 4" height="20rem" /> -->
      <charts-content-number v-if="type === 4" :x-axis="xAxis" :charts-data="chartsData" height="20rem" />
    </div>
  </div>
</template>

<script setup>
  import chartsPeopleNumber from './charts-people-numbers.vue'
  import chartsContentNumber from './charts-content-numbers.vue'
  import { ref, reactive, computed } from "vue"
  const d = defineProps({
    type: {
      type: Number,
      default: 1
    },
    title: {
      type: String,
      default: ''
    }
  })
  // const p = "+30%"
  // console.log(d.type)
  const p = computed(() => {
    switch (d.type) {
      case 1: return "+30%"
      case 2: return "+60%"
      case 3: return "+80%"
    }
  })
  const t = computed(() => {
    switch (d.type) {
        case 1: return 1000
        case 2: return 2000
        case 3: return 3000
      }
    })

  const data = [
    [12, 22, 32, 45, 32, 78, 89, 92],
    [1, 2, 43, 5, 67, 78, 89, 12],
    [1, 2, 43, 30, 67, 78, 89, 60],
    // [12, 22, 32, 45, 32, 78, 89, 92]
  ]

  const xAxis = reactive([
    '2024-1',
    '2024-2',
    '2024-3',
    '2024-4',
    '2024-5',
    '2024-6',
    '2024-7',
    '2024-8',
    '2024-9',
    '2024-10',
    '2024-11',
    '2024-12'
  ])
  const chartsData = reactive([12, 22, 32, 45, 32, 78, 89, 80, 85, 83, 92, 98])
</script>

<style scoped lang="scss"></style>
