<template>
    <div>
        <a-page-header style="border: 1px solid rgb(235, 237, 240); padding: 8px; margin: 0;" title="规则表单"
            @back="back" />
        <a-card>
            <a-form :model="model" :rules="vRules" :label-col="{ style: { width: '100px' } }">
                <a-form-item label="规则分组" name="execRemark">
                    <a-input v-model:value="model.execRemark" :readonly="this.$store.forms.action == 'detail'" />
                </a-form-item>
                <a-form-item label="规则名称" name="name">
                    <a-input v-model:value="model.name" :readonly="this.$store.forms.action == 'detail'" />
                </a-form-item>
                <a-form-item label="定时时间(cron)" name="execCron">
                    <a-input v-model:value="model.execCron" placeholder="cron表达式"
                        :readonly="this.$store.forms.action == 'detail'" />
                </a-form-item>
            </a-form>
            <div>
                <a-button type="link" @click="addField">
                    <PlusOutlined />
                </a-button>
                <a-divider style="margin: 10px;" />
                <a-table :columns="columns" :data-source="fields" row-key="index" bordered :pagination="false"
                    :scroll="{ x: 'max-content' }">
                    <template #bodyCell="{ column, record, index }">
                        <a-td :column="column" :record="record" :index="index">
                            <template v-if="column.key === 'taskId'">
                                <a-select v-model:value="record.taskId" style="min-width: 320px" :disabled="isReadOnly">
                                    <a-select-option v-for="task in taskList" :key="task.value" :value="task.value">{{
                                        task.label }}</a-select-option>
                                </a-select>
                            </template>
                            <template v-else-if="column.key === 'forSize'">
                                <a-input-number v-model:value="record.forSize" :readonly="isReadOnly" />
                            </template>
                            <template v-else-if="column.key === 'start'">
                                <a-input-number v-model:value="record.start" :readonly="isReadOnly" />
                            </template>
                            <template v-else-if="column.key === 'end'">
                                <a-input-number v-model:value="record.end" :readonly="isReadOnly" />
                            </template>
                            <template v-else-if="column.key === 'action'">
                                <a-button-group>
                                    <a-button @click="moveRow('up', index)" :disabled="index === 0">
                                        <UpOutlined />
                                    </a-button>
                                    <a-button @click="moveRow('down', index)" :disabled="index === fields.length - 1">
                                        <DownOutlined />
                                    </a-button>
                                    <a-button @click="deleteField(index)">
                                        <DeleteOutlined />
                                    </a-button>
                                </a-button-group>
                            </template>
                        </a-td>
                    </template>
                </a-table>
            </div>
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
import { Get, Add, Edit } from '../../../wailsjs/go/biz/PlanBiz.js';
import { GetList } from '../../../wailsjs/go/biz/TaskBiz.js';

export default {
    mounted() {
        this.getTaskList();
        this.id = this.$route.query.id;
        this.model.execRemark = this.$route.query.execRemark;
        if (this.id) {
            Get(this.id).then(res => {
                if (res.code === 200) {
                    this.model = res.data || {};
                    this.getFields();
                }
            });
        }
    },
    data() {
        return {
            id: "",
            taskList: [],
            model: {},
            fields: [],
            vRules: {
                name: [{ required: true, message: "请输入规则名称" }],
                execRemark: [{ required: true, message: "请输入规则分组" }],
            },
            columns: [
                { title: '执行操作', dataIndex: 'taskId', key: 'taskId' },
                { title: '重复次数', dataIndex: 'forSize', key: 'forSize' },
                { title: '开始位置', dataIndex: 'start', key: 'start' },
                { title: '结束位置', dataIndex: 'end', key: 'end' },
                { title: '操作', dataIndex: 'action', key: 'action', scopedSlots: { customRender: 'action' } },
            ],
        };
    },
    computed: {
        isReadOnly() {
            return this.$store.forms.action === 'detail';
        },
    },
    methods: {
        getTaskList() {
            GetList().then(res => {
                if (res.code === 200) {
                    this.taskList = res.data;
                } else {
                    this.$message.error(res.msg);
                }
            });
        },
        getFields() {
            this.fields = JSON.parse(this.model.content);
        },
        addField() {
            this.fields.push({
                taskId: "",
                forSize: 1,
                start: 0,
                end: -1,
            });
        },
        moveRow(direction, index) {
            if (direction === 'up' && index > 0) {
                [this.fields[index], this.fields[index - 1]] = [this.fields[index - 1], this.fields[index]];
            } else if (direction === 'down' && index < this.fields.length - 1) {
                [this.fields[index], this.fields[index + 1]] = [this.fields[index + 1], this.fields[index]];
            }
        },
        deleteField(index) {
            if (index !== -1) {
                this.fields.splice(index, 1);
            }
        },
        submitForm() {
            this.model.content = JSON.stringify(this.fields);
            if (this.$store.forms.action === "add") {
                this.model.id = "";
                this.model.status = "0";
                Add(this.model).then(res => {
                    if (res.code === 200) {
                        this.$message.success("提交成功");
                        this.$router.go(-1);
                    } else {
                        this.$message.warn(res.msg);
                    }
                });
            } else if (this.$store.forms.action === "edit") {
                Edit(this.model).then(res => {
                    if (res.code === 200) {
                        this.$message.success("提交成功");
                        this.$router.go(-1);
                    } else {
                        this.$message.warn(res.msg);
                    }
                });
            }
        },
        back() {
            this.$router.go(-1);
        },
    },
};
</script>

<style scoped>
.table-header {
    font-weight: bold;
}
</style>
