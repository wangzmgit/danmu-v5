<template>
    <n-input v-model:value="keyword" round placeholder="搜索" @keydown.enter="search">
        <template #suffix>
            <n-icon @click="search">
                <Search />
            </n-icon>
        </template>
    </n-input>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue';
import { NInput, NIcon } from 'naive-ui';
import { Search } from '@vicons/ionicons5';

export default defineComponent({
    emits: ['search'],
    props: {
        keyword: {
            type: String
        }
    },
    setup(props,ctx) {
        //搜索功能
        const keyword = ref(props.keyword);
        const search = () => {
            ctx.emit('search',keyword.value);
        }

        return {
            keyword,
            search,
        }
    },
    components: {
        NIcon,
        NInput,
        Search,
    }
});
</script>
