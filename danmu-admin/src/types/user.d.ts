export interface userType {
    uid: number,
    name: string,
    avatar: string,
    email: string,
    sign: string,
    created_at: string,
    role: number
}

export interface loginFormType {
    email: string,
    password: string,
}

export interface modifyUserType {
    id: number,
    name: string,
    email: string,
    sign: string
}

export interface modifyRoleType {
    uid: number,
    role: number
}
