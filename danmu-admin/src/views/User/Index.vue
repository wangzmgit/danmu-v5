<template>
    <n-card class="card">
        <div class="info-header">
            <search-box :keyword="keyword" @search="searchUser"></search-box>
        </div>
        <n-table class="table" striped>
            <thead class="table-head">
                <tr>
                    <th>UID</th>
                    <th>头像</th>
                    <th>昵称</th>
                    <th>签名</th>
                    <th>邮箱</th>
                    <th>注册时间</th>
                    <th>权限</th>
                    <th>操作</th>
                </tr>
            </thead>
            <tbody class="table-body">
                <tr v-for="item in userList">
                    <td>{{ item.uid }}</td>
                    <td>
                        <common-avatar :url="item.avatar"></common-avatar>
                    </td>
                    <td>{{ item.name }}</td>
                    <td>{{ item.sign }}</td>
                    <td>{{ item.email }}</td>
                    <td>
                        <n-time :time="new Date(item.created_at)" />
                    </td>
                    <td>
                        <n-select :disabled="disabledRole" v-model:value="item.role" :options="roleOptions"
                            @update:value="val => changeRole(val, item.uid)" />
                    </td>
                    <td class="btn-list">
                        <n-button text @click="showEditDrawer(item)">编辑信息</n-button>
                        <n-button text @click="deleteUser(item.uid)">删除</n-button>
                    </td>
                </tr>
            </tbody>
        </n-table>
        <div class="pagination">
            <n-pagination v-model:page="page" :item-count="count" :page-size="5" @update-page="pageChange" />
        </div>
    </n-card>
    <!--新增分区-->
    <n-drawer v-model:show="showEdit" :width="500" placement="right">
        <n-drawer-content title="编辑信息">
            <n-form ref="formRef" :rules="rules" :model="editForm" label-width="auto">
                <n-form-item label="昵称" path="name">
                    <n-input placeholder="请输入昵称" v-model:value="editForm.name" />
                </n-form-item>
                <n-form-item label="邮箱" path="email">
                    <n-input placeholder="请输入邮箱" v-model:value="editForm.email" />
                </n-form-item>
                <n-form-item label="个性签名">
                    <n-input v-model:value="editForm.sign" placeholder="请输入个性签名" maxlength="50" show-count
                        type="textarea" :autosize="{ minRows: 3, maxRows: 3 }" />
                </n-form-item>
                <n-button type="primary" @click="submitEdit">保存</n-button>
            </n-form>
        </n-drawer-content>
    </n-drawer>
</template>

<script lang="ts">
import { useRoute } from 'vue-router';
import { userType, modifyUserType } from '@/types/user';
import storage from '@/utils/stored-data';
import SearchBox from '@/components/SearchBox.vue';
import CommonAvatar from '@/components/CommonAvatar.vue';
import { defineComponent, onBeforeMount, reactive, ref } from 'vue';
import {
    NTable, NButton, NCard, NTime, NInput, NDrawer, NSelect, FormRules,
    FormInst, NDrawerContent, NForm, NFormItem, NPagination, useNotification
} from 'naive-ui';
import { getUserListAPI, searchUserAPI, modifyUserAPI, deleteUserAPI, modifyRoleAPI } from '@/api/user';

