<template>
    <div class="card-title">
        <img src="@/assets/logo.png" alt="logo" />
        <span>账号注册</span>
    </div>
    <n-form ref="formRef" :rules="rules" :model="registerForm" label-placement="left" label-width="auto">
        <n-form-item label="邮箱" path="email">
            <n-input placeholder="请输入邮箱" v-model:value="registerForm.email" />
        </n-form-item>
        <n-form-item label="密码" path="password">
            <n-input placeholder="请输入密码" v-model:value="registerForm.password" type="password" />
        </n-form-item>
        <n-form-item label="验证码" path="code">
            <n-input placeholder="请输入验证码" v-model:value="registerForm.code" />
            <n-button :disabled="disabledSend" @click="beforeSendCode">{{ sendBtnText }}</n-button>
        </n-form-item>
    </n-form>
    <div class="card-btn">
        <n-button @click="goLogin()">返回登录</n-button>
        <n-button type="primary" @click="registerClick">注册</n-button>
    </div>
</template>

<script lang="ts">
import { reactive, ref, defineComponent } from 'vue';
import { registerAPI } from '@/api/user';
import useSendCode from '@/hooks/send-code';
import { NForm, FormInst, FormRules, NFormItem, NButton, NInput, useNotification } from 'naive-ui';

export default defineComponent({
    emits: ['login'],
    setup(_props, ctx) {

        const formRef = ref<FormInst | null>(null);
        const notification = useNotification();//通知

        const registerForm = reactive({
            email: "",
            password: "",
            code: ""
        })

        const rules: FormRules = {
            email: [
                { required: true, message: "请输入邮箱", trigger: ['blur', 'input'] },
                { type: "email", message: "请输入正确的邮箱地址", trigger: ['blur', 'input'] },
            ],
            password: { required: true, message: '请输入密码', trigger: ['blur', 'input'] },
            code: { required: true, message: '请输入验证码', trigger: ['blur', 'input'] },
        }

        const registerClick = () => {
            formRef.value?.validate((errors: any) => {
                if (!errors) {
                    registerAPI(registerForm).then((res) => {
                        if (res.data.code === 2000) {
                            notification.success({
                                content: '注册成功',
                                duration: 5000,
                            })
                            goLogin();
                        }
                    }).catch((err) => {
                        notification.error({
                            title: '注册失败',
                            content: "原因:" + err.response.data.msg,
                            duration: 5000,
                        })
                    });
                } else {
                    notification.error({
                        title: '请检查输入的数据',
                        duration: 5000,
                    })
                }
            })
        }

        const goLogin = () => {
            ctx.emit("login");
        }

        //发送验证码
        const { disabledSend, sendBtnText, sendCode } = useSendCode();
        const beforeSendCode = () => {
            if (!registerForm.email) {
                return;
            }
            sendCode(registerForm.email);
        }

        return {
            rules,
            formRef,
            registerForm,
            sendBtnText,
            disabledSend,
            goLogin,
            sendCode,
            registerClick,
            beforeSendCode,
        }
    },
    components: {
        NForm,
        NFormItem,
        NButton,
        NInput,
    }
});
</script>

<style lang="less" scoped>
.card-title {
    display: flex;
    height: 30px;
    align-items: center;
    margin-bottom: 20px;

    img {
        width: 40px;
        height: 40px;
    }

    span {
        color: #18a058;
        font-size: 26px;
        margin-left: 10px;
    }
}

.card-btn {
    position: absolute;
    right: 30px;
    bottom: 26px;
    width: 360px;
    display: flex;
    justify-content: space-between;

    button {
        width: 140px;
        height: 36px;
    }
}
</style>