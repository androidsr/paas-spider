<template>
    <div style="padding: 10px;">
        <div>
            <a-flex>
                <a-form-item label="执行规则">
                    <a-select v-model:value="planId" style="width: 680px;" @change="resetHander">
                        <a-select-option v-for="(item, i) in planList" :value="item.value">{{ item.label
                            }}</a-select-option>
                    </a-select>
                </a-form-item>
            </a-flex>
            <a-flex gap="middle" horizontal>
                <a-form-item label="重复次数">
                    <a-input-number placeholder="重复次数" v-model:value="forSize"></a-input-number>
                </a-form-item>
                <a-form-item label="间隔">
                    <a-input-number placeholder="间隔时间(秒)" v-model:value="intervalTime"></a-input-number>
                </a-form-item>
                <a-form-item label="开始位置">
                    <a-input-number placeholder="开始位置" v-model:value="start"></a-input-number>
                </a-form-item>
                <a-form-item label="结束位置">
                    <a-input-number placeholder="结束位置" v-model:value="end"></a-input-number>
                </a-form-item>
                <a-button @click="execHandler" type="primary">
                    <PlaySquareOutlined v-if="!isRun" /> <sync-outlined spin v-if="isRun" />
                    开始
                </a-button>
                <a-button @click="resetHander">
                    重置
                </a-button>
            </a-flex>
            <a-divider style="height: 1px; margin: 0px;padding: 0px;" />
        </div>
        <a-steps style="padding: 20px;" direction="vertical" v-model:current="current" :items="items"></a-steps>
    </div>
</template>
<script>
import { message } from 'ant-design-vue';
import { Get, GetList } from '../../../wailsjs/go/biz/PlanBiz.js';
import { ExecStep, Get as GetTask } from '../../../wailsjs/go/biz/TaskBiz.js';

export default {
    mounted() {
        this.planId = this.$store.planId;
        this.getPlanList();
    },
    data() {
        return {
            isRun: false,
            forSize: 1,
            start: 0,
            end: -1,
            current: 0,
            planList: [],
            planId: "",
            planInfo: {},
            steps: [],
            items: [],
            forIndex: 0,
            taskIndex: 0,
            tIndex: 0,
            intervalTime: 1,
        }
    },
    methods: {
        getPlanList() {
            GetList().then(res => {
                if (res.code == 200) {
                    this.planList = res.data;
                } else {
                    message.error(res.msg);
                }
            });
        },
        resetHander() {
            this.forIndex = 0;
            this.taskIndex = 0;
            this.tIndex = 0;
        },
        async execHandler() {
            if (!this.planId) {
                message.error('请选择预案');
                return;
            }
            if (!this.isRun) {
                this.isRun = true;
            } else {
                this.isRun = false;
                return;
            }
            this.$store.setPlanId(this.planId);
            if (this.forIndex == 0) {
                this.current = 0;
                this.items = [];
            }

            let m = this;
            let res = await Get(this.planId);
            if (res.code != 200) {
                message.error(res.msg);
                return;
            }
            m.planInfo = res.data;
            let taskList = JSON.parse(m.planInfo.content);
            for (this.forIndex; this.forIndex < this.forSize; this.forIndex++) {
                if (!this.isRun) {
                    break;
                }
                let endSize = this.end == -1 ? taskList.length : this.end;
                if (this.tIndex == 0) {
                    this.tIndex = this.start;
                }
                for (this.tIndex; this.tIndex < endSize; this.tIndex++) {
                    let result = await this.execTask(taskList[this.tIndex]);
                    if (!result) {
                        this.isRun = false;
                        return false;
                    }
                }
            }
            this.forIndex = 0;
            this.isRun = false;
            setTimeout(() => {
            }, intervalTime * 1000);
        },
        async execTask(task) {
            let res = await GetTask(task.taskId);
            if (res.code != 200) {
                message.error(res.msg);
                return;
            }
            this.steps = JSON.parse(res.data.content);
            let endSize = task.end == -1 ? this.steps.length : task.end;
        
            for (this.taskIndex; this.taskIndex < task.forSize; this.taskIndex++) {
                if (!this.isRun) {
                    break;
                }
                
                for (var stepIndex = task.start; stepIndex < endSize; stepIndex++) {
                    let step = this.steps[stepIndex];
                    step.taskId = task.taskId;
                    this.items.push({ disabled: true, title: step.inputValue, subTitle: step.name, status: "wait", description: "" });
                    this.current = this.items.length - 1;
                    let res = await ExecStep(step);
                    let description = "";
                    if (res.code == 200) {
                        this.items[this.current].status = "finish";
                        if (!!res.data && res.data.length > 0) {
                            description = JSON.stringify(res.data);
                        } else {
                            description = res.msg;
                        }
                        this.items[this.current].description = description;
                    } else {
                        this.items[this.current].status = "error";
                        this.items[this.current].description = res.msg;
                        return false;
                    }
                }
            }
            this.taskIndex = 0;
            return true;
        }
    }
}
</script>