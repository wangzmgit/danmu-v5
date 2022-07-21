import request from '@/utils/request'

//发送验证码
export const sendRegisterCodeAPI = (email: string) => {
    return request.post('v1/user/register/code', { email });
}