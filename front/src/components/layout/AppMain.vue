<template>
  <v-main>
        <router-view></router-view>
  </v-main>
</template>

<script>
import pingMixin from "../../mixins/pingMixin";

export default {
  name: "AppMain",
  mixins: [pingMixin],
  created: function() {
    this.sendPing();
  },
  data: () => ({})
};
</script>
