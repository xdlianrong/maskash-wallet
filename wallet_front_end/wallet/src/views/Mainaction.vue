<template>
  <div style="height: 100%;">
    <navmenu @changecmp="changecmps" ref="n"></navmenu>
    <el-row type="flex" justify="center" id="om">
      <el-col :xs="20" :sm="15" :md="8" :lg="8" :xl="8" v-show="cmp !== 4">
        <div v-show="cmp === 1">
          <p>购买数额:</p>
          <el-input maxlength="10" v-model="money" oninput="value=value.replace(/[^\d]/g,'')"></el-input>
          <mybutton :buttonMsg="buy" @click.native="buym"></mybutton>
        </div>

        <div v-show="cmp === 2">
          <p><span class="t"></span>接收方公钥:</p>
          <el-input v-model="G1" placeholder="G1"/>
          <el-input v-model="G2" placeholder="G2" style="margin-top:10px;"/>
          <el-input v-model="P" placeholder="P" style="margin-top:10px;"/>
          <el-input v-model="pub" placeholder="pub" style="margin-top:10px;"/>
          <div style="margin-top:10px;">
            <a>或选择本地账户&emsp;</a>
            <el-select v-model="baccount" placeholder="请选择" clearable>
              <el-option v-for="account in accountList" :value="account" :key="account.key"
                         :label="account"/>
            </el-select>
          </div>
          <p><span class="t"/>使用承诺的数额</p>
          <el-input maxlength="10" v-model="transmoney" oninput="value=value.replace(/[^\d]/g,'')"></el-input>
          <p><span class="t"/>承诺cmv</p>
          <el-input v-model="cmv"/>
          <p><span class="t"/>随机数vor</p>
          <el-input v-model="r"/>
          <div style="margin-top:10px;">
            <a>或选择本地承诺&emsp;</a>
            <el-select v-model="bindCM" placeholder="请选择" clearable>
              <el-option v-for="cm in valuableCMList" :value="cm" :key="cm.key"
                         :label="cm.cmv.slice(0,8)+'... 价值：'+cm.amount"/>
            </el-select>
          </div>
          <p><span class="t"/>转出数额:</p>
          <el-input maxlength="10" v-model="spend" oninput="value=value.replace(/[^\d]/g,'')"/>
          <!-- 上面这些够了，可以返回东西了 -->
          <mybutton :buttonMsg="transfer" @click.native="transform" style="margin-bottom: 20px"/>
        </div>

        <div v-show="cmp === 3">
          <p>交易hash</p>
          <el-input v-model="hash"/>
          <mybutton :buttonMsg="recv" @click.native="recm"/>
        </div>

        <div v-show="cmp === 5">
          <mybutton :buttonMsg="showInfo" @click.native="showInfof" class="b1"/>
          <mybutton :buttonMsg="signout" @click.native="signoutf"/>
        </div>
      </el-col>

    </el-row>
    <el-row v-show="cmp === 4" type="flex" justify="center">
      <el-col :xs="24" :sm="20" :md="17" :lg="15" :xl="15">
        <el-table :data="hisList">
          <el-table-column
              prop="amount"
              label="数额"/>
          <el-table-column
              prop="hash"
              label="哈希hash"/>
          <el-table-column
              prop="cmv"
              label="承诺cmv"/>
          <el-table-column
              prop="vor"
              label="随机数vor"/>
        </el-table>
      </el-col>
    </el-row>

    <div id="progress" v-show="cmp === 1 && show">
      <el-timeline>
        <el-timeline-item
            v-for="(activity, index) in activities"
            :key="index"
            :icon="activity.icon"
            :type="activity.type"
            :color="activity.color"
            :size="activity.size"
            :timestamp="activity.timestamp">
          {{ activity.content }}
        </el-timeline-item>
      </el-timeline>
    </div>
  </div>
</template>
<script>
import navmenu from '../components/Navmenu'
import mybutton from '../components/Mybutton'
import globle from '../globle'

