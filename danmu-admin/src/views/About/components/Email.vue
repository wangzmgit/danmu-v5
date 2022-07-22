<template>
    <div class="database">
        <n-form :model="emailForm" label-width="auto">
            <n-form-item label="开启邮箱验证">
                <n-switch v-model:value="emailForm.open"></n-switch>
            </n-form-item>
            <div v-if="emailForm.open">
                <n-form-item label="发件人名称">
                    <n-input placeholder="发件人名称" v-model:value="emailForm.name" />
                </n-form-item>
                <n-form-item label="主机">
                    <n-input placeholder="主机" v-model:value="emailForm.host" />
                </n-form-item>
                <n-form-item label="端口">
                    <n-input-number placeholder="端口" v-model:value="emailForm.port" />
                </n-form-item>
                <n-form-item label="邮箱地址">
                    <n-input placeholder="邮箱地址" v-model:value="emailForm.address" />
                </n-form-item>
                <n-form-item label="授权码">
                    <n-input placeholder="授权码" type="password" v-model:value="emailForm.password" />
                </n-form-item>
            </div>
            <div class="submit">
                <span>如不修改密码请留空</span>
                <n-button type="primary" @click="setEmail">保存</n-button>
            </div>
        </n-form>
    </div>
</template>

<script lang="ts">
import { defineComponent, onBeforeMount, reactive } from "vue";
import { getEmailAPI, setEmailAPI } from '@/api/config';
import { NSwitch, NInput, NInputNumber, NForm, NFormItem, NButton, useNotification } from "naive-ui";

export default defineComponent({
    setup() {
        const notification = useNotification();

        const getEmail = () => {
            getEmailAPI().then((res) => {
                if (res.data.code === 2000) {
                    const resData = res.data.data;
                    emailForm.open = resData.open;
                    emailForm.name = resData.name;
                    emailForm.host = resData.host;
                    emailForm.port = resData.port;
                    emailForm.address = resData.address;
                }
            })
        }

        const emailForm = reactive({
            open: false,
            name: '',
            host: '',
            port: 465,
            address: '',
            password: ''
        });

        const setEmail = () => {
            setEmailAPI(emailForm).then((res) => {
                if (res.data.code === 2000) {
                    notification.success({
                        title: '修改成功',
                        duration: 3000
                    })
                }
            })
        }

        onBeforeMount(() => {
            getEmail();
        })

        return {
            emailForm,
            setEmail
        }
    },
    components: {
        NSwitch,
        NInput,
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