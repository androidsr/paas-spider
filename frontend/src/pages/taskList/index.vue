<template>
    <div>
        <div>
            <a-flex gap="middle" horizontal>
                <a-button @click="add">
                    <PlusOutlined />新增
                </a-button>
                <a-select placeholder="所属系统" allowClear v-model:value="query.system" :options="systemList"
                    style="width: 400px;"></a-select>
                <a-input placeholder="任务名称" v-model:value="query.name"></a-input>
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
            pages: { current: 1, size: 10 },
            systemList: [],
            records: [],
            record: {},
            columns: [
                { title: "所属系统", dataIndex: "system", width: 200 },
                { title: "任务名称", dataIndex: "name", width: 400 },
                { title: "执行时间", dataIndex: "execTime" },
                { title: '操作', key: 'action', fixed: 'right', width: 180 }
            ]
        };
    },
    methods: {
        async getSystemList() {
            let res = await GetSystemList();
            let data = res.data;
            if (res.code == 200) {
                this.systemList = data;
            }
            this.queryData();
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
            this.$store.setQueryData("taskList", this.query);
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