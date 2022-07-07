var storage = {
    //key保存键,value保存内容,*expired 失效时间,单位分钟
    set(key, value, expired) {
        const now = Date.now();
        let source = { key: key, value: value };

        if (expired) {
            source.value = JSON.stringify({ data: value, expired: now + (1000 * 60 * expired) });
        } else {
            source.value = JSON.stringify({ data: value });
        }
        localStorage.setItem(source.key, source.value);
    },
    get(key) {
        const now = Date.now();
        let source = { key: key, value: null };
        source.value = JSON.parse(localStorage.getItem(source.key));

        if (source.value) {
            //超过失效时 删除
            if (!source.value.expired) {
                return source.value.data;
            }
            else if (now >= source.value.expired) {
                this.remove(source.key);
                return null;
            } else {
                return source.value.data;
            }
        }
    },
    remove(key) {
        localStorage.removeItem(key);
    },
}

export default storage;
