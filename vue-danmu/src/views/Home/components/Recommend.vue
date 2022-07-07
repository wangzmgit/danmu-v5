<template>
    <div class="card-list" v-title :data-title="`${config.title} - 首页`">
        <div class="v-card" v-for="item in recommendVideo" @click="goVideo(item.vid)">
            <hover-mask>
                <img class="cover" :src="item.cover" alt="封面" />
                <template class="mask" v-slot:action>
                    <p class="mask-title">{{ item.title }}</p>
                    <p class="mask-up">UP:{{ item.author }}</p>
                    <p class="mask-clicks">播放:{{ item.clicks }}</p>
                </template>
                <template class="mask" v-slot:hide>
                    <div class="title-box">
                        <p class="title">{{ item.title }}</p>
                    </div>
                </template>
            </hover-mask>
        </div>
    </div>
</template>

<script>
import config from '@/config.js';
import { useRouter } from 'vue-router';
import { onBeforeMount, ref } from 'vue'
import { getRecommendVideoAPI } from "@/api/video.js"
import HoverMask from '@/components/HoverMask.vue'
export default {
    setup() {
        const router = useRouter()
        const recommendVideo = ref([]);

        //获取推荐视频
        const getRecommendVideo = () => {
            getRecommendVideoAPI().then((res) => {
                if (res.data.code === 2000) {
                    recommendVideo.value = res.data.data.videos;
                }
            });
        }

        //页面跳转
        const goVideo = (vid) => {
            let videoUrl = router.resolve({ name: "Video", params: { vid: vid } });
            window.open(videoUrl.href,'_blank');
        }

        onBeforeMount(() => {
            getRecommendVideo();
        })

        return {
            config,
            recommendVideo,
            goVideo
        }
    },
    components: {
        "hover-mask": HoverMask
    },
};
</script>

<style lang="less" scoped>
.card-list {
    width: 100%;
    height: 280px;
    display: grid;
    grid-template-rows: repeat(2, 50%);
    grid-template-columns: repeat(3, 33.3%);

    .v-card {
        padding: 5px;
        height: 130px;
        width: calc(100% - 10px);

        .cover {
            width: 100%;
            height: 130px;
        }
    }
}

.mask {
    color: #fff;
    height: 130px;

    .mask-title {
        height: 32px;
        font-size: 14px;
        padding: 10px;

        overflow: hidden;
        text-overflow: ellipsis;
        display: -webkit-box;
        -webkit-line-clamp: 2;
        -webkit-box-orient: vertical;
    }

    .mask-up {
        position: absolute;
        left: 10px;
        font-size: 12px;
        bottom: 20px;
    }

    .mask-clicks {
        position: absolute;
        left: 10px;
        font-size: 12px;
        bottom: 0;
    }

    .title-box {
        position: absolute;
        bottom: 0;
        height: 30px;
        width: 100%;
        background: linear-gradient(0deg, rgba(0, 0, 0, 0.5), transparent);

        .title {
            margin: 0;
            font-size: 14px;
            padding-left: 6px;
            line-height: 30px;
            font-weight: 500;
            overflow: hidden;
            text-overflow: ellipsis;
            display: -webkit-box;
            -webkit-line-clamp: 1;
            -webkit-box-orient: vertical;
        }
    }
}
</style>