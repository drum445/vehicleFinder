import Vue from 'vue'
import Router from 'vue-router'
import List from '@/components/List'
import Details from '@/components/Details.vue';

Vue.use(Router)

export default new Router({
  // remove # in URL
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'List',
      component: List
    },
    {
      path: '/vehicle/:vehicleID',
      name: 'Details',
      component: Details
    }
  ]
})
