<template>
    <div class="archive-data">
        <!--点赞收藏-->
        <div class="archive-item">
            <n-icon :class="[likeAnimation, archive.is_like ? 'active' : 'icon']" @click="likeClick">
                <svg t="1655271860793" class="icon" viewBox="0 0 1024 1024" version="1.1"
                    xmlns="http://www.w3.org/2000/svg" p-id="1802" width="200" height="200">
                    <path d="M939.517328 410.237436h-272.638296C772.478372 19.839876 591.9995 0 591.9995 0c-73.59954 0-58.239636 59.519628-63.9996 
                        69.759564 0 191.9988-202.878732 340.477872-202.878732 340.477872v541.436616C325.121168 1004.79372 400.0007 1023.9936 426.880532 
                        1023.9936h409.59744c38.39976 0 69.759564-101.119368 69.759564-101.119368 101.119368-344.957844 101.119368-447.9972 
                        101.119368-447.9972a63.9996 63.9996 0 0 0-68.479572-63.9996zM213.761864 410.237436H49.922888a33.279792 33.279792 0 0 0-33.919788 
                        33.279792l33.919788 545.916588a33.919788 33.919788 0 0 0 34.559784 34.559784h141.439116c29.439816 0 29.439816-23.039856 
                        29.439816-23.039856V451.837176a40.319748 40.319748 0 0 0-41.59974-41.59974z" p-id="1803">
                    </path>
                </svg>
            </n-icon>
            <p>{{ stat.like }}</p>
        </div>
        <div class="archive-item">
            <n-icon :class="archive.is_collect ? 'active' : 'icon'" @click="showCollect = true">
                <svg t="1655271803843" class="icon" viewBox="0 0 1025 1024" version="1.1"
                    xmlns="http://www.w3.org/2000/svg" p-id="1648" width="200" height="200">
                    <path d="M1024 384a103.04 103.04 0 0 0-72.32-64l-215.68-58.24L590.72 49.92A101.12 101.12 0 0 0 512 0a92.8 92.8 0 0 
                        0-78.72 51.84l-120.32 192-161.28 49.28L72.96 320a97.28 97.28 0 0 0-68.48 56.96 115.2 115.2 0 0 0 19.84 
                        97.28l140.8 178.56-7.68 256a117.76 117.76 0 0 0 17.92 88.96 64 64 0 0 0 49.92 21.12 181.76 181.76 0 0 0 
                        60.8-13.44l208-84.48 236.16 87.68a141.44 141.44 0 0 0 42.88 7.68c29.44 0 78.72-13.44 82.56-103.68l-12.8-230.4 
                        152.96-205.44A106.88 106.88 0 0 0 1024 384" p-id="1649"></path>
                </svg>
            </n-icon>
            <p>{{ stat.collect }}</p>
        </div>
    </div>
    <collection-list v-if="showCollect" :vid="vid" @close="closeCollection"></collection-list>
</template>

<script lang="ts">
import { NIcon } from 'naive-ui';
import { statType } from '@/types/archive';
import CollectionList from './CollectionList.vue';
import { reactive, ref, onBeforeMount, defineComponent } from 'vue';
import { getArchiveAPI, likeAPI, dislikeAPI } from '@/api/archive'

export default defineComponent({
    props: {
        vid: {
            type: Number,
            required: true
        },
        stat: {
            type: Object as () => statType,
            required: true
        },
    },
    setup(props) {
        const stat = ref(props.stat);//点赞收藏数据
        //点赞收藏数据
        const archive = reactive({
            is_collect: false,
            is_like: false
        })
        const showCollect = ref(false);
        const likeAnimation = ref('');

        //获取点赞收藏关注信息
        const getArchiveInfo = () => {
            getArchiveAPI(props.vid).then((res) => {
                if (res.data.code === 2000) {
                    const resArchive = res.data.data.archive;
                    archive.is_like = resArchive.is_like;
                    archive.is_collect = resArchive.is_collect;
                }
            })
        }

        //点赞点赞按钮
        const likeClick = () => {
            if (!archive.is_like) {
                //调用点赞接口
                likeAPI(props.vid);
                likeAnimation.value = 'like-active';
                stat.value.like++;
            } else {
                dislikeAPI(props.vid);
                stat.value.like--;
            }
            archive.is_like = !archive.is_like;
        }

        //关闭收藏夹回调
        const closeCollection = (val: number) => {
            if (val === 1) {
                stat.value.collect++;
                archive.is_collect = true;
            } else if (val === -1) {
                stat.value.collect--;
                archive.is_collect = false;
            }
            showCollect.value = false;
        }

        onBeforeMount(() => {
            getArchiveInfo();
        })


        return {
            stat,
            archive,
            showCollect,
            likeAnimation,
            likeClick,
            closeCollection
        }
    },
    components: {
        NIcon,
        CollectionList
    }
});
</script>

<style lang="less" scoped>
.archive-data {
    height: 30px;

    .archive-item {
        float: left;
        user-select: none;
        margin-right: 20px;

        i {
            font-size: 26px;
            line-height: 30px;
            cursor: pointer;
        }

        p {
            font-size: 16px;
            float: right;
            margin: 0 6px;
            line-height: 30px;
        }

        .icon:hover {
            color: #36ad6a;
        }

        .active {
            color: #36ad6a;
        }

        .like-active {
            animation: scaleDraw .3s ease-in-out;
        }
    }
}


@keyframes scaleDraw {
    0% {
        transform: scale(1);
        /*开始为原始大小*/
    }

    25% {
        transform: scale(1.2);
        /*放大1.1倍*/
    }

    100% {
        transform: scale(1);
    }
}
</style>