<template>
    <div class="collection" v-title data-title="收藏夹详情">
        <header-bar></header-bar>
        <div class="collection-main">
            <div v-if="loadingContent" class="content-left">
                <div class="card-item" v-for="(item, index) in videoList" :key="index">
                    <div class="card-left">
                        <img v-if="item.cover" :src="item.cover" alt="收藏夹封面">
                        <div class="card-no-cover" v-else>
                            <img src="@/assets/logo.png" alt="默认封面">
                        </div>
                    </div>
                    <div class="card-center">
                        <p class="title" @click="goVideo(item.vid)">{{ item.title }}</p>
                        <span class="desc">播放：{{ item.clicks }}</span>
                        <span class="desc">简介：{{ item.desc }}</span>
                    </div>
                    <div class="card-right" v-if="uid === collection!.author.uid">
                        <n-icon class="edit" size="20" @click="removeVideo(item.vid)">
                            <trash-outline />
                        </n-icon>
                    </div>
                </div>
                <div class="page-box">
                    <n-pagination v-model:page="page" :item-count="count" :page-size="8" @update-page="pageChange" />
                </div>
            </div>
            <div v-else class="content-left">
                <n-empty description="什么都没有"> </n-empty>
            </div>
            <div v-if="loadingInfo" class="content-right">
                <div class="cover">
                    <img v-if="collection!.cover" :src="collection!.cover" alt="收藏夹封面">
                    <div class="no-cover" v-else>
                        <img src="@/assets/logo.png" alt="默认封面">
                    </div>
                </div>
                <div class="info">
                    <span class="title">{{ collection!.name }}</span>
                    <span class="desc">简介：{{ collection!.desc }}</span>
                    <div class="desc">
                        <n-time type="date" :time="new Date(collection!.created_at)"></n-time>
                        <span>・</span>
                        <span class="open">{{ collection!.open ? '公开' : '私密' }}</span>
                        <span class="author" @click="goSpace(collection!.author.uid)">{{ collection!.author.name }}</span>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script lang="ts">
import storage from '@/utils/stored-data';
import { useRoute, useRouter } from 'vue-router';
import { TrashOutline } from "@vicons/ionicons5";
import HeaderBar from '@/components/HeaderBar.vue';
import { onBeforeMount, ref, defineComponent } from 'vue';
import { NTime, NIcon, NPagination, NEmpty, useNotification } from 'naive-ui';
import { getCollectionInfoAPI, getCollectVideoAPI, collectAPI } from '@/api/collect';
import { videoType } from '@/types/video';
import { collectionInfoType } from '@/types/collect';

export default defineComponent({
    setup() {
        const id = ref(0);
        const uid = ref(0);//用户ID
        const route = useRoute();
        const collection = ref<collectionInfoType>();
        const loadingInfo = ref(false);
        const loadingContent = ref(false);
        const notification = useNotification();

        const getCollectionInfo = (id: number) => {
            getCollectionInfoAPI(id).then((res) => {
                if (res.data.code === 2000) {
                    collection.value = res.data.data.collection;
                    loadingInfo.value = true;
                }
            })
        }

        const page = ref(1);
        const count = ref(0);
        const videoList = ref<Array<videoType>>([]);
        //获取收藏视频内容
        const getCollectVideo = () => {
            getCollectVideoAPI(id.value, page.value, 8).then((res) => {
                if (res.data.code === 2000) {
                    count.value = res.data.data.count;
                    videoList.value = res.data.data.videos;
                    if (count.value > 0) {
                        loadingContent.value = true;
                    }
                }
            })
        }

        //页码改变
        const pageChange = (target: number) => {
            page.value = target;
            getCollectVideo();
        }

        //移除视频
        const removeVideo = (vid: number) => {
            collectAPI({ vid, addList: [], cancelList: [id.value] }).then((res) => {
                if (res.data.code === 2000) {
                    getCollectVideo();
                }
            }).catch((err) => {
                notification.error({
                    title: '移除失败',
                    content: "原因:" + err.response.data.msg,
                    duration: 5000,
                })
            });
        }

        //页面跳转
        const router = useRouter();
        const goSpace = (uid: number) => {
            let userUrl = router.resolve({ name: "User", params: { uid: uid } });
            window.open(userUrl.href, '_blank');
        }

        const goVideo = (vid: number) => {
            let videoUrl = router.resolve({ name: "Video", params: { vid: vid } });
            window.open(videoUrl.href, '_blank');
        }

        onBeforeMount(() => {
            if (storage.get('userInfo') && storage.get('userInfo').uid) {
                uid.value = storage.get('userInfo').uid;
            }
            id.value = Number(route.params.id);
            getCollectionInfo(id.value);
            getCollectVideo();
        })

        return {
            uid,
            page,
            count,
            videoList,
            collection,
            loadingInfo,
            loadingContent,
            goSpace,
            goVideo,
            pageChange,
            removeVideo
        }
    },
    components: {
        NTime,
        NIcon,
        NEmpty,
        NPagination,
        HeaderBar,
        TrashOutline
    }
});
</script>

<style lang="less" scoped>
.collection {
    height: 100%;
    width: 100%;
    top: 0;
    bottom: 0;
    position: fixed;
    overflow-y: scroll;
}

.collection-main {
    width: calc(100% - 200px);
    background: #fff;
    margin: 20px auto 0 auto;
    display: flex;
    flex-wrap: nowrap;
    justify-content: space-between;
}

.content-left {
    width: calc(100% - 380px);
    min-width: 680px;
}

.card-item {
    width: 100%;
    height: 80px;
    margin-bottom: 12px;
    border-bottom: 1px solid #e6e6e6;
    padding-bottom: 12px;

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

        .card-no-cover {
            width: 100%;
            height: 100%;
            background-color: #e6e6e6;
            border-radius: 2px;

            img {
                position: absolute;
                width: 50px;
                height: 50px;
                margin: 15px 35px;
                filter: grayscale(100%);
                opacity: 50%;
            }
        }
    }

    .card-center {
        float: left;
        width: calc(100% - 230px);

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
        width: 90px;
        height: 100%;
        text-align: center;


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
}



/**右半部分 */
.content-right {
    width: 350px;
    height: 600px;
    min-width: 350px;

    .cover {
        width: 100%;
        height: 200px;
        margin-bottom: 12px;

        img {
            width: 100%;
            height: 100%;
            border-radius: 2px;
        }

        .no-cover {
            width: 100%;
            height: 100%;
            background-color: #e6e6e6;
            border-radius: 2px;

            img {
                position: absolute;
                width: 100px;
                height: 100px;
                margin: 50px 125px;
                filter: grayscale(100%);
                opacity: 50%;
            }
        }
    }

    .info {
        .title {
            display: block;
            height: auto;
            font-size: 16px;
            color: #212121;
            line-height: 18px;
            margin-bottom: 20px;
        }

        .desc {
            display: block;
            font-size: 12px;
            font-size: 14px;
            color: #666;
            margin: 2px 0;

            .author {
                float: right;
                cursor: pointer;

                &:hover {
                    color: #36ad6a;
                }
            }
        }


    }
}
</style>