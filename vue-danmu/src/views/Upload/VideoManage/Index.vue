<template>
    <div class="video-list">
        <p class="create-title">稿件管理</p>
        <div class="card-item" v-for="(item, index) in videoList" :key="index">
            <div class="card-left">
                <img v-if="item.cover" :src="item.cover" alt="视频封面">
            </div>
            <div class="card-center">
                <p class="title" @click="govideo(item.vid)">{{ item.title }}</p>
                <span class="desc">简介：{{ item.desc }}</span>
                <span class="desc">创建于：<n-time :time="new Date(item.created_at)"></n-time></span>
            </div>
            <div class="card-right">
                <n-button class="edit-item" text @click="modifyVideo(item.vid, 'info')">
                    <template #icon>
                        <n-icon>
                            <create-outline />
                        </n-icon>
                    </template>
                    编辑信息
                </n-button>
                <n-button class="edit-item" text @click="modifyVideo(item.vid, 'video')">
                    <template #icon>
                        <n-icon>
                            <create-outline />
                        </n-icon>
                    </template>
                    编辑内容
                </n-button>
                <n-popconfirm negative-text="取消" positive-text="确认" @positive-click="deleteVideo(item.vid)">
                    <template #trigger>
                        <n-icon class="edit" size="20">
                            <trash-outline />
                        </n-icon>
                    </template>
                    是否删除这个视频
                </n-popconfirm>
            </div>
        </div>
    </div>
    <div class="pagination">
        <n-pagination v-model:page="page" :item-count="count" :page-size="pageSize" @update-page="pageChange" />
    </div>
</template>

<script lang="ts">
import { useRouter } from 'vue-router';
import { ref, onBeforeMount, defineComponent } from 'vue';
import { getMyVideoAPI, deleteVideoAPI } from '@/api/video';
import { CreateOutline, TrashOutline } from '@vicons/ionicons5';
import { NTime, NIcon, NPagination, NButton, NPopconfirm, useNotification } from 'naive-ui';
import { videoType } from '@/types/video';

export default defineComponent({
    setup() {
        const pageSize = 5;
        const page = ref(1);
        const count = ref(0);
        const videoList = ref<Array<videoType>>([]);

        const notification = useNotification();

        //获取我的视频
        const getMyVideo = () => {
            getMyVideoAPI(page.value, pageSize).then((res) => {
                if (res.data.code === 2000) {
                    count.value = res.data.data.count;
                    videoList.value = res.data.data.videos;
                }
            }).catch((err) => {
                notification.error({
                    title: '获取失败',
                    content: "原因:" + err.response.data.msg,
                    duration: 5000,
                })
            });
        }

        //删除视频
        const deleteVideo = (id: number) => {
            deleteVideoAPI(id).then((res) => {
                if (res.data.code === 2000) {
                    getMyVideo();
                }
            }).catch((err) => {
                notification.error({
                    title: '删除失败',
                    duration: 5000,
                })
            });
        }

        //页码改变
        const pageChange = (target: number) => {
            page.value = target;
            getMyVideo();
        }

        //前往视频详情
        const router = useRouter();
        const govideo = (vid: number) => {
            let videoUrl = router.resolve({ name: "Video", params: { vid: vid } });
            window.open(videoUrl.href, '_blank');
        }

        //前往修改视频
        const modifyVideo = (vid: number, status: string) => {
            router.push({ name: "UploadVideoHome", params: { vid: vid }, query: { modify: status } });
        }

        onBeforeMount(() => {
            getMyVideo();
        })

        return {
            page,
            count,
            pageSize,
            videoList,
            govideo,
            pageChange,
            modifyVideo,
            deleteVideo
        }
    },
    components: {
        NTime,
        NIcon,
        NButton,
        NPagination,
        NPopconfirm,
        TrashOutline,
        CreateOutline
    }
});
</script>

<style lang="less" scoped>
.video-list {
    margin: 0 10px;
    height: 580px;

    .create-title {
        line-height: 30px;
        font-size: 20px;
        user-select: none;
    }

    .card-item {
        width: 100%;
        height: 80px;
        margin-bottom: 12px;
        border-bottom: 1px solid #e6e6e6;
        padding-bottom: 12px;
    }
}

.card-left {
    float: left;
    width: 120px;
    height: 80px;
    margin-right: 10px;

    img {
        width: 100%;
        height: 100%;
        border-radius: 2px;
    }
}

.card-center {
    float: left;
    width: calc(100% - 400px);

    .title {
        font-size: 14px;
        color: #212121;
        line-height: 18px;
        margin: 0 0 26px;
        cursor: pointer;
        overflow: hidden;
        text-overflow: ellipsis;
        display: -webkit-box;
        -webkit-line-clamp: 1;
        -webkit-box-orient: vertical;

        &:hover {
            color: #36ad6a;
        }
    }

    .desc {
        font-size: 12px;
        color: #999;
        overflow: hidden;
        text-overflow: ellipsis;
        display: -webkit-box;
        -webkit-line-clamp: 1;
        -webkit-box-orient: vertical;
    }
}

.card-right {
    float: left;
    width: 260px;
    height: 100%;
    text-align: center;

    .edit-item {
        padding: 0 6px;
    }

    .edit {
        color: #999;
        cursor: pointer;
        line-height: 90px;
        margin-right: 20px;

        &:hover {
            color: #36ad6a;
        }
    }
}

.pagination {
    margin-left: 20px;
    padding-bottom: 20px;
}
</style>
