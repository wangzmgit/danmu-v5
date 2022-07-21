<template>
    <header-bar></header-bar>
    <div class="home-box">
        <div class="home-top">
            <div class="home-top-left">
                <carousel></carousel>
            </div>
            <!--推荐视频-->
            <div class="home-top-right">
                <recommend></recommend>
            </div>
        </div>
        <div class="transition">
            <span>视频列表</span>
            <n-button @click="viewMore">查看更多</n-button>
        </div>
        <video-list></video-list>
    </div>
    <footer-bar class="footer"></footer-bar>
    <login v-if="showLogin" @close="closeLoginCard"></login>
</template>

<script lang="ts">
import { watch, ref, defineComponent } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { NButton } from 'naive-ui';
import Carousel from './components/Carousel.vue';
import Recommend from './components/Recommend.vue';
import HeaderBar from '@/components/HeaderBar.vue';
import Login from '@/components/LoginCard/Index.vue';
import VideoList from './components/VideoList.vue';
import FooterBar from '@/components/FooterBar.vue';

export default defineComponent({
    setup() {
        const route = useRoute();
        const router = useRouter();
        const showLogin = ref(false);
        const closeLoginCard = () => {
            showLogin.value = false;
            router.push({ name: 'Home', params: { login: "false" } });
        }

        const viewMore = () => {
            router.push({
                name: "VideoList",
                query: { partition: 0, subpartition: 0 }
            })
        }

        watch(() => route.params.login, (_newValue, _oldValue) => {
            if (route.params.login === 'true') {
                showLogin.value = true;
            }
        })

        return {
            showLogin,
            viewMore,
            closeLoginCard
        }
    },
    components: {
        Login,
        NButton,
        Carousel,
        Recommend,
        HeaderBar,
        VideoList,
        FooterBar
    }
});
</script>

<style lang="less" scoped>
.home-box {
    width: 80%;
    min-width: 1200px;
    margin: 10px auto;

    .home-top {
        display: flex;
        height: 300px;

        .home-top-left {
            padding-top: 10px;
            width: 40%;
            min-width: 500px;
        }

        .home-top-right {
            width: calc(60% - 20px);
            min-width: 680px;
            padding: 10px;
        }
    }
}

.transition {
    display: flex;
    align-items: center;
    justify-content: space-between;
    margin-left: 10px;

    span {
        color: #181818;
        font-size: 20px;
    }
}

.footer {
    margin-top: 20px;
}
</style>
