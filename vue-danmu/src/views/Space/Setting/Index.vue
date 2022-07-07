<template>
    <div>
        <div class="setting-title">
            <span @click="setRouter(true)" :class="isInfo ? 'active-text' : ''">基本信息</span>
            <span @click="setRouter(false)" :class="!isInfo ? 'active-text' : ''">账号安全</span>
            <div class="selected" v-bind:class="activeClass"></div>
        </div>
        <router-view class="router" />
    </div>
</template>

<script>
import { onMounted, computed, ref } from 'vue';
import { useRouter, useRoute } from 'vue-router';

export default {
    setup() {
        const isInfo = ref(true);
        const route = useRoute();
        const router = useRouter();

        const activeClass = computed(() => {
            return isInfo.value ? "" : "security-active";
        })

        const setRouter = (isInfoPage) => {
            if (isInfo.value !== isInfoPage) {
                //设置当前选中并修改页面
                isInfo.value = isInfoPage;
                router.push({ name: isInfoPage ? "InfoSetting" : "SecuritySetting" });
            }
        }

        onMounted(() => {
            if (route.name === "SecuritySetting") {
                isInfo.value = false;
            }
        })

        return {
            isInfo,
            setRouter,
            activeClass
        }
    },
};
</script>

<style lang="less" scoped>
.setting-title {
    position: relative;
    height: 100px;
    text-align: center;

    span {
        line-height: 100px;
        font-size: 20px;
        user-select: none;

        &:nth-child(2) {
            margin-left: 30px;
        }

    }
}

.selected {
    position: absolute;
    top: 70px;
    left: 344px;
    width: 72px;
    height: 4px;
    border-radius: 2px;
    background-color: #36ad6a;
    transition: left 0.3s;
    -moz-transition: left 0.3s;
    -webkit-transition: left 0.3s;
    -o-transition: left 0.3s;
}

/**激活时 */
.security-active {
    left: 454px;
}

.active-text {
    color: #36ad6a;
}

.router {
    width: 80%;
    margin: 0 auto;
}
</style>