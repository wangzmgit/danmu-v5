import { RouteRecordRaw } from "vue-router";
const upload = () => import("../views/Upload/Index.vue");
const uploadVideo = () => import("../views/Upload/UploadVideo/Index.vue");
const videoManage = () => import("../views/Upload/VideoManage/Index.vue");
const commentManage = () => import("../views/Upload/CommentManage/Index.vue");

const uploadRouter: Array<RouteRecordRaw> = [
    {
        path: "/upload",
        name: "Upload",
        component: upload,
        redirect: '/upload/video',
        children: [
            {
                path: '/upload/video/:vid?',
                name: 'UploadVideoHome',
                meta: { auth: true },
                component: uploadVideo,
            },
            {
                path: '/upload/video/manage',
                name: 'VideoManage',
                meta: { auth: true },
                component: videoManage,
            },
            {
                path: '/upload/comment/manage',
                name: 'CommentManage',
                meta: { auth: true },
                component: commentManage,
            },
        ]
    }
]

export default uploadRouter;