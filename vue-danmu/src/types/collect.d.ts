export interface collectionType {
    id: number,
    name?: string,
    cover?: string,
    desc?: string,
    open?: boolean,
    created_at?: string
}

export interface collectType {
    vid: number,
    addList: Array<number>,
    cancelList: Array<number>,
}

export interface modifyCollectionType {
    id: number,
    name: string,
    cover: string,
    desc: string,
    open: boolean,
}

export interface collectionInfoType {
    id: number,
    name: string,
    cover: string,
    desc: string,
    open: boolean,
    created_at: string
    author: {
        uid: number,
        name: string
    }
}