import { addAnnounceType } from '@/types/announce';
import request from '@/utils/request';

//获取公告列表
export const getAnnounceListAPI = (page: number, page_size: number) => {
  return request.get(`v1/message/announce/list?page=${page}&page_size=${page_size}`);
}

//添加公告
export const addAnnounceAPI = (addAnnounce: addAnnounceType) => {
  return request.post('v1/message/announce/add', addAnnounce);
}

//删除公告
export const deleteAnnounceAPI = (id: number) => {
  return request.post('v1/message/announce/delete', { id });
}
