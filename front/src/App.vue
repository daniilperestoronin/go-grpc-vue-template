<template>
  <v-app id="inspire">
    <app-nav-bar></app-nav-bar>
    <app-header></app-header>
    <app-main></app-main>
  </v-app>
</template>

<script>
import AppNavBar from "./components/layout/AppNavBar";
import AppHeader from "./components/layout/AppHeader";
import AppMain from "./components/layout/AppMain";

export default {
  components: {
    AppNavBar,
    AppHeader,
    AppMain
  },
  data: () => ({})
};
</script>
