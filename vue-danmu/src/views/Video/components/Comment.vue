<template>
    <div class="comment-box">
        <!--头像-->
        <common-avatar :url="userInfo.avatar"></common-avatar>
        <!--输入框-->
        <n-input class="comment-input" v-model:value="commentForm.content" placeholder="在这里写点什么吧~" maxlength="255"
            show-count type="textarea" :autosize="descSize" />
        <n-button v-if="!userInfo" disabled>未登录</n-button>
        <n-button v-else type="primary" @click="addComment">发表</n-button>
    </div>
    <div class="comment-item" v-for="(item, index) in commentList" :key="index">
        <!--头像和昵称-->
        <div class="comment-user">
            <common-avatar :url="item.avatar"></common-avatar>
            <div class="user-name">
                <span @click="goUserSpace(item.uid)">{{ item.name }}</span>
            </div>
            <n-time class="comment-time" type="relative" :time="new Date(item.created_at)"></n-time>
            <div class="comment-btn">
                <n-button v-if="userInfo" text @click="showReply(index, '')">回复</n-button>
                <n-button v-if="item.uid === userInfo.uid" text @click="deleteClick(item.id, index)">删除</n-button>
            </div>
        </div>
        <!--评论内容-->
        <div class="comment-content">
            <div v-if="item.content.indexOf('@') === -1">
                <span>{{ item.content }}</span>
            </div>
            <div class="content-text" v-else v-for="content in handleMention(item.content)">
                <span :class="content.class" @click="goMention(content.key)">{{ content.value }}</span>
            </div>
        </div>
        <!--动态回复框-->
        <div v-if="showReplyFlag[index]">
            <div class="reply-box">
                <common-avatar :url="userInfo.avatar"></common-avatar>
                <!--输入框-->
                <n-input class="reply-input" v-model:value="replyForm.content" :placeholder="replyTip" maxlength="200"
                    show-count type="textarea" :autosize="descSize" />
                <n-button type="primary" @click="submitReply(item)">回复</n-button>
            </div>
        </div>
        <!--回复-->
        <div class="reply-list" v-if="item.reply">
            <div class="reply-item" v-for="(reply, i) in item.reply" :key="i">
                <!--头像和昵称-->
                <div class="reply-user">
                    <div class="info">
                        <common-avatar :url="reply.avatar" :size="26" :iconsize="12"></common-avatar>
                        <div class="reply-user-name">
                            <span @click="goUserSpace(reply.uid)">{{ reply.name }}</span>
                        </div>
                        <n-time class="reply-time" type="relative" :time="new Date(reply.created_at)"></n-time>
                    </div>
                    <div class="reply-btn">
                        <n-button v-if="userInfo" text @click="showReply(index, reply.name)">回复</n-button>
                        <n-button v-if="reply.uid === userInfo.uid" text @click="deleteClick(reply.id, index, i)">删除
                        </n-button>
                    </div>
                </div>
                <!--二级评论内容-->
                <div class="reply-content">
                    <div v-if="reply.content.indexOf('@') === -1">
                        <span>{{ reply.content }}</span>
                    </div>
                    <div class="content-text" v-else v-for="content in handleMention(reply.content)">
                        <span :class="content.class" @click="goMention(content.key)">{{ content.value }}</span>
                    </div>
                </div>
            </div>
            <div class="more">
                <n-button :disabled="item.noMore" text @click="getMoreReply(item.id, index)">加载更多</n-button>
            </div>
        </div>
    </div>
    <!-- 加载更多 -->
    <div v-if="!noMore" class="more-btn">
        <n-button @click="getMore">加载更多</n-button>
    </div>
</template>

<script lang="ts">
import useMention from '@/hooks/mention';
import { useRouter } from 'vue-router';
import { onBeforeMount, reactive, ref, computed, defineComponent } from "vue";
import storage from "@/utils/stored-data";
import { addCommentAPI } from '@/api/comment';
import useComment from '@/hooks/comment';
import CommonAvatar from '@/components/CommonAvatar.vue';
import { NButton, NInput, NIcon, NTime, useNotification } from "naive-ui";
import { postCommentType, commentType } from '@/types/comment';

