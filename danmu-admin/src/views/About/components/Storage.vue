<template>
    <div class="database">
        <n-form :model="storageForm" label-width="auto">
            <n-form-item label="使用https">
                <n-switch v-model:value="storageForm.https"></n-switch>
            </n-form-item>
            <n-form-item label="最大图片文件大小(MB)">
                <n-input-number placeholder="文件域名" v-model:value="storageForm.maxImgSize" />
            </n-form-item>
            <n-form-item label="最大视频文件大小(MB)">
                <n-input-number placeholder="文件域名" v-model:value="storageForm.maxVideoSize" />
            </n-form-item>
            <n-form-item label="文件域名">
                <n-input placeholder="文件域名" v-model:value="storageForm.domain" />
            </n-form-item>
            <n-form-item label="使用OSS存储">
                <n-switch v-model:value="storageForm.oss"></n-switch>
            </n-form-item>
            <div v-if="storageForm.oss">
                <n-form-item label="OSS存储空间(Bucket)">
                    <n-input placeholder="存储空间(Bucket)" v-model:value="storageForm.bucket" />
                </n-form-item>
                <n-form-item label="OSS存储区域(Endpoint)">
                    <n-input placeholder="存储区域(Endpoint)" v-model:value="storageForm.endpoint" />
                </n-form-item>
                <n-form-item label="accesskeyId">
                    <n-input placeholder="accesskeyId" v-model:value="storageForm.accesskeyId" />
                </n-form-item>
                <n-form-item label="accesskeySecret">
                    <n-input placeholder="accesskeySecret" type="password"
                        v-model:value="storageForm.accesskeySecret" />
                </n-form-item>
            </div>
            <div class="submit">
                <span>如不修改密码请留空</span>
                <n-button type="primary" @click="setStorage">保存</n-button>
            </div>
        </n-form>
    </div>
</template>

<script lang="ts">
import { defineComponent, onBeforeMount, reactive } from "vue";
import { getStorageAPI, setStorageAPI } from '@/api/config';
import { NInput, NSwitch, NInputNumber, NForm, NFormItem, NButton, useNotification } from "naive-ui";
import { storageType } from "@/types/config";

export default defineComponent({
    setup() {
        const notification = useNotification();

        const getStorage = () => {
            getStorageAPI().then((res) => {
                if (res.data.code === 2000) {
                    const resData = res.data.data;
                    storageForm.https = resData.https;
                    storageForm.oss = resData.oss;
                    storageForm.domain = resData.domain;
                    storageForm.maxImgSize = resData.max_img_size;
                    storageForm.maxVideoSize = resData.max_video_size;
                    storageForm.bucket = resData.bucket;
                    storageForm.endpoint = resData.endpoint;
                    storageForm.accesskeyId = resData.accesskeyId;
                }
            })
        }

        const storageForm = reactive<storageType>({
            https: false,
            oss: false,
            domain: '',
            maxImgSize: 5,
            maxVideoSize: 500,
            bucket: '',
            endpoint: '',
            accesskeyId: '',
            accesskeySecret: '',
        });

        const setStorage = () => {
            setStorageAPI(storageForm).then((res) => {
                if (res.data.code === 2000) {
                    notification.success({
                        title: '修改成功',
                        duration: 3000
                    })
                }
            })
        }

        onBeforeMount(() => {
            getStorage();
        })

        return {
            storageForm,
            setStorage
        }
    },
    components: {
        NInput,
        NSwitch,
        NForm,
        NButton,
        NFormItem,
        NInputNumber
    }
});
</script>

<style lang="less" scoped>
.submit {
    display: flex;
    align-items: center;
    justify-content: space-between;

    span {
        color: #666;
    }
}
</style>