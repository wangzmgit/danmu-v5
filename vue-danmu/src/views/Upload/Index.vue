<template>
    <div class="upload" v-title data-title="创作中心">
        <header-bar></header-bar>
        <div class="upload-box">
            <n-menu class="upload-menu" :options="menuOptions" :default-value="defaultOption" />
            <div class="upload-router">
                <router-view></router-view>
            </div>
        </div>
    </div>
</template>

<script>
import { h, ref, onBeforeMount } from "vue";
import { NIcon, NMenu } from "naive-ui";
import { RouterLink, useRoute } from 'vue-router';
import HeaderBar from '@/components/HeaderBar.vue';
import { BookOutline, VideocamOutline, ListOutline } from "@vicons/ionicons5";


export default {
    setup() {
        const route = useRoute();
        const defaultOption = ref('');//默认激活菜单
        const menuOptions = [
            {
                label: () =>
                    h(
                        RouterLink,
                        {
                            to: {
                                name: "UploadVideoHome",
                            }
                        },
                        { default: () => '发布视频' }
                    ),
                key: "upload",
                icon: renderIcon(VideocamOutline),
            },
            {
                label: () =>
                    h(
                        RouterLink,
                        {
                            to: {
                                name: "VideoManage",
                            }
                        },
                        { default: () => '稿件管理' }
                    ),
                key: "content",
                icon: renderIcon(BookOutline),
            },
            {
                label: () =>
                    h(
                        RouterLink,
                        {
                            to: {
                                name: "CommentManage",
                            }
                        },
                        { default: () => '评论管理' }
                    ),
                key: "comment",
                icon: renderIcon(ListOutline),
            },
        ];

        function renderIcon(icon) {
            return () => h(NIcon, null, { default: () => h(icon) });
        }

        onBeforeMount(() => {
            switch (route.name) {
                case 'UploadVideoHome':
                    defaultOption.value = 'upload';
                    break;
                case 'VideoManage':
                    defaultOption.value = 'content';
                    break;
                case 'CommentManage':
                    defaultOption.value = 'comment';
                    break;
                default:
                    defaultOption.value = 'upload';
                    break;
            }
        })

        return {
            menuOptions,
            defaultOption,
            renderIcon
        }
    },
    components: {
        NMenu,
        HeaderBar,
    }
}
</script>

<style lang="less" scoped>
.upload {
    top: 0;
    bottom: 0;
    width: 100%;
    height: 100%;
    position: fixed;
    overflow-y: scroll;
    background: #f7f7f7;

    .upload-box {
        display: flex;
        margin: 20px 120px;
    }

    .upload-menu {
        width: 220px;
        height: 500px;
        min-width: 200px;
        background-color: #fff;
    }

    .upload-router {
        width: 100%;
        min-width: 700px;
        min-height: 500px;
        margin-left: 20px;
        background-color: #fff;
    }
}
</style>