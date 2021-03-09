<template>
    <div style="height: 100%;">
        <navmenu @changecmp="changecmps" ref="nav"></navmenu>  
        <el-row type="flex" justify="center" id="o">
            <el-col :xs="20" :sm="15" :md="12" :lg="8" :xl="7">  
            <div v-if="cmp == 1">
                <p>购买金额:</p>
                <el-input maxlength="10" v-model="money" ></el-input>
                <mybutton :buttonMsg="buy" @click.native="buym"></mybutton>
            </div>

            <div v-if="cmp == 2" style="top: 10%;">
                <p>收款方公钥:</p>
                <el-input maxlength="10" v-model="recPub" ></el-input>
                <p>转账金额:</p>
                <el-input maxlength="10" v-model="transmoney" ></el-input>
                <p>代币承诺</p>
                <el-input maxlength="10" v-model="moneyProm" ></el-input>
                <!-- 上面这些够了，可以返回东西了 -->
                <mybutton :buttonMsg="transfer" @click.native="transferm"></mybutton>
            </div>

            <div v-if="cmp == 3">
                <p>收款方公钥</p>
                <el-input maxlength="10" v-model="money" ></el-input>
                <mybutton :buttonMsg="recv" @click.native="recm"></mybutton>
            </div>

            <div v-if="cmp == 4">
                
            </div>

            <div v-if="cmp == 5">
                <mybutton :buttonMsg="showImfo" @click.native="showImfof"></mybutton>
                <mybutton :buttonMsg="signout" @click.native="signoutf"></mybutton>
            </div>
            </el-col>
        </el-row>
    </div>
</template>
<script>
import navmenu from '../components/Navmenu'
import mybutton from '../components/Mybutton'
var account;
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
            signout: '登出',
            showImfo: '显示账户信息',
            money: '',
            cmp: '1', // 用来改变显示的组件
            account,
            recPub: '',
            transmoney: '',
            moneyProm: '',
        }
    },
    created: function () {
        account = this.$route.query.account;
        if (account == undefined) {
            this.$message.error({
                message: '请登录账户',
                duration: 1400
            }); 
            setTimeout(() => {
                this.$router.push({
                    path: '/',
                    name: 'Main',
                })
            }, 1500);
            
        }
    },
    methods: {
        getPri() {
            var pri = JSON.parse(window.localStorage.getItem(account)).bi;
            console.log(pri);
            return pri;
        },
        transferm() {
            console.log("我要转账");
            // var pri = this.getPri();
            // this.axios.post('http://localhost:1998/wallet/register', {
            //         pri: pri,
            //     }).then((response)=>{
            //     // 一堆赋值，我也忘了要存啥了
            //     console.log(response);
            // })
        },
        buym() {
            console.log("我要购币");
             this.getPri();
            // this.axios.post('http://localhost:1998/wallet/register', {
            //         pri: pri,
            //     }).then((response)=>{
            //     // 一堆赋值，我也忘了要存啥了
            //     console.log(response);
            // })
        },
        recm() {
            console.log("我要收款");
        },
        changecmps(index) {
            this.cmp = index;
        },
        signoutf() {
            account = undefined;
            this.$router.push({
                path: '/'
            })
        },
        showImfof() {
            var imfo = (JSON.parse(window.localStorage.getItem(account))).bi;
            this.$alert(JSON.stringify(imfo), {
                confirmButtonText: '确定',});
        }
    }
}
</script>
<style>

</style>