<template>
    <n-card class="card">
        <n-table striped>
            <thead class="table-head">
                <tr>
                    <th>VID</th>
                    <th>封面</th>
                    <th>标题</th>
                    <th>分区</th>
                    <th>简介</th>
                    <th>上传时间</th>
                    <th>允许转载</th>
                    <th>操作</th>
                </tr>
            </thead>
            <tbody class="table-body">
                <tr v-for="item in videos">
                    <td>{{ item.vid }}</td>
                    <td>
                        <img class="cover" :src="item.cover" alt="视频封面">
                    </td>
                    <td>{{ item.title }}</td>
                    <td>{{ item.partition }}</td>
                    <td>{{ item.desc }}</td>
                    <td>
                        <n-time :time="new Date(item.created_at)" />
                    </td>
                    <td>{{ toCopyright(item.copyright) }}</td>
                    <td>
                        <div class="btn-list">
                            <n-button text @click="getResourceList(item.vid)">查看视频</n-button>
                            <n-button text @click="reviewVideo(item.vid, true)">通过</n-button>
                            <n-button text @click="reviewVideo(item.vid, false)">不通过</n-button>
                        </div>
                    </td>
                </tr>
            </tbody>
        </n-table>
    </n-card>
    <n-drawer v-model:show="activeDrawer" :width="502">
        <n-drawer-content title="视频列表">
            <video-list :list="resources"></video-list>
        </n-drawer-content>
    </n-drawer>
</template>

<script lang="ts">
import VideoList from './components/VideoList.vue';
import { NDrawer, NDrawerContent } from 'naive-ui';
import { NTable, NButton, NCard, NTime, useNotification } from 'naive-ui';
import { getReviewListAPI, getResourceListAPI, reviewVideoAPI } from '@/api/review';
import { defineComponent, onBeforeMount, reactive, ref } from 'vue';
import { videoType } from '@/types/video';
export default defineComponent({
    setup() {
        const pagination = reactive({
            current: 1,
            pageSize: 8,
        })

        const videos = ref<Array<videoType>>([]);
        const notification = useNotification();//通知

        const getVideoList = () => {
            getReviewListAPI(pagination.current, pagination.pageSize).then((res) => {
                if (res.data.code === 2000) {
                    videos.value = res.data.data.videos;
                }
            }).catch((err) => {
                notification.error({
                    title: '获取视频列表失败',
                    content: "原因:" + err.response.data.msg,
                    duration: 5000,
                })
            });
        }

        const resources = ref([]);//视频资源
        const activeDrawer = ref(false);//是否打开审核抽屉
        //获取审核资源
        const getResourceList = (vid: number) => {
            getResourceListAPI(vid).then((res) => {
                if (res.data.code === 2000) {
                    resources.value = res.data.data.video;
                    activeDrawer.value = true;
                }
            }).catch((err) => {
                notification.error({
                    title: '获取资源失败',
                    content: "原因:" + err.response.data.msg,
                    duration: 5000,
                })
            });
        }

        //审核视频
        const reviewVideo = (vid: number, status: boolean) => {
            reviewVideoAPI(vid, status).then((res) => {
                if (res.data.code === 2000) {
                    getVideoList();
                }
            }).catch((err) => {
                notification.error({
                    title: '审核调用失败',
                    content: "原因:" + err.response.data.msg,
                    duration: 5000,
                })
            });
        }

        //允许转载
        const toCopyright = (copyright: boolean) => {
            if (copyright) return "否";
            else return "是";
        }

        onBeforeMount(() => {
            getVideoList();
        })

        return {
            videos,
            resources,
            activeDrawer,
            toCopyright,
            reviewVideo,
            getResourceList,

        }
    },
    components: {
        NTable,
        NCard,
        NTime,
        NButton,
        NDrawer,
        NDrawerContent,
        VideoList
    }
});
</script>

<style lang="less" scoped>
.card {
    margin: 10px;
    width: calc(100% - 20px);
    border-radius: 10px;
}

.table-head {
    text-align: center;
}

.table-body {
    text-align: center;

    .cover {
        height: 60px;
        width: 100px;
    }

    .btn-list {
        button {
            margin: 0 6px;
        }
    }
}
</style>