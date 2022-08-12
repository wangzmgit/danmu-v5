<template>
    <div class="emoji" :style="{ width: emojiWidth }">
        <div class="emoji-body">
            <ul class="emoji-items">
                <li class="emoji-item" v-for="(item, index) in emojiList" :key="index" @click="chooseEmoji(item.name)">
                    <span class="emoji-img" :class="item.className"></span>
                </li>
            </ul>
        </div>
    </div>
</template>
 
<script lang="ts">
import { defineComponent } from "vue";
import { emojiList } from "../ts/emoji-list";

export default defineComponent({
    props: {
        emojiWidth: {
            type: String,
            default: "560px",
        },
    },
    setup(_props, ctx) {
        const chooseEmoji = (name: string) => {
            ctx.emit("chooseEmoji", name);
        }

        return {
            emojiList,
            chooseEmoji
        }
    }
});
</script>
 
<style lang="less" scoped>
.emoji {
    position: relative;
    z-index: 999;
    color: black;
}

.emoji-body {
    height: 200px;
    position: absolute;
    background: #fff;
    border: 1px solid #ddd;
    border-radius: 0 4px 4px 4px;
    overflow: hidden;

    .emoji-items {
        max-height: 197px;
        overflow-y: scroll;
        padding: 10px;
        
        .emoji-item {
            background: #f7f7f7;
            padding: 5px 10px;
            border-radius: 5px;
            display: inline-block;
            margin: 0 10px 12px 0;
            transition: 0.3s;
            line-height: 19px;
            font-size: 20px;
            cursor: pointer;


            &:hover {
                background: #eee;
                box-shadow: 0 2px 2px 0 rgba(0, 0, 0, 0.14), 0 3px 1px -2px rgba(0, 0, 0, 0.2),
                    0 1px 5px 0 rgba(0, 0, 0, 0.12);
            }

            .emoji-img {
                width: 72px;
                height: 72px;
                display: block;
                zoom: 50%;
            }
        }


        &::-webkit-scrollbar {
            width: 6px;
        }

        &::-webkit-scrollbar-thumb {
            border-radius: 3px;
            background: #b1b1b1;
        }

        &::-webkit-scrollbar-track {
            border-radius: 5px;
        }

    }

}
</style>