export interface messageType {
    fid: number,
    content: string,
}

export interface messageListType {
    uid: number,
    name: string,
    avatar: string,
    created_at: string
}

export interface messageDetailsType {
    fid: number,
    from_id: number,
    content: string,
    created_at: string
}