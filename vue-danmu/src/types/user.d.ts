export interface userInfoType {
    uid: number,
    name: string,
    avatar: string,
    email?: string,
    gender?: number,
    sign?: string,
    birthday?: string
}

export interface registerFormType {
    email: string,
    password: string,
    code: string
}

export interface loginFormType {
    email: string,
    password: string,
}

export interface modifyInfoType {
    name: string,
    gender: number,
    birthday: string,
    sign: string
}
