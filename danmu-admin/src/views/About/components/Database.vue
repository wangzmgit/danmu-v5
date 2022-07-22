<template>
    <div class="database">
        <n-form :model="dbForm" label-width="auto">
            <n-form-item label="数据库名">
                <n-input placeholder="数据库名" v-model:value="dbForm.dbName" />
            </n-form-item>
            <n-form-item label="MySQL地址">
                <n-input placeholder="MySQL地址" v-model:value="dbForm.dbHost" />
            </n-form-item>
            <n-form-item label="MySQL端口">
                <n-input-number placeholder="MySQL端口" v-model:value="dbForm.dbPort" />
            </n-form-item>
            <n-form-item label="MySQL用户名">
                <n-input placeholder="MySQL用户名" v-model:value="dbForm.dbUser" />
            </n-form-item>
            <n-form-item label="MySQL密码">
                <n-input placeholder="MySQL密码" type="password" v-model:value="dbForm.dbPwd" />
            </n-form-item>
            <n-form-item label="Redis地址">
                <n-input placeholder="Redis地址" v-model:value="dbForm.cacheHost" />
            </n-form-item>
            <n-form-item label="Redis端口">
                <n-input-number placeholder="Redis端口" v-model:value="dbForm.cachePort" />
            </n-form-item>
            <n-form-item label="Redis密码">
                <n-input placeholder="Redis密码" type="password" v-model:value="dbForm.cachePwd" />
            </n-form-item>
            <div class="submit">
                <span>如不修改密码请留空</span>
                <n-button type="primary" @click="setDatabase">保存</n-button>
            </div>
        </n-form>
    </div>
</template>

<script lang="ts">
import { defineComponent, onBeforeMount, reactive } from "vue";
import { getDatabaseAPI, setDatabaseAPI } from '@/api/config';
import { NInput, NInputNumber, NForm, NFormItem, NButton, useNotification } from "naive-ui";

export default defineComponent({
    setup() {
        const notification = useNotification();

        const getDatabase = () => {
            getDatabaseAPI().then((res) => {
                if (res.data.code === 2000) {
                    const resData = res.data.data;
                    dbForm.dbName = resData.db_name;
                    dbForm.dbHost = resData.db_host;
                    dbForm.dbPort = resData.db_port;
                    dbForm.dbUser = resData.db_user;
                    dbForm.cacheHost = resData.cache_host;
                    dbForm.cachePort = resData.cache_port;
                }
            })
        }

        const dbForm = reactive({
            dbName: '',
            dbHost: '',
            dbPort: 3306,
            dbUser: '',
            dbPwd: '',
            cacheHost: '',
            cachePort: 6379,
            cachePwd: '',
        });

        const setDatabase = () => {
            setDatabaseAPI(dbForm).then((res) => {
                if (res.data.code === 2000) {
                    notification.success({
                        title: '修改成功',
                        duration: 3000
                    })
                }
            })
        }

        onBeforeMount(() => {
            getDatabase();
        })

        return {
            dbForm,
            setDatabase
        }
    },
    components: {
        NInput,
        NForm,
        NButton,
        NFormItem,
        NInputNumber
    }
});
</script>

<style lang="less" scoped>
.submit {
    display: flex;
    align-items: center;
    justify-content: space-between;

    span {
        color: #666;
    }
}
</style>