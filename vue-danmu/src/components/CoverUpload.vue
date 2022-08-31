<template>
    <div class="upload-cover">
        <n-upload name="cover" multiple directory-dnd :show-file-list="false" :action="coverUrl" :headers="headers"
            @before-upload="beforeUploadCover" @change="handleChange">
            <img v-if="cover" class="cover" :src="cover" alt="封面" />
            <n-upload-dragger v-else>
                <div v-if="!uploading">
                    <div style="margin-bottom: 12px">
                        <n-icon size="48" :depth="3">
                            <archive-icon />
                        </n-icon>
                    </div>
                    <n-text style="font-size: 16px">
                        点击或拖拽图片到此处上传封面
                    </n-text>
                    <n-p depth="3" style="margin: 8px 0 0 0">
                        上传文件大小需小于{{ config.maxImgSize }}M,仅支持.jpg .jpeg .png格式文件
                    </n-p>
                </div>
                <n-progress v-else type="circle" :percentage="percent" />
            </n-upload-dragger>
        </n-upload>
    </div>
</template>

<script lang="ts">
import { defineComponent, onBeforeMount, ref } from "vue";
import config from "@/config";
import storage from "@/utils/stored-data";
import { getAccessToken } from "@/api/token";
import { CoverUrl as coverUrl } from "@/utils/request";
import { ArchiveOutline as ArchiveIcon } from "@vicons/ionicons5";
import { NIcon, NUpload, NUploadDragger, NText, NP, NProgress, useNotification } from 'naive-ui';

export default defineComponent({
    props: {
        cover: {
            type: String
        }
    },
    emits: ["finish"],
    setup(props, ctx) {
        const cover = ref(props.cover);
        //文件上传请求头
        const headers = {
            Authorization: "Bearer " + storage.get("access_token"),
        }
        const percent = ref(0);//上传百分比
        const uploading = ref(false);//是否上传中
        const notification = useNotification();//通知

        //上传之前的回调
        const beforeUploadCover = async (options: any) => {
            const file = options.file;
            const isJpgOrPng =
                file.type === "image/jpeg" || file.type === "image/png";
            if (!isJpgOrPng) {
                notification.error({
                    title: '上传失败',
                    content: "文件只支持jpg/jpeg/png格式",
                    duration: 5000,
                });
            }
            const isLtMaxSize = file.file.size / 1024 / 1024 < config.maxImgSize;
            if (!isLtMaxSize) {
                notification.error({
                    title: '上传失败',
                    content: `图片大小不能超过${config.maxImgSize}M`,
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
                const res = JSON.parse(options.event.currentTarget.response);
                cover.value = res.data.url;
                ctx.emit("finish", cover.value);
                notification.success({
                    title: '上传完成',
                    duration: 3000,
                });
            } else if (status === "error") {
                notification.error({
                    title: '文件上传失败',
                    duration: 5000,
                });
            }

            //上传进度
            const event = options.event;
            if (event) {
                percent.value = Math.floor((event.loaded / event.total) * 100);
            }
        }

        // 刷新token
        const refreshToken = () => {
            getAccessToken().then((res) => {
                if (res.data.code === 2000) {
                    headers.Authorization = `Bearer ${res.data.data.token}`;
                    storage.set("access_token", res.data.data.token, 5);
                }
            })
        }

        onBeforeMount(() => {
            // 先执行一次刷新token，每经过三分钟再执行一次
            refreshToken();
            setInterval(() => {
                refreshToken();
            }, 180000);//3 * 60 * 1000 = 3分钟
        })

        return {
            cover,
            config,
            headers,
            percent,
            coverUrl,
            uploading,
            handleChange,
            beforeUploadCover
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
.upload-cover {
    width: 350px;
    margin: 20px auto;

    .cover {
        width: 350px;
        height: 200px;
    }

}
</style>