<template>
    <div v-title data-title="搜索">
        <header-bar></header-bar>
        <div class="search">
            <n-input round placeholder="搜索关键词~" v-model:value="keywords" @keydown.enter="searchVideo">
                <template #suffix>
                    <n-icon @click="searchVideo">
                        <Search />
                    </n-icon>
                </template>
            </n-input>
        </div>
        <div class="card-list">
            <div class="card-box" v-for="(item, index) in videoList" :key="index">
                <div class="card" @click="govideo(item.vid)">
                    <img class="cover" :src="item.cover" />
                    <div class="info">
                        <p class="title" v-dompurify-html="keyHighlight(item.title)"></p>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script lang="ts">
import { searchVideoAPI } from "@/api/search";
import HeaderBar from "@/components/HeaderBar.vue";
import { defineComponent, onBeforeMount, ref } from "vue";
import { NInput, NIcon, useNotification } from "naive-ui";
import { useRoute, useRouter } from "vue-router";
import { Search } from '@vicons/ionicons5';
import { baseVideoType } from "@/types/video";

export default defineComponent({
    setup() {
        const route = useRoute();
        const router = useRouter();
        const keywords = ref("");
        const videoList = ref<Array<baseVideoType>>([]);
        const notification = useNotification();

        const searchVideo = () => {
            searchVideoAPI(keywords.value).then((res) => {
                if (res.data.code === 2000) {
                    videoList.value = res.data.data.videos;
                }
            }).catch((err) => {
                notification.error({
                    title: "获取失败",
                    duration: 5000,
                })
            });
        }

        //关键词高亮
        const keyHighlight = (title:string) => {
            let res = '';
            let indexArr:Array<number> = []; // 需要标红的字的下标数组
            const keywordsArray = keywords.value.split(" ");
            const getReplaceStr = (str:string) => `<font color="#409EFF">${str}</font>`;
            keywordsArray.forEach((keyword) => {
                let filterStr = title;
                let stopFlag = false;
                while (!stopFlag && filterStr && keyword) {
                    const index = filterStr.indexOf(keyword); // 返回匹配的第一个字符的下标
                    if (index === -1) stopFlag = true;
                    else {
                        keyword.split("").forEach((s, i) => {
                            indexArr.push(index + Number(i));
                        });
                        filterStr =
                            filterStr.substring(0, index) +
                            " " +
                            filterStr.substring(index + 1);
                    }
                }
            });
            indexArr = Array.from(new Set(indexArr)); // 去重
            title.split("").forEach((char, charIndex) => {
                res += indexArr.includes(charIndex) ? getReplaceStr(char) : char;
            });
            return res;
        }

        const govideo = (vid:number) => {
            let videoUrl = router.resolve({ name: "Video", params: { vid: vid } });
            window.open(videoUrl.href, '_blank');
        }

        onBeforeMount(() => {
            keywords.value = route.params.keywords.toString();
            searchVideo();
        })

        return {
            keywords,
            videoList,
            govideo,
            searchVideo,
            keyHighlight
        }
    },
    components: {
        NIcon,
        NInput,
        Search,
        HeaderBar,
    },
});
</script>

<style lang="less" scoped>
.search {
    width: 100%;
    margin: 30px auto;
    width: 500px;
}

.card-list {
    width: calc(100% - 100px);
    margin-left: 50px;
    height: 500px;
    min-width: 700px;
    display: grid;
    grid-template-rows: repeat(3, 33.3%);
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