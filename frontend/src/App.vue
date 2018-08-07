<template>
  <div id="app">

    <nav class="navbar navbar-expand-sm bg-dark navbar-dark">
      <router-link class="navbar-brand" to="/">SiteName</router-link>
      <!-- Links -->
      <ul class="navbar-nav">
        <li class="nav-item">
          <router-link class="nav-link" to="/">Home</router-link>
        </li>
      </ul>
    </nav>

    <div class="er-container">
      <keep-alive :include="['list']">
        <router-view></router-view>
      </keep-alive>
    </div>

  </div>
</template>

<script>
export default {
  name: 'app',
  data () {
    return {
      baseURL: "http://localhost:5000/api/vehicle"
    }
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

.er-container {
  padding-top: 20px;
  padding-bottom: 40px;
  padding-left: 20px;
  padding-right: 20px;
}

[v-cloak] {
  display: none
}
</style>
