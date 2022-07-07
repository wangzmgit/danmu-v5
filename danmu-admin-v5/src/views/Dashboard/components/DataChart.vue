<template>
    <div class="echarts-box">
        <div class="echart" id="echart"></div>
    </div>
</template>

<script>
import { ref } from "vue";
import * as echarts from "echarts";
import theme from './theme.json'
import { onMounted, onUnmounted } from 'vue';
export default {
    props: {
        data: {
            type: Object
        }
    },
    setup(props) {
        const echart = echarts;
        const date = ref([]);
        const newUser = ref([]);
        const newVideo = ref([]);

        const initChart = () => {
            echart.registerTheme('westeros', theme);//注册主题
            let chart = echart.init(document.getElementById("echart"), "westeros");
            chart.setOption({
                legend: {
                    data: ['新增用户', '新增视频']
                },
                xAxis: {
                    type: "category",
                    data: date.value
                },
                tooltip: {
                    trigger: "axis"
                },
                yAxis: {
                    type: "value"
                },
                series: [
                    {
                        name: '新增用户',
                        data: newUser.value,
                        type: "line",
                        smooth: true
                    },
                    {
                        name: '新增视频',
                        data: newVideo.value,
                        type: "line",
                        smooth: true
                    }
                ]
            });
            window.onresize = function () {
                //自适应大小
                chart.resize();
            };
        }

        onMounted(() => {
            props.data.forEach(item => {
                date.value.push(item.date);
                newUser.value.push(item.user);
                newVideo.value.push(item.video);
            });
            initChart();
        });

        onUnmounted(() => {
            echart.dispose;
        });

        return {
            initChart
        };
    }
};
</script>

<style lang="less">
.echarts-box {
    width: 100%;
    height: 360px;

    .echart {
        width: 100%;
        height: 100%;
    }
}
</style>