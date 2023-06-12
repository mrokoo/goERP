<script setup lang="ts">
import { NCarousel, NStatistic, NGi, NGrid, NCard } from "naive-ui";
import { ref } from "vue";
import * as echarts from "echarts";

const chartRef = ref<HTMLElement | null>(null);

const myChart = echarts.init(chartRef);
const option = {
  xAxis: {
    type: "category",
    data: ["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"],
  },
  yAxis: {
    type: "value",
  },
  series: [
    {
      data: [820, 932, 901, 934, 1290, 1330, 1320],
      type: "line",
      smooth: true,
    },
  ],
};

option && myChart.setOption(option);
</script>

<template>
  <n-carousel show-arrow style="margin-bottom: 20px">
    <img
      class="carousel-img"
      src="https://naive-ui.oss-cn-beijing.aliyuncs.com/carousel-img/carousel1.jpeg"
    />
    <img
      class="carousel-img"
      src="https://naive-ui.oss-cn-beijing.aliyuncs.com/carousel-img/carousel2.jpeg"
    />
    <img
      class="carousel-img"
      src="https://naive-ui.oss-cn-beijing.aliyuncs.com/carousel-img/carousel3.jpeg"
    />
    <img
      class="carousel-img"
      src="https://naive-ui.oss-cn-beijing.aliyuncs.com/carousel-img/carousel4.jpeg"
    />
  </n-carousel>
  <n-grid :x-gap="12" :y-gap="8" :cols="4">
    <n-gi>
      <n-card>
        <n-statistic label="统计数据" :value="99">
          <template #suffix> / 100 </template>
        </n-statistic>
      </n-card>
    </n-gi>
    <n-gi>
      <n-card>
        <n-statistic label="统计数据" :value="99">
          <template #suffix> / 100 </template>
        </n-statistic>
      </n-card>
    </n-gi>
    <n-gi>
      <n-card>
        <n-statistic label="统计数据" :value="99">
          <template #suffix> / 100 </template>
        </n-statistic>
      </n-card>
    </n-gi>
    <n-gi>
      <n-card>
        <n-statistic label="统计数据" :value="99">
          <template #suffix> / 100 </template>
        </n-statistic>
      </n-card>
    </n-gi>
  </n-grid>

  <n-grid :x-gap="12" :y-gap="8" :cols="2">
    <n-gi>
      <n-card>
        <div ref="chartRef"></div>
      </n-card>
    </n-gi>
    <n-gi>
      <n-card> </n-card>
    </n-gi>
  </n-grid>
</template>

<style scoped>
.carousel-img {
  width: 100%;
  height: 300px;
  object-fit: cover;
}
</style>
