<template>
    <div style="height: 100%;">
        <navmenu></navmenu>  
        <el-row type="flex" justify="center" id="o" style="top: 35%">
            <el-col :xs="20" :sm="15" :md="12" :lg="8" :xl="7">  
            <div>
                <p>转账金额:</p>
                <el-input maxlength="10" v-model="money" ></el-input>
                <mybutton :buttonMsg="transfer" @click.native="transferm"></mybutton>
            </div>
            </el-col>
        </el-row>
    </div>
</template>
<script>
import navmenu from '../components/Navmenu'
import mybutton from '../components/Mybutton'
export default {
    components: {
        navmenu,
        mybutton
    },
    data() {
        return {
            transfer: '发起转账',
            buy: '购买',
            recv: '收款',
            money: ''
        }
    },
    methods: {
        transferm() {
            console.log("我要转账啦");
        },
        buym() {

        },
        recm() {

        }
    }
}
</script>
<style>
    .ma {
        
    }

    
</style>