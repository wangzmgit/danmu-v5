<template>
    <div id="followBox" class="follow-list" :style="`height:${height}px`">
        <div class="follow-card" v-for="(item, index) in followList" :key="index">
            <!--头像-->
            <div class="follow-avatar">
                <n-avatar v-if="item.avatar" round :size="60" :src="item.avatar" />
                <n-avatar v-else round :size="60">
                    <n-icon size="36">
                        <Person />
                    </n-icon>
                </n-avatar>
            </div>
            <!--昵称和个签-->
            <span class="follow-name" @click="goUserSpace(item.uid)">{{ item.name }}</span>
            <span class="follow-sign">{{ item.sign }}</span>
        </div>
    </div>
</template>

<script>
import { useRouter } from "vue-router";
import { onBeforeMount, ref, reactive, onBeforeUnmount } from "vue";
import { Person } from '@vicons/ionicons5';
import { getFollowingAPI, getFollowersAPI } from '@/api/follow';
import { NAvatar, NIcon, NScrollbar, useNotification } from "naive-ui";

export default {
    props: {
        uid: {
            type: Number,
            default: 0,
        },
        following: {
            type: Boolean,
            default: false,
        },
        height: {
            type: Number,
            default: 600
        }
    },
    setup(props) {
        const pageInfo = reactive({
            current: 1,
            pageSize: 9
        });
        const noMore = ref(false);//没有更多
        const loading = ref(false);//加载中
        const followList = ref([]);
        const router = useRouter();
        const notification = useNotification();//通知

        //进入用户空间
        const goUserSpace = (uid) => {
            let userUrl = router.resolve({ name: "User", params: { uid: uid } });
            window.open(userUrl.href, '_blank');
        }

        //获取关注
        const getFollowingList = () => {
            getFollowingAPI(props.uid, pageInfo.current, pageInfo.pageSize).then((res) => {
                if (res.data.code === 2000) {
                    const resUser = res.data.data.users;
                    if (resUser) {
                        followList.value.push(...resUser);
                    }
                    if (!resUser || resUser.length < pageInfo.pageSize) {
                        noMore.value = true;
                    }
                    loading.value = false;
                }
            }).catch((err) => {
                notification.error({
                    title: '获取失败',
                    content: "原因:" + err.response.data.msg,
                    duration: 5000,
                })
            });
        }

        //获取粉丝
        const getFollowersList = () => {
            getFollowersAPI(props.uid, pageInfo.current, pageInfo.pageSize).then((res) => {
                if (res.data.code === 2000) {
                    const resUser = res.data.data.users;
                    if (resUser) {
                        followList.value.push(...resUser);
                    }
                    if (!resUser || resUser.length < pageInfo.pageSize) {
                        noMore.value = true;
                    }
                    loading.value = false;
                }
            }).catch((err) => {
                notification.error({
                    title: '获取失败',
                    content: "原因:" + err.response.data.msg,
                    duration: 5000,
                })
            });
        }

        const lazyLoading = (e) => {
            if (e.target.id === "followBox") {
                const scrollTop = document.documentElement.scrollTop || document.body.scrollTop;
                const clientHeight = document.documentElement.clientHeight;
                const scrollHeight = document.documentElement.scrollHeight;
                if (scrollTop + clientHeight >= scrollHeight) {
                    if (!noMore.value && !loading.value) {
                        pageInfo.current++;
                        loading.value = true;
                        if (props.following) {
                            getFollowingList();
                        } else {
                            getFollowersList();
                        }
                    }
                }
            }
        }

        onBeforeMount(() => {
            if (props.following) {
                getFollowingList();
            } else {
                getFollowersList();
            }
            window.addEventListener('scroll', lazyLoading, true);
        })

        onBeforeUnmount(() => {
            window.removeEventListener('scroll', lazyLoading);
        })

        return {
            noMore,
            loading,
            followList,
            lazyLoading,
            goUserSpace
        }
    },
    components: {
        NIcon,
        NAvatar,
        NScrollbar,
        Person
    },
};
</script>

<style lang="less" scoped>
.follow-list {
    overflow-y: scroll;

    /**修改滚动条样式 */
    &::-webkit-scrollbar {
        width: 6px;
    }

    &::-webkit-scrollbar-thumb {
        /*滚动条里面小方块*/
        border-radius: 3px;
        background: #a3a3a3;
    }

    &::-webkit-scrollbar-track {
        /*滚动条里面轨道*/
        border-radius: 5px;
    }
}

.follow-card {
    height: 70px;
    position: relative;
    margin: 0 10px;
    border-bottom: 1px solid #d1d1d1;

    /**移除最后一个的底部边框 */
    &:last-child {
        border-bottom: none;
    }
}

.follow-avatar>span {
    margin-top: 5px;
    margin-left: 20px;
}

.follow-name {
    color: #333;
    position: absolute;
    top: 10px;
    left: 100px;
    font-weight: 600;
    cursor: pointer;
}

.follow-sign {
    position: absolute;
    top: 38px;
    left: 100px;
    font-size: 12px;
    color: #666;
}
</style>