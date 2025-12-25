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
        <charts-people-number :data="msg_data[type - 1]" height="100%" />
      </div>
      <!-- <charts-content-number v-if="type === 4" height="20rem" /> -->
      <charts-content-number v-if="type === 4" :x-axis="xAxis" :charts-data="chartsData" height="20rem" />
    </div>
  </div>
</template>

<script setup>
  import service from '@/utils/request'
  import chartsPeopleNumber from './charts-people-numbers.vue'
  import chartsContentNumber from './charts-content-numbers.vue'
  import { ref, reactive, computed, watch} from "vue"
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

  //   [12, 22, 32, 45, 32, 78, 89, 92],
  //   [1, 2, 43, 5, 67, 78, 89, 12],
  //   [1, 2, 43, 30, 67, 78, 89, 60],
  //   // [12, 22, 32, 45, 32, 78, 89, 92]
  // ]

  // const chartsData = ref([])

  const load_ChartsData = () => {
      return service({
        url: '/z/List',
        method: 'get'
      })
    }

  const load_MsgData = () => {
      return service({
        url: '/z/Msg_Data',
        method: 'get'
      })
    }

  // const xAxis = ref([
  //   '2024-1',
  //   '2024-2',
  //   '2024-3',
  //   '2024-4',
  //   '2024-5',
  //   '2024-6',
  //   '2024-7',
  //   '2024-8',
  //   '2024-9',
  //   '2024-10',
  //   '2024-11',
  //   '2024-12'
  // ])
  // const chartsData1 = ref([12, 22, 32, 45, 32, 78, 89, 80, 85, 83, 92, 98])
  // console.log(chartsData1)
  // const chartsData = ref([])
  const chartsData = ref([0])
  const xAxis = ref([])
  const msg_data = ref([[],[],[]])

  const p = computed(() => {
    switch (d.type) {
      case 1: return `${(msg_data.value[0].reduce((a, b) => a + b, 0)/5).toFixed(0)}%`
      case 2: return `${(msg_data.value[1].reduce((a, b) => a + b, 0)/5).toFixed(0)}%`
      case 3: return `${(msg_data.value[2].reduce((a, b) => a + b, 0)/5).toFixed(0)}%`
    }
  })
  const t = computed(() => {
    switch (d.type) {
        case 1: return msg_data.value[0].reduce((a, b) => a + b, 0)
        case 2: return msg_data.value[1].reduce((a, b) => a + b, 0)
        case 3: return msg_data.value[2].reduce((a, b) => a + b, 0)
      }
    })

  watch(
    () => d.type,
    async (val) => {
      if (val === 4) {
        const ele = await load_ChartsData()
        chartsData.value = ele.data.map(item => Number(item.title))
        xAxis.value = ele.data.map(item => item.status)
        // console.log(chartsData)
      }
      else{
        const ele = await load_MsgData()
        msg_data.value[0] =(ele.data.map(item => Number(item.vistor)))
        msg_data.value[1] =(ele.data.map(item => Number(item.member)))
        msg_data.value[2] =(ele.data.map(item => Number(item.qnum)))
        console.log(msg_data)
      }
    },
    { immediate: true }
  )

</script>

<style scoped lang="scss"></style>
