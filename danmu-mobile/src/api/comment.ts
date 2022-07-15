import request from '@/utils/request';

//获取评论
export const getCommentListAPI = (vid: number, reply: number, page: number, page_size: number) => {
    return request.get(`v1/comment/list?vid=${vid}&reply=${reply}&page=${page}&page_size=${page_size}`);
}

//获取回复
export const getReplyListAPI = (cid: number, offset: number, page: number, page_size: number) => {
    return request.get(`v1/comment/reply/list?cid=${cid}&offset=${offset}&page=${page}&page_size=${page_size}`);
}