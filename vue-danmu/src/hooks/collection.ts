import { ref } from 'vue';
import { useNotification } from 'naive-ui';
import { getCollectionListAPI, addCollectionAPI } from '@/api/collect';
import { collectionType } from '@/types/collect';

export default function useCollection() {
    const collections = ref<Array<collectionType>>([]);
    const notification = useNotification();//通知

    //获取收藏夹列表
    const getCollectionList = () => {
        getCollectionListAPI().then((res) => {
            if (res.data.code === 2000) {
                collections.value = res.data.data.collections;
            }
        })
    }

    //创建收藏夹
    const createCollection = (name:string) => {
        addCollectionAPI(name).then((res) => {
            if (res.data.code === 2000) {
                collections.value.push({
                    id: res.data.data.id,
                    cover: "",
                    name: name,
                    desc: "",
                    open: false,
                })
            }
        }).catch((err) => {
            notification.error({
                title: '创建失败',
                content: "原因:" + err.response.data.msg,
                duration: 5000,
            })
        });
    }

    return {
        collections,
        getCollectionList,
        createCollection
    }
}

