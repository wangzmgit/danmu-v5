<template>
    <div class="user-card">
        <div class="card-avatar">
            <n-avatar v-if="userInfo.avatar" round :size="100" :src="userInfo.avatar" />
            <n-avatar v-else round :size="100">
                <n-icon size="60">
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
        <!--关注粉丝信息-->
        <div class="card-info">
            <div>
                <p class="info-title">投稿</p>
                <p>{{ count }}</p>
            </div>
            <div>
                <p class="info-title">关注</p>
                <p class="info-content" @click="goPage('Following')">{{ following }}</p>
            </div>
            <div>
                <p class="info-title">粉丝</p>
                <p class="info-content" @click="goPage('Followers')">{{ followers }}</p>
            </div>
        </div>
    </div>
    <div class="video-list">
        <user-video-list @setCount="setCount"></user-video-list>
    </div>
</template>

<script>
import { ref, computed, onMounted } from "vue";
import storage from "@/utils/stored-data.js";
import { NIcon, NAvatar } from 'naive-ui';
import { useRouter } from "vue-router";
import { getFollowDataAPI } from '@/api/follow';
import { Person, Female, Male } from '@vicons/ionicons5';
import UserVideoList from './components/UserVideoList'

export default {
    setup() {
        const count = ref(0);
        const followers = ref(0);
        const following = ref(0);
        const router = useRouter();

        const userInfo = computed(() => {
            return storage.get("userInfo");
        })

        //前往关注和粉丝页面
        const goPage = (name) => {
            router.push({ name: name });
        }

        //设置视频数量
        const setCount = (value) => {
            count.value = value;
        }

        //获取关注数和粉丝数
        const getFollowData = (id) => {
            getFollowDataAPI(id).then((res) => {
                if (res.data.code === 2000) {
                    followers.value = res.data.data.followers
                    following.value = res.data.data.following;
                }
            }).catch((err) => {
                notification.error({
                    title: '获取失败',
                    content: "原因:" + err.response.data.msg,
                    duration: 5000,
                })
            });
        }

        onMounted(() => {
            getFollowData(userInfo.value.uid);
        })


        return {
            count,
            followers,
            following,
            userInfo,
            goPage,
            setCount
        }
    },
    components: {
        NIcon,
        NAvatar,
        Person,
        Male,
        Female,
        UserVideoList
    },
};
</script>
<style lang="less" scoped>
.user-card {
    display: flex;

    .card-avatar {
        width: 160px;
        height: 160px;

        span {
            margin: 30px 30px;
        }
    }

    .card-name {
        width: 500px;
        margin: 0 10px;

        .name {
            margin-top: 45px;
            font-size: 20px;
            font-weight: 500;
        }

        .sign {
            font-size: 14px;
            color: #2e2e2e;
        }
    }

    .card-info {
        width: 210px;
        display: flex;
        margin-right: 40px;

        div {
            width: 70px;
            text-align: center;

            .info-title {
                margin-top: 56px;
                margin-bottom: 6px;
                font-weight: bold;
            }

            .info-content:hover {
                cursor: pointer;
                color: #36ad6a;
            }
        }
    }
}

.video-list {
    margin-left: 20px;
}
</style>