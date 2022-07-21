<template>
    <div class="security-info">
        <div>
            <p>邮箱:&nbsp;&nbsp;{{ userInfo.email }}</p>
            <n-button @click="modify">修改邮箱</n-button>
        </div>
        <div>
            <p>密码:&nbsp;&nbsp;********</p>
            <n-button @click="modify">修改密码</n-button>
        </div>
    </div>
</template>

<script lang="ts">
import { defineComponent, onBeforeMount, ref } from 'vue';
import { useNotification, NButton } from 'naive-ui';
import storage from "@/utils/stored-data";
import { userInfoType } from '@/types/user';


export default defineComponent({
    setup() {
        const userInfo = ref<userInfoType>({
            uid: 0,
            name: "",
            avatar: ""
        });
        const notification = useNotification();//通知

        const modify = () => {
            notification.info({
                title: '当前版本暂不支持修改',
                duration: 3000,
            });
        }

        onBeforeMount(() => {
            userInfo.value = storage.get('userInfo') as userInfoType;
        })

        return {
            modify,
            userInfo
        }
    },
    components: {
        NButton
    }
});
</script>

<style lang="less" scoped>
.security-info {
    div {
        height: 60px;
        width: 300px;
        display: flex;
        margin-left: 30px;
        justify-content: space-between;
        align-items: center;

        p {
            margin: 0;
        }

    }
}
</style>