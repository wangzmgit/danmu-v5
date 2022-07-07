import request from '@/utils/request'

//发送弹幕
export const sendDanmakuAPI = ({ vid, time, color, type, text, part }) => {
    return request.post('v1/danmaku/send', { vid, time, color, type, text, part });
}

//获取弹幕
export const getDanmakuAPI = (vid, part) => {
    return request.get(`v1/danmaku/list?vid=${vid}&part=${part}`);
}
