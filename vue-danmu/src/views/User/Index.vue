<template>
    <div class="space" v-title data-title="个人空间">
        <header-bar></header-bar>
        <div class="space-box">
            <div class="user-card">
                <div class="card-avatar">
                    <n-avatar v-if="userInfo.avatar" round :size="70" :src="userInfo.avatar" />
                    <n-avatar v-else round :size="70">
                        <n-icon size="50">
                            <Person />
                        </n-icon>
                    </n-avatar>
                </div>
                <div class="card-name">
                    <p class="name">{{ userInfo.name }}
                        <n-icon v-if="userInfo.gender === 1" color="#1890ff">
                            <Male />
                        </n-icon>
                        <n-icon v-else-if="userInfo.gender === 2" color="#eb2f96">
                            <Female />
                        </n-icon>
                    </p>
                    <p class="sign">{{ userInfo.sign }}</p>
                </div>
                <div v-if="login" class="card-btn">
                    <n-button size="small" @click="goMsg">私信</n-button>
                    <n-button size="small" v-if="isFollow" type="primary" @click="followClick">已关注</n-button>
                    <n-button size="small" v-else type="error" @click="followClick">关注</n-button>
                </div>
            </div>
            <n-menu class="menu" mode="horizontal" :options="menuOptions" :default-value="defaultOption" />
            <div class="user-content">
                <router-view :key="routerKey"></router-view>
            </div>
        </div>
    </div>
</template>

<script>
import { storeToRefs } from 'pinia';
import { h, ref, watch, onBeforeMount } from "vue";
import { getUserInfoByIDAPI } from '@/api/user';
import useUserFollow from '@/hooks/user-follow';
import { RouterLink, useRoute, useRouter } from 'vue-router';
import { useLoginStore } from '@/store/login-store';
import HeaderBar from '@/components/HeaderBar.vue';
import { Male, Female, Person } from "@vicons/ionicons5";
import { NIcon, NMenu, NButton, NAvatar, useNotification } from "naive-ui";

export default {
    setup() {
        const userInfo = ref({
            avatar: "",
            name: "",
            sing: "",
            gender: 0,
        });
        const menuOptions = [
            {
                label: () =>
                    h(
                        RouterLink,
                        {
                            to: {
                                name: "UserVideo",
                            }
                        },
                        { default: () => `${getGender()}的作品` }
                    ),
                key: "video",
            },
            {
                label: () =>
                    h(
                        RouterLink,
                        {
                            to: {
                                name: "UserFollowing",
                            }
                        },
                        { default: () => '关注' }
                    ),
                key: "following",
            },

            {
                label: () =>
                    h(
                        RouterLink,
                        {
                            to: {
                                name: "UserFollowers",
                            }
                        },
                        { default: () => '粉丝' }
                    ),
                key: "followers",
            },
        ];
        const defaultOption = ref('');//默认激活菜单
        const route = useRoute();
        const routerKey = ref(new Date());
        const loginStore = useLoginStore();
        const { login } = storeToRefs(loginStore);
        const notification = useNotification();//通知

        //获取用户信息
        const getUserInfoByID = (uid) => {
            getUserInfoByIDAPI(uid).then((res) => {
                if (res.data.code === 2000) {
                    userInfo.value = res.data.data.user;
                    document.title = `${userInfo.value.name}的空间`;
                }
            }).catch((err) => {
                notification.error({
                    title: '修改失败',
                    content: "原因:" + err.response.data.msg,
                    duration: 5000,
                })
            });
        }

        //获取性别
        const getGender = () => {
            switch (userInfo.value.gender) {
                case 1:
                    return '他';
                case 2:
                    return '她';
                default:
                    return 'TA';
            }
        }

        //监听路由改变
        watch(() => route.name, (_newValue, _oldValue) => {
            routerKey.value = new Date();
        })

        //跳转私信页面
        const router = useRouter();
        const goMsg = () => {
            router.push({ name: "Message", params: { fid: userInfo.value.uid } });
        }

        //关注相关
        const { isFollow, getFollowStatus, follow, unfollow } = useUserFollow();
        const followClick = () => {
            const uid = Number(route.params.uid)
            if (isFollow.value) {
                unfollow(uid);
            } else {
                follow(uid);
            }
            isFollow.value = !isFollow.value;
        }

        onBeforeMount(() => {
            getUserInfoByID(route.params.uid);
            switch (route.name) {
                case 'UserFollowing':
                    defaultOption.value = 'following';
                    break;
                case 'UserFollowers':
                    defaultOption.value = 'followers';
                    break;
                default:
                    defaultOption.value = 'video';
                    break;
            }
            getFollowStatus(route.params.uid);
        })

        return {
            login,
            isFollow,
            userInfo,
            routerKey,
            menuOptions,
            defaultOption,
            goMsg,
            followClick,
        }
    },
    components: {
        NIcon,
        NMenu,
        NButton,
        NAvatar,
        HeaderBar,
        Male,
        Female,
        Person
    }
}
</script>

<style lang="less" scoped>
.space {
    height: 100%;
    width: 100%;
    top: 0;
    bottom: 0;
    position: fixed;
    overflow-y: scroll;
    background: #e9e9e9;

    .space-box {
        width: 1100px;
        background-color: #fff;
        margin: 10px auto 30px;
        height: 600px;
    }
}

.user-card {
    display: flex;

    .card-avatar {
        width: 120px;
        height: 120px;

        span {
            margin: 30px 30px;
        }
    }

    .card-name {
        width: 760px;
        margin: 0 10px;

        .name {
            color: #333;
            margin-top: 35px;
            margin-bottom: 10px;
            font-size: 20px;
            font-weight: 500;

            i {
                font-size: 22px;
            }
        }

        .sign {
            color: #666;
            margin: 0;
            font-size: 14px;
        }
    }
}

.menu {
    border-bottom: 1px solid #e9e9e9;
    width: 100%;
}

.card-btn {
    width: 200px;
    text-align: center;

    button {
        margin-top: 70px;
        margin-right: 20px;
    }
}
</style>
