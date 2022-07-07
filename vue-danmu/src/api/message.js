import request from '@/utils/request'

//获取私信列表
export const getMsgListAPI = () => {
    return request.get('v1/message/list')
}

//获取私信列表
export const getMsgDetailsAPI = (fid, page, page_size) => {
    return request.get(`v1/message/details?fid=${fid}&page=${page}&page_size=${page_size}`)
}

//发送私信
export const sendMsgAPI = ({ fid, content }) => {
    return request.post('v1/message/send', { fid, content })
}

//已读私信
export const readMsgAPI = (id) => {
    return request.post('v1/message/read', { id })
}