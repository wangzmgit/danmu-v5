import request from '@/utils/request';
import storage from "@/utils/stored-data";

export const getAccessToken = async () => {
    const res = await request.get('v1/user/token/refresh', {
        headers: {
            Authorization: `Bearer ${storage.get('admin_refresh_token')}`
        }
    });

    return res
}