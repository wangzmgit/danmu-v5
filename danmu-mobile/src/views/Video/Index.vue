<template>
    <div class="video" v-title :data-title="`视频播放-${title}`">
        <header-bar></header-bar>
        <div class="video-main">
            <div v-if="!loading" class="video-player">
                <!-- <w-player :key="part" :vid="videoInfo!.vid" :part="part" :resource="options" /> -->
                <video-player :vid="vid" :part="part" :resources="resources"></video-player>
                <!-- 标题 -->
                <div class="video-title">
                    <span :class="['title', fold ? 'title-fold' : '']">{{ videoInfo!.title }}</span>
                    <div class="fold" @click="fold = !fold">
                        <n-icon color='#666'>
                            <caret-down v-if="fold" />
                            <caret-up v-else />
                        </n-icon>
                    </div>
                </div>
                <!-- 其他信息 -->
                <div v-if="fold" class="author-fold">
                    <div class="info-fold">
                        <common-avatar :size="32" :url="videoInfo!.author.avatar"></common-avatar>
                        <span class="name">{{ videoInfo!.author.name }}</span>
                    </div>
                    <n-time type="relative" :time="new Date(videoInfo!.created_at)"></n-time>
                </div>
                <div class="video-info" v-else>
                    <div class="user-info">
                        <common-avatar :size="50" :url="videoInfo!.author.avatar"></common-avatar>
                        <p class="name">{{ videoInfo!.author.name }}</p>
                        <span class="date">发布于<n-time :time="new Date(videoInfo!.created_at)"></n-time></span>
                    </div>
                    <p v-show="videoInfo!.copyright" class="info-item">
                        <n-icon color='#fd6d6f'>
                            <ban-sharp />
                        </n-icon>
                        <span>未经作者授权，禁止转载</span>
                    </p>
                    <p class="info-item">{{ videoInfo!.desc }}</p>
                </div>
                <!-- 视频分集 -->
                <div v-if="resources.length > 1">
                    <part-list :resources="resources" :active="part" @change="changePart"></part-list>
                </div>
                <n-divider></n-divider>
                <!--发表评论-->
                <comment :vid="videoInfo!.vid"></comment>
            </div>
            <div v-else class="player-skeleton">
                <n-skeleton height="200px" />
                <n-skeleton text :repeat="2" />
                <n-skeleton text width=" 60%" />
            </div>
        </div>
    </div>
</template>

<script lang="ts">
import config from '@/config';
import { useRoute, useRouter } from 'vue-router';
import { NIcon, NTime, NSkeleton, NDivider } from 'naive-ui';
import { defineComponent, onBeforeMount, reactive, ref } from 'vue';
import { getVideoInfoAPI } from '@/api/video';
import { BanSharp, CaretUp, CaretDown } from '@vicons/ionicons5';
import Comment from './components/Comment.vue';
import VideoPlayer from './components/VideoPlayer.vue';
import PartList from './components/PartList.vue';
import HeaderBar from '@/components/HeaderBar.vue';
import CommonAvatar from '@/components/CommonAvatar.vue';
import { videoType } from '@/types/video';

export default defineComponent({
    setup() {
        const route = useRoute();
        const router = useRouter();
        const vid = parseInt(route.params.vid.toString());
        const options = reactive({
            type: "hls",
            data: {}
        })
        const title = config.title;
        const part = ref(1);//当前分集
        const fold = ref(true);//折叠
        const loading = ref(true);
        const resources = ref([]);
        const videoInfo = ref<videoType | null>(null);

        //获取视频信息
        const getVideoInfo = (vid: number) => {
            getVideoInfoAPI(vid).then((res) => {
                if (res.data.code === 2000) {
                    videoInfo.value = res.data.data.video;
                    resources.value = res.data.data.video.resources;
                    //设置播放的资源
                    if (!resources.value[part.value - 1]) {
                        part.value = 1;
                    }
                    options.data = resources.value[part.value - 1];

                    //修改网站标题
                    document.title = `${res.data.data.video.title}-${title}`
                    loading.value = false;
                }
            })
        }

        const changePart = (target: number) => {
            if (resources.value[target - 1]) {
                part.value = target;
            }
            options.data = resources.value[part.value - 1];
            router.replace({ query: { p: part.value } });
        }

        onBeforeMount(() => {
            getVideoInfo(parseInt(route.params.vid.toString()));
            if (route.query.p) {
                part.value = Number(route.query.p);
            }
        })


        return {
            vid,
            fold,
            part,
            title,
            options,
            loading,
            resources,
            videoInfo,
            changePart
        }
    },
    components: {
        VideoPlayer,
        Comment,
        NTime,
        NIcon,
        NDivider,
        NSkeleton,
        CaretUp,
        CaretDown,
        BanSharp,
        PartList,
        HeaderBar,
        CommonAvatar,
    }
});
</script>

<style lang="less" scoped>
.video {
    height: 100%;
    width: 100%;
    top: 0;
    bottom: 0;
    position: fixed;
    overflow-y: scroll;
}

.video-main {
    width: 100%;
    background: #fff;
    // margin: 20px auto 0 auto;
    display: flex;
    flex-wrap: nowrap;

    .video-player {
        width: 100%;
    }

    //骨架占位
    .player-skeleton {
        margin: 0 auto;
        width: 100%;
    }
}

//标题
.video-title {
    display: flex;
    justify-content: space-between;
    margin: 10px 0;

    .title {
        color: #333;
        line-height: 22px;
        font-weight: 700;
        font-size: 20px;
        margin-left: 6px;
    }

    .title-fold {
        overflow: hidden;
        text-overflow: ellipsis;
        display: -webkit-box;
        -webkit-line-clamp: 1;
        -webkit-box-orient: vertical;
    }

    .fold {
        width: 40px;
        text-align: center;
    }
}



/**折叠状态 */
.author-fold {
    width: calc(100% -20px);
    padding: 6px 10px;
    display: flex;
    align-items: center;
    justify-content: space-between;

    .info-fold {
        display: flex;
        align-items: center;

        .name {
            padding-left: 6px;
        }
    }
}

/**展开 */
.video-info {
    .user-info {
        position: relative;
        margin-left: 6px;

        .name {
            position: absolute;
            top: -12px;
            left: 60px;
            color: #383838;
            font-size: 16px;
        }


        .date {
            position: absolute;
            color: #666;
            top: 30px;
            left: 60px;
            font-size: 14px;
        }
    }

    .info-item {
        color: #666;
        margin: 4px 0 2px 6px;
    }
}
</style>