<template>
  <div id="o">
    <el-row>
            <!-- 图片也可以换成 element 自带的 -->
            <el-row type="flex" justify="center"><el-col :xs="7" :sm="5" :md="3" :lg="2" :xl="2"><img src="../assets/wallet-outline.svg"></el-col></el-row> 
            <el-row type="flex" justify="center" style="margin-top: 0.5rem;">
                <el-col :xs="10" :sm="8" :md="5" :lg="4" :xl="4">
                    <router-link to="/Newwallet" style="text-decoration: none;">
                        <mybutton :buttonMsg="newbutton"></mybutton>
                    </router-link>
                </el-col>
            </el-row>
            <el-row type="flex" justify="center">
                <el-col :xs="10" :sm="8" :md="5" :lg="4" :xl="4">
                    <router-link to="/Loadwallet" style="text-decoration: none;">
                        <mybutton :buttonMsg="addbutton" style="margin-top: -10px"></mybutton>
                    </router-link>
                </el-col>
            </el-row>
        </el-row>  
    </div>
</template>
<script>
import mybutton from "../components/Mybutton"
export default {
    components: {
        mybutton
    },
    data() {
        return {
            addbutton: "加载钱包",
            newbutton: "新建钱包"
        }
    }
}
</script>
