<template>
    <div>
        <div>
            <a-flex gap="middle" horizontal>
                <a-button @click="add">
                    <PlusOutlined />新增
                </a-button>
                <a-select placeholder="所属系统" v-model:value="query.execRemark" allowClear :options="groupList" style="width: 400px;"></a-select>
                <a-input placeholder="规则名称" v-model:value="query.name"></a-input>
                <a-button @click="queryData">
                    <SearchOutlined />查询
                </a-button>
            </a-flex>
        </div>
        <a-divider style="margin-top: 10px;margin-bottom: 10px;" />
        <div>
            <a-table ref="table" :data-source="records" :columns="columns" @change="pageClick" :pagination="pages"
                row-key="id" :scroll="{ x: 300, y: 500 }" :customRow="customRow">
                <template #bodyCell="{ column, record }">
                    <template v-if="column.key === 'action'">
                        <a-space warp>
                            <a key="list-loadmore-more" @click="copy(record)">复制</a>
                            <a key="list-loadmore-more" @click="detail(record)">查看</a>
                            <a key="list-loadmore-more" @click="edit(record)">编辑</a>
                            <a key="list-loadmore-more" @click="del(record)">删除</a>
                        </a-space>
                    </template>
                </template>
            </a-table>
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
                { title: "所属系统", dataIndex: "execRemark", width: 200 },
                { title: "规则名称", dataIndex: "name", width: 400 },
                {
                    title: "定时状态", dataIndex: "status",
                    customRender: (text) => {
                        return text === '1' ? '启用' : '未启用';
                    },
                },
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
                    this.pages.total = data.total;
                    this.records = data.rows;
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
        pageClick(pagination) {
            this.pages = pagination;
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