export interface partitionType {
    id: number,
    content: string
    subpartition?: Array<partitionType>
}

export interface addPartitionType {
    content: string,
    parentId: number
}