export interface postCommentType {
    vid: number,
    content: string,
    parentId: number,
}


export interface commentType {
    id: number,
    content: string,
    uid: number,
    name: string,
    avatar: string,
    reply: Array<replyType> | null,
    created_at: string,
    page: number,//回复页码
    noMore: boolean
}

export interface replyType {
    id: number,
    content: string,
    uid: number,
    name: string,
    avatar: string,
    created_at: string,
}