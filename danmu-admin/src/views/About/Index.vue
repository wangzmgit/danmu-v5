<template>
    <n-card class="card">
        <n-tabs type="line" default-value="info" animated>
            <n-tab-pane name="info" tab="状态">
                <info></info>
            </n-tab-pane>
            <n-tab-pane name="database" tab="数据库">
                <database></database>
            </n-tab-pane>
            <n-tab-pane name="email" tab="邮箱">
                <email></email>
            </n-tab-pane>
            <n-tab-pane name="storage" tab="存储">
                <storage></storage>
            </n-tab-pane>
            <n-tab-pane name="other" tab="其他">
                <other></other>
            </n-tab-pane>
        </n-tabs>
    </n-card>
</template>

<script lang="ts">
import Info from './components/Info.vue';
import Email from './components/Email.vue';
import Storage from './components/Storage.vue';
import Database from './components/Database.vue';
import Other from './components/Other.vue';
import { NCard, NTabs, NTabPane } from "naive-ui";
import { defineComponent } from 'vue';

export default defineComponent({
    components: {
        Info,
        Email,
        Other,
        Storage,
        Database,
        NCard,
        NTabs,
        NTabPane
    }
});
</script>

<style lang="less" scoped>
.card {
    margin: 10px;
    width: calc(100% - 20px);
    border-radius: 10px;
}
</style>