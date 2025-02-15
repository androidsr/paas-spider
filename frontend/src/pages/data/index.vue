<template>
    <div>
        <div>
            <a-flex gap="middle" horizontal>
                <a-button @click="export">
                    <CopyOutlined />导出
                </a-button>
                <a-select v-model:value="query.taskId" style="width: 600px;">
                    <a-select-option v-for="(item, i) in taskList" :value="item.value">
                        {{ item.label }}
                    </a-select-option>
                </a-select>
                <a-input placeholder="所属网站" v-model:value="query.system"></a-input>
                <a-input placeholder="操作名称" v-model:value="query.name"></a-input>
                <a-button @click="queryData">
                    <SearchOutlined />查询
                </a-button>
            </a-flex>
        </div>
        <a-divider style="margin: 10px;" />
        <div class="scrollable-list">
            <a-list bordered :data-source="records">
                <template #renderItem="{ item }">
                    <a-list-item style="width: 100%;">
                        <template #actions>
                            <a key="list-loadmore-more" @click="detail(item)">查看</a>
                            <a key="list-loadmore-more" @click="del(item)">删除</a>
                        </template>
                        <a-list-item-meta>
                            <template #title>
                                <strong>{{ item.name }}</strong>
                            </template>
                            <template #description>
                                <div>{{ item.execTime }}</div>
                                <div>
                                    <a-typography-paragraph :ellipsis="true" :content="item.outputValue" />
                                </div>
                            </template>
                        </a-list-item-meta>
                        <div>{{ item.system }}</div>
                    </a-list-item>
                </template>
                <template #footer v-if="isMore">
                    <a-flex justify="center" align="center"><a @click="morePage">加载更多</a></a-flex>
                </template>
            </a-list>
        </div>
    </div>
</template>
<script>
import { message, Modal } from 'ant-design-vue';
import { Delete, Page } from '../../../wailsjs/go/biz/RecordBiz.js';
import { GetList } from '../../../wailsjs/go/biz/TaskBiz.js';


export default {
    mounted() {
        this.getTaskList();
        this.loadPage();
    },
    data() {
        return {
            query: {},
            taskList: [],
            isMore: false,
            pages: { current: 1, size: 10 },
            records: [],
            columns: [
                { title: "所属系统", dataIndex: "system" },
                { title: "任务名称", dataIndex: "name" },
                { title: "输入", dataIndex: "inputValue" },
                { title: "输出", dataIndex: "outputValue" },
                { title: "执行时间", dataIndex: "execTime" },
                { title: '操作', key: 'action', fixed: 'right', width: 180 }
            ]
        };
    },
    methods: {
        getTaskList() {
            GetList().then(res => {
                if (res.code == 200) {
                    this.taskList = res.data;
                } else {
                    message.error(res.msg);
                }
            });
        },
        loadPage() {
            Page({ page: this.pages, ...this.query }).then(res => {
                let data = res.data;
                if (res.code == 200) {
                    if (data.total > data.current * data.size) {
                        this.isMore = true;
                    } else {
                        this.isMore = false;
                    }
                    this.records.push(...data.rows)
                } else {
                    this.$message.error(res.msg)
                }
            })
        },
        queryData() {
            this.pages.current = 1;
            this.records = [];
            this.loadPage();
        },
        morePage() {
            this.pages.current++;
            this.loadPage();
        },
        export() {

        },
        detail(data) {
            this.$store.setAction("detail");
            this.$router.push({
                path: "/data/form",
                query: {
                    id: data.id,
                },
            });
        },
        del(data) {
            let m = this;
            Modal.confirm({
                title: '确定删除？',
                okText: '确定',
                cancelText: '取消',
                onOk() {
                    Delete(data.id).then(res => {
                        if (res.code == 200) {
                            message.success('删除成功!');
                            m.queryData();
                        } else {
                            message.error(res.msg);
                        }
                    });
                },
                onCancel() {
                },
            });
        },
    },
};
</script>

<style scoped>
.scrollable-list {
    max-height: 86vh;
    overflow-y: auto;
}
</style>