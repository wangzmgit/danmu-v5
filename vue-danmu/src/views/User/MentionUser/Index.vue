<template>
    <n-result title="用户不存在">
        <template #icon>
            <div style="width: 300px">
                <img class="image" src="@/assets/not-found-user.png" alt="404">
            </div>
        </template>
    </n-result>
</template>

<script lang="ts">
import { NResult } from 'naive-ui';
import { defineComponent, onBeforeMount } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { getUidByNameAPI } from '@/api/user';
export default defineComponent({
    setup() {
        const route = useRoute();
        const router = useRouter();

        onBeforeMount(() => {
            if (route.params.name) {
                getUidByNameAPI(route.params.name.toString()).then((res) => {
                    if (res.data.code === 2000) {
                        const uid = res.data.data.uid;
                        if (uid) router.push({ name: 'User', params: { uid: uid} });        
                    }
                })
            }
        })
    },
    components: {
        NResult
    }
});
</script>

<style lang="less" scoped>
.image {
    width: 100%;
}
</style>