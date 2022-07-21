<template>
    <div class="space" v-title :data-title="`${userInfo.name}的个人中心`">
        <header-bar></header-bar>
        <div class="space-box">
            <n-menu class="space-menu" :options="menuOptions" :default-value="defaultOption" />
            <div class="space-router">
                <router-view></router-view>
            </div>
        </div>
    </div>
</template>


<script lang="ts">
import { h, computed, ref, onBeforeMount, defineComponent } from "vue";
import { NIcon, NMenu } from "naive-ui";
import { RouterLink, useRoute } from 'vue-router';
import storage from "@/utils/stored-data";
import HeaderBar from '@/components/HeaderBar.vue';
import {
    ChatbubbleOutline, VideocamOutline, SettingsOutline,
    BookmarkOutline, CloudUploadOutline, MailOutline
} from "@vicons/ionicons5";
import { userInfoType } from "@/types/user";
export default defineComponent({
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
                                name: "SpaceInfo",
                            }
                        },
                        { default: () => '首页' }
                    ),
                key: "home",
                icon: renderIcon(VideocamOutline, '#609a8b'),
            },
            {
                label: () =>
                    h(
                        RouterLink,
                        {
                            to: {
                                name: "Collection",
                            }
                        },
                        { default: () => '收藏夹' }
                    ),
                key: "collection",
                icon: renderIcon(BookmarkOutline, '#e3c0aa'),
            },
            {
                label: () =>
                    h(
                        RouterLink,
                        {
                            to: {
                                name: "Upload",
                            }
                        },
                        { default: () => '投稿' }
                    ),
                key: "upload",
                icon: renderIcon(CloudUploadOutline, '#7daebd'),
            },
            {
                label: () =>
                    h(
                        RouterLink,
                        {
                            to: {
                                name: "Announce",
                            }
                        },
                        { default: () => '公告' }
                    ),
                key: "announce",
                icon: renderIcon(MailOutline, '#8484c3'),
            },
            {
                label: () =>
                    h(
                        RouterLink,
                        {
                            to: {
                                name: "Message",
                            }
                        },
                        { default: () => '消息' }
                    ),
                key: "message",
                icon: renderIcon(ChatbubbleOutline, '#c79fa7'),
            },
            {
                label: () =>
                    h(
                        RouterLink,
                        {
                            to: {
                                name: "Setting",
                            }
                        },
                        { default: () => '设置' }
                    ),
                key: "setting",
                icon: renderIcon(SettingsOutline, '#808080'),
            },
        ];

        function renderIcon(icon: any, color: string) {
            return () => h(NIcon, { color: color }, { default: () => h(icon) });
        }

        const userInfo = computed(() => {
            return storage.get("userInfo");
        })

        onBeforeMount(() => {
            switch (route.name) {
                case 'SpaceInfo':
                    defaultOption.value = 'home';
                    break;
                case 'Collection':
                    defaultOption.value = 'collection';
                    break;
                case 'Announce':
                    defaultOption.value = 'announce';
                    break;
                case 'Message':
                    defaultOption.value = 'message';
                    break;
                case 'InfoSetting': case 'SecuritySetting':
                    defaultOption.value = 'setting';
                    break;
                default:
                    defaultOption.value = 'home';
                    break;
            }
        })

        return {
            userInfo,
            menuOptions,
            defaultOption,
            renderIcon
        }
    },
    components: {
        NMenu,
        HeaderBar,
    },
});
</script>

<style lang="less" scoped>
.space {
    height: 100%;
    width: 100%;
    top: 0;
    bottom: 0;
    position: fixed;
    overflow-y: scroll;
    background: #e9e9e9;
}

.space-box {
    width: 1100px;
    background-color: #fff;
    display: flex;
    margin: 10px auto 30px;
    min-height: 630px;

    .space-menu {
        width: 200px;
        min-height: 100%;
        border: 1px solid #efeff5;
    }

    .space-router {
        width: calc(100% - 200px);
        height: 100%;
    }
}
</style>