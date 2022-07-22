import { loginFormType, modifyRoleType, modifyUserType } from '@/types/user';
import request from '@/utils/request';

//登录
export const loginAPI = (login: loginFormType) => {
  return request.post('v1/user/login', login);
}

export const rootLoginAPI = (login: loginFormType) => {
  return request.post('v1/user/root/login', login);
}


//获取用户列表
export const getUserListAPI = (page: number, page_size: number) => {
  return request.get(`v1/user/list?page=${page}&page_size=${page_size}`);
}

//删除用户
export const deleteUserAPI = (id: number) => {
  return request.post('v1/user/delete', { id });
}

//修改用户信息
export const modifyUserAPI = (modify: modifyUserType) => {
  return request.post('v1/user/info/modify/admin', modify);
}

//搜索用户
export const searchUserAPI = (keyword: string) => {
  return request.get(`v1/user/search/admin?keywords=${keyword}`);
}

//修改用户权限
export const modifyRoleAPI = (modify: modifyRoleType) => {
  return request.post('v1/user/role/modify', modify);
}
