import request from '@/utils/request';

//获取公告
export const getAnnounceAPI = (page, page_size) => {
    return request.get(`v1/message/announce/list?page=${page}&page_size=${page_size}`);
}
