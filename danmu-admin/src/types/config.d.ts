export interface databaseType {
    dbName: string,
    dbHost: string,
    dbPort: number,
    dbUser: string,
    dbPwd: string,
    cacheHost: string,
    cachePort: number,
    cachePwd: string
}

export interface emailType {
    open: boolean,
    name: string,
    host: string,
    port: number,
    address: string,
    password: string
}

export interface storageType {
    https: boolean,
    maxImgSize: number,
    maxVideoSize: number,
    oss: boolean,
    domain: string,
    bucket: string,
    endpoint: string,
    accesskeyId: string,
    accesskeySecret: string
}

export interface otherType {
    prefix: string,
    maxRes: number,
    fePath: string,
    email: string,
    password: string
}