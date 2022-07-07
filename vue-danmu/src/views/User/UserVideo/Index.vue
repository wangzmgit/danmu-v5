<template>
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
    <div class="page-box">
        <n-pagination v-model:page="page" :item-count="count" :page-size="8" @update-page="pageChange" />
    </div>
</template>

<script>
import { onBeforeMount, ref } from "vue";
import { NPagination } from "naive-ui";
import { useRouter, useRoute } from "vue-router";
import useUserVideo from "@/hooks/user-video";

export default {


    setup() {
        const uid = ref(0);
        const page = ref(1);
        const route = useRoute();
        const router = useRouter();
        const { count, videoList, getVideoListByUid } = useUserVideo();

        const govideo = (vid) => {
            let videoUrl = router.resolve({ name: "Video", params: { vid: vid } });
            window.open(videoUrl.href, '_blank');
        }

        onBeforeMount(() => {
            uid.value = route.params.uid;
            getVideoListByUid(uid.value, page.value, 8);
        })

        //页码改变
        const pageChange = (target) => {
            page.value = target;
            getVideoListByUid(uid.value, page.value, 8);
        }


        return {
            page,
            count,
            videoList,
            govideo,
            pageChange
        }

    },
    components: {
        NPagination
    }
}
</script>

<style lang="less" scoped>
.card-list {
    width: calc(100% - 20px);
    margin-left: 10px;
    height: 360px;
    display: grid;
    grid-template-rows: repeat(2, 50%);
    grid-template-columns: repeat(4, 25%);

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

.page-box {
    margin: 10px;
}
</style>