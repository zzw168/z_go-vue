<!--
    @auther: bypanghu<bypanghu@163.com>
    @date: 2024/5/8
!-->

<template>
  <el-carousel>
    <el-carousel-item
      v-for="(item,index) in banners"
      :key="index"
      @click="openLink(item.link)"
      class="w-full h-auto cursor-pointer lg:h-40"
    >
      <div class="flex justify-center">
        <el-image
          :src="item.img"
          fit="fill"
          class="w-4/5 h-auto"
          :href="item.link"
        />
      </div>
    </el-carousel-item>
  </el-carousel>
</template>

<script setup>
  import banner from '@/assets/banner.png'
  import banner1 from '@/assets/banner1.png'
  import banner2 from '@/assets/banner2.png'

  const openLink = (link) => {
    window.open(link, '_blank')
  }

  const banners = [
    {
      img: banner,
      link: 'https://v3.cn.vuejs.org/guide/introduction.html'
    },
    {
      img: banner2,
      link: 'https://v3.cn.vuejs.org/guide/introduction.html'
    },
    {
      img: banner1,
      link: 'https://v3.cn.vuejs.org/guide/introduction.html'
    }
  ]
</script>

<style scoped lang="scss"></style>
