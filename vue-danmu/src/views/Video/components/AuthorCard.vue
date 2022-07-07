<template>
    <div v-if="loading" class="video-author">
        <div class="video-author-box">
            <!--头像-->
            <div class="author-avatar">
                <n-avatar round :size="50">
                    <n-icon size="26">
                        <Person />
                    </n-icon>
                </n-avatar>
            </div>
            <div class="author-info">
                <n-skeleton height="20px" width="70px" />
                <n-skeleton height="16px" width="80%" />
            </div>
        </div>
    </div>
    <div v-else class="video-author">
        <div class="video-author-box">
            <!--头像-->
            <div class="author-avatar">
                <n-avatar v-if="author.avatar" round :size="50" :src="author.avatar" />
                <n-avatar v-else round :size="50">
                    <n-icon size="26">
                        <Person />
                    </n-icon>
                </n-avatar>
            </div>
            <!--昵称和个签-->
            <div class="author-info">
                <p @click="goUserSpace(author.uid)">{{ author.name }}</p>
                <p>{{ author.sign }}</p>
            </div>
            <div class="follow-btn">
                <n-button size="small" v-if="isFollow" type="primary" @click="unfollow(author.uid)">已关注</n-button>
                <n-button size="small" v-else :disabled="!login" type="error" @click="follow(author.uid)">关注</n-button>
            </div>
        </div>
    </div>
</template>

<script>
import { watch } from 'vue';
import { storeToRefs } from 'pinia';
import { Person } from '@vicons/ionicons5';
import { useRouter } from 'vue-router';
import useUserFollow from '@/hooks/user-follow';
import { useLoginStore } from '@/store/login-store';
import { NIcon, NButton, NAvatar, NSkeleton } from 'naive-ui';

export default {
    props: {
        author: {
            type: Object
        },
        loading: {
            type: Boolean
        }
    },
    setup(props) {
        const router = useRouter();
        const loginStore = useLoginStore();
        const { login } = storeToRefs(loginStore);

        const goUserSpace = (uid) => {
            let userUrl = router.resolve({ name: "User", params: { uid: uid } });
            window.open(userUrl.href, '_blank');
        }

        const { isFollow, getFollowStatus, follow, unfollow } = useUserFollow();
        watch(() => props.loading, (newValue, _oldValue) => {
            if (!newValue) {
                getFollowStatus(props.author.uid)
            }
        })

        return {
            login,
            isFollow,
            goUserSpace,
            follow,
            unfollow
        }
    },
    components: {
        NIcon,
        NButton,
        NAvatar,
        NSkeleton,
        Person
    }
}
</script>

<style lang="less" scoped>
// .author-loading {

// }

.video-author {
    width: 100%;
    height: 90px;
    border-radius: 6px;
    background-color: #f1f2f3;

    .video-author-box {
        display: flex;
        height: 100%;
        align-items: center;

        .author-avatar {
            width: 50px;
            height: 50px;
            padding-left: 10px;
        }

        .author-info {
            height: 50px;
            display: flex;
            flex-direction: column;
            justify-content: space-around;
            width: calc(100% - 160px);
            padding-left: 16px;

            p {
                width: 100%;
                margin: 0;
                overflow: hidden;
                text-overflow: ellipsis;
                display: -webkit-box;
                -webkit-line-clamp: 1;
                -webkit-box-orient: vertical;

                &:nth-child(1) {
                    color: #222;
                    font-size: 14px;
                    cursor: pointer;
                    font-weight: 500;
                }

                &:nth-child(2) {
                    color: #999;
                    font-size: 12px;
                }
            }
        }

        .follow-btn {
            width: 80px;
            text-align: center;
        }
    }
}
</style>