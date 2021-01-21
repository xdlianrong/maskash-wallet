<template>
    <el-menu mode="horizontal" class="emenu" background-color="rgb(5, 126, 214)" active-text-color="white" text-color="rgb(89, 162, 214)">
        <el-menu-item index="1">购币</el-menu-item>
        <el-menu-item index="2">转账</el-menu-item>
        <el-menu-item index="3">收款</el-menu-item>
        <el-menu-item index="4">我的</el-menu-item>
    </el-menu>
</template>
<style>
    .el-menu-item:hover {
        background-color: rgb(2, 100, 170) !important;
        text-align: center;
    }
    .el-menu--horizontal>.el-menu-item {
        width: 25%;
        text-align: center;
    }
    .emenu {
        display: flex;
    }
</style>
