<template>
    <div class="upload-video">
        <n-upload name="video" :data="uploadData" multiple directory-dnd :show-file-list="false" :action="videoUrl"
            :headers="headers" @before-upload="beforeUploadVideo" @change="handleChange">
            <n-upload-dragger>
                <div v-if="!uploading">
                    <div style="margin-bottom: 12px">
                        <n-icon size="48" :depth="3">
                            <archive-icon />
                        </n-icon>
                    </div>
                    <n-text style="font-size: 16px">
                        点击或拖拽视频到此处上传视频
                    </n-text>
                    <n-p depth="3" style="margin: 8px 0 0 0">
                        上传文件大小需小于{{ config.maxVideoSize }}M,仅支持.mp4格式文件
                    </n-p>
                </div>
                <n-progress v-else type="circle" :percentage="percent" />
            </n-upload-dragger>
        </n-upload>
    </div>
</template>

<script lang="ts">
import { defineComponent, ref } from "vue";
import config from "@/config";
import storage from "@/utils/stored-data";
import { VideoUrl as videoUrl } from "@/utils/request";
import { ArchiveOutline as ArchiveIcon } from "@vicons/ionicons5";
import { NIcon, NUpload, NUploadDragger, NText, NP, NProgress, useNotification } from 'naive-ui';

export default defineComponent({
    emits: ["finish"],
    props: {
        vid: {
            type: Number,
            required: true
        }
    },
    setup(props, ctx) {
        //文件上传请求头
        const headers = {
            Authorization: "Bearer " + storage.get("token"),
        }
        const uploadData: any = {
            vid: props.vid
        };
        const percent = ref(0);//上传百分比
        const uploading = ref(false);//是否上传中
        const notification = useNotification();//通知

        //上传之前的回调
        const beforeUploadVideo = async (options: any) => {
            const file = options.file;
            const isJpgOrPng = file.type === "video/mp4";
            if (!isJpgOrPng) {
                notification.error({
                    title: '上传失败',
                    content: "文件只支持mp4格式",
                    duration: 5000,
                });
            }
            const isLtMaxSize = file.file.size / 1024 / 1024 < config.maxVideoSize;
            if (!isLtMaxSize) {
                notification.error({
                    title: '上传失败',
                    content: `视频大小不能超过${config.maxVideoSize}M`,
                    duration: 5000,
                });
            }
            return isJpgOrPng && isLtMaxSize;
        }

        //上传变化的回调
        const handleChange = (options: any) => {
            uploading.value = true;
            const status = options.file.status;
            if (status === "finished") {
                ctx.emit("finish");
                notification.success({
                    title: '上传完成',
                    duration: 3000,
                });
                uploading.value = false;
            } else if (status === "error") {
                notification.error({
                    title: '文件上传失败',
                    duration: 5000,
                });
                uploading.value = false;
            }

            //上传进度
            const event = options.event;
            if (event) {
                percent.value = Math.floor((event.loaded / event.total) * 100);
            }
        }

        return {
            config,
            headers,
            percent,
            videoUrl,
            uploading,
            uploadData,
            handleChange,
            beforeUploadVideo
        }
    },
    components: {
        NP,
        NIcon,
        NText,
        NUpload,
        NProgress,
        NUploadDragger,
        //图标
        ArchiveIcon
    }
});
</script>

<style lang="less" scoped>
.upload-video {
    width: 350px;
    margin: 50px auto;
}
</style>