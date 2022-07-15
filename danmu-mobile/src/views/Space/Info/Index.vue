<template>
    <div v-title data-title="个人中心"></div>
    <header-bar></header-bar>
    <div class="user-card">
        <div class="card-avatar">
            <common-avatar :url="userInfo.avatar" :size="76" :iconsize="50"></common-avatar>
        </div>
        <!--关注粉丝信息-->
        <div class="card-info-box">
            <div class="card-info">
                <div>
                    <p class="info-title">投稿</p>
                    <p class="info-content">{{ count }}</p>
                </div>
                <div>
                    <p class="info-title">关注</p>
                    <p class="info-content">{{ following }}</p>
                </div>
                <div>
                    <p class="info-title">粉丝</p>
                    <p class="info-content">{{ followers }}</p>
                </div>
            </div>
            <n-button type="primary" ghost @click="goPage('EditInfo')">编辑信息</n-button>
        </div>
    </div>
    <div class="card-name">
        <p class="name">{{ userInfo.name }}</p>
        <p class="sign">{{ userInfo.sign }}</p>
    </div>

    <div ref="scrollBox" class="video-list" @scroll="scrollList">
        <video-list :videos="videoList"></video-list>
        <div v-show="loading" class="loading">
            <n-spin></n-spin>
        </div>
    </div>
</template>

<script lang="ts">
import { getMyVideoAPI } from '@/api/video';
import { ref, computed, onMounted, defineComponent } from "vue";
import storage from "@/utils/stored-data";
import { NButton, NSpin, useMessage } from 'naive-ui';
import { useRouter } from "vue-router";
import { getFollowDataAPI } from '@/api/follow';
import { userInfoType } from "@/types/user";
import CommonAvatar from "@/components/CommonAvatar.vue";
import HeaderBar from "@/components/HeaderBar.vue";
import { videoType } from '@/types/video';
import VideoList from '@/components/VideoList.vue'

export default defineComponent({
    setup() {
        const count = ref(0);
        const followers = ref(0);
        const following = ref(0);
        const router = useRouter();
        const message = useMessage();

        const userInfo = computed((): userInfoType => {
            return storage.get("userInfo") as userInfoType;
        })

        //前往关注和粉丝页面
        const goPage = (name: string) => {
            router.push({ name: name });
        }

        //获取关注数和粉丝数
        const getFollowData = (id: number) => {
            getFollowDataAPI(id).then((res) => {
                if (res.data.code === 2000) {
                    followers.value = res.data.data.followers
                    following.value = res.data.data.following;
                }
            }).catch(() => {
                message.error('获取失败');
            });
        }

        //视频相关
        const page = ref(1);
        const pageSize = 10;
        const noMore = ref(false);
        const loading = ref(false);
        const videoList = ref<Array<videoType>>([]);
        const scrollBox = ref<HTMLElement | null>(null);
        const scrollList = (() => {
            //节流
            var timer: number | null = null;
            return () => {
                if (timer) {
                    return
                }

                timer = setTimeout(() => {
                    const scrollTop = scrollBox.value?.scrollTop || 0;
                    const clientHeight = scrollBox.value?.clientHeight || 0;
                    const scrollHeight = scrollBox.value?.scrollHeight || 0;
                    if (scrollTop + clientHeight >= scrollHeight - 50) {
                        if (!noMore.value && !loading.value) {
                            page.value++;
                            getMyVideo();
                        }
                    }
                    timer = null;
                }, 500);
            }
        })();

        //获取我的视频
        const getMyVideo = () => {
            getMyVideoAPI(page.value, pageSize).then((res) => {
                const resVideos = res.data.data.videos;
                if (resVideos.length < pageSize) {
                    noMore.value = true;
                    message.info("没有更多了");
                }
                count.value = res.data.data.count;
                videoList.value = videoList.value.concat(resVideos);
            }).catch(() => {
                message.error('获取失败');
            });
        }

        onMounted(() => {
            getFollowData(userInfo.value!.uid);
            getMyVideo();
        })


        return {
            count,
            loading,
            followers,
            following,
            userInfo,
            scrollBox,
            videoList,
            goPage,
            scrollList
        }
    },
    components: {
        NSpin,
        NButton,
        CommonAvatar,
        VideoList,
        HeaderBar
    },
});
</script>
<style lang="less" scoped>
.user-card {
    display: flex;
    height: 120px;
    align-items: flex-start;

    .card-avatar {
        width: 160px;

        span {
            margin: 16px 30px;
        }
    }

    .card-info-box {
        .card-info {
            width: 210px;
            display: flex;
            justify-content: space-around;

            div {
                width: 70px;
                text-align: center;

                .info-title {
                    margin-bottom: 2px;
                    // font-weight: bold;
                }

                .info-content {
                    margin: 2px 0 6px 0;
                }
            }
        }

        button {
            width: 100%;
        }
    }
}

.card-name {
    margin: 0 10px;

    .name {
        margin: 0;
        font-size: 18px;
        font-weight: 500;
    }

    .sign {
        margin: 3px 0;
        font-size: 14px;
        color: #666;
    }
}

.video-list {
    height: calc(100vh - 200px);
    overflow: scroll;

    .loading {
        padding: 20px 0;
        text-align: center;
    }
}
</style>