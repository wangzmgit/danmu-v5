import { danmakuType } from "./danmaku";

export interface qualityType {
    [key: number]: {
        name?: string,
        url: string,
        type?: string
    }
}

export interface danmakuOptionsType {
    open: boolean,
    placeholder?: string,
    data?: Array<danmakuType>,
    send?: (danmaku: danmakuType) => void
}


export interface optionsType {
    resource: string | qualityType,
    cover?: string,
    type?: string,//视频类型
    mobile?: boolean,//移动端
    blob?: boolean,//mp4视频是否使用blob
    customType?: (player: HTMLVideoElement, src: string) => void,
    theme?: string,//主题色,
    danmaku?: danmakuOptionsType,
    playbackSpeed?: Array<number>,// 播放速度
}

export interface resourceType {
    [key: string]: string,
}