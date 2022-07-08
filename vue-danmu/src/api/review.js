import request from '@/utils/request';

//获取视频状态
export const getVideoStatusAPI = (vid) => {
    return request.get('v1/review/status?vid=' + vid);
}

//提交审核
export const submitReviewAPI = (id) => {
    return request.post('v1/review/submit', { id });
}
