import request from '@/utils/request'

//获取关注数据
export const getFollowDataAPI = (uid: number) => {
    return request.get(`v1/follow/count?uid=${uid}`)
}