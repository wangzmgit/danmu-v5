export interface baseResourceType {
    id: number,
    title: string,
}

export interface reviewResourceType extends baseResourceType {
    modify: boolean,
    review: number
}