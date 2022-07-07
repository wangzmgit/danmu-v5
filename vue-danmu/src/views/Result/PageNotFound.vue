<template>
    <div v-title data-title="404">
        <n-result title="404 资源不存在">
            <template #icon>
                <div style="width: 300px">
                    <img class="image" src="@/assets/not-found-page.png" alt="404">
                </div>
            </template>
            <template #footer>
                <n-button @click="goHome">返回首页</n-button>
            </template>
        </n-result>
    </div>
</template>

<script>
import { useRouter } from 'vue-router';
import { NResult, NButton } from 'naive-ui';
export default {
    setup() {
        const router = useRouter();
        const goHome = () => {
            router.push({ name: "Home" });
        }

        return {
            goHome
        }
    },
    components: {
        NResult,
        NButton
    }
}
</script>

<style lang="less" scoped>
.image {
    width: 100%;
}
</style>