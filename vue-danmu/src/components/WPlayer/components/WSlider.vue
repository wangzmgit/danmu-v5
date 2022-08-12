<template>
    <div :class="vertical ? 'slider-line-vertical' : 'slider-line'">
        <div ref="sliderRef" class="player-bar" @click="vertical ? clickSliderVertical($event) : clickSlider($event)">
            <div class="player-played"
                :style="vertical ? `height: ${activePercentage}%` : `width: ${activePercentage}%`">
                <div ref="blockRef" class="player-block"></div>
            </div>
        </div>
    </div>
</template>

<script lang="ts">
import { defineComponent, onMounted, ref } from 'vue';
export default defineComponent({
    emits: ['changeValue'],
    props: {
        value: {
            type: Number,
            default: 0,
        },
        max: {
            type: Number,
            default: 100,
        },
        min: {
            type: Number,
            default: 0,
        },
        vertical: {
            type: Boolean,
            default: false
        }
    },
    setup(props, ctx) {
        const activePercentage = ref(props.value);
        const blockRef = ref<HTMLElement | null>(null);
        const sliderRef = ref<HTMLElement | null>(null);


        //点击滑动条
        const clickSlider = (e: MouseEvent) => {
            const activeHeight = (e.pageX - sliderRef.value!.getBoundingClientRect().left);
            let percentage = Math.round((activeHeight / sliderRef.value!.clientWidth) * 100) / 100;

            activePercentage.value = percentage * 100;
            ctx.emit('changeValue', Math.round((props.max - props.min) * percentage * 100) / 100);
        }

        // 滑动滑动条
        const slideTimeLine = () => {
            blockRef.value!.onmousedown = function () {
                document.onmousemove = function (e) {
                    const activeHeight = (e.pageX - sliderRef.value!.getBoundingClientRect().left);
                    let percentage = Math.round((activeHeight / sliderRef.value!.clientWidth) * 100) / 100;
                    percentage = Math.max(0, percentage);
                    percentage = Math.min(percentage, 1);

                    activePercentage.value = percentage * 100;
                    ctx.emit('changeValue', Math.round((props.max - props.min) * percentage * 100) / 100);
                };
                document.onmouseup = function () {
                    document.onmousemove = document.onmouseup = null;
                };
            };
        }

        //点击滑动条(垂直)
        const clickSliderVertical = (e: MouseEvent) => {
            const activeHeight = sliderRef.value!.clientHeight - (e.pageY - sliderRef.value!.getBoundingClientRect().top);
            let percentage = Math.round((activeHeight / sliderRef.value!.clientHeight) * 100) / 100;

            activePercentage.value = percentage * 100;
            ctx.emit('changeValue', Math.round((props.max - props.min) * percentage * 100) / 100);
        }

        // 滑动滑动条(垂直)
        const slideTimeLineVertical = () => {
            blockRef.value!.onmousedown = function () {
                document.onmousemove = function (e) {
                    //计算新的百分比
                    const activeHeight = sliderRef.value!.clientHeight - (e.pageY - sliderRef.value!.getBoundingClientRect().top);
                    let percentage = Math.round((activeHeight / sliderRef.value!.clientHeight) * 100) / 100;
                    percentage = Math.max(0, percentage);
                    percentage = Math.min(percentage, 1);

                    activePercentage.value = percentage * 100;
                    ctx.emit('changeValue', Math.round((props.max - props.min) * percentage * 100) / 100);
                };
                document.onmouseup = function () {
                    document.onmousemove = document.onmouseup = null;
                };
            };
        }

        onMounted(() => {
            props.vertical ? slideTimeLineVertical() : slideTimeLine();
            activePercentage.value = Math.round((props.value / (props.max - props.min)) * 100);
        })

        return {
            blockRef,
            sliderRef,
            activePercentage,
            clickSlider,
            clickSliderVertical,
        }
    }
});
</script>

<style lang="less" scoped>
.slider-line {
    margin-left: 6px;
    width: calc(100% - 12px);

    .player-bar {
        position: relative;
        width: 100%;
        background: rgba(107, 107, 107, 0.6);
        height: 4px;
        cursor: pointer;
    }

    .player-played {
        position: absolute;
        background: #18a058;
        height: 4px;

        .player-block {
            position: absolute;
            top: -6px;
            right: -6px;
            width: 12px;
            height: 12px;
            border-radius: 50%;
            border: 2px solid #18a058;
            background: #fff;
            transition: 0.2s all;
            user-select: none;
        }
    }
}


.slider-line-vertical {
    height: calc(100% - 24px);
    margin: 12px 0;

    .player-bar {
        position: relative;
        width: 4px;
        height: 100%;
        cursor: pointer;
        background: rgba(107, 107, 107, 0.6);
    }

    .player-played {
        position: absolute;
        background: #36ad6a;
        width: 100%;
        bottom: 0;

        .player-block {
            position: absolute;
            top: -4px;
            right: -4px;
            width: 8px;
            height: 8px;
            border-radius: 50%;
            border: 2px solid #36ad6a;
            background: #fff;
            transition: 0.2s all;
            user-select: none;
        }
    }
}
</style>