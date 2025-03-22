<template>
  <a-config-provider :locale="locale" :theme="{
    token: {
      borderRadius: 1,
      sizeStep: 4,
      sizeUnit: 4,
      wireframe: true,
      margin: '8px',
      fontSize: 14,
      colorPrimary: '#007bff'
    },
    components: {

    },
  }">
    <a-layout>
      <a-layout-sider :style="{ minHeight: '100vh' }" :trigger="null" :collapsed="true" :collapsible="true"
        :collapsedWidth="40" theme="light">
        <a-menu :style="{ minHeight: '100vh' }" mode="inline" @click="menuClick" theme="light"
        :selectedKeys="[selectedKey]" v-show="$store.isAbout">
          <a-menu-item key="/taskList">
            <HomeOutlined />&nbsp;&nbsp;操作配置
          </a-menu-item>
          <a-menu-item key="/taskPlan">
            <FieldTimeOutlined />&nbsp;&nbsp;执行规则
          </a-menu-item>
          <a-menu-item key="/taskExec">
            <OrderedListOutlined />&nbsp;&nbsp;开始执行
          </a-menu-item>
          <a-menu-item key="/data">
            <DatabaseOutlined />&nbsp;&nbsp;数据中心
          </a-menu-item>
          <a-menu-item key="/config">
            <SettingOutlined />&nbsp;&nbsp;设置
          </a-menu-item>
          <a-menu-item key="/about">
            <InfoCircleOutlined />&nbsp;&nbsp;关于
          </a-menu-item>
        </a-menu>
      </a-layout-sider>
      <a-layout>
        <a-layout-content style="background-color: white;padding: 10px;">
          <router-view></router-view>
        </a-layout-content>
      </a-layout>
    </a-layout>
  </a-config-provider>
</template>

<script>
import { InitBrowser } from '../wailsjs/go/biz/TaskBiz.js';
import { GetConfig } from '../wailsjs/go/main/App';

export default {
  mounted() {
    this.$router.push(this.selectedKey);
    GetConfig().then((res) => {
      console.log(res)
      if (res.code != 200) {
        this.$store.setIsAbout(false)
        this.$router.push("/about");
      } else if (res.data.legalStatement == '1') {
        this.$store.setIsAbout(true);
        InitBrowser()
      }
    })
  },
  watch: {
    '$route': function (newRoute) {
      if (newRoute.path.split('/').length - 1 == 1) {
        this.selectedKey = newRoute.path;
      }
    }
  },
  data() {
    return {
      selectedKey: "/taskList",
      menus: [],
    };
  },
  methods: {
    menuClick(data) {
      if (!data.key) {
        return;
      }
      this.$store.setAction("list");
      this.$router.push(data.key)
    }
  },
};
</script>
<style scoped>
.iframeStyle {
  width: 100%;
  border: none;
  height: 89vh;
}
</style>