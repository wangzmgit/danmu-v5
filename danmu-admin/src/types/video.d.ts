import { userInfoType } from "./user"

export interface videoType {
    vid: number,
    uid: number,
    title: string,
    cover: string,
    desc: string,
    created_at: string,
    copyright: boolean,
    partition: number
}