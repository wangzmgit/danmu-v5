import { defineStore } from 'pinia';

export const useLoginStore = defineStore({
    id: "login",
    state: () => {
        return {
            login: false,
        }
    },
    actions: {
        setLoginState(val: boolean) {
            this.login = val;
        }
    },
    persist: {
        enabled: true,
        strategies: [{
            key: 'logged',
            storage: localStorage,
        }]
    }
})
