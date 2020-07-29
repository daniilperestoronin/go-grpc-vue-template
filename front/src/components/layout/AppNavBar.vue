<template>
  <v-navigation-drawer v-model="drawer" :clipped="$vuetify.breakpoint.lgAndUp" app>
    <v-list dense>
      <template v-for="item in items">
        <v-row v-if="item.heading" :key="item.heading" align="center">
          <v-col cols="6">
            <v-subheader v-if="item.heading">{{ item.heading }}</v-subheader>
          </v-col>
          <v-col cols="6" class="text-center">
            <a href="#!" class="body-2 black--text">EDIT</a>
          </v-col>
        </v-row>
        <v-list-group
          v-else-if="item.children"
          sub-group
          :key="item.text"
          v-model="item.model"
          :prepend-icon="item.model ? item.icon : item['icon-alt']"
          append-icon
        >
          <template v-slot:activator>
            <v-list-item-content>
              <v-list-item-title>{{ item.text }}</v-list-item-title>
            </v-list-item-content>
          </template>
          <v-list-item v-for="(child, i) in item.children" :key="i" link :to="child.href">
            <v-list-item-action v-if="child.icon">
              <v-icon>{{ child.icon }}</v-icon>
            </v-list-item-action>
            <v-list-item-content>
              <v-list-item-title>{{ child.text }}</v-list-item-title>
            </v-list-item-content>
          </v-list-item>
        </v-list-group>
        <v-list-item v-else :key="item.text" :to="item.href">
          <v-list-item-action>
            <v-icon>{{ item.icon }}</v-icon>
          </v-list-item-action>
          <v-list-item-content>
            <v-list-item-title>{{ item.text }}</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </template>
    </v-list>
  </v-navigation-drawer>
</template>

<script>
export default {
  name: "AppNavBar",
  data: () => ({
    items: [
      {
        icon: "mdi-file-table-box-outline",
        "icon-alt": "mdi-file-table-box-outline",
        text: "Tables",
        model: false,
        children: [
          { text: "Simple", href: "/tables/simple" },
          { text: "Pivot", href: "/tables/pivot" },
        ],
      },
      { icon: "mdi-map-marker-radius-outline", text: "Map", href: "/map" },
      { icon: "mdi-calendar-month-outline", text: "Calendar", href: "/calendar" },
      { icon: "mdi-view-dashboard-outline", text: "Card" , href: "/card"},
      { icon: "mdi-view-carousel-outline", text: "Carousel" , href: "/carousel"},
      { icon: "mdi-form-select", text: "Form" , href: "/form"},
      { icon: "mdi-format-list-bulleted", text: "List" , href: "/list"},
      { icon: "mdi-clipboard-list-outline", text: "Stepper" , href: "/stepper"},
      { icon: "mdi-timeline-clock-outline", text: "Timeline" , href: "/timeline"},
      { icon: "mdi-cog", text: "Setting" , href: "/setting"},
    ],
  }),
};
</script>
