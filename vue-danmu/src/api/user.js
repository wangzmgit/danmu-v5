import request from '@/utils/request';

//注册
export const registerAPI = ({ email, password, code }) => {
    return request.post('v1/user/register', { email, password, code });
}

//登录
export const loginAPI = ({ email, password }) => {
    return request.post('v1/user/login', { email, password });
}

//修改用户信息
export const modifyUserInfoAPI = ({ name, gender, birthday, sign }) => {
    return request.post('v1/user/info/modify', { name, gender, birthday, sign });
}

//获取用户信息(自己的)
export const getUserInfoAPI = () => {
    return request.get('v1/user/info/get');
}

//通过ID获取用户信息
export const getUserInfoByIDAPI = (uid) => {
    return request.get(`v1/user/info/other?uid=${uid}`);
}

//通过ID获取用户信息
export const getUidByNameAPI = (name) => {
    return request.get(`v1/user/info/uid?name=${name}`);
}