<template>
    <div>
        <div>
            <a-flex gap="middle" horizontal>
                <a-button type="primary" @click="add">
                    <PlusOutlined />新建规则
                </a-button>
                <a-input placeholder="规则名称" v-model:value="query.name"></a-input>
                <a-button @click="queryData">
                    <SearchOutlined />查询
                </a-button>
            </a-flex>
        </div>
        <a-tabs v-model:activeKey="activeKey" @change="changeTabs" style="margin-top: 10px;">
            <a-tab-pane v-for="v in groupList" :key="v" :tab="v"></a-tab-pane>
        </a-tabs>
        <div class="scrollable-list">
            <a-list bordered :data-source="records">
                <template #renderItem="{ item }">
                    <a-list-item style="width: 100%;">
                        <template #actions>
                            <a key="list-loadmore-more" @click="statusChange(item)">{{ item.status == '1' ? '停止' : '启动'
                                }}</a>
                            <a key="list-loadmore-more" @click="copy(item)">复制</a>
                            <a key="list-loadmore-more" @click="detail(item)">查看</a>
                            <a key="list-loadmore-more" @click="edit(item)">编辑</a>
                            <a key="list-loadmore-more" @click="del(item)">删除</a>
                        </template>
                        <a-list-item-meta>
                            <template #title>
                                <strong>{{ item.name }}</strong>
                            </template>
                            <template #description>
                                <div>
                                    <a-typography-paragraph :ellipsis="true" :content="item.execRemark" />
                                </div>
                                <div>{{ item.execCron }}</div>
                            </template>
                        </a-list-item-meta>
                        <div>{{ item.status == '1' ? '启动' : '未启动' }}</div>

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
import { Page, Delete, Edit, GetGroupList } from '../../../wailsjs/go/biz/PlanBiz.js';
import { message, Modal } from 'ant-design-vue';


export default {
    mounted() {
        this.query = this.$store.queryData["taskPlan"] || {};
        this.getGroupList();
    },
    data() {
        return {
            query: {},
            isMore: false,
            activeKey: "",
            groupList: [],
            pages: { current: 1, size: 10 },
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
            this.query.execRemark = this.activeKey;
            this.queryData();
        },
        getGroupList() {
            GetGroupList().then(res => {
                let data = res.data;
                if (res.code == 200) {
                    if (data.length > 0) {
                        this.groupList = data;
                        if (!this.query.execRemark) {
                            this.activeKey = data[0];
                            this.query.execRemark = this.activeKey;
                        } else {
                            this.activeKey = this.query.execRemark;
                        }
                        this.queryData();
                    }
                } else {
                    this.$message.error(res.msg)
                }
            })
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
            this.$store.setQueryData("taskPlan", this.query);
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
                path: "/taskPlan/form",
                query: {
                    id: data.id,
                    execRemark: this.activeKey,
                },
            });
        },
        copy(data) {
            this.$store.setAction("add");
            this.$router.push({
                path: "/taskPlan/form",
                query: {
                    id: data.id,
                },
            });
        },
        edit(data) {
            this.$store.setAction("edit");
            this.$router.push({
                path: "/taskPlan/form",
                query: {
                    id: data.id,
                },
            });
        },
        detail(data) {
            this.$store.setAction("detail");
            this.$router.push({
                path: "/taskPlan/form",
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
        statusChange(data) {
            if (data.status == "0") {
                data.status = "1";
            } else {
                data.status = "0";
            }
            Edit(data).then(res => {
                if (res.code == 200) {
                    this.$message.success("成功");
                    this.queryData();
                } else {
                    this.$message.warn(res.msg);
                }
            })
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