import request from '@/utils/request';
import { danmakuType } from '@/components/WPlayer/types/danmaku';

//发送弹幕
export const sendDanmakuAPI = (danmaku: danmakuType) => {
    return request.post('v1/danmaku/send', danmaku);
}

//获取弹幕
export const getDanmakuAPI = (vid: number, part: number) => {
    return request.get(`v1/danmaku/list?vid=${vid}&part=${part}`);
}
