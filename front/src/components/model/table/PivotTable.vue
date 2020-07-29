<template>
  <div id="app" class="container mt-5">
    <div class="mb-5">
      <pivot
        :data="asyncData"
        :fields="fields"
        :available-field-keys="availableFieldKeys"
        :row-field-keys="rowFieldKeys"
        :col-field-keys="colFieldKeys"
        :reducer="reducer"
        :default-show-settings="defaultShowSettings"
      >
        <template v-slot:value="{ value }">
          {{ value | number }}
        </template>
        <template v-slot:countryFlagHeader="{ value }">
          <span :class="`flag-icon flag-icon-${countryCode(value)}`"></span>
        </template>
        <template v-slot:countryNameHeader="{ value }">
          {{ value | capitalize }}
        </template>
        <template v-slot:genderHeader="{ value }">
          {{ value | capitalize }}
        </template>
        <template v-slot:computing>
          <div class="position-absolute w-100 h-100 text-center" style="z-index: 1;">
            <div class="position-sticky bg-white d-inline-block mt-5 p-3" style="top: 0;">
              Loading table values...
            </div>
          </div>
        </template>
      </pivot>
    </div>
  </div>
</template>

<script>
import Pivot from '@marketconnect/vue-pivot-table'
import data from './data'

export default {
  name: 'PivotTable',
  components: { Pivot },
  data: () => {
    return {
      data: data,
      asyncData: [],

      // Pivot params
      fields: [{
        key: 'country',
        getter: item => item.country,
        label: 'Country',
        headers: [{
          slotName: 'countryFlagHeader',
          label: 'Flag',
          checked: true
        }, {
          slotName: 'countryNameHeader',
          label: 'Name',
          checked: true
        }],
        headerAttributeFilter: true,
        valueFilter: true
      }, {
        key: 'gender',
        getter: item => item.gender,
        label: 'Gender',
        headerSlotName: 'genderHeader',
        valueFilter: true,
        valueFilterSlotName: 'genderHeader'
      }, {
        key: 'year',
        getter: item => item.year,
        label: 'Year',
        valueFilter: true
      }],
      availableFieldKeys: [],
      rowFieldKeys: ['country', 'gender'],
      colFieldKeys: ['year'],
      reducer: (sum, item) => sum + item.count,
      defaultShowSettings: true,
      isDataLoading: false,

      // Pivot table standalone field params
      rowFields: [{
        getter: item => item.country,
        label: 'Country',
        headerSlotNames: ['countryFlagHeader', 'countryNameHeader']
      }, {
        getter: item => item.gender,
        label: 'Gender',
        headerSlotName: 'genderHeader'
      }],
      colFields: [{
        getter: item => item.year,
        label: 'Year'
      }]
    }
  },
  methods: {
    countryCode: function(country) {
      switch (country) {
        case 'Australia':
          return 'au'
        case 'China':
          return 'cn'
        case 'France':
          return 'fr'
        case 'India':
          return 'in'
        case 'United States':
          return 'us'
      }
    }
  },
  created: function() {
    // Simulate async data loading
    this.isDataLoading = true
    setTimeout(() => {
      this.asyncData = data
      this.isDataLoading = false
    }, 3000)
  },
  filters: {
    number: function(value) {
      return value.toLocaleString()
    },
    capitalize: function(value) {
      return value.replace(/\b\w/g, l => l.toUpperCase())
    }
  }
}
</script>

<style lang="scss">
$enable-rounded: false;
$table-cell-padding: .5rem; // default in bs5
$table-cell-padding-sm: .25rem; // default in bs5
@import '~bootstrap/scss/bootstrap.scss';

/* Table with aria-busy attribute */
table[aria-busy='true'] {
  opacity: 0.6;
}

/* FontAwesome icons */
.svg-inline--fa.fa-fw {
  width: 1.25em;
}

.svg-inline--fa {
  display: inline-block;
  font-size: inherit;
  height: 1em;
  overflow: visible;
  vertical-align: -.125em !important;
}

.fa-pulse {
  animation: fa-spin 1s infinite steps(8);
}

@keyframes fa-spin {
  0% {
    transform: rotate(0deg);
  }

  to {
    transform: rotate(1turn);
  }
}
</style>
