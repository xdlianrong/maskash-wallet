<template>
  <div>
    <h4 style="text-align: center; margin-top: 0;">我们将通过您的个人信息为您生成钱包公钥并在本地储存相关信息</h4>
            <p style="text-align: center;">为保证您的账户安全，并让您在多个设备上打开您的钱包</p>
            <p style="text-align: center;">请及时备份并安全保管相关文件</p>
            <p style="text-align: center;">文件储存地址为.....</p>
            <p style="text-align: center;">文件名为.....</p>
            <el-row type="flex" justify="center">
                <el-col :xs="20" :sm="15" :md="12" :lg="8" :xl="7"  @click="register" id="hh">
                    <p>姓名：</p>
                    <el-input maxlength="12" v-model="name" minlength="1"></el-input>
                    <p>身份证号：</p>
                    <el-input maxlength="18" minlength="18" v-model="id"></el-input>
                    <p>自定义字符串：</p>
                    <el-input maxlength="255" v-model="string" minlength="1"></el-input>
                    <mybutton :buttonMsg="bm" style="margin-top: 2.5rem;" @click.native="register">创建钱包</mybutton>
                </el-col>
            </el-row>
  </div>
</template>

<script>
// @ is an alias to /src
import mybutton from '../components/Mybutton.vue'

export default {
    components: {
        mybutton
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
            console.log("?"); 
        }
    }
}

</script>
