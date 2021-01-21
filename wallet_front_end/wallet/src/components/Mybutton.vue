<template>
    <div class="border">
        <div class="button">{{ buttonMsg }}</div>
    </div>
</template>
<style>
    .button {
        font-size: 1.2rem;
        text-align: center;
        background: #59aafa;
        border-radius: 4px;
        color: white;
        line-height: 3.5rem;
        box-shadow: 0 2px 4px rgba(0, 0, 0, .12), 0 0 6px rgba(0, 0, 0, .04);
    }

    .border {
        border-radius: 4px;
        border: #f2f2f2 9px solid;
        margin-top: 2.5rem;
    }
</style>
<script>
export default {
    props: {
        buttonMsg: {
            type: String,
            require: true
        }
    }
}
</script>