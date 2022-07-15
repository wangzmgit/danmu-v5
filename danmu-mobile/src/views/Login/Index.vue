<template>
	<div class="bg">
		<div class="head-box">
			<div class="head-text">
				<span>您好，</span>
				<br>
				<span>欢迎回到 {{ config.title }}</span>
			</div>
		</div>
		<div class="login-div">
			<div class="login-box">
				<n-form ref="formRef" :rules="rules" :model="loginForm">
					<n-form-item label="邮箱" path="email">
						<n-input placeholder="请输入邮箱" v-model:value="loginForm.email" />
					</n-form-item>
					<n-form-item label="密码" path="password">
						<n-input placeholder="请输入密码" v-model:value="loginForm.password" type="password" />
					</n-form-item>
					<div class="card-btn">
						<n-button type="primary" @click="loginClick()">登录</n-button>
						<n-button @click="goRegister">注册</n-button>
					</div>
				</n-form>
			</div>
		</div>
	</div>
</template>
<script lang="ts">
import config from "@/config";
import { reactive, ref, defineComponent } from 'vue';
import { loginAPI } from '@/api/user';
import storage from '@/utils/stored-data';
import { loginFormType } from '@/types/user';
import { NForm, FormInst, FormRules, NFormItem, NButton, NInput, useMessage } from 'naive-ui';
import router from "@/router";

export default defineComponent({
	setup() {

		const formRef = ref<FormInst | null>(null);
		const message = useMessage();//通知

		const loginForm = reactive<loginFormType>({
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

		const loginClick = () => {
			formRef.value?.validate((errors) => {
				if (!errors) {
					loginAPI(loginForm).then((res) => {
						if (res.data.code === 2000) {
							storage.set("token", res.data.data.token, 14 * 24 * 60);
							storage.set('userInfo', res.data.data.user, 14 * 24 * 60);
							router.push({ name: 'Home' });
						}
					}).catch(() => {
						message.error('登录失败');
					});
				} else {
					message.error('请检查输入的数据');
				}
			})
		}

		const goRegister = () => {
			router.push({ name: 'Register' });
		}

		return {
			rules,
			config,
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
		NInput
	}
});
</script>

<style lang="less" scoped>
.bg {
	position: fixed;
	top: 0;
	bottom: 0;
	width: 100%;
	background-color: #36ad6a;
}

.head-box {
	width: 100%;
	height: 30rem;

	.head-text {
		text-align: left;
		font-size: 1.2rem;
		color: #fff;
		padding: 3rem 0 0 1rem;
		font-weight: bold;
		line-height: 2.2rem;
	}
}

.login-div {
	width: 100%;
	height: 100%;
	position: relative;
	margin-top: -20rem;
	border-radius: 2rem 2rem 0 0;
	background-color: #fff;

	.login-box {
		padding: 2rem 1rem;

		.card-btn {
			text-align: center;

			button {
				width: 100%;
				margin-top: 1rem;
			}
		}
	}
}
</style>
