import { ref } from 'vue';


export default function useMsg() {
    const msg = ref("");
    const showMsg = ref(false);
    //显示消息
    const changeMsg = (message: string) => {
        msg.value = message;
        showMsg.value = true;
        //定时隐藏
        setTimeout(() => {
            showMsg.value = false;
        }, 3000);
    }
    return {
        msg,
        showMsg,
        changeMsg
    }
}

