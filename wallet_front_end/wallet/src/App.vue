<template>
  <div id="app">
    <router-view/>
  </div>
</template>

<style>
    #app {
            position: relative;
            top: 40%;
            transform: translateY(-50%);
    }

    html, body {
        height: 100%;
        margin: 0; 
        padding: 0;
    }

    router-link {
          text-decoration: none;
    }
</style>
