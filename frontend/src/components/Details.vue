<template>
  <div class="post">
    <div class="row" v-cloak>
		  <div class="col-lg-12">
        <h1>{{vehicle.derivative}}</h1>
        <p>{{vehicle.id}}</p>
        <hr>
      </div>
    </div>
    <transition name="fade">
      <div class="row" v-if="loaded">
        <div class="col-md-6">
          <h4>Vehicle Information</h4>
          <ul class="list-group">
            <li class="list-group-item">Make: {{vehicle.make}}</li>
            <li class="list-group-item">Short Model: {{vehicle.shortModel}}</li>
            <li class="list-group-item">Long Model: {{vehicle.longModel}}</li>
            <li class="list-group-item">Trim: {{vehicle.trim}}</li>
            <li class="list-group-item">Derivative: {{vehicle.derivative}}</li>
            <li class="list-group-item">Introduced: {{vehicle.introduced}}</li>
            <li class="list-group-item">Discontinued: {{vehicle.discontinued}}</li>
          </ul>
        </div>
        <div class="col-md-6">
          <img v-bind:src="vehicle.image" style="width:60%;">
        </div>
        <center>
          <router-link class="nav-link" to="/">Back To List</router-link>
        </center>
      </div>
      <div class="row" v-if="loaded == false">
        <h1>Failed to load lead</h1>
      </div>
    </transition>
	</div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'vehicle',
  data () {
    return {
      loaded: null,
      vehicle: {},
    }
  },
  methods: {
    getLead() {
  		axios.get(`${this.$parent.baseURL}/${this.$route.params.vehicleID}`).then(response => {
  			this.vehicle = response.data;
        this.loaded = true;
  		}).catch(response => {
        this.loaded = false;
      });
    }
  },
  mounted() {
   this.getLead();
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1, h2 {
  font-weight: normal;
}

ul {
  list-style-type: none;
  padding: 0;
}

li {
  display: inline-block;
  margin: 0 10px;
}

a {
  color: #42b983;
}

table, th, td {
    border: 1px solid black;
}

th, td {
    padding: 15px;
}

table {
    border-spacing: 5px;
}

.fade-enter-active, .fade-leave-active {
  transition: opacity .5s;
}
.fade-enter, .fade-leave-to /* .fade-leave-active below version 2.1.8 */ {
  opacity: 0;
}
</style>
