<template>
<!-- 其实可以用 Tab 实现 -->
    <div>
        <img src="../assets/m.png" class="s">
        <el-card>
            
            <span style="font-size: 0.8rem;">余额</span>
            <span style="font-size: 1.5rem; ">&emsp;{{nowm}}</span>
        </el-card>
        <el-menu mode="horizontal" class="emenu" background-color="rgb(5, 126, 214)" active-text-color="white" text-color="rgb(89, 162, 214)">
            <el-menu-item index="1" @click="con(1)" ref="no">购币</el-menu-item>
            <el-menu-item index="2" @click="con(2)">转账</el-menu-item>
            <el-menu-item index="3" @click="con(3)">收款</el-menu-item>
            <el-menu-item index="4" @click="con(4)">历史</el-menu-item>
            <el-menu-item index="5" @click="con(5)">我的</el-menu-item>
        </el-menu>
    </div>
</template>
<script>
export default {

    data() {
        return {
            nowm: 0
        }
    },
    methods: {
        con(index) {
            console.log(index);
            // 修改父组件的值
            this.$emit('changecmp', index);    
        },
        changenm(vm) {
            this.nowm = vm;
        }
    },
    mounted: function () {
        // 初始化点击，不能放在 created 里。因为还没创建出来是拿不到的
        this.$refs.no.$el.click();
    }
}
</script>
<style>
/* 开发者模式拿到的 css 类名 */
    .el-menu-item:hover, .el-menu-item:focus {
        background-color: rgb(2, 100, 170) !important;
        text-align: center;
    }
    .el-menu--horizontal>.el-menu-item {
        width: 20%;
        text-align: center;
    }
    .emenu {
        display: flex;
    }
    .el-menu-item.is-active {
        background-color: rgb(2, 100, 170) !important;
    }
    .el-card {
        border: none !important;
        background-color:#1656b9 !important;
        color: white !important;
        text-align: center;
        /* display: flex;
        justify-content: center; */
    }
    .el-card, .el-message {
        border-radius: 0 !important;
    }
    .s {
        position: absolute;
        height: 4.5rem;
    }
    .el-card__body {
        padding: 1.5rem !important;
    }
</style>
