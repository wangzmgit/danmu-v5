<template>
    <div class="msg">
        <!--左边-->
        <div class="msg-left">
            <div class="left-top"></div>
            <div v-for="(item, index) in msgList" :key="index">
                <div class="msg-user-item" @click="getMsgContent(item, index)">
                    <common-avatar class="msg-avatar" :url="item.avatar" :size="45"></common-avatar>
                    <span class="msg-name">{{ item.name }}</span>
                    <n-time class="msg-date" :time="new Date(item.created_at)" type="relative"></n-time>
                    <div class="dot" v-if="!item.status"></div>
                </div>
            </div>
        </div>
        <!--右侧-->
        <div class="msg-right">
            <div class="left-top right-top">{{ targetUser.name }}</div>
            <div id="msgBox" class="msg-main">
                <div class="content-box" v-for="(item, index) in msgDetails" :key="index">
                    <!--自己发送的-->
                    <div v-if="item.from_id == userInfo.uid">
                        <common-avatar class="avatar-right" :url="userInfo.avatar" :size="45"></common-avatar>
                        <span class="content-right">{{ item.content }}</span>
                    </div>
                    <!--收到的-->
                    <div v-else>
                        <common-avatar class="avatar-left" :url="targetUser.avatar" :size="45"></common-avatar>
                        <div class="content-left-box">
                            <span class="content-left">{{ item.content }}</span>
                        </div>
                    </div>
                    <!-- 解决div无法撑开的问题 -->
                    <div style="clear:both;" />
                </div>
            </div>

            <div class="msg-input">
                <n-input v-model:value="msgForm.content" type="textarea" placeholder="发个消息呗~" maxlength="255" show-count
                    :autosize="descSize" @keydown.enter="sendMsg"></n-input>
                <div class="btn-box">
                    <div>
                        <!-- 此处可以放图片/emoji图标 -->
                    </div>
                    <n-button type="primary" :loading="sendLoading" @click="sendMsg">发送</n-button>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import { Base64 } from 'js-base64';
import { useRoute } from 'vue-router';
import storage from '@/utils/stored-data';
import { MsgSocketURL } from '@/utils/request.js';
import { getUserInfoByIDAPI } from '@/api/user.js';
import { getMsgListAPI, getMsgDetailsAPI, readMsgAPI, sendMsgAPI } from '@/api/message.js';
import CommonAvatar from '@/components/CommonAvatar.vue';
import { NInput, NButton, NTime, useNotification } from 'naive-ui';
import { onBeforeMount, onBeforeUnmount, reactive, ref, computed, nextTick } from 'vue';

export default {
    setup() {
        const descSize = {
            minRows: 4,
            maxRows: 4
        }
        const msgList = ref([]);
        const msgDetails = ref([]);
        const msgForm = reactive({
            fid: 0,
            content: ''
        })
        const notification = useNotification();//通知
        //用户信息
        const userInfo = computed(() => {
            const userInfo = storage.get("userInfo");
            if (userInfo) {
                return userInfo;
            }
            return {};
        })

        //获取消息列表
        const getMsgList = () => {
            getMsgListAPI().then((res) => {
                if (res.data.code === 2000) {
                    msgList.value = res.data.data.messages;
                    if (msgForm.fid !== 0) {
                        initSendUser(msgForm.fid)
                    } else if (msgList.value.length > 0) {
                        getMsgContent(msgList.value[0], 0);
                    }
                }
            });
        }

        //初始化发送的用户
        const initSendUser = (fid) => {
            //遍历当前消息列表查找用户
            for (let i = 0; i < msgList.value.length; i++) {
                if (msgList.value[i].uid === fid) {
                    getMsgContent(msgList.value[i], i);
                    return;
                }
            }
            //获取用户信息并添加到用户列表
            getUserInfoByIDAPI(fid).then((res) => {
                if (res.data.code === 2000) {
                    const user = res.data.data.user;
                    msgList.value.unshift({
                        uid: user.uid,
                        avatar: user.avatar,
                        name: user.name,
                        created_at: new Date(),
                        status: true
                    });
                    targetUser.name = user.name;
                    targetUser.avatar = user.avatar;
                }
            })
        }

        //获取消息详情
        const page = ref(1);
        const loading = ref(false);
        const noMore = ref(false);//没有更多
        const allowToBottom = ref(true);//是否允许自动跳转到底部
        const targetUser = reactive({
            name: "",
            avatar: ""
        })

        const getMsgContent = (item, index) => {
            page.value = 1;
            noMore.value = false;
            msgDetails.value = [];
            msgForm.fid = item.uid;
            targetUser.name = item.name;
            targetUser.avatar = item.avatar;
            getMoreMsgContent(item.uid);
            //已读消息
            readMsgAPI(item.uid);
            msgList.value[index].status = true;
        }

        //获取更多消息
        const getMoreMsgContent = (fid) => {
            loading.value = true;
            getMsgDetailsAPI(fid, page.value, 10).then((res) => {
                if (res.data.code === 2000) {
                    const resMsg = res.data.data.messages;
                    if (resMsg.length < 10) {
                        noMore.value = true;
                    }
                    msgDetails.value = resMsg.concat(msgDetails.value);
                    nextTick(() => {
                        toBottom()
                    })
                }
                loading.value = false;
            });
        }

        //加载更多
        const lazyLoading = (e) => {
            if (e.target.id === "msgBox") {
                let scroll = document.getElementById("msgBox");
                if (scroll.scrollTop < 30 && !loading.value && !noMore.value) {
                    page.value++;
                    allowToBottom.value = false;
                    getMoreMsgContent(msgForm.fid);
                }
            }
        }

        //到达底部
        const toBottom = () => {
            let scroll = document.getElementById("msgBox");
            if (allowToBottom.value) {
                scroll.scrollTop = scroll.scrollHeight;
            } else {
                scroll.scrollTop = 100;
                allowToBottom.value = true;
            }
        }

        //发送消息
        const sendLoading = ref(false);
        const sendMsg = () => {
            if (!msgForm.content) {
                notification.error({
                    title: '不能发送空消息',
                    duration: 3000,
                });
                return;
            }
            sendLoading.value = true;
            sendMsgAPI(msgForm).then((res) => {
                if (res.data.code === 2000) {
                    msgDetails.value.push({
                        from_id: userInfo.value.uid,
                        content: msgForm.content
                    });
                    msgForm.content = "";
                    sendLoading.value = false;
                    nextTick(() => {
                        toBottom()
                    })
                }
            }).catch((err) => {
                sendLoading.value = false;
                notification.error({
                    title: '发送失败',
                    description: err.response.data.msg,
                    duration: 3000,
                });

            });
        }

        //websocket
        const SocketURL = ref("");
        const websocket = ref(null);
        //初始化weosocket
        const initWebSocket = () => {
            const wsProtocol = window.location.protocol === 'http:' ? 'ws://' : 'wss://';
            if (MsgSocketURL === "/api/v1/message/ws") {
                SocketURL.value = wsProtocol + window.location.host + "/api/v1/message/ws";
            } else {
                //处理协议部分
                let reg = new RegExp('^http(s)?:')
                SocketURL.value = MsgSocketURL.replace(reg, wsProtocol);
            }
            SocketURL.value += "?token=" + storage.get("token");
            websocket.value = new WebSocket(SocketURL.value);
            websocket.value.onmessage = websocketOnmessage;
        }

        //数据接收
        const websocketOnmessage = (e) => {
            const res = JSON.parse(Base64.decode(e.data));
            if (msgForm.fid === res.fid) {
                msgDetails.value.push({
                    from_id: res.fid,
                    content: res.content,
                });
                nextTick(() => {
                    toBottom()
                })
            } else {
                msgForm.fid = fid;
                initSendUser(res.fid);
            }
        }

        //加载和卸载页面
        const route = useRoute();
        onBeforeMount(() => {
            if (route.params.fid) {
                msgForm.fid = Number(route.params.fid);
            }
            getMsgList();
            initWebSocket();
            window.addEventListener('scroll', lazyLoading, true);
        })

        onBeforeUnmount(() => {
            if (websocket.value) {
                websocket.value.close();
            }
            window.removeEventListener('scroll', lazyLoading);
        })

        return {
            userInfo,
            descSize,
            msgList,
            msgDetails,
            targetUser,
            msgForm,
            sendLoading,
            sendMsg,
            getMsgContent
        }
    },
    components: {
        NTime,
        NInput,
        NButton,
        CommonAvatar
    }
};
</script>

