<template>
    <div class="database">
        <n-form :model="otherForm" label-width="auto">
            <n-form-item label="默认用户名前缀">
                <n-input v-model:value="otherForm.prefix" maxlength="7" show-count/>
            </n-form-item>
            <n-form-item label="前端文件路径">
                <n-input v-model:value="otherForm.fePath" />
            </n-form-item>
            <n-form-item label="最大分辨率">
                <n-select v-model:value="otherForm.maxRes" :options="resOptions"></n-select>
            </n-form-item>
            <n-form-item label="超级管理员邮箱">
                <n-input v-model:value="otherForm.email" />
            </n-form-item>
            <n-form-item label="超级管理员密码">
                <n-input type="password" v-model:value="otherForm.password" />
            </n-form-item>
            <div class="submit">
                <span>如不修改密码请留空</span>
                <n-button type="primary" @click="setOther">保存</n-button>
            </div>
        </n-form>
    </div>
</template>

<script lang="ts">
import { defineComponent, onBeforeMount, reactive } from "vue";
import { getOtherAPI, setOtherAPI } from '@/api/config';
import { NInput, NSelect, NForm, NFormItem, NButton, useNotification } from "naive-ui";
import { otherType } from "@/types/config";

export default defineComponent({
    setup() {
        const notification = useNotification();
        const resOptions = [
            {
                label: '不处理',
                value: 0
            },
            {
                label: '1080P',
                value: 1080
            },
            {
                label: '720P',
                value: 720
            },
            {
                label: '480P',
                value: 480
            },
            {
                label: '360P',
                value: 360
            },
        ];

        const getOther = () => {
            getOtherAPI().then((res) => {
                if (res.data.code === 2000) {
                    const resData = res.data.data;
                    otherForm.prefix = resData.prefix;
                    otherForm.fePath = resData.fe_path;
                    otherForm.maxRes = resData.max_res;
                    otherForm.email = resData.email;
                }
            })
        }

        const otherForm = reactive<otherType>({
            prefix: '',
            maxRes: 0,
            fePath: '',
            email: '',
            password: ''
        });

        const setOther = () => {
            setOtherAPI(otherForm).then((res) => {
                if (res.data.code === 2000) {
                    notification.success({
                        title: '修改成功',
                        duration: 3000
                    })
                }
            })
        }

        onBeforeMount(() => {
            getOther();
        })

        return {
            resOptions,
            otherForm,
            setOther
        }
    },
    components: {
        NInput,
        NSelect,
        NForm,
        NButton,
        NFormItem,
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