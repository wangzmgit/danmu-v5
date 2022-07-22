<template>
    <n-card class="card">
        <div class="user-info">
            <div class="avatar-box">
                <common-avatar :url="userInfo.avatar"></common-avatar>
                <span class="greet">{{ greet }}</span>
            </div>
            <span class="role">{{ toRoleText(userInfo.role) }}</span>
        </div>
    </n-card>
    <n-card class="card">
        <data-chart v-if="!loadingChart" :data="recentData"></data-chart>
    </n-card>
    <n-card class="card">
        <div class="total-data">
            <div v-for="(item, index) in todayData" :key="index" class="data-card">
                <div class="data-icon">
                    <component :is="item.icon"></component>
                </div>
                <div class="data-content">
                    <p>{{ item.title }}</p>
                    <n-number-animation :to="item.data"></n-number-animation>
                </div>
            </div>
        </div>
    </n-card>
</template>

<script lang="ts">
import storage from '@/utils/stored-data';
import { onBeforeMount, reactive, ref, computed, defineComponent } from 'vue';
import DataChart from './components/DataChart.vue';
import { getTotalDataAPI, getRecentDataAPI } from '@/api/dashboard';
import { NIcon, NCard, NNumberAnimation } from 'naive-ui';
import { PersonOutline, VideocamOutline, CreateOutline, MailOutline } from '@vicons/ionicons5';
import CommonAvatar from '@/components/CommonAvatar.vue';
export default defineComponent({
    setup() {
        //今日数据
        const todayData = reactive([
            {
                icon: "PersonOutline",
                title: "用户数",
                data: 0,
            },
            {
                icon: "VideocamOutline",
                title: "视频数",
                data: 0,
            },
            {
                icon: "CreateOutline",
                title: "待审核视频",
                data: 0,
            },
            {
                icon: "MailOutline",
                title: "消息",
                data: 0,
            },
        ])
        const totalData = ref([]);
        const recentData = ref([]);
        const loadingChart = ref(true);
        const getTotalData = () => {
            getTotalDataAPI().then((res) => {
                if (res.data.code === 2000) {
                    const resData = res.data.data.data;
                    todayData[0].data = resData.user;
                    todayData[1].data = resData.video;
                    todayData[2].data = resData.review;
                }
            })
        }

        const getRecentData = () => {
            getRecentDataAPI().then((res) => {
                if (res.data.code === 2000) {
                    recentData.value = res.data.data.data.reverse();
                    loadingChart.value = false;
                }
            })
        }

        //用户信息
        const userInfo = computed(() => {
            if (storage.get("adminInfo")) {
                return storage.get("adminInfo");
            } else {
                return {
                    uid: 0,
                    avatar: ''
                }
            }
        });

        //打招呼
        const greet = computed(() => {
            const time = new Date();
            const hour = time.getHours();
            const name = userInfo.value.name;
            if (hour >= 0 && hour <= 6) {
                return `早点休息${name}( - . - )`;
            } else if (hour > 6 && hour <= 9) {
                return `早上好！${name}`;
            } else if (hour > 9 && hour <= 11) {
                return `上午好！${name} 今天也要有所收获呀~`;
            } else if (hour > 11 && hour <= 13) {
                return `中午好！${name}`;
            } else if (hour > 13 && hour <= 19) {
                return `下午好！${name}`;
            } else if (hour > 19 && hour <= 23) {
                return `晚上好！${name}`;
            }
        })

        const toRoleText = (role: number) => {
            switch (role) {
                case 1:
                    return '审核';
                case 2:
                    return '管理员';
                case 3:
                    return '超级管理员';
            }
        }

        onBeforeMount(() => {
            getTotalData();
            getRecentData();
        })

        return {
            greet,
            totalData,
            recentData,
            loadingChart,
            todayData,
            userInfo,
            toRoleText
        }
    },
    components: {
        NIcon,
        NCard,
        DataChart,
        NNumberAnimation,
        PersonOutline,
        VideocamOutline,
        CreateOutline,
        MailOutline,
        CommonAvatar
    }
});
</script>

<style lang="less" scoped>
.card {
    margin: 10px;
    width: calc(100% - 20px);
    border-radius: 10px;
}

.user-info {
    display: flex;
    align-items: center;
    justify-content: space-between;

    .avatar-box {
        display: flex;
        align-items: center;

        .greet {
            color: #333;
            font-size: 18px;
            margin-left: 20px;
        }
    }

    .role {
        color: #666;
    }
}

.total-data {
    display: flex;
    margin-top: 20px;
    justify-content: space-around;

    .data-card {
        width: 20%;
        display: flex;
        align-items: center;
        border-radius: 10px;

        &:nth-child(1) {
            background: #40c9c6;
        }

        &:nth-child(2) {
            background: #36a3f7;
        }

        &:nth-child(3) {
            background: #f4516c;
        }

        &:nth-child(4) {
            background: #34bfa3;
        }

        &:nth-child(5) {
            background: #7daebd;
        }

        //卡片图标
        .data-icon {
            width: 25%;
            color: #fff;
            font-size: 50px;
            margin-left: 5%;
        }

        //卡片文字内容
        .data-content {
            width: 70%;
            height: 60px;
            color: #fff;
            font-size: 16px;
            margin-left: 20px;

            p {
                margin: 0 0 10px 0;
            }
        }
    }
}
</style>