<style lang="less" scoped>
.msg {
    display: flex;
    height: 656px;
}

.msg-left {
    width: 220px;
    border-right: 1px solid #e1e2e3;

    .left-top {
        height: 40px;
        border-bottom: 1px solid #e1e2e3;
    }

    .msg-user-item {
        width: 100%;
        height: 60px;
        background-color: #fff;
        position: relative;
        border-bottom: 1px solid #e0e0e0;

        &:hover {
            background-color: #f7f7f7;
        }

        .msg-avatar {
            margin: 8px 0 0 6px;
        }

        .dot {
            top: 10px;
            left: 42px;
            width: 10px;
            height: 10px;
            position: absolute;
            border-radius: 50%;
            background-color: #f5222d;
        }

        .msg-name {
            position: absolute;
            top: 8px;
            left: 60px;
            color: #333;
            font-size: 14px;
        }

        .msg-date {
            position: absolute;
            top: 32px;
            left: 60px;
            color: #999;
            font-size: 12px;
        }
    }
}

.msg-right {
    width: calc(100% - 220px);

    .right-top {
        color: #333;
        text-align: center;
        font-size: 16px;
        line-height: 40px;
        border-bottom: 1px solid #e1e2e3;
    }

    .msg-main {
        height: 450px;
        background-color: #f7f7f7;
        border-bottom: 1px solid #e1e2e3;
        overflow-y: auto;

        /**修改滚动条样式 */
        &::-webkit-scrollbar {
            width: 8px;
        }

        &::-webkit-scrollbar-thumb {
            /*滚动条里面小方块*/
            border-radius: 2px;
            background: #999999;
        }

        &::-webkit-scrollbar-track {
            /*滚动条里面轨道*/
            border-radius: 5px;
        }
    }

}

/**聊天部分 */
.content-box {
    min-height: 70px;
    margin: 6px 20px;

    &:nth-child(1) {
        margin-top: 10px;
    }

    .avatar-right {
        float: right;
    }

    .content-right {
        float: right;
        max-width: 80%;
        margin-right: 10px;
        margin-top: 6px;
        background-color: #40a9ff;
        color: #fff;
        font-size: 16px;
        border-radius: 3px;
        padding: 5px 10px 5px 10px;
    }

    .avatar-left {
        float: left;
    }

    .content-left-box {
        float: left;
        margin-left: 10px;
        margin-top: 10px;
        max-width: 80%;
        background: #fff;
        padding: 5px 10px 5px 10px;
        border-radius: 3px;
    }

    .content-left {
        font-size: 1rem;
    }
}

.msg-input {
    padding: 10px;

    .btn-box {
        height: 50px;
        display: flex;
        margin-left: 10px;
        align-items: center;
        justify-content: space-between;

        button {
            width: 100px;
            height: 30px;
        }
    }
}
</style>