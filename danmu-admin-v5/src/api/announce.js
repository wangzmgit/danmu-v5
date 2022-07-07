import request from '@/utils/request';

//获取公告列表
export const getAnnounceListAPI = (page, page_size) => {
  return request.get(`v1/message/announce/list?page=${page}&page_size=${page_size}`);
}

//添加公告
export const addAnnounceAPI = ({ title, content, url }) => {
  return request.post('v1/message/announce/add', { title, content, url });
}

//删除公告
export const deleteAnnounceAPI = (id) => {
  return request.post('v1/message/announce/delete', { id });
}