let account;
const accountList = []
// 本账户可用的承诺
let valuableCMList = [];
export default {
  components: {
    navmenu,
    mybutton,
  },
  data() {
    return {
      showLog: globle.showLog,
      transfer: '发起转出',
      buy: '兑换',
      recv: '接收',
      signout: '登出',
      showInfo: '显示账户信息',
      money: '',
      cmp: '1', // 用来改变显示的组件
      transmoney: '',
      // 转账时要花掉的承诺
      cmv: '',
      r: '',
      G1: '',
      G2: '',
      P: '',
      pub: '',
      hash: '',
      hisList: '',
      spend: '',
      nowm: '',
      accountList,
      valuableCMList,
      baccount: '',
      bindCM: '',
      loading: true,
      show: false,
      activities: [{
        content: '开始转出'
      }, {
        content: '生成会计平衡证明'
      }, {
        content: '生成相等证明'
      }, {
        content: '生成范围证明'
      }, {
        content: '生成格式正确证明'
      }, {
        content: '挖矿共识'
      }, {
        content: '转出成功'
      }]
    }
  },
  created: function () {
    account = this.$route.query.account;
    if (account === undefined) {
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
    this.hisList = JSON.parse(window.localStorage.getItem(account)).history;
  },
  mounted: function () {
    this.showCoin();
    // 获取本地账户列表
    for (let i = 0; i < window.localStorage.length; i++) {
      const name = window.localStorage.key(i);
      if (name !== 'loglevel:webpack-dev-server' && name !== account) {
        accountList.push(name);
      }
    }
    // 获取本账户所有可用承诺
    if (valuableCMList.length > 0)
      valuableCMList = []
    for (let i = 0; i < this.hisList.length; i++) {
      if (this.hisList[i].valuable)
        valuableCMList.push(this.hisList[i]);
    }
    this.$refs.n.changeName(account);
  },
  watch: {
    baccount(val) {
      if (!val) {
        this.G1 = '';
        this.G2 = '';
        this.P = '';
        this.pub = '';
      } else {
        const b = JSON.parse(window.localStorage.getItem(val)).info;
        this.G1 = b.G1;
        this.G2 = b.G2;
        this.P = b.P;
        this.pub = b.publickey;
      }
    },
    bindCM(val) {
      if (!val) {
        this.transmoney = ''
        this.cmv = ''
        this.r = ''
      } else {
        this.transmoney = val.amount
        this.cmv = val.cmv
        this.r = val.vor
      }
    }
  },
  methods: {
    getPri() {
      return JSON.parse(window.localStorage.getItem(account)).info;
    },
    storeInfo(response, amount) {
      // 更新信息
      // 取出 history 并修改
      const old = JSON.parse(window.localStorage.getItem(account));
      const newRecord = response.data.coin;
      newRecord.vm = amount;
      old.history.push(newRecord); // 喜加一
      window.localStorage.setItem(account, JSON.stringify(old));
      this.showCoin();
    },
    //废除承诺，将承诺值为cmv的承诺在localStorage中的valuable值改为false
    abolitionCM(cmv) {
      const findCmvIndex = function (cmv, history) {
        for (let i = 0; i < history.length; i++) {
          if (history[i].cmv === cmv) {
            return i
          }
        }
        return -1
      }
      const localStorage = JSON.parse(window.localStorage.getItem(account));
      const index = findCmvIndex(cmv, localStorage.history)
      if (index === -1 || !localStorage.history[index].valuable) return
      localStorage.history[index].valuable = false
      window.localStorage.setItem(account, JSON.stringify(localStorage));
    },
    Pub(G1, G2, P, H) {
      this.G1 = G1;
      this.G2 = G2;
      this.P = P;
      this.H = H;
    },
    transform() {
      if (this.showLog)
        console.log("我要转账");
      if (!this.G1 || !this.G2 || !this.P || !this.pub) {
        this.$message.error('请完整输入接收方公钥账户');
        return
      }
      if (!this.transmoney) {
        this.$message.error('请填写使用承诺的数额');
        return
      }
      const amount = parseInt(this.transmoney)
      if (amount === 0) {
        this.$message.error('使用承诺的数额需大于0');
        return
      }
      if (!this.spend) {
        this.$message.error('请填写要转出的数额');
        return
      }
      const spend = parseInt(this.spend)
      if (spend === 0) {
        this.$message.error('要转出的数额需大于0');
        return
      }
      if (this.transmoney < this.spend) {
        this.$message.error('转出数额不可大于使用承诺的数额');
        return
      }
      if (!this.cmv) {
        this.$message.error('请填写承诺cmv');
        return
      }
      if (!this.r) {
        this.$message.error('请填写承诺随机数vor');
        return
      }
      const pri = this.getPri();
      this.$message('正在生成：会计平衡证明、监管相等证明、范围证明、密文格式正确证明');
      this.axios.post('http://' + globle.serverIp + '/wallet/exchange', {
        sg1: pri.G1,
        sg2: pri.G2,
        sp: pri.P,
        sh: pri.publickey,
        sx: pri.privatekey,
        amount: amount,
        rg1: this.G1,
        rg2: this.G2,
        rp: this.P,
        rh: this.pub,
        cmv: this.cmv,
        vor: this.r,
        spend: spend
      }).then((response) => {
        this.abolitionCM(this.cmv)
        response.data.coin.valuable = true
        this.storeInfo(response, -this.spend);
      }).catch((response) => {
        this.$message.error(response);
        console.error(response);
      });
    },
    buym() {
      if (this.showLog)
        console.log("我要兑换");
      if (!this.money) {
        this.$message.error("请填写要兑换的金额");
        return
      }
      const amount = parseInt(this.money)
      if (amount === 0) {
        this.$message.error("要兑换的金额需大于0");
        return
      }
      const pri = this.getPri();
      if (this.showLog)
        console.log(pri.privatekey);
      this.axios({
        url: 'http://' + globle.serverIp + '/wallet/buycoin',
        method: 'post',
        data: {
          g1: pri.G1,
          g2: pri.G2,
          p: pri.P,
          h: pri.publickey,
          x: pri.privatekey,
          amount: amount
        },
        timeout: '600000'
      }).then((response) => {
        response.data.coin.valuable = true
        console.log(response)
        this.storeInfo(response, this.money);
      }).catch((response) => {
        this.$message.error(response);
        console.error(response);
      });
      this.show = true;
      this.progress();
      // this.$message.success({
      //   message: '金额加密正确',
      //   duration: 1000
      // });
      // setTimeout(() => {
      //   this.$message.success({
      //     message: '公钥加密正确',
      //     duration: 1000
      //   });
      // }, 2000);
    },
    recm() {
      if (this.showLog)
        console.log("我要收款");
      if (this.hash.length === 0) {
        this.$message.error("请填写交易hash");
        return
      }
      if (this.hash.length !== 66) {
        this.$message.error("未找到指定交易");
        return
      }
      const pri = this.getPri();
      this.axios({
        url: 'http://' + globle.serverIp + '/wallet/receive',
        method: 'post',
        data: {
          g1: pri.G1,
          g2: pri.G2,
          p: pri.P,
          h: pri.publickey,
          x: pri.privatekey,
          hash: this.hash
        },
        timeout: '600000'
      }).then((response) => {
        if (isNaN(response.data.coin.amount)) {
          this.$message.error("交易解密失败");
          return
        }
        this.storeInfo(response, response.data.coin.amount);
        this.$message.info("接收成功！");
      }).catch((response) => {
        this.$message.error(response);
        console.error(response);
      });
    },
    changecmps(index) {
      this.cmp = index;
      // 防止小手机转账界面崩坏
      if (this.cmp === 2) {
        document.getElementById("om").style.top = "10px";
        document.getElementById("om").style.transform = "none";
      } else {
        document.getElementById("om").style.top = "30%";
        document.getElementById("om").style.transform = "translateY(-50%)";
      }
      // 加载历史
      if (this.cmp === 4) {
        // 更新
        this.hisList = JSON.parse(window.localStorage.getItem(account)).history;
      }
    },
    signoutf() {
      account = undefined;
      accountList.length = 0;
      this.$router.push({
        path: '/'
      })
    },
    showInfof() {
      this.showCoin();
      const accountInfo = (JSON.parse(window.localStorage.getItem(account))).info
      var G1 = JSON.stringify(accountInfo.G1);
      var G2 = JSON.stringify(accountInfo.G2);
      var P = JSON.stringify(accountInfo.P);
      var pub = JSON.stringify(accountInfo.publickey);
      var pri = JSON.stringify(accountInfo.privatekey);
      this.$alert("<p>G1:" + G1 + "</p>" +
          "<p>G2:" + G2 + "</p>" +
          "<p>P:" + P + "</p>" +
          "<p>pub:" + pub + "</p>" +
          "<p>pri:" + pri + "</p>", {
        confirmButtonText: '确定',
        dangerouslyUseHTMLString: true,
        customClass: 'message_box_alert'
      });

      // var twqee = {
      //     hash: "ASBWJAKFA",
      //     cmv: "SDUIFUISAASK",
      //     r: "DSFSAFSA",
      //     amount: 100,
      //     vm: 100
      // };
      // var old = JSON.parse(window.localStorage.getItem(account));            old.history.push(twqee); // 喜加一
      // window.localStorage.setItem(account, JSON.stringify(old));
      // if(this.showLog)
      //  console.log(old);
    },
    showCoin() {
      // 刷新余额
      if (this.showLog)
        console.log("改余额");
      var sum = 0;
      var his = JSON.parse(window.localStorage.getItem(account)).history;
      for (var i = 0; i < his.length; i++) {
        if (his[i].vm !== undefined) {
          sum = sum + parseInt(his[i].vm);
          if (this.showLog)
            console.log(sum);
        }
      }
      // 上当了,因为少了个s找了好久问题
      this.$refs.n.changeSum(sum);
    },
    // 改变展示状态
    progress() {
      var time = 0;
      for (var i = 0; i < 7; i++) {
        time = time + Math.random() * 2000 + 1000;
        this.pc(i, time);
      }
    },
    pc(i, time) {
      var state = document.getElementById("progress");
      setTimeout(() => {
        var ojArr = state.getElementsByTagName("li")[i].getElementsByTagName("div");
        ojArr[0].style.borderLeft = "2px solid limegreen";
        ojArr[1].style.backgroundColor = "limegreen";
        if (this.showLog)
          console.log(ojArr[1]);
        if (i === 7) {
          setTimeout(() => {
            this.show = false;
          }, 1000);
        }
      }, time);
    }
  }
}
</script>
<style>
.message_box_alert {
  word-break: break-all !important;
}

.el-message-box {
  width: 80%;
}

#om {
  position: relative;
  top: 35%;
  transform: translateY(-50%);
}

.el-table td, .el-table th {
  text-align: center !important;
}

#progress {
  display: flex;
  margin-left: 70%;
  align-items: flex-end;
  flex-direction: column;
  transform: translateY(-25%);
  width: 20%;
}

.el-timeline-item {
  padding-bottom: 50px !important;
}

.el-timeline {
  font-size: 16px !important;
}
</style>
<style scoped>
.el-col p {
  margin-top: 25px !important;
}
</style>