import VueRouter from 'vue-router'

import SimpleTable from "../components/model/table/SimpleTable.vue"
import PivotTable from "../components/model/table/PivotTable.vue"
import Map from "../components/model/map/Map.vue"
import Calendar from "../components/model/calendar/Calendar.vue"
import Carousel from "../components/model/carousel/Carousel.vue"
import List from "../components/model/list/List.vue"
import Timeline from "../components/model/timeline/Timeline.vue"
import Stepper from "../components/model/stepper/Stepper.vue"
import Card from "../components/model/card/Card.vue"
import Form from "../components/model/form/Form.vue"

const routes = [
  { path: '/tables/simple', component: SimpleTable },
  { path: '/tables/pivot', component: PivotTable },
  { path: '/map', component: Map },
  { path: '/calendar', component: Calendar },
  { path: '/carousel', component: Carousel },
  { path: '/list', component: List },
  { path: '/timeline', component: Timeline },
  { path: '/stepper', component: Stepper },
  { path: '/card', component: Card },
  { path: '/form', component: Form }
];

export default new VueRouter({
  mode: 'history',
  routes: routes
});