export default defineComponent({
    props: {
        vid: {
            type: Number,
            required: true
        }
    },
    setup(props) {
        const replyTip = ref('留下条评论吧');
        //输入框大小
        const descSize = {
            minRows: 3,
            maxRows: 3
        }
        const commentForm = reactive<postCommentType>({
            vid: props.vid,
            content: "",
            parentId: 0,
        })

        const replyUser = ref("");
        const replyForm = reactive({
            vid: props.vid,
            content: "",
            parentId: 0,
        })

        const userInfo = computed(() => {
            if (storage.get("userInfo")) {
                return storage.get("userInfo");
            } else {
                return {
                    uid: 0,
                    avatar: ''
                }
            }
        });
        const notification = useNotification();//通知

        //评论回复
        const replyCount = 2;//获取评论时查询的回复数
        const { noMore, commentList, getCommentList, getReplyListSync, deleteCommentSync } = useComment();
        const addComment = () => {
            addCommentAPI(commentForm).then((res) => {
                if (res.data.code === 2000) {
                    //加载评论
                    if (noMore.value) {
                        commentList.value.push({
                            id: res.data.data.id,
                            uid: userInfo.value.uid,
                            name: userInfo.value.name,
                            avatar: userInfo.value.avatar,
                            content: commentForm.content,
                            created_at: new Date().toDateString(),
                            reply: [],
                            page: 1,
                            noMore: true
                        })
                    }
                    notification.success({
                        title: '发布成功',
                        duration: 5000,
                    });
                    commentForm.content = "";
                }
            }).catch((err) => {
                notification.error({
                    title: '发布失败',
                    content: "原因:" + err.response.data.msg,
                    duration: 5000,
                })
            });
        }

        //提交回复
        const submitReply = (comment: commentType) => {
            replyForm.parentId = comment.id;
            if (replyUser.value) {
                replyForm.content = `回复 @${replyUser.value} :${replyForm.content}`;
            }
            if (comment.reply.length < replyCount) {
                comment.noMore = true;
            }
            addCommentAPI(replyForm).then((res) => {
                if (res.data.code === 2000) {
                    notification.success({
                        title: '发布成功',
                        duration: 5000,
                    });

                    if (comment.noMore) {
                        const newReply = {
                            id: res.data.data.id,
                            uid: userInfo.value.uid,
                            name: userInfo.value.name,
                            avatar: userInfo.value.avatar,
                            content: replyForm.content,
                            created_at: new Date().toDateString(),
                        }
                        comment.reply.push(newReply);
                    }
                    replyForm.content = "";
                    //关闭动态回复框
                    showReplyFlag.value.forEach((item,index) => {
                        if (item) showReplyFlag.value[index] = false;
                    });
                }
            }).catch((err) => {
                notification.error({
                    title: '发布失败',
                    content: "原因:" + err.response.data.msg,
                    duration: 5000,
                })
            });

        }

        //页面跳转
        const router = useRouter();
        const goUserSpace = (uid: number) => {
            let userUrl = router.resolve({ name: "User", params: { uid: uid } });
            window.open(userUrl.href, '_blank');
        }

        //显示隐藏动态回复
        const showReplyFlag = ref<Array<boolean>>([]);
        const showReply = (index: number, name: string) => {
            if (showReplyFlag.value[index]) {
                showReplyFlag.value[index] = false;
                return;
            }
            for (let i = 0; i < commentList.value.length; i++) {
                showReplyFlag.value[i] = false;
            }
            if (name) {
                replyTip.value = `回复 @${name}: `;
                replyUser.value = name;
            }
            showReplyFlag.value[index] = true;
        }

        //获取评论回复
        const page = ref(1);
        //加载更多回复
        const getMoreReply = (cid: number, index: number) => {
            const pageSize = 2;
            if (!commentList.value[index].page) {
                commentList.value[index].page = 1;
            }
            //获取回复内容
            getReplyListSync(cid, replyCount, commentList.value[index].page, pageSize).then((res) => {
                if (res === null) {
                    commentList.value[index].noMore = true;
                    return;
                }

                commentList.value[index].page += 1;
                commentList.value[index].reply.push(...res);
            })
        }

        //加载更多
        const getMore = () => {
            page.value++;
            getCommentList(props.vid, replyCount, page.value, 8);
        }


        //删除评论回复
        const deleteClick = (id: number, index: number, replyIndex: number | null = null) => {
            deleteCommentSync(id).then((res) => {
                if (res) {
                    if (replyIndex !== null) {
                        (commentList.value[index].reply)?.splice(replyIndex, 1);
                    } else {
                        commentList.value.splice(index, 1);
                    }
                }
            })
        }

        //处理@
        const { handleMention } = useMention();

        //前往@的用户
        const goMention = (name: string | null) => {
            if (name) {
                let userUrl = router.resolve({ name: 'MentionUser', params: { name: name } });
                window.open(userUrl.href, '_blank');
            }
        }

        onBeforeMount(() => {
            getCommentList(props.vid, replyCount, page.value, 8);
        })

        return {
            noMore,
            replyTip,
            userInfo,
            descSize,
            replyForm,
            commentForm,
            commentList,
            showReplyFlag,
            getMore,
            goMention,
            showReply,
            addComment,
            deleteClick,
            submitReply,
            goUserSpace,
            getMoreReply,
            handleMention
        }
    },
    components: {
        NIcon,
        NTime,
        NInput,
        NButton,
        CommonAvatar,
    }
});
</script>

