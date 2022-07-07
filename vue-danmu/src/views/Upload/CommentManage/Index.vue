<template>
    <div class="select-box">
        <span class="title">评论管理</span>
        <n-select class="select" label-field="title" value-field="vid" remote :options="videoList"
            @update:value="changeVideo" />
    </div>
    <div class="comment-list">
        <div class="comment-item" v-for="(item, index) in commentList" :key="index">
            <!--头像和昵称-->
            <div class="comment-user">
                <common-avatar :url="item.avatar"></common-avatar>
                <div class="user-name">
                    <span @click="goUserSpace(item.uid)">{{ item.name }}</span>
                </div>
                <n-time class="comment-time" type="relative" :time="new Date(item.created_at)"></n-time>
                <div class="comment-btn">
                    <n-button text @click="deleteClick(item.id, index)">删除</n-button>
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
        </div>
    </div>
    <div class="pagination">
        <n-pagination v-model:page="page" :item-count="total" :page-size="6" @update-page="pageChange" />
    </div>
</template>

<script>
import { getMyVideoAPI } from '@/api/video';
import useMention from '@/hooks/mention';
import { useRouter } from 'vue-router';
import { onBeforeMount, ref } from "vue";
import useComment from '@/hooks/comment';
import CommonAvatar from '@/components/CommonAvatar.vue';
import { NButton, NTime, NPagination, NSelect } from "naive-ui";

export default {
    setup() {
        const page = ref(1);
        const currentVid = ref(0);
        const { total, commentList, deleteCommentSync, getManageCommentList } = useComment();

        //页码改变
        const pageChange = (target) => {
            page.value = target;
            getManageCommentList(currentVid.value, page.value, 6);
        }

        //页面跳转
        const router = useRouter();
        const goUserSpace = (uid) => {
            let userUrl = router.resolve({ name: "User", params: { uid: uid } });
            window.open(userUrl.href, '_blank');
        }


        //删除评论回复
        const deleteClick = (id, index) => {
            deleteCommentSync(id).then((res) => {
                if (res) {
                    commentList.value.splice(index, 1);
                }
            })
        }

        const { handleMention } = useMention();
        //前往@的用户
        const goMention = (name) => {
            let userUrl = router.resolve({ name: 'MentionUser', params: { name: name } });
            window.open(userUrl.href, '_blank');
        }

        //选择视频
        const videoPage = ref(1);//分页
        const videoList = ref([]);//视频列表
        const videoPageSize = 15;
        //获取分区列表
        const getVideoList = () => {
            getMyVideoAPI(videoPage.value, videoPageSize).then((res) => {
                if (res.data.code === 2000) {
                    const resVideo = res.data.data.videos;
                    videoList.value = videoList.value.concat(resVideo);
                    if (resVideo.length === videoPageSize) {
                        videoPage.value++;
                        getVideoList();
                    }
                }
            })
        }

        const changeVideo = (vid) => {
            currentVid.value = vid;
            pageChange(1);
        }


        onBeforeMount(() => {
            getVideoList();
        })

        return {
            page,
            total,
            videoPage,
            videoList,
            currentVid,
            commentList,
            changeVideo,
            pageChange,
            goMention,
            deleteClick,
            goUserSpace,
            handleMention
        }
    },
    components: {
        NTime,
        NSelect,
        NButton,
        NPagination,
        CommonAvatar,
    }
}
</script>

<style lang="less" scoped>
.select-box {
    height: 60px;
    padding: 0 20px;
    display: flex;
    align-items: center;
    justify-content: space-between;

    .title {
        color: #181818;
        font-size: 20px;
    }

    .select {
        width: 200px;
    }
}

.comment-list {
    height: 600px;
    padding: 10px 30px;
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

        .content-text {
            display: inline-block;

            .mention-user {
                color: #36ad6a;
                cursor: pointer;
                padding: 0 2px;
            }
        }

    }
}


.pagination {
    margin-left: 20px;
    padding-bottom: 20px;
}
</style>