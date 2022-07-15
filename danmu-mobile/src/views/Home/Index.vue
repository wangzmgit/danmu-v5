<template>
    <div v-title :data-title="`${title} - 首页`"></div>
    <header-bar class="header"></header-bar>
    <div ref="scrollBox" class="content" @scroll="scrollList">
        <div class="carousel">
            <carousel></carousel>
        </div>
        <video-list :videos="videoDataList"></video-list>
        <div v-show="loading" class="loading">
            <n-spin></n-spin>
        </div>
        <footer-bar></footer-bar>
    </div>
</template>

<script lang="ts">
import { ref, defineComponent, onBeforeMount } from 'vue';
import HeaderBar from '@/components/HeaderBar.vue';
import FooterBar from '@/components/FooterBar.vue';
import Carousel from './components/Carousel.vue';
import VideoList from '@/components/VideoList.vue';
import { getVideoListAPI } from "@/api/video";
import { baseVideoType } from '@/types/video';
import { NSpin, useMessage } from 'naive-ui';
import config from '@/config';


export default defineComponent({
    setup() {
        const title = config.title;
        const message = useMessage();

        //视频列表
        const page = ref(1);
        const pageSize = 10;
        const noMore = ref(false);
        const loading = ref(false);
        const videoDataList = ref(Array<baseVideoType>());
        const getVideoList = () => {
            loading.value = true;
            getVideoListAPI(page.value, pageSize, 0).then((res) => {
                if (res.data.code === 2000) {
                    const resVideos = res.data.data.videos;
                    if (resVideos.length < pageSize) {
                        noMore.value = true;
                        message.info("没有更多了");
                    }
                    videoDataList.value = videoDataList.value.concat(resVideos);
                    loading.value = false;
                }
            })
        }

        const scrollBox = ref<HTMLElement | null>(null);
        const scrollList = (() => {
            //节流
            var timer: number | null = null;
            return () => {
                if (timer) {
                    return
                }

                timer = setTimeout(() => {
                    const scrollTop = scrollBox.value?.scrollTop || 0;
                    const clientHeight = scrollBox.value?.clientHeight || 0;
                    const scrollHeight = scrollBox.value?.scrollHeight || 0;
                    if (scrollTop + clientHeight >= scrollHeight - 50) {
                        if (!noMore.value && !loading.value) {
                            page.value++;
                            getVideoList();
                        }
                    }
                    timer = null;
                }, 500);
            }
        })();

        onBeforeMount(() => {
            getVideoList();
        })

        return {
            title,
            loading,
            scrollBox,
            videoDataList,
            scrollList
        }
    },
    components: {
        NSpin,
        HeaderBar,
        Carousel,
        FooterBar,
        VideoList
    }
})
</script>

<style lang="less" scoped>
.header {
    top: 0;
    z-index: 999;
    position: fixed;
    height: 50px;
}

.content {
    overflow-y: scroll;
    margin-top: 50px;
    height: calc(100vh - 50px);

    .carousel {
        height: 220px;
    }

    .loading {
        padding: 20px 0;
        text-align: center;
    }
}
</style>