<style lang="less" scoped>
.comment-box {
    margin-top: 10px;
    padding-bottom: 20px;
    display: flex;
    align-items: center;
    justify-content: space-between;
    border-bottom: 1px solid #e3e4e5;

    .comment-input {
        width: calc(100% - 200px);
    }

    button {
        width: 90px;
    }

}

.comment-item {
    margin-top: 10px;

    .comment-user {
        position: relative;

        .user-name {
            position: absolute;
            top: 0;
            left: 60px;
            font-weight: 500;
            cursor: pointer;
            color: #6d757a;
            font-size: 14px;
        }

        .comment-time {
            position: absolute;
            color: #969996;
            top: 26px;
            left: 60px;
            font-size: 12px;
        }

        .comment-btn {
            float: right;

            button {
                font-size: 10px;
                color: #969996;
                margin-left: 6px;

                &:hover {
                    color: #36ad6a;
                }
            }
        }
    }

    .comment-content {
        margin: 6px 0 6px 60px;
        font-size: 14px;
        line-height: 16px;
        padding-bottom: 10px;
        border-bottom: 1px solid #e3e4e5;
    }
}

/**动态回复框 */
.reply-box {
    display: flex;
    align-items: center;
    justify-content: space-around;
    margin: 10px 0 10px 60px;
    padding-bottom: 10px;
    border-bottom: 1px solid #e3e4e5;

    .reply-input {
        width: calc(100% - 180px);
    }

    button {
        width: 80px;
    }
}

.reply-list {
    border-bottom: 1px solid #e3e4e5;

    .reply-item {
        .reply-user {
            height: 30px;
            display: flex;
            align-items: center;
            justify-content: space-between;
            margin-left: 60px;

            .info {
                display: inline-flex;
                align-items: center;

                .reply-user-name {
                    cursor: pointer;
                    color: #666;
                    font-size: 12px;
                    padding: 0 10px;
                }

                .reply-time {
                    color: #969996;
                    font-size: 10px;
                    line-height: 30px;
                }
            }

            .reply-btn {
                float: right;

                button {
                    font-size: 10px;
                    color: #969996;
                    margin-left: 6px;

                    &:hover {
                        color: #36ad6a;
                    }
                }
            }
        }

        .reply-content {
            color: #222;
            display: block;
            margin: 6px 0 0 96px;
            padding-bottom: 10px;
        }
    }

    .more {
        button {
            font-size: 12px;
            color: #969996;
            margin: 10px 60px;

            &:hover {
                color: #36ad6a;
            }
        }
    }
}

.content-text {
    display: inline-block;

    .mention-user {
        color: #36ad6a;
        cursor: pointer;
        padding: 0 2px;
    }
}

.more-btn {
    text-align: center;
}
</style>