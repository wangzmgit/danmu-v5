<template>
    <cover-upload class="upload" v-if="showUpload" :cover="videoInfo.cover" @finish="finishUpload" />
    <n-form class="info-form" label-placement="left" label-width="auto">
        <n-form-item label="标题">
            <n-input v-model:value="videoInfo.title" placeholder="请输入标题" maxlength="50" show-count />
        </n-form-item>
        <n-form-item label="视频简介">
            <n-input v-model:value="videoInfo.desc" placeholder="简单介绍一下视频~" maxlength="200" show-count type="textarea"
                :autosize="descSize" />
        </n-form-item>
        <n-form-item label="禁止转载">
            <n-switch v-model:value="videoInfo.copyright" />
        </n-form-item>
        <n-form-item label="分区">
            <n-input v-if="partitionText" disabled :value="partitionText"></n-input>
            <select-partition v-else @selected="selectedPartition"></select-partition>
        </n-form-item>
        <div class="upload-next-btn">
            <n-button v-if="isModify" type="primary" @click="modifyVideoInfo">修改</n-button>
            <n-button v-else type="primary" @click="uploadInfo">下一步</n-button>
        </div>
    </n-form>
</template>

<script>
import { ref, onMounted, reactive } from "vue";
import CoverUpload from "@/components/CoverUpload";
import { uploadVideoInfoAPI, modifyVideoInfoAPI } from '@/api/video';
import SelectPartition from "./SelectPartition.vue";
import { NIcon, useNotification } from 'naive-ui';
import { NForm, NFormItem, NButton, NInput, NSwitch } from 'naive-ui';//表单相关

export default {
    emits: ["finish"],
    props: {
        info: {
            type: Object,
        }
    },
    setup(props, ctx) {
        //简介输入框大小
        const descSize = {
            minRows: 3,
            maxRows: 5
        }

        const notification = useNotification();//通知

        const videoInfo = reactive({
            vid: 0,
            title: "",
            cover: "",
            desc: "",
            copyright: true,
            partition: 0,
        })

        const isModify = ref(false);
        const showUpload = ref(false);//显示上传
        const partitionText = ref("");//分区


        //封面上传完成
        const finishUpload = (cover) => {
            videoInfo.cover = cover;
        }

        //选中分区
        const selectedPartition = (value) => {
            videoInfo.partition = value;
        }

        //上传视频信息
        const uploadInfo = () => {
            if (!videoInfo.cover) {
                notification.error({
                    title: '上传失败',
                    content: "请上传视频封面",
                    duration: 5000,
                });
                return;
            }
            if (!videoInfo.title) {
                notification.error({
                    title: '上传失败',
                    content: "请填写视频标题",
                    duration: 5000,
                });
                return;
            }
            if (!videoInfo.partition) {
                notification.error({
                    title: '上传失败',
                    content: "请选择视频分区",
                    duration: 5000,
                });
                return;
            }
            uploadVideoInfoAPI(videoInfo).then((res) => {
                if (res.data.code === 2000) {
                    ctx.emit('finish', res.data.data.vid);
                }
            })
        }

        //修改视频信息
        const modifyVideoInfo = () => {
            if (!videoInfo.cover) {
                notification.error({
                    title: '请上传视频封面',
                    duration: 5000,
                });
                return;
            }
            if (!videoInfo.title) {
                notification.error({
                    title: '请填写视频标题',
                    duration: 5000,
                });
                return;
            }
            modifyVideoInfoAPI(videoInfo).then((res) => {
                if (res.data.code === 2000) {
                    notification.success({
                        title: '修改成功',
                        duration: 5000,
                    });
                }
            })
        }

        onMounted(() => {
            if (props.info.vid) {
                isModify.value = true;
                const video = props.info;
                videoInfo.vid = video.vid;
                videoInfo.title = video.title;
                videoInfo.desc = video.desc;
                videoInfo.cover = video.cover;
                videoInfo.copyright = video.copyright;
                partitionText.value = video.partition;
            }
            showUpload.value = true;
        })

        return {
            descSize,
            isModify,
            videoInfo,
            showUpload,
            partitionText,
            modifyVideoInfo,
            uploadInfo,
            finishUpload,
            selectedPartition
        }
    },
    components: {
        NIcon,
        NForm,
        NInput,
        NButton,
        NSwitch,
        NFormItem,
        CoverUpload,
        SelectPartition,
    }
}
</script>

<style lang="less" scoped>
.upload {
    margin: 50px auto;
}

.info-form {
    width: calc(100% - 260px);
    margin-left: 130px;

    .upload-next-btn {
        float: right;
        margin-bottom: 20px;

        button {
            width: 160px;
            height: 40px;
        }
    }
}
</style>