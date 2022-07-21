export interface userInfoType {
    uid: number,
    name: string,
    avatar: string,
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

