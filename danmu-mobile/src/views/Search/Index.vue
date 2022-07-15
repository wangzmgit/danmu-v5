<template>
    <div v-title data-title="搜索">
        <header-bar></header-bar>
        <video-list :videos="videoList"></video-list>
    </div>
</template>

<script lang="ts">
import { useMessage } from "naive-ui";
import { searchVideoAPI } from "@/api/search";
import { useRoute, useRouter } from "vue-router";
import VideoList from "@/components/VideoList.vue";
import HeaderBar from "@/components/HeaderBar.vue";
import { defineComponent, onBeforeMount, ref, watch } from "vue";

export default defineComponent({
    setup() {
        const route = useRoute();
        const router = useRouter();
        const keywords = ref("");
        const videoList = ref([]);
        const message = useMessage();

        const searchVideo = () => {
            searchVideoAPI(keywords.value).then((res) => {
                if (res.data.code === 2000) {
                    videoList.value = res.data.data.videos;
                }
            }).catch(() => {
                message.error("获取失败");
            });
        }

        //关键词高亮
        const keyHighlight = (title: string) => {
            let res = '';
            let indexArr: Array<number> = []; // 需要标红的字的下标数组
            const keywordsArray = keywords.value.split(" ");
            const getReplaceStr = (str: string) => `<font color="#409EFF">${str}</font>`;
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

        const govideo = (vid: number) => {
            router.push({ name: "Video", params: { vid: vid } });
        }

        watch(() => route.query.keywords, (newValue) => {
            keywords.value = (newValue || "").toString();
            searchVideo();
        })

        onBeforeMount(() => {
            keywords.value = (route.query.keywords || "").toString();
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
        HeaderBar,
        VideoList
    },
});
</script>