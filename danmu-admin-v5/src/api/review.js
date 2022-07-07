import request from '@/utils/request';

//获取审核视频列表
export const getReviewListAPI = (page, page_size) => {
    return request.get(`v1/review/list?page=${page}&page_size=${page_size}`);
}

//获取审核资源列表
export const getResourceListAPI = (vid) => {
    return request.get('v1/review/resource?vid=' + vid);
}

//审核视频
export const reviewVideoAPI = (id, status) => {
    return request.post('v1/review/video/add', { id, status });
}

//审核资源
export const reviewResourceAPI = (id, status) => {
    return request.post('v1/review/resource/add', { id, status });
}