<template>
    <div style="height: 100%;">
        <navmenu @changecmp="changecmps" ref="nav"></navmenu>  
        <el-row type="flex" justify="center" id="o">
            <el-col :xs="20" :sm="15" :md="12" :lg="8" :xl="7">  
            <div v-show="cmp == 1">
                <p>购买金额:</p>
                <el-input maxlength="10" v-model="money" ></el-input>
                <mybutton :buttonMsg="buy" @click.native="buym"></mybutton>
            </div>

            <div v-show="cmp == 2">
                <p>收款方公钥:</p>
                <el-input v-model="G1" placeholder="G1"></el-input>
                <el-input v-model="G2" placeholder="G2" style="margin-top:10px;"></el-input>
                <el-input v-model="P" placeholder="P" style="margin-top:10px;"></el-input>
                <el-input v-model="H" placeholder="H" style="margin-top:10px;"></el-input>
                <p>转账金额:</p>
                <el-input maxlength="10" v-model="transmoney" ></el-input>
                <p>代币承诺</p>
                <el-input maxlength="10" v-model="moneyProm" ></el-input>
                <p>随机数</p>
                <el-input maxlength="10" v-model="r" ></el-input>
                <!-- 上面这些够了，可以返回东西了 -->
                <mybutton :buttonMsg="transfer" @click.native="transferm"></mybutton>
            </div>

            <div v-show="cmp == 3">
                 <p>交易hash</p>
                <el-input v-model="hash" ></el-input>
                <mybutton :buttonMsg="recv" @click.native="recm"></mybutton>
            </div>

            <div v-show="cmp == 4">
                <el-table  :data="his" style="width: 100%">
                    <el-table-column
                        prop="amount"
                        label="金额"
                        width="180">
                    </el-table-column>
                    <el-table-column
                        prop="cmv"
                        label="姓名"
                        width="180">
                    </el-table-column>
                    <el-table-column
                        prop="r"
                        label="地址">
                    </el-table-column>
                    <el-table-column
                        prop="spent"
                        label="使用">
                    </el-table-column>
                </el-table>
            </div>

            <div v-show="cmp == 5">
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
var his = new Array();
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
            transmoney: '',
            moneyProm: '',
            r: '',
            G1: '',
            G2: '',
            P: '',
            H: '',
            hash: '',
            his: his
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
            return pri;
        },
        storeImfo(response, amount) {
            // 更新信息
            // 取出 history 并修改
            var old = JSON.parse(window.localStorage.getItem(account));
            var neww = response.data;
            neww.vm = amount;
            old.history.push(neww); // 喜加一
            window.localStorage.account = JSON.stringify(old);
            console.log(window.localStorage.account);
        },
        Pub(G1, G2, P, H) {
            this.G1 = G1;
            this.G2 = G2;
            this.P = P;
            this.H = H;
        },
        transferm() {
            console.log("我要转账");
            var pri = this.getPri();
            this.axios.post('http://localhost:4396/wallet/buyCoin', {
                priA: pri,
                amount: this.transmoney,
                pubB: new this.Pub(this.G1, this.G2, this.P, this.H),
                cmv: this.moneyProm,
                r: this.r
            }).then((response)=>{
                this.storeImfo(response, -this.transmoney);
            }).catch((response)=>{
                    this.$message.error(response);
                    console.log(response);
            });
        },
        buym() {
            console.log("我要购币");
            var pri = this.getPri();
            this.axios.post('http://localhost:4396/wallet/buyCoin', {
                pri: pri,
                amount: this.buym
            }).then((response)=>{
                this.storeImfo(response, this.buym);
            }).catch((response)=>{
                    this.$message.error(response);
                    console.log(response);
            });
        
        },
        recm() {
            console.log("我要收款");
            var pri = this.getPri();
            this.axios.post('http://localhost:4396/wallet/buyCoin', {
                pri: pri,
                hash: this.hash
            }).then((response)=>{
                this.storeImfo(response, response.data.amount);
            }).catch((response)=>{
                    this.$message.error(response);
                    console.log(response);
            });
        },
        changecmps(index) {
            this.cmp = index;
            // 防止小手机转账界面崩坏
            if (this.cmp == 2) {
                document.getElementById("o").style.top = "10px";
                document.getElementById("o").style.transform = "none";
            } else {
                document.getElementById("o").style.top = "45%";
                document.getElementById("o").style.transform = "translateY(-50%)";
            }
            // 加载历史
            if (this.cmp == 4) {
                his = JSON.parse(window.localStorage.getItem(account)).history;
            }
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