<template>
    <video id="player" class="video" preload="auto" controls></video>
    <div class="video-box">
        <n-scrollbar style="max-height: 300px;">
            <div class="video-item" v-for="(item, index) in resources">
                <div class="item-left" @click="playVideo(item)">
                    <span>P{{ index + 1 }} {{ item.title }}</span>
                    <n-tag class="tag" :type="toTagType(item.review)">{{ toTagText(item.review) }}</n-tag>
                </div>
                <div v-show="item.review === 1000" class="item-right">
                    <span class="btn" @click="reviewResource(item.id, index, true)">通过</span>
                    <span class="btn" @click="reviewResource(item.id, index, false)">不通过</span>
                </div>
            </div>
        </n-scrollbar>
    </div>
</template>

<script>
import Hls from "hls.js";
import { onUnmounted, ref } from "vue";
import { reviewResourceAPI } from "@/api/review";
import { NScrollbar, NIcon, NTag, useNotification } from "naive-ui";
export default {
    props: {
        list: {
            type: Array
        }
    },
    setup(props) {
        const resources = ref(props.list);
        const notification = useNotification();//通知

        const toTagType = (state) => {
            switch (state) {
                case 2000:
                    return "success";
                case 200: case 1000:
                    return "info";
                default:
                    return "error";
            }
        }

        const toTagText = (state) => {
            switch (state) {
                case 200:
                    return "视频处理中";
                case 1000:
                    return "等待审核";
                case 2000:
                    return "审核通过";
                case 3200:
                    return "视频内容存在问题";
                case 3300:
                    return "视频处理失败";
                default:
                    return "未知";
            }
        }

        const hls = ref(null);
        const initPlayer = (src, player) => {
            if (!hls.value) hls.value = new Hls();
            hls.value.loadSource(src);
            hls.value.attachMedia(player);
            hls.value.on(Hls.Events.ERROR, () => {
                notification.error({
                    title: '资源加载失败',
                    duration: 5000,
                })
            });
        }

        //播放视频
        const playVideo = (resource) => {
            const src = resource.original ? resource.original : resource.res360;
            initPlayer(src, document.getElementById("player"));
        }

        //审核资源
        const reviewResource = (id, index, status) => {
            reviewResourceAPI(id, status).then((res) => {
                if (res.data.code === 2000) {
                    resources.value.splice(index, 1);
                }
            }).catch((err) => {
                notification.error({
                    title: '审核调用失败',
                    content: "原因:" + err.response.data.msg,
                    duration: 5000,
                })
            });
        }

        onUnmounted(() => {
            if (hls.value) {
                hls.value.stopLoad();
                hls.value.destroy();
                hls.value = null;
            }
        })

        return {
            resources,
            toTagType,
            toTagText,
            playVideo,
            reviewResource
        }
    },
    components: {
        NTag,
        NIcon,
        NScrollbar,
    }
}
</script>

<style lang="less" scoped>
.video {
    width: 100%;
    height: 270px;
    background-color: black;
}

.video-box {
    width: 100%;
    padding-bottom: 30px;

    .video-item {
        height: 50px;
        padding: 0 10px;
        display: flex;
        align-items: center;
        justify-content: space-between;
        border: 1px solid #efeff5;
        margin: 10px 10px;

        .item-left {
            cursor: pointer;

            .tag {
                margin-left: 10px;
            }
        }

        .item-right {
            width: 80px;
            display: flex;
            align-items: center;
            justify-content: space-between;

            .btn {
                cursor: pointer;
                color: #767c82;

                &:hover {
                    color: #18a058;
                }
            }
        }
    }
}
</style>