export default defineComponent({
    setup() {
        const page = ref(1);
        const count = ref(0);
        const userList = ref<Array<userType>>([]);
        const notification = useNotification();//通知

        const getUserList = () => {
            getUserListAPI(page.value, 6).then((res) => {
                if (res.data.code === 2000) {
                    count.value = res.data.data.count;
                    userList.value = res.data.data.users;
                }
            }).catch(() => {
                notification.error({
                    title: '加载失败',
                    duration: 5000,
                })
            });
        }

        //页码改变
        const pageChange = (target: number) => {
            page.value = target;
            getUserList();
        }

        //权限
        const roleOptions = [
            {
                label: '用户',
                value: 0
            },
            {
                label: '审核',
                value: 1
            },
            {
                label: '管理',
                value: 2
            },
        ]
        const disabledRole = storage.get("adminInfo").role === 3 ? false : true;
        const changeRole = (val: number, uid: number) => {
            modifyRoleAPI({ uid: uid, role: val }).then((res) => {
                if (res.data.code === 2000) {
                    getUserList();
                }
            }).catch(() => {
                notification.error({
                    title: '修改失败',
                    duration: 5000,
                })
            });
        }


        //搜索
        const keyword = ref('');
        const searchUser = (keyword: string) => {
            page.value = 1;
            count.value = 0;
            if (!keyword) {
                getUserList();
                return;
            }
            searchUserAPI(keyword).then((res) => {
                if (res.data.code === 2000) {
                    userList.value = res.data.data.users;
                }
            })
        }

        //编辑用户信息
        const editForm = reactive<modifyUserType>({
            id: 0,
            name: '',
            email: '',
            sign: ''
        });
        const showEdit = ref(false);
        const formRef = ref<FormInst | null>(null);
        const rules: FormRules = {
            email: [
                { required: true, message: "请输入邮箱", trigger: ['blur', 'input'] },
                { type: "email", message: "请输入正确的邮箱地址", trigger: ['blur', 'input'] },
            ],
            name: { required: true, message: '请输入昵称', trigger: ['blur', 'input'] },
        }
        //打开修改抽屉
        const showEditDrawer = (item: any) => {
            editForm.id = item.uid;
            editForm.name = item.name;
            editForm.email = item.email;
            editForm.sign = item.sign;
            showEdit.value = true;
        }

        const submitEdit = () => {
            formRef.value?.validate((errors) => {
                if (!errors) {
                    modifyUserAPI(editForm).then((res) => {
                        if (res.data.code === 2000) {
                            notification.success({
                                title: '修改成功',
                                duration: 5000,
                            })
                            getUserList();
                            showEdit.value = false;
                        }
                    }).catch((err) => {
                        notification.error({
                            title: '修改失败',
                            content: "原因:" + err.response.data.msg,
                            duration: 5000,
                        })
                    });
                } else {
                    notification.error({
                        title: '请检查输入的数据',
                        duration: 5000,
                    })
                }
            })
        }

        // 删除
        const deleteUser = (id: number) => {
            deleteUserAPI(id).then((res) => {
                if (res.data.code === 2000) {
                    getUserList();
                }
            }).catch(() => {
                notification.error({
                    title: '删除失败',
                    duration: 5000,
                })
            });
        }

        const route = useRoute();
        onBeforeMount(() => {
            if (route.query.uid) {
                keyword.value = route.query.uid.toString();
                searchUser(keyword.value);
            } else {
                searchUser("");
            }
        })

        return {
            page,
            count,
            rules,
            formRef,
            keyword,
            userList,
            showEdit,
            editForm,
            disabledRole,
            roleOptions,
            changeRole,
            submitEdit,
            searchUser,
            pageChange,
            deleteUser,
            showEditDrawer
        }
    },
    components: {
        NForm,
        NInput,
        NTable,
        NCard,
        NTime,
        NSelect,
        NButton,
        NDrawer,
        NFormItem,
        SearchBox,
        CommonAvatar,
        NPagination,
        NDrawerContent,
    }
});
</script>

<style lang="less" scoped>
.card {
    margin: 10px;
    width: calc(100% - 20px);
    border-radius: 10px;

    .info-header {
        float: right;
        height: 30px;
        width: 300px;
        margin-bottom: 20px;
    }

    .table {

        .table-head {
            text-align: center;
        }

        .table-body {
            text-align: center;

            .btn-list {
                button {
                    margin: 0 6px;
                }
            }
        }
    }
}

.pagination {
    margin-top: 20px;
}
</style>