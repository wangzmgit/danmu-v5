import request from '@/utils/request';

//评论
export const addCommentAPI = ({ vid, content, parentId }) => {
    return request.post('v1/comment/add', { vid, content, parentId });
}

//获取评论
export const getCommentListAPI = (vid, reply, page, page_size) => {
    return request.get(`v1/comment/list?vid=${vid}&reply=${reply}&page=${page}&page_size=${page_size}`);
}

//获取回复
export const getReplyListAPI = (cid, offset, page, page_size) => {
    return request.get(`v1/comment/reply/list?cid=${cid}&offset=${offset}&page=${page}&page_size=${page_size}`);
}

//删除评论回复
export const deleteCommentAPI = (id) => {
    return request.post('v1/comment/delete', { id });
}

//获取管理评论
export const getManageCommentListAPI = (vid, page, page_size) => {
    return request.get(`v1/comment/list/manage?vid=${vid}&page=${page}&page_size=${page_size}`);
}