import { userInfoType } from "./user"

export interface baseVideoType {
    vid: number,
    title: string,
    cover: string,
    created_at: string
}

export interface videoType extends baseVideoType {
    desc: string,
    copyright: boolean,
    clicks: number
    author: userInfoType
}

// 上传视频
export interface uploadVideoType {
    title: string,
    cover: string,
    desc: string,
    copyright: boolean,
    partition: number
}

//上传的视频
export interface myUploadVideoType extends baseVideoType {
    review: number,
    clicks: number
}

// 修改视频信息
export interface modifyVideoType {
    vid: number,
    title: string,
    cover: string,
    desc: string,
    copyright: boolean
}

//分P
export interface partListType {
    title: string,
    duration: number
}