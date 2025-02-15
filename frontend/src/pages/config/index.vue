<template>
    <a-layout style="min-height: 92vh;background-color: white;">
        <a-layout-content>
            <a-row justify="center" align="middle" style="height: 100%; padding: 20px;">
                <a-col :span="24">
                    <a-card title="设置" bordered>
                        <a-form :form="form" :label-col="{ span: 4 }" :wrapper-col="{ span: 20 }"
                            @submit.prevent="handleSubmit">
                            <a-form-item label="浏览器程序">
                                <a-input v-model:value="form.browser" placeholder="c:/xxx/chrome.exe" />
                            </a-form-item>

                            <a-form-item label="显示浏览器">
                                <a-switch v-model:checked="form.showBrowser" />
                            </a-form-item>
                            <a-form-item label="页面超时(毫秒)">
                                <a-input-number v-model:value="form.timeout" placeholder="" />
                            </a-form-item>
                            <a-form-item label="执行间隔(毫秒)">
                                <a-input-number v-model:value="form.startInterval" placeholder="最小值" />
                                <a-input-number v-model:value="form.endInterval" placeholder="最大值" />
                            </a-form-item>
                            <a-form-item label="页面上限">
                                <a-input-number v-model:value="form.pageLimit" placeholder="" />
                            </a-form-item>
                            <a-form-item label="链接最少文本">
                                <a-input-number v-model:value="form.linkText" placeholder="" />
                            </a-form-item>
                            <a-form-item label="纯净模式">
                                <a-switch v-model:checked="form.pureType" />
                            </a-form-item>

                            <a-form-item :wrapper-col="{ span: 3, offset: 20 }">
                                <a-button type="primary" html-type="submit">保存设置</a-button>
                            </a-form-item>
                        </a-form>
                    </a-card>
                </a-col>
            </a-row>
        </a-layout-content>
    </a-layout>
</template>

<script>
import { GetConfig, SetConfig } from '../../../wailsjs/go/main/App'
import { InitBrowser } from '../../../wailsjs/go/biz/TaskBiz.js';

export default {
    mounted() {
        GetConfig().then(res => {
            if (res.code == 200) {
                this.form = JSON.parse(res.data.content);
            } else {
                this.$message.error("获取设置失败！");
            }
        })
    },
    data() {
        return {
            form: {
                browser: "",
                showBrowser: false,
                timeout: 5000,
                pageLimit: 10,
                pureType: false,
                startInterval: 700,
                endInterval: 700,
                linkText: 6,
            },
        };
    },
    methods: {
        // 保存设置
        handleSubmit() {
            SetConfig({ id: "1", legalStatement: "1", content: JSON.stringify(this.form) }).then(res => {
                if (res.code == 200) {
                    InitBrowser()
                    this.$message.success("设置成功");
                } else {
                    this.$message.error("设置保存失败！");

                }
            })
        },
    },
};
</script>

<style scoped>
.a-layout-content {
    padding: 50px 0;
}

.a-card {
    padding: 20px;
}
</style>