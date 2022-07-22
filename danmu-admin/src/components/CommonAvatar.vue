<template>
    <n-avatar v-if="url" round :size="size" :src="url" />
    <n-avatar v-else round :size="size">
        <n-icon :size="iconsize">
            <Person />
        </n-icon>
    </n-avatar>
</template>

<script lang="ts">
import { NAvatar, NIcon } from "naive-ui";
import { Person } from '@vicons/ionicons5';
import { defineComponent } from "vue";
export default defineComponent({
    props: {
        url: {
            type: String,
        },
        size: {
            type: Number,
            default: 50
        },
        iconsize: {
            type: Number,
            default: 26
        }
    },
    components: {
        NIcon,
        NAvatar,
        Person
    }
});
</script>