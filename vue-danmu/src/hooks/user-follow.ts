import { ref } from 'vue';
import { useNotification } from 'naive-ui';
import { getFollowStatusAPI, followAPI, unfollowAPI } from "@/api/follow";

export default function useUserFollow() {
    const isFollow = ref(false);//是否关注

    const notification = useNotification();//通知

    const getFollowStatus = (fid: number) => {
        getFollowStatusAPI(fid).then((res) => {
            if (res.data.code === 2000) {
                isFollow.value = res.data.data.follow;
            }
        })
    }

    //关注
    const follow = (id: number) => {
        followAPI(id).then((res) => {
            if (res.data.code === 2000) {
                notification.success({
                    title: '关注成功',
                })
            }
        }).catch((err) => {
            notification.error({
                title: '关注失败',
                content: "原因:" + err.response.data.msg,
                duration: 5000,
            })
        });
    }

    //取关
    const unfollow = (id: number) => {
        unfollowAPI(id).then((res) => {
            if (res.data.code === 2000) {
                notification.success({
                    title: '取关成功',
                })
            }
        }).catch((err) => {
            notification.error({
                title: '取关失败',
                content: "原因:" + err.response.data.msg,
                duration: 5000,
            })
        });
    }

    return {
        isFollow,
        follow,
        unfollow,
        getFollowStatus,
    }
}

