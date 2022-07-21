<template>
    <div class="video-box">
        <p class="video-title">我的视频</p>
        <div class="card-list">
            <div class="v-card" v-for="(item, index) in videoList" :key="index">
                <div class="card-item" @click="goVideo(item.review, item.vid)">
                    <img class="cover" :src="item.cover" />
                    <div class="info">
                        <p class="title">{{ item.title }}</p>
                        <!--播放量-->
                        <p class="clicks">播放&ensp;{{ item.clicks ? item.clicks : 0 }}</p>
                    </div>
                </div>
                <div class="my-upload-card-btn">
                    <n-button text v-if="item.review !== 2000" @click="viewStatus(item.vid)">
                        <template #icon>
                            <n-icon>
                                <information-circle-outline />
                            </n-icon>
                        </template>
                        查看状态
                    </n-button>
                    <n-button text v-else @mouseover="showMenu(index, true)" @mouseleave="showMenu(index, false)">
                        <template #icon>
                            <n-icon>
                                <create-outline />
                            </n-icon>
                        </template>
                        修改内容
                    </n-button>
                    <n-button text @click="deleteVideo(item.vid)">
                        <template #icon>
                            <n-icon>
                                <trash-outline />
                            </n-icon>
                        </template>
                        删除
                    </n-button>
                    <!--修改视频-->
                    <div v-show="modifyMenu[index]" class="modify-menu" @mouseleave="showMenu(index, false)">
                        <div class="menu-item">
                            <span class="modify-btn" @click="modifyVideo(item.vid, 'video')">修改视频</span>
                        </div>
                        <div class="menu-item">
                            <span class="modify-btn" @click="modifyVideo(item.vid, 'info')">修改信息</span>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <n-pagination v-model:page="page" :item-count="count" :page-size="pageSize" @update-page="pageChange" />
    </div>
</template>

<script lang="ts">
import { useRouter } from 'vue-router';
import { myUploadVideoType } from '@/types/video';
import { ref, onBeforeMount, defineComponent } from 'vue';
import { getMyVideoAPI, deleteVideoAPI } from '@/api/video';
import { NButton, NIcon, NPagination, useNotification } from 'naive-ui';
import { CreateOutline, InformationCircleOutline, TrashOutline } from "@vicons/ionicons5"
export default defineComponent({
    emits: ['setCount'],
    setup(_props, ctx) {
        const pageSize = 8;
        const page = ref(1);
        const count = ref(0);

        const videoList = ref<Array<myUploadVideoType>>([]);
        const router = useRouter();
        const notification = useNotification();//通知

        //获取我的视频
        const getMyVideo = () => {
            getMyVideoAPI(page.value, pageSize).then((res) => {
                if (res.data.code === 2000) {
                    count.value = res.data.data.count;
                    videoList.value = res.data.data.videos;
                    ctx.emit('setCount', count.value);
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
            }).catch(() => {
                notification.error({
                    title: '删除失败',
                    duration: 5000,
                })
            });
        }

        //修改菜单
        const modifyMenu = ref<Array<boolean>>([]);
        const showMenu = (index: number, show: boolean) => {
            if (modifyMenu.value[index] && !show) {
                modifyMenu.value[index] = false;
                return;
            }
            for (let i = 0; i < videoList.value.length; i++) {
                modifyMenu.value[i] = false;
            }
            modifyMenu.value[index] = true;
        }

        //查看状态
        const viewStatus = (vid: number) => {
            router.push({ name: "UploadVideoHome", params: { vid: vid } });
        }

        const modifyVideo = (vid: number, status: string) => {
            router.push({ name: "UploadVideoHome", params: { vid: vid }, query: { modify: status } });
        }

        const goVideo = (review: number, vid: number) => {
            if (review === 2000) {
                router.push({ name: "Video", params: { vid: vid } });
            }
        }

        //页码改变
        const pageChange = (target: number) => {
            page.value = target;
            getMyVideo();
        }

        onBeforeMount(() => {
            getMyVideo();
        })

        return {
            page,
            count,
            pageSize,
            videoList,
            modifyMenu,
            goVideo,
            showMenu,
            pageChange,
            viewStatus,
            deleteVideo,
            modifyVideo
        }
    },
    components: {
        NIcon,
        NButton,
        NPagination,
        CreateOutline,
        TrashOutline,
        InformationCircleOutline
    }
});
</script>

<style lang="less" scoped>
.video-box {
    height: 430px;

    .video-title {
        font-size: 18px;
        margin-top: 20px;
    }
}

.card-list {
    width: 100%;
    height: 360px;
    display: grid;
    grid-template-rows: repeat(2, 50%);
    grid-template-columns: repeat(4, 25%);

    .v-card {
        padding: 5px;
        height: 170px;
        width: calc(100% - 10px);

        .card-item {
            position: relative;
            width: 100%;
            height: 130px;

            .cover {
                width: 100%;
                height: 100%;
                border-top-left-radius: 6px;
                border-top-right-radius: 6px;
            }

            .info {
                color: #fff;
                position: absolute;
                bottom: 0;
                width: 100%;
                background: linear-gradient(0deg, rgba(0, 0, 0, 0.7), transparent);

                p {
                    margin: 0;
                    padding-left: 10px;
                }

                .title {
                    font-size: 14px;
                    line-height: 16px;
                    overflow: hidden;
                    text-overflow: ellipsis;
                    display: -webkit-box;
                    -webkit-line-clamp: 1;
                    -webkit-box-orient: vertical;
                }

                .clicks {
                    font-size: 10px;
                }
            }
        }
    }
}

.my-upload-card-btn {
    width: calc(100% - 2px);
    border: 1px solid #d0d0d0;
    border-top: 0;
    border-bottom-left-radius: 6px;
    border-bottom-right-radius: 6px;

    button {
        width: 50%;
        color: #6f6f6f;
    }

    .modify-menu {
        width: 160px;
        height: 92px;
        z-index: 999;
        position: absolute;
        background-color: #fff;
        border-radius: 10px;
        box-shadow: 0 0 30px rgba(18, 18, 18, 0.2);

        .menu-item {
            margin-top: 7px;
            width: 120px;
            height: 36px;
            margin-left: 20px;
            -webkit-user-select: none;
            -moz-user-select: none;
            -ms-user-select: none;
            user-select: none;

            .modify-btn {
                display: block;
                color: #18191b;
                line-height: 36px;
                text-align: left;
                border-radius: 6px;
                padding-left: 6px;

                &:hover {

                    background-color: #c9ccd0;
                }
            }
        }
    }
}
</style>