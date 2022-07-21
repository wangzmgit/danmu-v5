import { ref } from 'vue';
import { useNotification } from 'naive-ui';
import { commentType } from '@/types/comment';
import { getCommentListAPI, getReplyListAPI, deleteCommentAPI, getManageCommentListAPI } from '@/api/comment';


export default function useComment() {
    const total = ref(0);
    const noMore = ref(false);
    const commentList = ref<Array<commentType>>([]);
    const notification = useNotification();//通知

    const getCommentList = (vid: number, reply: number, page: number, pageSize: number) => {
        getCommentListAPI(vid, reply, page, pageSize).then((res) => {
            if (res.data.code === 2000) {
                const resComment = res.data.data.comments;
                if (resComment) {
                    commentList.value.push(...resComment);
                    if (resComment.length < pageSize) {
                        noMore.value = true;
                    }
                } else {
                    noMore.value = true;
                }
            }
        }).catch((err) => {
            notification.error({
                title: '获取失败',
                content: "原因:" + err.response.data.msg,
                duration: 5000,
            })
        });
    }

    const getManageCommentList = (vid: number, page: number, pageSize: number) => {
        getManageCommentListAPI(vid, page, pageSize).then((res) => {
            if (res.data.code === 2000) {
                total.value = res.data.data.count;
                commentList.value = res.data.data.comments;
            }
        }).catch((err) => {
            notification.error({
                title: '获取失败',
                content: "原因:" + err.response.data.msg,
                duration: 5000,
            })
        });
    }

    const getReplyListSync = async (vid: number, offset: number, page: number, pageSize: number) => {
        try {
            const res = await getReplyListAPI(vid, offset, page, pageSize);
            if (res.data.code === 2000) {
                return res.data.data.comments;
            }
            return [];
        } catch (err: any) {
            notification.error({
                title: '获取失败',
                content: "原因:" + err.response.data.msg,
                duration: 5000,
            });
            return [];
        }
    }

    const deleteCommentSync = async (id: number) => {
        try {
            const res = await deleteCommentAPI(id);
            if (res.data.code === 2000) {
                return true;
            }
            return false;
        } catch (err: any) {
            notification.error({
                title: '删除失败',
                content: "原因:" + err.response.data.msg,
                duration: 5000,
            });
            return false;
        }
    }


    return {
        total,
        noMore,
        commentList,
        getCommentList,
        getReplyListSync,
        deleteCommentSync,
        getManageCommentList
    }
}

