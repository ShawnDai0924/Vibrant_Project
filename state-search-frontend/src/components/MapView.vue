<template>
  <div id="map" class="map-container"></div>
</template>

<script>
// Import the Loader from '@googlemaps/js-api-loader'
import { Loader } from '@googlemaps/js-api-loader';

export default {
  props: {
    selectedState: {
      type: String,
      required: false,
    },
  },
  data() {
    return {
      map: null,
      stateLayer: null, // To store the state boundaries
    };
  },
  watch: {
    selectedState(newVal) {
      if (newVal) {
        this.highlightState(newVal); // Highlight the new state whenever it changes
      }
    }
  },
  mounted() {
    const loader = new Loader({
      apiKey: 'AIzaSyB1-psa7VQjYCGqGVYsLJVrOfdnYOYQuFs',
      version: 'weekly',
    });

    loader.load().then(() => {
      // eslint-disable-next-line no-undef
      const mapOptions = {
        center: { lat: 37.1, lng: -95.7 }, // Centered in the US
        zoom: 4,
      };

      // eslint-disable-next-line no-undef
      this.map = new google.maps.Map(document.getElementById('map'), mapOptions);

      // Initialize the data layer
      // eslint-disable-next-line no-undef
      this.stateLayer = new google.maps.Data({ map: this.map });
    });
  },
  methods: {
    // Highlight the selected state by loading GeoJSON and clearing old data
    highlightState(stateName) {
      if (!this.map || !this.stateLayer) return;

      // Remove all existing highlights
      this.stateLayer.forEach((feature) => {
        this.stateLayer.remove(feature); 
      });

      // Load state boundary GeoJSON
      const geoJsonUrl = `https://raw.githubusercontent.com/PublicaMundi/MappingAPI/master/data/geojson/us-states.json`; // Example GeoJSON source

      // Load GeoJSON data for the new state
      this.stateLayer.loadGeoJson(geoJsonUrl, null, () => {
        // Filter and remove all non-matching features
        this.stateLayer.forEach((feature) => {
          if (feature.getProperty('name') !== stateName) {
            this.stateLayer.remove(feature); // Remove non-matching states
          }
        });

        this.stateLayer.setStyle({
          fillColor: 'yellow',
          strokeWeight: 2,
        });

        // Zoom in on the selected state
        this.map.fitBounds(this.getFeatureBounds(this.stateLayer));
      });
    },

    // Helper function to get the bounds of a GeoJSON feature
    getFeatureBounds(layer) {
      // eslint-disable-next-line no-undef
      const bounds = new google.maps.LatLngBounds();
      layer.forEach((feature) => {
        const geometry = feature.getGeometry();
        geometry.forEachLatLng((latLng) => {
          bounds.extend(latLng);
        });
      });
      return bounds;
    }
  }
};
</script>

<style scoped>
#map {
  width: 100%;
  height: 500px;
  margin-top: 50px;
}
</style>
