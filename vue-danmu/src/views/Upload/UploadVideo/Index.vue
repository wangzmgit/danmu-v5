<template>
    <div class="upload-video" v-title data-title="视频投稿">
        <n-steps class="step" :current="current" :status="currentStatus">
            <n-step title="视频信息" />
            <n-step title="上传视频" />
            <n-step title="审核" />
            <n-step title="上传成功" />
        </n-steps>
        <div class="upload-center">
            <upload-info v-if="current === 1" :info="videoInfo" @finish="infoFinish"></upload-info>
            <upload-video :vid="videoInfo.vid" v-else-if="current === 2"></upload-video>
        </div>

    </div>
</template>

<script lang="ts">
import { onBeforeMount, ref, defineComponent } from 'vue';
import { NStep, NSteps } from "naive-ui";
import { getVideoStatusAPI } from '@/api/review';
import { useRoute } from 'vue-router';
import { useNotification } from 'naive-ui';
import UploadInfo from './components/UploadInfo.vue';
import UploadVideo from './components/UploadVideo.vue';
export default defineComponent({
    setup() {
        const current = ref(1);
        const videoInfo = ref({
            vid: 0,
            state: 0
        });
        const currentStatus = ref<"wait" | "error" | "finish" | "process">("process");
        const route = useRoute();
        const notification = useNotification();//通知

        const infoFinish = (vid: number) => {
            videoInfo.value.vid = vid;
            current.value = 2;
        }

        onBeforeMount(() => {
            const vid = Number(route.params.vid);
            const modify = (route.query.modify || "").toString();
            if (vid) {
                current.value = 4; //默认结果页
                getVideoStatusAPI(vid).then((res) => {
                    if (res.data.code === 2000) {
                        videoInfo.value = res.data.data.video;
                        switch (videoInfo.value.state) {
                            case 100:
                                current.value = 2;
                                break;
                            case 2000:
                                modifyUpload(modify);
                                break;
                            default:
                                break;
                        }
                    }
                }).catch((err) => {
                    console.log(err)
                    notification.error({
                        title: '获取失败',
                        content: "原因:" + err.response.data.msg,
                        duration: 5000,
                    })
                });
            }
        })

        //修改上传内容
        const modifyUpload = (modify: string) => {
            switch (modify) {
                case "video":
                    current.value = 2;
                    break;
                case "info":
                    current.value = 1;
                    break;
                default:
                    current.value = 4;
                    break;
            }
        }

        return {
            current,
            videoInfo,
            currentStatus,
            infoFinish
        }
    },
    components: {
        NStep,
        NSteps,
        UploadInfo,
        UploadVideo
    }
});
</script>

<style lang="less" scoped>
.step {
    width: calc(100% - 100px);
    margin-left: 100px;
    padding-top: 30px;
}
</style>
