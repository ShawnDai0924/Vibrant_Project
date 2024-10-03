<template>
  <div class="search-container">
    <div class="search-box">
      <input
        v-model="searchTerm"
        @input="onSearch"
        placeholder="Search"

      />
      <!-- Dropdown list -->
      <ul v-if="filteredStates.length" class="dropdown">
        <li
          v-for="state in filteredStates"
          :key="state.id"
          v-html="highlightMatch(state.name)"
          @click="selectState(state.name)"
        >
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      searchTerm: '',      
      states: [],          
      filteredStates: [],  
      loading: false,      
      error: null,         
    };
  },
  methods: {
    // Triggers when the user types into the search input
    async onSearch() {
      // Clear filteredStates if the search term is too short
      if (this.searchTerm.length < 1) {
        this.filteredStates = [];
        return;
      }

      // Set loading to true and clear any previous error
      this.loading = true;
      this.error = null;

      // GraphQL query to fetch the list of states
      const query = `
        query {
          states {
            id
            name
          }
        }
      `;
      // Send GraphQL request to the server
      try {
        const response = await fetch('http://localhost:8080/graphql', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({ query }),
        });


        // Parse the JSON response and update the states array
        const { data } = await response.json();
        this.states = data.states;
        this.filteredStates = this.filterStates();  
      } catch (err) {
        // Handle any errors that occur during fetching
        this.error = 'Failed to fetch states: ' + err.message;
      } finally {
        // Set loading to false once the operation completes
        this.loading = false;
      }
    },

    // Method to filter the states based on the search term
    filterStates() {
      const searchTermLower = this.searchTerm.toLowerCase();  // Convert search term to lowercase for case-insensitive matching
      return this.states.filter(state =>
        state.name.toLowerCase().startsWith(searchTermLower)  // Filter states that start with the search term
      );
    },

    // Method to highlight matching text in the state names
    highlightMatch(stateName) {
      const searchTerm = this.searchTerm;

      // Create a regular expression to match the search term in a case-insensitive way
      const regex = new RegExp(`(${searchTerm})`, 'gi');
      return stateName.replace(regex, '<strong>$1</strong>');  //Bold effect
    },

    // Method to handle when a state is selected from the dropdown
    selectState(stateName) {
      this.$emit('state-selected', stateName);  
      this.searchTerm = stateName;              
      this.filteredStates = [];                 
    }
  }
};
</script>
