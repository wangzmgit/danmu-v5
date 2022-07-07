<template>
    <div class="card-title">
        <img src="@/assets/logo.png" alt="logo" />
        <span>账号登录</span>
    </div>
    <n-form ref="formRef" :rules="rules" :model="loginForm" label-placement="left" label-width="auto">
        <n-form-item label="邮箱" path="email">
            <n-input placeholder="请输入邮箱" v-model:value="loginForm.email" />
        </n-form-item>
        <n-form-item label="密码" path="password">
            <n-input placeholder="请输入密码" v-model:value="loginForm.password" type="password" />
        </n-form-item>
        <div class="card-btn">
            <n-button @click="goRegister()">注册</n-button>
            <n-button type="primary" @click="loginClick()">登录</n-button>
        </div>
    </n-form>
</template>

<script>
import { reactive, ref } from 'vue';
import { loginAPI } from '@/api/user';
import storage from '@/utils/stored-data';
import { useLoginStore } from '@/store/login-store';
import { NForm, NFormItem, NButton, NInput, useNotification } from 'naive-ui';


export default {
    emits: ['register','close'],
    setup(_props, ctx) {

        const formRef = ref(null);
        const store = useLoginStore();
        const notification = useNotification();//通知

        const loginForm = reactive({
            email: "",
            password: "",
        })

        const rules = {
            email: [
                { required: true, message: "请输入邮箱", trigger: ['blur', 'input'] },
                { type: "email", message: "请输入正确的邮箱地址", trigger: ['blur', 'input'] },
            ],
            password: { required: true, message: '请输入密码', trigger: ['blur', 'input'] },
        }

        const loginClick = () => {
            formRef.value.validate((errors) => {
                if (!errors) {
                    loginAPI(loginForm).then((res) => {
                        if (res.data.code === 2000) {
                            storage.set("token", res.data.data.token, 14 * 24 * 60);
                            storage.set('userInfo', res.data.data.user, 14 * 24 * 60);
                            store.setLoginState(true);
                            ctx.emit('close');
                        }
                    }).catch((err) => {
                        notification.error({
                            title: '登录失败',
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

        const goRegister = () => {
            ctx.emit("register");
        }

        return {
            rules,
            formRef,
            loginForm,
            loginClick,
            goRegister
        }
    },
    components: {
        NForm,
        NFormItem,
        NButton,
        NInput,
    }
}
</script>

<style lang="less" scoped>
.card-title {
    display: flex;
    height: 50px;
    align-items: center;
    margin-bottom: 30px;

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
    width: 100%;
    display: flex;
    margin-top: 10px;
    justify-content: space-between;

    button {
        width: 150px;
        height: 40px;
    }
}
</style>