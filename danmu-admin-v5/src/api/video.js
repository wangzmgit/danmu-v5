import request from '@/utils/request';

//删除视频
export const deleteVideoAPI = (id) => {
    return request.post('v1/video/delete/admin', { id });
}

//获取视频列表
export const getVideoListAPI = (page, page_size) => {
    return request.get(`v1/video/list/admin?page=${page}&page_size=${page_size}`);
}

//搜索视频
export const searchVideoAPI = (keyword) => {
    return request.get(`v1/video/search?keywords=${keyword}`);
}