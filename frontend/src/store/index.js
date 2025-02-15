import { defineStore } from 'pinia';

const useStore = defineStore('store', {
    persist: {
        enabled: true,
    },
    state: () => ({
        isAbout: true,
        loading: true,
        forms: {
            action: "",
        },
        queryData: {},
        planId: {}
    }),
    getters: {

    },
    actions: {
        setAction(value) {
            this.forms.action = value;
        },
        setIsAbout(value) {
            this.isAbout = value;
        },
        setQueryData(key, value) {
            this.queryData[key] = value;
        },
        setPlanId(value) {
            this.planId = value;
        }
    }
})
export default useStore