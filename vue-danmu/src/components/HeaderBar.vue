<template>
    <div class="header-bar">
        <div class="header-content">
            <div class="header-left" @click="goPage('Home')">
                <img class="logo" src="@/assets/logo.png" alt="logo">
                <span class="title">{{ title }}</span>
            </div>
            <div v-show="route.name !== 'Search'" class="header-center">
                <n-input v-model:value="keywords" round placeholder="搜索" @keydown.enter="search">
                    <template #suffix>
                        <n-icon @click="search">
                            <Search />
                        </n-icon>
                    </template>
                </n-input>
            </div>
            <div class="header-right">
                <!-- 用户头像 -->
                <div v-if="login" class="avatar-box">
                    <div @click="goPage('Space')">
                        <common-avatar :url="userInfo.avatar" :size="40" :iconSize="22"></common-avatar>
                    </div>
                    <div class="header-menu">
                        <div class="menu-item">
                            <span class="btn" @click="goPage('Space')">{{ userInfo.name }}</span>
                        </div>
                        <div class="menu-item">
                            <span class="btn" @click="logout">退出登录</span>
                        </div>
                    </div>
                </div>
                <div v-else>
                    <n-button text @click="goLogin">登录/注册</n-button>
                </div>
                <!-- 图形按钮 -->
                <div class="icon-btn" @click="goPage('Message')">
                    <n-icon class="icon">
                        <MailOutline />
                    </n-icon>
                    <div class="icon-action">
                        <div class="icon-action-msg">消息</div>
                    </div>
                </div>
                <div class="icon-btn" @click="goPage('Collection')">
                    <n-icon class="icon">
                        <BookmarkOutline />
                    </n-icon>
                    <div class="icon-action">
                        <div class="icon-action-msg">收藏</div>
                    </div>
                </div>
                <!-- 投稿按钮 -->
                <n-button class="upload-btn" color="#36ad6a" @click="goPage('Upload')">
                    <template #icon>
                        <n-icon>
                            <CloudUploadOutline />
                        </n-icon>
                    </template>
                    投稿
                </n-button>
            </div>
        </div>
    </div>
    <Login v-if="showLogin" @close="closeCard" />
</template>

<script lang="ts">

import config from '@/config';
import { storeToRefs } from 'pinia';
import { computed, defineComponent, ref } from 'vue';
import storage from '@/utils/stored-data';
import { useRouter, useRoute } from 'vue-router';
import { useLoginStore } from '@/store/login-store';
import Login from '@/components/LoginCard/Index.vue';
import CommonAvatar from './CommonAvatar.vue';
import { NInput, NIcon, NButton } from 'naive-ui';
import { Search, CloudUploadOutline, MailOutline, BookmarkOutline } from '@vicons/ionicons5';

export default defineComponent({
    setup() {
        const title = config.title;
        const showLogin = ref(false);
        const router = useRouter();
        const route = useRoute();
        const loginStore = useLoginStore();
        const { login } = storeToRefs(loginStore);

        const closeCard = () => {
            showLogin.value = false;
        }

        const goLogin = () => {
            showLogin.value = true;
        }

        const logout = () => {
            storage.remove('token');
            storage.remove('userInfo');
            loginStore.setLoginState(false);
        }

        const userInfo = computed(() => {
            const userInfo = storage.get("userInfo");
            if (userInfo) {
                return userInfo;
            }
            loginStore.setLoginState(false);
            return {};
        })

        const goPage = (name: string) => {
            router.push({ name: name });
        }

        //搜索功能
        const keywords = ref("");
        const search = () => {
            let searchUrl = router.resolve({ name: "Search", params: { keywords: keywords.value } });
            window.open(searchUrl.href, '_blank');
        }


        return {
            title,
            login,
            route,
            keywords,
            userInfo,
            showLogin,
            search,
            logout,
            goLogin,
            goPage,
            closeCard,
        }
    },
    components: {
        Login,
        NIcon,
        NInput,
        NButton,
        CommonAvatar,
        //图标
        Search,
        MailOutline,
        BookmarkOutline,
        CloudUploadOutline,
    }
});
</script>

<style lang="less" scoped>
.header-bar {
    display: flex;
    align-items: center;
    width: 100%;
    height: 60px;
    min-width: 1200px;
    background-color: #fff;
    -webkit-box-shadow: 0px 0px 3px #c8c8c8;
    -moz-box-shadow: 0px 0px 3px #c8c8c8;
    box-shadow: 0px 0px 3px #c8c8c8;

    .header-content {
        margin: 0 auto;
        width: 1200px;
        height: 50px;
        display: flex;
        justify-content: space-between;

        .header-left {
            display: flex;
            align-items: center;
            padding-top: 5px;

            .logo {
                height: 50px;
                width: 50px;
            }

            .title {
                color: rgba(0, 0, 0, .85);
                display: inline-block;
                margin: 0 0 0 12px;
                font-size: 16px;
                font-weight: 500;
                vertical-align: top;
            }
        }


        .header-center {
            width: 300px;
            margin-top: 8px;
        }

        .header-right {
            display: flex;
            align-items: center;

            .avatar-box {
                padding-top: 5px;
                margin-right: 6px;
            }

            .upload-btn {
                margin-left: 6px;
            }
        }

    }
}


.icon-btn {
    width: 50px;
    height: 50px;
    margin: 0 3px;
    position: relative;
    display: inline-flex;

    display: flex;
    align-items: center;
    justify-content: center;

    .icon {
        cursor: pointer;
        font-size: 20px;
        transition: font .46s cubic-bezier(.4, 0, .2, 1), transform .46s cubic-bezier(.4, 0, .2, 1);
    }

    .icon-action {
        position: absolute;
        top: calc(50% + 10px);
        opacity: 0;
        transition: opacity .46s cubic-bezier(.4, 0, .2, 1), transform .46s cubic-bezier(.4, 0, .2, 1);

        .icon-action-msg {
            cursor: pointer;
            font-size: 12px;
            color: #36ad6a;
        }
    }

    &:hover {
        .icon-action {
            opacity: 1;
            transform: translateY(-10px);
        }

        .icon {
            font-size: 16px;
            color: #36ad6a;
            transform: translateY(-10px);
        }
    }
}

.avatar-box {
    position: relative;
    cursor: pointer;
    margin: 0 10px;

    &:hover {
        .header-menu {
            display: block;
        }
    }

    .header-menu {
        display: none;
        width: 200px;
        height: 100px;
        top: 40px;
        left: -80px;
        position: absolute;
        z-index: 999;
        background-color: #fff;
        border-radius: 10px;
        animation: menu .3s ease-in;
        box-shadow: 0 0 30px rgba(0, 0, 0, .1);
    }

    .menu-item {
        margin-top: 7px;
        width: 160px;
        height: 36px;
        margin-left: 20px;
        -webkit-user-select: none;
        -moz-user-select: none;
        -ms-user-select: none;
        user-select: none;

        .btn {
            display: block;
            color: #18191b;
            line-height: 36px;
            text-align: left;
            border-radius: 6px;
            padding-left: 6px;

            &:hover {
                background-color: #c9ccd0;
            }
        }

    }
}

@keyframes menu {
    0% {
        opacity: 0;
        transform: translateY(-10px);
    }

    100% {
        opacity: 1;
    }
}
</style>