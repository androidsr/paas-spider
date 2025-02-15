<template>
    <div>
        <a-page-header style="border: 1px solid rgb(235, 237, 240); padding: 8px; margin: 0;" title="操作表单"
            @back="back" />
        <a-card>
            <a-form :model="model" :rules="vRules" :label-col="{ style: { width: '80px' } }">
                <a-form-item label="操作分类" name="system">
                    <a-input v-model:value="model.system" :readonly="this.$store.forms.action == 'detail'" />
                </a-form-item>
                <a-form-item label="操作名称" name="name">
                    <a-input v-model:value="model.name" :readonly="this.$store.forms.action == 'detail'" />
                </a-form-item>
                <a-form-item label="执行时间" name="execTime" v-if="$store.forms.action == 'detail'">
                    <a-input v-model:value="model.execTime" :readonly="this.$store.forms.action == 'detail'" />
                </a-form-item>

                <a-tabs v-model:activeKey="activeKey" @change="tabChange">
                    <a-tab-pane key="config" tab="配置模型">
                        <div v-for="(field, index) in fields" :key="index">
                            <a-divider style="margin: 8px;">{{ field.name }}</a-divider>
                            <a-row>
                                <a-col :span="12">
                                    <a-form-item label="操作名称">
                                        <a-input v-model:value="field.name" style="width: 100%;"
                                            :readonly="isReadOnly" />
                                    </a-form-item>
                                </a-col>
                                <a-col :span="12">
                                    <a-form-item label="选择器">
                                        <a-mentions v-model:value="field.selector" autofocus :options="options"
                                            placeholder="输入/提示查找规则" prefix="/" :disabled="isReadOnly"></a-mentions>
                                    </a-form-item>
                                </a-col>
                            </a-row>

                            <a-row>
                                <a-col :span="12">
                                    <a-form-item label="事件类型">
                                        <a-select v-model:value="field.eventType" style="width: 100%;"
                                            :readonly="isReadOnly">
                                            <a-select-option value="0">打开页面</a-select-option>
                                            <a-select-option value="1">点击事件</a-select-option>
                                            <a-select-option value="2">输入参数</a-select-option>
                                            <a-select-option value="3">点击键盘</a-select-option>
                                            <a-select-option value="19">获取链接</a-select-option>
                                            <a-select-option value="9">获取文本</a-select-option>
                                            <a-select-option value="7">获取Cookie</a-select-option>
                                            <a-select-option value="8">获取Header</a-select-option>
                                            <a-select-option value="4">截幕</a-select-option>
                                            <a-select-option value="5">等待元素</a-select-option>
                                            <a-select-option value="6">滚动页面</a-select-option>
                                            <a-select-option value="10">获取图片</a-select-option>
                                            <a-select-option value="11">获取音频</a-select-option>
                                            <a-select-option value="12">获取视频</a-select-option>
                                        </a-select>
                                    </a-form-item>
                                </a-col>
                                <a-col :span="12">
                                    <a-form-item label="输入参数">
                                        <a-input v-model:value="field.inputValue" style="width: 100%;"
                                            :readonly="isReadOnly" />
                                    </a-form-item>
                                </a-col>
                            </a-row>

                            <a-row>
                                <a-col :span="12">
                                    <a-form-item label="等待时间">
                                        <a-input-number v-model:value="field.sleepTime" style="width: 100%;"
                                            :readonly="isReadOnly" />
                                    </a-form-item>
                                </a-col>
                                <a-col :span="12">
                                    <a-form-item label="操作">
                                        <a-button-group>
                                            <a-button @click="moveRow('up', index)" :disabled="index === 0">
                                                <UpOutlined />
                                            </a-button>
                                            <a-button @click="moveRow('down', index)"
                                                :disabled="index === fields.length - 1">
                                                <DownOutlined />
                                            </a-button>
                                            <a-button @click="deleteField(index)">
                                                <DeleteOutlined />
                                            </a-button>
                                        </a-button-group>
                                    </a-form-item>
                                </a-col>
                            </a-row>
                        </div>

                        <a-row>
                            <a-col>
                                <a-button type="link" @click="addField">
                                    <PlusOutlined /> 添加行
                                </a-button>
                            </a-col>
                        </a-row>
                    </a-tab-pane>

                    <a-tab-pane key="text" tab="文本模式" style="padding: 10px 0 0 0;">
                        <a-form-item label="脚本内容" name="content">
                            <a-textarea v-model:value="model.content" :rows="15" style="width: 100%;"
                                :readonly="this.$store.forms.action == 'detail'" />
                        </a-form-item>
                    </a-tab-pane>
                </a-tabs>
            </a-form>

            <template #actions v-if="this.$store.forms.action != 'detail'">
                <a-space :size="16">
                    <a-button @click="back">返回</a-button>
                    <a-button type="primary" @click="submitForm">确认</a-button>
                </a-space>
            </template>
        </a-card>
    </div>
