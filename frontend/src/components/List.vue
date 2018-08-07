<template>
  <div class="list">
    <div class="row">
      <div class="col-sm-12">
        <form v-on:submit.prevent="filterVehicles">
          <input v-model="filters.free" v-on:keyup="startSearch" type="text" name="freeSearch" value="" placeholder="Free Search">
        </form>
        <hr>
        <table class="table-hover">
          <tr>
            <th>ID</th>
            <th>Make</th>
            <th>Short Model</th>
            <th>Long Model</th>
            <th>Trim</th>
            <th>Derivative</th>
            <th>Introduced</th>
            <th>Discontinued</th>
          </tr>
          <tr v-on:click="loadVehicle(vehicle.id)" class="vehicle-row" v-for="vehicle in vehicles">
            <td>{{vehicle.id}}</td>
            <td>{{vehicle.make}}</td>
            <td>{{vehicle.shortModel}}</td>
            <td>{{vehicle.longModel}}</td>
            <td>{{vehicle.trim}}</td>
            <td>{{vehicle.derivative}}</td>
            <td>{{vehicle.introduced}}</td>
            <td>{{vehicle.discontinued ? vehicle.discontinued : 'N/A'}}</td>
          </tr>
        </table>
        <hr>
        <button class="btn btn-primary" v-if="filters.page > 1" v-on:click="previousPage">Previous Page</button>
        <button class="btn btn-primary" v-if="filters.page*10 < count" v-on:click="nextPage">Next Page</button>
        <br><br>
        <p>Page: {{filters.page}}</p>
        <p>Total: {{count}}</p>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: "list",
  data () {
    return {
      filters: {"page": 1},
      vehicles: [],
      count: 0,
      typingtimer: null
    }
  },
  methods: {
    getVehicles() {
      clearTimeout(this.typingtimer);

  		axios.get(`${this.$parent.baseURL}`, {"params": this.filters}).then(response => {
        this.count = response.data.count;
  			this.vehicles = response.data.vehicles;
  		});
    },
    loadVehicle(id) {
      this.$router.push('/vehicle/' + id);
    },
    nextPage() {
      this.filters["page"]++;
      this.getVehicles();
    },
    previousPage() {
      this.filters["page"]--;
      this.getVehicles();
    },
    filterVehicles() {
      this.filters["page"] = 1;
      this.getVehicles();
    },
    startSearch() {
      var self = this;
      clearTimeout(this.typingtimer);
      this.typingtimer = setTimeout(function () {
        self.filterVehicles();
      }, 500);

    }
  },
  mounted() {
    this.getVehicles();
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
  color: blue !important;
  cursor: pointer;
}

.vehicle-row:hover {
  background-color: cyan !important;
  cursor: pointer !important;
}

table {
    font-family: arial, sans-serif;
    border-collapse: collapse;
    width: 100%;
}

td, th {
    border: 1px solid #dddddd;
    text-align: left;
    padding: 8px;
}

tr:nth-child(even) {
    background-color: #dddddd;
}
</style>
