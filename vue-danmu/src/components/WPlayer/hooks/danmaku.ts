import { sendDanmakuAPI, getDanmakuAPI } from '@/api/danmaku';

export default function useDanmakuAPI() {

    return {
        getDanmakuAPI,
        sendDanmakuAPI
    }
}
