import request from '@/utils/request';

//登录
export const loginAPI = ({ email, password }) => {
  return request.post('v1/user/login', { email, password });
}

export const rootLoginAPI = ({ email, password }) => {
  return request.post('v1/user/root/login', { email, password });
}


//获取用户列表
export const getUserListAPI = (page, page_size) => {
  return request.get(`v1/user/list?page=${page}&page_size=${page_size}`);
}

//删除用户
export const deleteUserAPI = (id) => {
  return request.post('v1/user/delete', { id });
}

//修改用户信息
export const modifyUserAPI = ({ id, name, email, sign }) => {
  return request.post('v1/user/info/modify/admin', { id, name, email, sign });
}

//搜索用户
export const searchUserAPI = (keyword) => {
  return request.get(`v1/user/search/admin?keywords=${keyword}`);
}

//修改用户权限
export const modifyRoleAPI = ({ uid, role }) => {
  return request.post('v1/user/role/modify', { uid, role });
}
