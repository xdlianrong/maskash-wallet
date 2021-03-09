<template>
    <div>
        <ul class="infinite-list"  style="overflow:auto;padding:0;">
            <li v-for="name in list" :key="name.id" class="infinite-list-item" @click="signin(name)">{{ name }}</li>
        </ul>
        <backbutton></backbutton>
    </div>
</template>
<script>
import backbutton from '../components/Backbutton'
var list = new Array(); 
export default {    
    components: {
        backbutton
    },
    data () {
        return {
            list
        }
    },
    mounted: function () {
        // 填充 list 数组        
        var storage = window.localStorage;
        // 防止再次填充 list，清空数组
        list.length = 0;
        for(var i = 0; i < storage.length; i++) {
            if (storage.key(i) != 'loglevel:webpack-dev-server') {
                list.push(storage.key(i));
            }
        }        
    },
    methods: {
        signin (name) {
            this.$router.push({
                name: 'Mainaction', // 没有这句会 undefined
                path: '/Mainaction',
                query: {
                    account: name
                }
            })
        }
    }
}
</script>
<style>
    .infinite-list .infinite-list-item {
        display: flex;
        align-items: center;
        justify-content: center;
        height: 50px;
        background: #e8f3fe;
        margin: 10px;
        color: #7dbcfc;
    }
</style>