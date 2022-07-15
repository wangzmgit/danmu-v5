<template>
    <div class="comment-item" v-for="(item, index) in commentList" :key="index">
        <!--头像和昵称-->
        <div class="comment-user">
            <common-avatar :url="item.avatar" :size="38" :iconsize="26"></common-avatar>
            <div class="user-name">
                <span>{{ item.name }}</span>
            </div>
            <n-time class="comment-time" type="relative" :time="new Date(item.created_at)"></n-time>
        </div>
        <!--评论内容-->
        <div class="comment-content">
            <div v-if="item.content.indexOf('@') === -1">
                <span>{{ item.content }}</span>
            </div>
            <div v-else>
                <div class="content-text" v-for="(content, c_key) in handleMention(item.content)" :key="c_key">
                    <span :class="content.class">{{ content.value }}</span>
                </div>
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
                            <span>{{ reply.name }}</span>
                        </div>
                        <n-time class="reply-time" type="relative" :time="new Date(reply.created_at)"></n-time>
                    </div>
                </div>
                <!--二级评论内容-->
                <div class="reply-content">
                    <div v-if="reply.content.indexOf('@') === -1">
                        <span>{{ reply.content }}</span>
                    </div>
                    <div v-else>
                        <div class="content-text" v-for="(content, c_key) in handleMention(reply.content)" :key="c_key">
                            <span :class="content.class">{{ content.value }}</span>
                        </div>
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
import useComment from '@/hooks/comment';
import useMention from '@/hooks/mention';
import { defineComponent, onBeforeMount, ref } from "vue";
import CommonAvatar from '@/components/CommonAvatar.vue';
import { NButton, NTime } from "naive-ui";

export default defineComponent({
    props: {
        vid: {
            type: Number,
        }
    },
    setup(props: any) {

        //获取评论回复
        const page = ref(1);
        const replyCount = 2;//获取评论时查询的回复数
        const { noMore, commentList, getCommentList, getReplyListSync } = useComment();

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
                (commentList.value[index].reply)?.push(...res);
            })
        }

        //加载更多
        const getMore = () => {
            page.value++;
            getCommentList(props.vid, replyCount, page.value, 8);
        }

        //处理@
        const { handleMention } = useMention();

        onBeforeMount(() => {
            getCommentList(props.vid, replyCount, page.value, 8);
        })

        return {
            noMore,
            commentList,
            getMore,
            getMoreReply,
            handleMention
        }
    },
    components: {
        NTime,
        NButton,
        CommonAvatar,
    }
});
</script>

<style lang="less" scoped>
.comment-item {
    margin-top: 10px;

    .comment-user {
        position: relative;
        padding-left: 12px;

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
    }

    .comment-content {
        margin: 6px 0 6px 60px;
        font-size: 14px;
        line-height: 16px;
        padding-bottom: 10px;
        border-bottom: 1px solid #e3e4e5;
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