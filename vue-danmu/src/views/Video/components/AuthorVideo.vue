<template>
    <div class="video-head">
        <span class="head-title">TA的视频</span>
    </div>
    <div class="card-box" v-for="item in videoList" @click="govideo(item.vid)">
        <img class="cover" :src="item.cover" />
        <div class="info">
            <span class="title">{{ item.title }}</span>
            <n-time class="time" :time="new Date(item.created_at)"></n-time>
        </div>
    </div>
</template>

<script lang="ts">
import { defineComponent, onBeforeMount, ref } from "vue";
import { useRouter } from "vue-router";
import useUserVideo from "@/hooks/user-video";
import { NTime } from "naive-ui";

export default defineComponent({
    props: {
        uid: {
            type: Number,
            required: true
        }
    },
    setup(props) {
        const uid = ref(props.uid);
        const router = useRouter();
        const { videoList, getVideoListByUid } = useUserVideo();

        const govideo = (vid: number) => {
            let videoUrl = router.resolve({ name: "Video", params: { vid: vid } });
            window.open(videoUrl.href, '_blank');
        }

        onBeforeMount(() => {
            getVideoListByUid(uid.value, 1, 6);
        })


        return {
            videoList,
            govideo,
        }
    },
    components: {
        NTime
    }

});
</script>

<style lang="less" scoped>
.video-head {
    margin: 12px 0;

    .head-title {
        color: #333;
        font-size: 16px;
        font-weight: 500;
    }
}

.card-box {
    display: flex;
    margin-bottom: 6px;

    .cover {
        position: relative;
        width: 140px;
        height: 80px;
        border-radius: 2px;
        background: #f4f5f7;
    }

    .info {
        margin-left: 10px;

        .title {
            color: #333;
            cursor: pointer;
            height: 36px;
            line-height: 18px;
            font-size: 16px;
            font-weight: 500;
            overflow: hidden;
            text-overflow: ellipsis;
            word-break: break-all;
            display: -webkit-box;
            -webkit-line-clamp: 2;
            -webkit-box-orient: vertical;

            &:hover {
                color: #36ad6a;
            }
        }

        .time {
            margin-top: 6px;
            display: inline-block;
            height: 16px;
            width: 100%;
            color: #999;
            font-size: 12px;
        }
    }
}
</style>