<template>
  <div>
    <v-container>
      <v-row class="pb-6 pt-6" align="center" justify="center">
        <v-col cols="12" sm="6">
         <v-text-field v-model="userId" label="User ID" outlined clearable></v-text-field>
         <v-row class="pb-6" align="center" justify="center">
           <v-btn v-if="loading" elevation="2" large loading></v-btn>
           <v-btn v-else elevation="2" large @click="searchUser">Search user info</v-btn>
         </v-row>
        </v-col>
      </v-row>
    </v-container>
    <v-container v-if="showInfo">
      <v-row justify="center" class="mb-6">
          <h2>Shopping History</h2>
      </v-row>
      <v-row justify="center">
          <v-expansion-panels accordion>
            <v-expansion-panel
              v-for="(item,i) in 5"
              :key="i"
            >
                <v-expansion-panel-header>Item</v-expansion-panel-header>
                <p class="text-left ml-6">{{loading}}</p>
                <v-expansion-panel-content>
                  Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.
                </v-expansion-panel-content>
            </v-expansion-panel>
          </v-expansion-panels>
      </v-row>
    </v-container>  
  </div>
</template>

<script>
  export default {
    data () {
      return {
        userId: 'Enter the user ID',
        loading: false,
        showInfo: true,
        transactions: [],
        otherUsersIp: [],
        recommendations: []
      }
    },
    methods: {
      async searchUser(){
        this.loading = true;
        const response = await fetch(`http://localhost:81/user_info/${this.userId}`);
        const resp_data = await response.json();
        this.transactions = resp_data
        this.loading = false;
        console.log(resp_data)
      }
    }
  }
</script>