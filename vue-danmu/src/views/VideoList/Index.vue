<template>
    <div v-title data-title="全部视频">
        <n-affix class="header" :top="0">
            <header-bar></header-bar>
        </n-affix>
        <!-- 分区选择 -->
        <div class="content">
            <div class="partitions">
                <p :class="selectedPartition === 0 ? 'selected' : ''" @click="selectPartition(0)">
                    <span>全部</span>
                </p>
                <p v-for="item in partitions" :key="item.id" @click="selectPartition(item.id)">
                    <span :class="selectedPartition === item.id ? 'selected' : ''">{{ item.content }}</span>
                </p>
            </div>
            <div class="partitions">
                <p :class="selectedSubpartition == 0 ? 'selected' : ''" @click="selectSubartition(0)">
                    <span>全部</span>
                </p>
                <p v-for="item in subpartition" :key="item.id" @click="selectSubartition(item.id)">
                    <span :class="selectedSubpartition == item.id ? 'selected' : ''">{{ item.content }}</span>
                </p>
            </div>
        </div>
        <!-- 视频列表 -->
        <div class="card-list">
            <div class="card-box" v-for="(item, index) in videoList" :key="index">
                <div class="card" @click="govideo(item.vid)">
                    <img class="cover" :src="item.cover" />
                    <div class="info">
                        <p class="title">{{ item.title }}</p>
                    </div>
                </div>
            </div>
        </div>
        <!-- 加载更多 -->
        <div v-if="!noMore" class="more-btn">
            <n-button @click="getMore">加载更多</n-button>
        </div>
    </div>
</template>

<script lang="ts">
import { baseVideoType } from '@/types/video';
import { partitionType } from '@/types/partition';
import { getVideoListAPI } from "@/api/video";
import { getPartitionAPI } from "@/api/partition";
import { useRoute, useRouter } from "vue-router";
import { reactive, ref, toRefs, onBeforeMount, defineComponent } from "vue";
import HeaderBar from "@/components/HeaderBar.vue";
import { NAffix, NButton, useNotification } from "naive-ui";

export default defineComponent({
    setup() {
        const partitionInfo = reactive({
            partitions: [] as partitionType[], //分区内容
            subpartition: [] as partitionType[], //子分区内容
            selectedPartition: 0,//选中分区
            selectedSubpartition: 0,//选中子分区
        });
        const route = useRoute();
        const router = useRouter();
        const notification = useNotification();

        //获取分区列表
        const getPartitionList = (fid: number) => {
            getPartitionAPI(fid).then((res) => {
                if (res.data.code === 2000) {
                    if (fid === 0) {
                        partitionInfo.partitions = res.data.data.partitions;
                    } else {
                        partitionInfo.subpartition = res.data.data.partitions;
                    }
                }
            }).catch(() => {
                notification.error({
                    title: "分区加载失败",
                    duration: 5000
                })
            });
        }

        //设置分区
        const selectPartition = (id: number) => {
            if (partitionInfo.selectedPartition === id) return;
            let newQuery = JSON.parse(JSON.stringify(route.query));
            newQuery.partition = id;
            partitionInfo.selectedPartition = id;
            partitionInfo.selectedSubpartition = 0;
            router.replace({ query: newQuery });
            if (id === 0) {
                partitionInfo.subpartition = [];
            } else {
                getPartitionList(id);
            }
            getVideoList(true);
        }

        //设置子分区
        const selectSubartition = (id: number) => {
            if (partitionInfo.selectedSubpartition === id) return;
            let newQuery = JSON.parse(JSON.stringify(route.query));
            newQuery.subpartition = id;
            partitionInfo.selectedSubpartition = id;
            router.replace({ query: newQuery });

            getVideoList(true);
        }

        const page = ref(1);
        const videoList = ref<Array<baseVideoType>>([]);
        const noMore = ref(false);
        const getVideoList = (init = false) => {
            if (init) {
                page.value = 1;
                videoList.value = [];
                noMore.value = false;
            }
            const current = partitionInfo.selectedSubpartition || partitionInfo.selectedPartition;
            getVideoListAPI(page.value, 15, current).then((res) => {
                if (res.data.code === 2000) {
                    if (res.data.data.videos) {
                        videoList.value.push(...res.data.data.videos);
                        if (res.data.data.videos.length < 15) {
                            noMore.value = true;
                        }
                    } else {
                        noMore.value = true;
                    }
                }
            });
        }

        //加载更多
        const getMore = () => {
            page.value++;
            getVideoList();
        }

        //前往视频详情
        const govideo = (vid: number) => {
            let videoUrl = router.resolve({ name: "Video", params: { vid: vid } });
            window.open(videoUrl.href, '_blank');
        }

        onBeforeMount(() => {
            getPartitionList(0);
            partitionInfo.selectedPartition = Number(route.query.partition);
            partitionInfo.selectedSubpartition = Number(route.query.subpartition);
            if (partitionInfo.selectedPartition !== 0) {
                getPartitionList(partitionInfo.selectedPartition);
            }
            getVideoList();
        })

        return {
            noMore,
            videoList,
            getMore,
            govideo,
            ...toRefs(partitionInfo),
            selectPartition,
            selectSubartition,
            getPartitionList
        }
    },
    components: {
        NAffix,
        NButton,
        HeaderBar,
    },
});
</script>

<style lang="less" scoped>
.header {
    z-index: 999;
    width: 100%;
}

.content {
    width: 90%;
    min-width: 700px;
    margin: 90px auto 10px auto;
}

.partitions {
    display: flex;
    margin-top: 20px;

    p {
        cursor: pointer;
        margin: 0 20px;

        span {
            padding: 3px 6px;
        }
    }

    .selected {
        color: #fff;
        border-radius: 6px;
        background-color: rgba(24, 143, 255, 0.5);
    }
}

.card-list {
    width: 90%;
    margin: 10px auto;
    display: grid;
    grid-template-columns: repeat(5, 20%);

    .card-box {
        margin: 0 auto;
        height: 100%;
        width: calc(100% - 10px);

        .card {
            position: relative;
            width: 100%;
            height: 150px;
            margin-top: 15px;
            border-radius: 6px;

            .cover {
                width: 100%;
                height: 100%;
                border-radius: 6px;
            }

            .info {
                color: #fff;
                overflow: hidden;
                position: absolute;
                bottom: 0;
                width: 100%;
                border-bottom-left-radius: inherit;
                border-bottom-right-radius: inherit;
                background: linear-gradient(0deg, rgba(0, 0, 0, 0.7), transparent);

                .title {
                    padding-left: 10px;
                    font-size: 14px;
                    line-height: 16px;
                    overflow: hidden;
                    text-overflow: ellipsis;
                    display: -webkit-box;
                    -webkit-line-clamp: 1;
                    -webkit-box-orient: vertical;
                }
            }
        }
    }
}

.more-btn {
    text-align: center;
}
</style>