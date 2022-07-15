import request from '@/utils/request';
import { messageType } from '@/types/message';

//获取私信列表
export const getMsgListAPI = () => {
    return request.get('v1/message/list')
}

//获取私信列表
export const getMsgDetailsAPI = (fid: number, page: number, page_size: number) => {
    return request.get(`v1/message/details?fid=${fid}&page=${page}&page_size=${page_size}`)
}

//发送私信
export const sendMsgAPI = (message: messageType) => {
    return request.post('v1/message/send', message)
}