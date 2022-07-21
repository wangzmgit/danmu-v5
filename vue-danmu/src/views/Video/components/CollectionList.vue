<template>
    <div class="login-bg">
        <div class="login-card">
            <div class="card-head">
                <span>添加收藏</span>
                <n-icon class="close-icon" @click="closeCard">
                    <Close />
                </n-icon>
            </div>
            <n-scrollbar style="max-height: 270px;margin-top: 10px;">
                <n-checkbox-group v-if="!loading" :default-value="defaultChecked" @update:value="handleUpdateValue">
                    <div class="collention-item" v-for="item in collections">
                        <n-checkbox :value="item.id" :label="item.name" />
                        <span>{{ item.open ? '' : '私密' }}</span>
                    </div>
                </n-checkbox-group>
                <div class="add-collection">
                    <n-button v-if="!isAdd" type="primary" ghost @click="changeAdd">新建收藏夹</n-button>
                    <n-input-group v-else>
                        <n-input ref="addInput" v-model:value="name" placeholder="收藏夹名称" maxlength="20" show-count
                            @blur="changeAdd" />
                        <!-- 使用mousedown触发而不是click触发，因为input的blur要早于click -->
                        <n-button style="width: 30%;" type="primary" @mousedown="addCollection">创建</n-button>
                    </n-input-group>
                </div>
            </n-scrollbar>
            <div class="save-btn">
                <n-button type="primary" @mousedown="save">完成</n-button>
            </div>
        </div>
    </div>
</template>

<script lang="ts">
import { onBeforeMount, defineComponent, nextTick, ref } from 'vue';
import {
    NIcon, NButton, NInput, NInputGroup,
    NCheckbox, NCheckboxGroup, NScrollbar, useNotification
} from 'naive-ui';
import { Close } from '@vicons/ionicons5';
import { getCollectInfoAPI, collectAPI } from '@/api/collect';
import useCollection from '@/hooks/collection';
export default defineComponent({
    emits: ['close'],
    props: {
        vid: {
            type: Number,
            required: true
        }
    },
    setup(props, ctx) {
        const notification = useNotification();

        const loading = ref(true);//加载中
        const defaultChecked = ref<Array<number>>([]);//默认选中
        const { collections, getCollectionList, createCollection } = useCollection();

        const closeCard = () => {
            ctx.emit('close', 0);
        }

        // 获取收藏信息
        const getCollectInfo = () => {
            getCollectInfoAPI(props.vid).then((res) => {
                if (res.data.code === 2000) {
                    defaultChecked.value = res.data.data.collect;
                    loading.value = false;
                }
            })
        }

        //新建收藏夹
        const name = ref("")
        const isAdd = ref(false);
        const addInput = ref<HTMLElement | null>(null);
        const changeAdd = () => {
            isAdd.value = !isAdd.value;
            if (isAdd.value) {
                //此时input不存在，无法修改焦点
                nextTick(() => {
                    addInput.value!.focus();
                })
            }
        }

        //收藏相关
        const checkedValue = ref<Array<number>>([])
        const handleUpdateValue = (value: Array<number | string>) => {
            checkedValue.value = value as Array<number>;
        }

        //新建收藏夹
        const addCollection = () => {
            if (!name.value) {
                notification.error({
                    title: '请输入收藏夹名',
                });
                return;
            }

            createCollection(name.value);
            name.value = "";
        }

        //保存收藏
        const save = () => {
            //原数组不存在新数组存在表示添加
            const addList = checkedValue.value.filter((v) => {
                return defaultChecked.value.indexOf(v) == -1
            })

            //原数组存在新数组不存在表示删除
            const cancelList = defaultChecked.value.filter((v) => {
                return checkedValue.value.indexOf(v) == -1
            })
            collectAPI({ vid: props.vid, addList: addList, cancelList: cancelList }).then((res) => {
                if (res.data.code === 2000) {
                    var count = 0;
                    //否则收藏量不变
                    if (defaultChecked.value.length === 0 && checkedValue.value.length !== 0)
                        count = 1; //之前没有收藏，之后收藏了，收藏量+1
                    else if (defaultChecked.value.length !== 0 && checkedValue.value.length === 0)
                        count = -1;//之前收藏了，之后没有收藏，收藏量-1
                    ctx.emit('close', count);
                }
            }).catch((err) => {
                notification.error({
                    title: '保存失败',
                    content: "原因:" + err.response.data.msg,
                    duration: 5000,
                })
            });
        }

        onBeforeMount(() => {
            getCollectionList();
            getCollectInfo();
        })

        return {
            name,
            isAdd,
            loading,
            addInput,
            collections,
            defaultChecked,
            save,
            closeCard,
            changeAdd,
            addCollection,
            handleUpdateValue,
            getCollectionList
        }
    },
    components: {
        NIcon,
        NButton,
        NInput,
        NInputGroup,
        NScrollbar,
        NCheckbox,
        NCheckboxGroup,
        Close
    }
});
</script>

<style lang="less" scoped>
.login-bg {
    top: 0;
    left: 0;
    z-index: 999;
    position: fixed;
    width: 100%;
    height: 100%;
    background-color: rgba(0, 0, 0, 0.45);
}

.login-card {
    position: absolute;
    top: 50%;
    left: 50%;
    margin: -220px 0 0 -163px;
    width: 330px;
    height: 360px;
    padding: 6px 12px;
    background-color: #fff;
    border-radius: 10px;
    animation: fadein .3s ease-in;
    box-shadow: 16px 16px 50px -12px rgba(0, 0, 0, 0.8);
}

.card-head {
    display: flex;
    height: 30px;
    align-items: center;
    justify-content: space-between;

    // 关闭按钮
    .close-icon {
        font-size: 26px;
        color: #adadad;
        cursor: pointer;

        &:hover {
            color: #18a058;
        }
    }

}

.collention-item {
    height: 36px;
    padding: 0 6px;
    line-height: 36px;

    span {
        float: right;
        font-size: 12px;
        color: #999;
    }
}

.add-collection {
    width: 100%;

    button {
        width: 100%;
    }
}

.save-btn {
    width: 100%;
    text-align: center;

    button {
        width: 160px;
        margin-top: 10px;
    }
}
</style>