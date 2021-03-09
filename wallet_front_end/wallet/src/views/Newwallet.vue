<template>
    <div id="o">
        <!-- 输入信息，需要用 el-col 实现响应式布局 -->
        <el-row type="flex" justify="center">
            <el-col :xs="20" :sm="15" :md="12" :lg="8" :xl="7"  @click="register">
                <span class="imfo">
                    <h4 style="margin-top: 0;">我们将通过您的个人信息为您生成钱包公钥并在本地储存相关信息</h4>
                    <p >为保证您的账户安全，并让您在多个设备上打开您的钱包，请及时备份并安全保管相关文件</p>
                    <p>文件储存地址为.....</p>
                    <p>文件名为.....</p>
                </span>
                <p>姓名：</p>
                <el-input maxlength="12" v-model="name" minlength="1"></el-input>
                <p>身份证号：</p>
                <el-input maxlength="18" minlength="18" v-model="id"></el-input>
                <p>自定义字符串：</p>
                <el-input maxlength="255" v-model="string" minlength="1"></el-input>
                <mybutton :buttonMsg="bm" @click.native="register">创建钱包</mybutton>
            </el-col>
        </el-row>
        <backbutton></backbutton>
    </div>
</template>

<script>
// @ is an alias to /src
import mybutton from '../components/Mybutton.vue'
import backbutton from '../components/Backbutton.vue'

export default {
    components: {
        mybutton,
        backbutton
    },
    data() {
        return {
            bm: '创建钱包',
            id: '',
            name: '',
            string: ''
        }
    },
    methods: {
        register() {
            if (this.id == '' || this.name == '' || this.string == '') {
                this.$message.error ('提交的信息不能为空');
            } else {
                this.axios.post('http://localhost:1998/wallet/register', {
                name: this.name,
                id: this.id,
                str: this.string
                }).then((response)=>{
                    this.$message.success({
                        message: '创建成功',
                        duration: 1500
                    }); 
                    // 创建成功后加入 localstorage
                    var storage = window.localStorage;
                    var imfo = {};
                    imfo.bi = response.data;
                    storage.setItem(this.string, JSON.stringify(imfo));
                    console.log(storage);
                    // 跳转
                    setTimeout(() => {
                        this.$router.push({
                            name: 'Mainaction',
                            path: '/Mainaction',
                            query: {
                                account: this.string,
                            }
                        })
                    }, 1500);
                }).catch((response)=>{
                    this.$message.error('创建失败，请重试');
                    console.log(response);
                })
            }
        }
    }
}
</script>
<style>
    #o {
        position: relative;
        top: 45%;
        transform: translateY(-50%);
    }   
    .imfo {
        text-align: center;    
    }
    .imfo p {
        margin: 0.2rem;
        font-size: 0.8rem;
    }
</style>