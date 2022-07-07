import request from '@/utils/request';

//获取视频状态
export const getArchiveAPI = (vid) => {
    return request.get('v1/archive/video?vid=' + vid);
}

//点赞
export const likeAPI = (id) => {
    return request.post('v1/archive/like/add', { id });
}

//取消赞
export const dislikeAPI = (id) => {
    return request.post('v1/archive/like/cancel', { id })
}
