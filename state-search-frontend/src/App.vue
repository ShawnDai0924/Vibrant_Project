<template>
  <div id="app">
    <h1>American States and Territories</h1>
    <!-- StateSearch component to search and select a state -->
    <StateSearch v-on:state-selected="handleStateSelected" />

    <!-- Display Google Map and highlight the selected state -->
    <MapView v-bind:selectedState="selectedState" />
  </div>
</template>


<script>
// Import the components
import StateSearch from './components/StateSearch.vue';
import MapView from './components/MapView.vue';

export default {
  components: {
    StateSearch,
    MapView,
  },
  data() {
    return {
      selectedState: '', // Holds the selected state from StateSearch.vue
    };
  },
  methods: {
    // Handle the event emitted from StateSearch.vue
    handleStateSelected(stateName) {
      this.selectedState = stateName; // Update the selected state
    },
  },
};
</script>

<style>
#app {
  text-align: center;
}
</style>
