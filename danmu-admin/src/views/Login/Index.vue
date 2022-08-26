<template>
    <div v-title data-title="管理员登录">
        <div class="login-bg">
            <img src="@/assets/login.svg" />
        </div>
        <div class="login-card">
            <p class="title">管理员登录</p>
            <n-form ref="formRef" :rules="rules" :model="loginForm" label-placement="left" label-width="auto">
                <n-form-item label="邮箱" path="email">
                    <n-input placeholder="请输入邮箱" v-model:value="loginForm.email" />
                </n-form-item>
                <n-form-item label="密码" path="password">
                    <n-input placeholder="请输入密码" v-model:value="loginForm.password" type="password" />
                </n-form-item>
                <div class="card-btn">
                    <n-checkbox @update-checked="loginChange">root登录</n-checkbox>
                    <n-button type="primary" @click="loginClick()">登录</n-button>
                </div>
            </n-form>
        </div>
    </div>
</template>

<script lang="ts">
import { defineComponent, reactive, ref } from 'vue';
import { loginAPI, rootLoginAPI } from '@/api/user';
import storage from '@/utils/stored-data';
import { useRouter } from 'vue-router';
import { NForm, FormInst, FormRules, NFormItem, NButton, NInput, NCheckbox, useNotification } from 'naive-ui';

export default defineComponent({
    setup(_props) {

        const formRef = ref<FormInst | null>(null);
        const router = useRouter();
        const notification = useNotification();//通知

        const loginForm = reactive({
            email: "",
            password: "",
        })

        const rules: FormRules = {
            email: [
                { required: true, message: "请输入邮箱", trigger: ['blur', 'input'] },
                { type: "email", message: "请输入正确的邮箱地址", trigger: ['blur', 'input'] },
            ],
            password: { required: true, message: '请输入密码', trigger: ['blur', 'input'] },
        }

        //改变登录方式
        const root = ref(false);
        const loginChange = (val: boolean) => {
            root.value = val;
        }

        const loginClick = () => {
            formRef.value?.validate((errors) => {
                if (!errors) {
                    if (root.value) {
                        rootLogin();
                    } else {
                        adminLogin();
                    }
                } else {
                    notification.error({
                        title: '请检查输入的数据',
                        duration: 5000,
                    })
                }
            })
        }

        const adminLogin = () => {
            loginAPI(loginForm).then((res) => {
                if (res.data.code === 2000) {
                    if (res.data.data.user.role in [1, 2, 3]) {
                        storage.set("admin_access_token", res.data.data.access_token, 5);
                        storage.set("admin_refresh_token", res.data.data.refresh_token, 14 * 24 * 60);
                        storage.set('adminInfo', res.data.data.user, 14 * 24 * 60);
                        router.push({ name: "Home" });
                    } else {
                        notification.error({
                            title: '权限不足',
                            duration: 5000,
                        })
                    }
                }
            }).catch((err) => {
                notification.error({
                    title: '登录失败',
                    content: "原因:" + err.response.data.msg,
                    duration: 5000,
                })
            });
        }

        const rootLogin = () => {
            rootLoginAPI(loginForm).then((res) => {
                if (res.data.code === 2000) {
                    storage.set("admin_access_token", res.data.data.access_token, 5);
                    storage.set("admin_refresh_token", res.data.data.refresh_token, 14 * 24 * 60);
                    storage.set('adminInfo', res.data.data.user, 14 * 24 * 60);
                    router.push({ name: "Home" });
                }
            }).catch((err) => {
                notification.error({
                    title: '登录失败',
                    content: "原因:" + err.response.data.msg,
                    duration: 5000,
                })
            });
        }

        return {
            rules,
            formRef,
            loginForm,
            loginClick,
            loginChange

        }
    },
    components: {
        NForm,
        NFormItem,
        NButton,
        NInput,
        NCheckbox
    }
});
</script>

<style lang="less" scoped>
.login-bg {
    position: fixed;
    top: 0;
    bottom: 0;
    right: 0;
    left: 0;
    z-index: -1;

    img {
        width: 100%;
        height: 100%;
        border: 0px;
    }
}

.login-card {
    position: absolute;
    top: 50%;
    left: 50%;
    padding: 0 20px;
    margin: -160px 0 0 -220px;
    width: 440px;
    height: 260px;
    background-color: #fff;
    border-radius: 20px;
    box-shadow: 16px 16px 50px -12px rgba(0, 0, 0, 0.8);

    .title {
        margin: 20px 0 20px 30px;
        font-size: 22px;
    }

    .card-btn {
        padding-top: 10px;
        display: flex;
        align-items: center;
        justify-content: space-between;

        button {
            width: 120px;
        }
    }
}
</style>