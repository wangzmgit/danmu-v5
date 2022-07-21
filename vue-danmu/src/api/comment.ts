import request from '@/utils/request';
import { postCommentType } from '@/types/comment';

//评论
export const addCommentAPI = (postComment: postCommentType) => {
    return request.post('v1/comment/add', postComment);
}

//获取评论
export const getCommentListAPI = (vid: number, reply: number, page: number, page_size: number) => {
    return request.get(`v1/comment/list?vid=${vid}&reply=${reply}&page=${page}&page_size=${page_size}`);
}

//获取回复
export const getReplyListAPI = (cid: number, offset: number, page: number, page_size: number) => {
    return request.get(`v1/comment/reply/list?cid=${cid}&offset=${offset}&page=${page}&page_size=${page_size}`);
}

//删除评论回复
export const deleteCommentAPI = (id: number) => {
    return request.post('v1/comment/delete', { id });
}

//获取管理评论
export const getManageCommentListAPI = (vid: number, page: number, page_size: number) => {
    return request.get(`v1/comment/list/manage?vid=${vid}&page=${page}&page_size=${page_size}`);
}