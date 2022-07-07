const upload = () => import("../views/Upload/Index");
const uploadVideo = () => import("../views/Upload/UploadVideo/Index");
const videoManage = () => import("../views/Upload/VideoManage/Index");
const commentManage = () => import("../views/Upload/CommentManage/Index");

const uploadRouter = [
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