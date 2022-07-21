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
</template>

<script lang="ts">
import { defineComponent, onBeforeMount, ref } from "vue";
import { useRouter } from "vue-router";
import { getVideoListAPI } from "@/api/video";
import { baseVideoType } from "@/types/video";

export default defineComponent({
    setup() {
        const router = useRouter();
        const videoList = ref<Array<baseVideoType>>([]);

        const getVideoList = () => {
            getVideoListAPI(1, 10, 0).then((res) => {
                if (res.data.code === 2000) {
                    videoList.value = res.data.data.videos;
                }
            })
        }

        const govideo = (vid:number) => {
            let videoUrl = router.resolve({ name: "Video", params: { vid: vid } });
            window.open(videoUrl.href, '_blank');
        }

        onBeforeMount(() => {
            getVideoList();
        })

        return {
            videoList,
            govideo,
        }
    },

});
</script>

<style lang="less" scoped>
.card-list {
    width: 100%;
    height: 360px;
    display: grid;
    grid-template-rows: repeat(2, 50%);
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
</style>