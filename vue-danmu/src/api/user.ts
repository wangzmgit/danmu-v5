import { loginFormType, modifyInfoType, registerFormType } from '@/types/user';
import request from '@/utils/request';

//注册
export const registerAPI = (register: registerFormType) => {
    return request.post('v1/user/register', register);
}

//登录
export const loginAPI = (login: loginFormType) => {
    return request.post('v1/user/login', login);
}

//修改用户信息
export const modifyUserInfoAPI = (modify: modifyInfoType) => {
    return request.post('v1/user/info/modify', modify);
}

//获取用户信息(自己的)
export const getUserInfoAPI = () => {
    return request.get('v1/user/info/get');
}

//通过ID获取用户信息
export const getUserInfoByIDAPI = (uid: number) => {
    return request.get(`v1/user/info/other?uid=${uid}`);
}

//通过ID获取用户信息
export const getUidByNameAPI = (name: string) => {
    return request.get(`v1/user/info/uid?name=${name}`);
}