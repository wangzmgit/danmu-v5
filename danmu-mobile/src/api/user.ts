import request from '@/utils/request';
import { loginFormType, registerFormType } from '@/types/user';

//注册
export const registerAPI = (registerForm: registerFormType) => {
    return request.post('v1/user/register', registerForm);
}

//登录
export const loginAPI = (loginForm: loginFormType) => {
    return request.post('v1/user/login', loginForm);
}
