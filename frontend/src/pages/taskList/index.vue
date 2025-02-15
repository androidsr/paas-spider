<template>
    <div>
        <div>
            <a-flex gap="middle" horizontal>
                <a-button type="primary" @click="add">
                    <PlusOutlined />创建操作
                </a-button>
                <a-input placeholder="操作分类" v-model:value="query.system"></a-input>
                <a-input placeholder="任务名称" v-model:value="query.name"></a-input>
                <a-button @click="queryData">
                    <SearchOutlined />查询
                </a-button>
            </a-flex>
        </div>

        <a-tabs v-model:activeKey="activeKey" @change="changeTabs" style="margin-top: 10px;">
            <a-tab-pane v-for="v in systemList" :key="v" :tab="v"></a-tab-pane>
        </a-tabs>
        <div class="scrollable-list">
            <a-list bordered :data-source="records">
                <template #renderItem="{ item }">
                    <a-list-item style="width: 100%;">
                        <template #actions>
                            <a key="list-loadmore-more" @click="copy(item)">复制</a>
                            <a key="list-loadmore-more" @click="detail(item)">查看</a>
                            <a key="list-loadmore-more" @click="edit(item)">编辑</a>
                            <a key="list-loadmore-more" @click="del(item)">删除</a>
                        </template>
                        <a-list-item-meta :description="item.system">
                            <template #title>
                                <strong>{{ item.name }}</strong>
                            </template>
                        </a-list-item-meta>
                        <div>{{ item.execTime }}</div>
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
import { Delete, GetSystemList, Page } from '../../../wailsjs/go/biz/TaskBiz.js';


export default {
    mounted() {
        this.query = this.$store.queryData["taskList"] || {};
        this.getSystemList();
    },
    data() {
        return {
            query: {},
            isMore: false,
            activeKey: "",
            pages: { current: 1, size: 10 },
            systemList: [],
            records: [],
            record: {},
            columns: [
                { title: "所属系统", dataIndex: "system" },
                { title: "任务名称", dataIndex: "name" },
                { title: "执行时间", dataIndex: "updateTime" },
                { title: '操作', key: 'action', fixed: 'right', width: 180 }
            ]
        };
    },
    methods: {
        changeTabs(data) {
            this.query.system = this.activeKey;
            this.queryData();
        },
        async getSystemList() {
            let res = await GetSystemList();
            let data = res.data;
            if (res.code == 200) {
                this.systemList = data;
            }
            if (!this.query.system) {
                this.activeKey = data[0];
                this.query.system = this.activeKey;
            } else {
                this.activeKey = this.query.system;
            }
            this.queryData();
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
            this.$store.setQueryData("taskList", this.query);
            this.pages.current = 1;
            this.records = [];
            this.loadPage();
        },
        morePage() {
            this.pages.current++;
            this.loadPage();
        },
        add(data) {
            this.$store.setAction("add");
            this.$router.push({
                path: "/taskList/form",
                query: {
                    id: data.id,
                    system: this.activeKey,
                },
            });
        },
        copy(data) {
            this.$store.setAction("add");
            this.$router.push({
                path: "/taskList/form",
                query: {
                    id: data.id,
                },
            });
        },
        edit(data) {
            this.$store.setAction("edit");
            this.$router.push({
                path: "/taskList/form",
                query: {
                    id: data.id,
                },
            });
        },
        detail(data) {
            this.$store.setAction("detail");
            this.$router.push({
                path: "/taskList/form",
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