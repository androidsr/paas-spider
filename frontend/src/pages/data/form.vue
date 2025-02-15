<template>
    <div>
        <a-page-header style="border: 1px solid rgb(235, 237, 240);padding: 8px;margin: 0;" title="数据详情" @back="back" />
        <a-card>
            <a-tabs v-model:activeKey="activeKey">
                <a-tab-pane v-for="(v, k) in dataMap" :key="k" :tab="k">
                    <a-typography>
                        <a-typography-title>{{ k }}</a-typography-title>
                        <a-typography-paragraph>{{ v }}</a-typography-paragraph>
                    </a-typography>
                </a-tab-pane>
            </a-tabs>
        </a-card>
    </div>
</template>
<script>
import { Get } from '../../../wailsjs/go/biz/RecordBiz.js';

export default {
    mounted() {
        this.taskId = this.$route.query.id;
        if (this.taskId) {
            Get(this.taskId).then(res => {
                if (res.code == 200) {
                    this.model = res.data || {};
                    try {
                        this.dataMap = JSON.parse(this.model.outputValue);
                    } catch (e) {
                        this.dataMap["数据内容"] = this.model.outputValue;
                    }
                }
            })
        }
    },
    data() {
        return {
            taskId: "",
            model: {},
            dataMap: {}
        }
    },
    methods: {
        back() {
            this.$router.go(-1);
        }
    },
}
</script>