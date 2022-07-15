import request from '@/utils/request';

//获取公告
export const getAnnounceAPI = (page: number, page_size: number) => {
    return request.get(`v1/message/announce/list?page=${page}&page_size=${page_size}`);
}
