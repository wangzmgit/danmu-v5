<template>
    <n-form label-placement="left" label-width="auto">
        <div class="avatar-box">
            <p class="avatar-title">头像:</p>
            <div @mouseover="mask = true" @mouseleave="mask = false">
                <n-upload name="avatar" :show-file-list="false" :action="avatarUrl" :headers="headers"
                    :before-upload="beforeUploadAvatar" @change="handleChange">
                    <div v-show="mask" class="mask">
                        <p>更换头像</p>
                    </div>
                    <n-avatar class="avatar" v-if="userInfo.avatar" round :size="60" :src="userInfo.avatar" />
                    <n-avatar class="avatar" v-else round :size="60">
                        <n-icon size="30">
                            <Person />
                        </n-icon>
                    </n-avatar>
                </n-upload>
            </div>
        </div>
        <n-form-item label="UID:">
            <p class="uid form-item">{{ userInfo.uid }}</p>
        </n-form-item>
        <n-form-item label="昵称:">
            <n-input class="form-input name" v-model:value="userInfo.name" placeholder="请输入昵称" maxlength="20"
                show-count />
        </n-form-item>
        <n-form-item label="性别:">
            <n-radio-group class="form-item" v-model:value="userInfo.gender">
                <n-radio-button value="0">未知</n-radio-button>
                <n-radio-button value="1">男</n-radio-button>
                <n-radio-button value="2">女</n-radio-button>
            </n-radio-group>
        </n-form-item>
        <n-form-item label="生日:">
            <n-date-picker class="form-item" v-model:value="userInfo.birthday" placeholder="选择出生日期" type="date" />
        </n-form-item>
        <n-form-item label="个性签名:">
            <n-input class="form-input" v-model:value="userInfo.sign" placeholder="请输入个性签名" maxlength="50" show-count
                type="textarea" :autosize="descSize" />
        </n-form-item>
        <div class="setting-save-btn">
            <n-button type="primary" @click="modifyUserInfo">修改</n-button>
        </div>
    </n-form>
</template>

<script>
import config from "@/config";
import { Person } from '@vicons/ionicons5';
import storage from "@/utils/stored-data.js";
import { timestampToTime } from '@/utils/time.js';
import { ref, reactive, onBeforeMount } from "vue";
import {
    NForm, NIcon, NFormItem, NButton, NUpload, NAvatar,
    NRadioGroup, NRadioButton, NInput, NDatePicker, useNotification
} from 'naive-ui';
import { AvatarUrl as avatarUrl } from '@/utils/request';
import { modifyUserInfoAPI, getUserInfoAPI } from "@/api/user";

export default {

    props: {
        info: {
            type: Object,
        }
    },
    setup() {
        //简介输入框大小
        const descSize = {
            minRows: 3,
            maxRows: 3
        }

        //文件上传请求头
        const headers = {
            Authorization: "Bearer " + storage.get("token"),
        }

        const mask = ref(false);//上传头像遮罩
        const userInfo = reactive({
            uid: 0,
            avatar: "",
            name: "",
            sign: "",
            gender: "",
            birthday: null,
        })

        const notification = useNotification();//通知


        //上传之前的回调
        const beforeUploadAvatar = (options) => {
            const file = options.file;
            const isJpgOrPng =
                file.type === "image/jpeg" || file.type === "image/png";
            if (!isJpgOrPng) {
                notification.error({
                    title: '上传失败',
                    content: "文件只支持jpg/jpeg/png格式",
                    duration: 5000,
                });
            }
            const isLtMaxSize = file.file.size / 1024 / 1024 < config.maxImgSize;
            if (!isLtMaxSize) {
                notification.error({
                    title: '上传失败',
                    content: `图片大小不能超过${config.maxImgSize}M`,
                    duration: 5000,
                });
            }
            return isJpgOrPng && isLtMaxSize;
        }

        //上传变化的回调
        const handleChange = (options) => {
            const status = options.file.status;
            if (status === "finished") {
                const res = JSON.parse(options.event.currentTarget.response);
                userInfo.avatar = res.data.url;
                getUserInfo();//获取用户信息
                notification.success({
                    title: '上传完成',
                    duration: 3000,
                });
            } else if (status === "error") {
                notification.error({
                    title: '文件上传失败',
                    duration: 5000,
                });
            }
        }

        const modifyUserInfo = () => {
            const modifyForm = {
                name: userInfo.name,
                sign: userInfo.sign,
                gender: Number(userInfo.gender),
                birthday: timestampToTime(userInfo.birthday, 'Y-M-D'),
            }
            if (!userInfo.name) {
                notification.error({
                    content: "请填写昵称",
                    duration: 5000,
                });
                return;
            }
            modifyUserInfoAPI(modifyForm).then((res) => {
                if (res.data.code === 2000) {
                    getUserInfo();//获取用户信息
                    notification.success({
                        title: '修改成功',
                        duration: 5000,
                    })
                }
            }).catch((err) => {
                notification.error({
                    title: '修改失败',
                    content: "原因:" + err.response.data.msg,
                    duration: 5000,
                })
            });
        }

        //获取用户信息
        const getUserInfo = () => {
            getUserInfoAPI().then((res) => {
                if (res.data.code === 2000) {
                    storage.update("userInfo", res.data.data.user);
                }
            })
        }

        onBeforeMount(() => {
            const info = storage.get('userInfo');
            userInfo.uid = info.uid;
            userInfo.name = info.name;
            userInfo.sign = info.sign;
            userInfo.avatar = info.avatar;
            userInfo.gender = info.gender.toString();
            userInfo.birthday = Date.parse(info.birthday);
        })

        return {
            mask,
            headers,
            descSize,
            userInfo,
            avatarUrl,
            getUserInfo,
            handleChange,
            modifyUserInfo,
            beforeUploadAvatar
        }
    },
    components: {
        NIcon,
        NForm,
        NInput,
        NAvatar,
        NUpload,
        NButton,
        NFormItem,
        NDatePicker,
        NRadioGroup,
        NRadioButton,
        Person
    }
}
</script>

<style lang="less" scoped>
.avatar-box {
    height: 60px;
    margin: 20px 0;
    display: flex;
    position: relative;

    .avatar-title {
        width: 60px;
        margin: 0;
        text-align: right;
        font-size: 14px;
        line-height: 60px;
    }

    .mask {
        position: absolute;
        width: 60px;
        height: 60px;
        z-index: 9999;
        margin-left: 36px;
        border-radius: 30px;
        text-align: center;
        cursor: pointer;
        background-color: rgba(0, 0, 0, 0.3);

        p {
            color: #fff;
            margin: 0;
            font-size: 10px;
            line-height: 60px;
        }
    }

    .avatar {
        margin-left: 36px;
    }
}

.uid {
    margin: 0;
}

.form-item {
    padding-left: 20px;
}

.name {
    &:deep(.n-input__input-el) {
        height: auto;
    }
}

.form-input {
    margin-left: 20px;
}

.setting-save-btn {
    float: right;
    margin-top: 30px;

    button {
        width: 120px;
        height: 40px;
    }
}
</style>