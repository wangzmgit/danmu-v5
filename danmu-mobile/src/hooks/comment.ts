import { ref } from 'vue';
import { useMessage } from 'naive-ui';
import { getCommentListAPI, getReplyListAPI } from '@/api/comment';
import { commentType } from '@/types/comment';


export default function useComment() {
    const total = ref(0);
    const noMore = ref(false);
    const commentList = ref<Array<commentType>>([]);
    const message = useMessage();//通知

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
        }).catch(() => {
            message.error('获取失败');
        });
    }



    const getReplyListSync = async (vid: number, offset: number, page: number, pageSize: number) => {
        try {
            const res = await getReplyListAPI(vid, offset, page, pageSize);
            if (res.data.code === 2000) {
                return res.data.data.comments;
            }
            return [];
        } catch {
            message.error('获取失败');
            return [];
        }
    }


    return {
        total,
        noMore,
        commentList,
        getCommentList,
        getReplyListSync,
    }
}

