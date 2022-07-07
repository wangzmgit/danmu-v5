import { ref } from 'vue';
import { useNotification } from 'naive-ui';
import { sendRegisterCodeAPI } from '@/api/code.js';

export default function useSendCode() {

    const disabledSend = ref(false);//禁用发送按钮
    const sendBtnText = ref("发送验证码");//发送按钮文字
    const notification = useNotification();//通知

    const sendCode = (email) => {
        //禁用发送按钮
        disabledSend.value = true;
        sendRegisterCodeAPI(email).then((res) => {
            if (res.data.code === 2000) {
                if (res.data.data.code) {
                    notification.success({
                        title: '发送成功',
                        content: "验证码为:" + res.data.data.code,
                        duration: 10000,
                    })
                } else {
                    notification.success({
                        content: '发送成功',
                        duration: 5000,
                    })
                }

                //开启倒计时
                let count = 0;
                let tag = setInterval(() => {
                    if (++count >= 60) {
                        clearInterval(tag);
                        disabledSend.value = false;
                        sendBtnText.value = "发送验证码";
                        return;
                    }
                    sendBtnText.value = `${60 - count}秒后获取`;
                }, 1000);
            }
        }).catch((err) => {
            disabledSend = false;
            sendBtnText.value = "发送验证码";
            notification.error({
                title: '发送失败',
                content: "原因:" + err.response.data.msg,
                duration: 5000,
            });
        });
    }

    return {
        disabledSend,
        sendBtnText,
        sendCode
    }
}