</template>

<script>
import { Add, Edit, Get } from '../../../wailsjs/go/biz/TaskBiz.js';

export default {
    mounted() {
        this.taskId = this.$route.query.id;
        this.model.system = this.$route.query.system;
        if (this.taskId) {
            Get(this.taskId).then(res => {
                if (res.code == 200) {
                    this.model = res.data || {};
                    this.getFields();
                }
            })
        }
    },
    data() {
        return {
            taskId: "",
            activeKey: "config",
            model: {},
            fields: [],
            options: [{
                "value": "/div[contains(normalize-space(text()), 'xxx')]",
                "label": "查找去除空格后包含指定文本的div元素"
            }, {
                "value": "/div[contains(@class, 'xxx')]//span[text()='yyy']",
                "label": "查找指定类名的div元素中，文本为'yyy'的span元素"
            }, {
                "value": "/h2[text()='标题']/following-sibling::p",
                "label": "查找文本为'标题'的h2元素后的p兄弟元素"
            }, {
                "value": "/*[@id='app']/div/div[1]",
                "label": "查找id为'app'元素的第一个子级div"
            }, {
                "value": "/span[text()='yyy']/parent::div",
                "label": "查找文本为'yyy'的span元素的父div元素"
            }, {
                "value": "/div[@id='xxx']//span",
                "label": "查找id为'xxx'的div元素中的所有span元素"
            }, {
                "value": "/input[@name='username' and @type='password']",
                "label": "查找name为'username'且type为'password'的input元素"
            }, {
                "value": "/ul/li[last()]",
                "label": "查找ul中最后一个li元素"
            }, {
                "value": "/a[@class='btn' and contains(@href, 'submit')]",
                "label": "查找class为'btn'且href包含'submit'的a元素"
            }, {
                "value": "/div[@id and @class]",
                "label": "查找同时具有id和class属性的div元素"
            }, {
                "value": "/div[.//span[text()='xxx']]",
                "label": "查找包含文本为'xxx'的span元素的div"
            }],
            vRules: {
                system: [{ required: true, message: "请选择分组类型" }],
                name: [{ required: true, message: "请输入分组名称" }],
            },
        }
    },
    computed: {
        isReadOnly() {
            return this.$store.forms.action === 'detail';
        }
    },
    methods: {
        getFields() {
            this.fields = JSON.parse(this.model.content);
        },
        tabChange(data) {
            if (data == "config") {
                this.fields = JSON.parse(this.model.content);
            } else {
                this.model.content = JSON.stringify(this.fields, null, 2);
            }
        },
        moveRow(direction, index) {
            if (direction === 'up' && index > 0) {
                [this.fields[index], this.fields[index - 1]] = [this.fields[index - 1], this.fields[index]];
            } else if (direction === 'down' && index < this.fields.length - 1) {
                [this.fields[index], this.fields[index + 1]] = [this.fields[index + 1], this.fields[index]];
            }
        },
        addField() {
            if (this.fields.length > 0) {
                let field = this.fields[this.fields.length - 1];
                this.fields.push({
                    groupName: field.groupName,
                    eventType: field.eventType,
                    name: field.name,
                    inputValue: "",
                    selector: "",
                    sleepTime: 0,
                });
            } else {
                this.fields.push({
                    groupName: "",
                    eventType: "1",
                    name: "",
                    inputValue: "",
                    selector: "",
                    sleepTime: 0,
                });
            }
        },
        deleteField(index) {
            if (index !== -1) {
                this.fields.splice(index, 1);
            }
        },
        submitForm() {
            if (this.activeKey == "config") {
                this.model.content = JSON.stringify(this.fields);
            } else {
                this.fields = JSON.parse(this.model.content);
            }
            this.model.taskId = this.taskId;
            if (this.$store.forms.action == "add") {
                this.model.id = "";
                Add(this.model).then(res => {
                    if (res.code == 200) {
                        this.$message.success(res.msg);
                        this.back();
                    }
                })
            } else {
                Edit(this.model).then(res => {
                    if (res.code == 200) {
                        this.$message.success(res.msg);
                        this.back();
                    }
                })
            }
        },
        back() {
            this.$router.back();
        }
    }
}
</script>

<style scoped>
.table-header {
    font-weight: bold;
}
</style>
