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
