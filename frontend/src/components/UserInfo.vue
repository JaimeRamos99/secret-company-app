<template>
  <div>
    <v-parallax height="200" src="https://cdn.vuetifyjs.com/images/parallax/material2.jpg"></v-parallax>
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
      <v-row justify="center" class="mb-15">
          <h1 class="pr-6">Name: {{userName}}</h1>
          <h1>ID: {{ID}}</h1>
      </v-row>
      <v-row justify="center" class="mb-6">
          <h2>Shopping History</h2>
      </v-row>
      <v-row justify="center">
          <v-expansion-panels accordion>
            <v-expansion-panel
              v-for="(item,i) in transactions"
              :key="i"
            >
                <v-expansion-panel-header>Transaction ID: {{item.transactionId}}</v-expansion-panel-header>
                <p class="text-left ml-6">IP: {{item.ip}}</p>
                <p class="text-left ml-6">Device: {{item.device}}</p>
                <v-expansion-panel-content>
                  Products:
                  <v-expansion-panels inset>
                      <v-expansion-panel
                        v-for="(product,i) in item.includes"
                        :key="i"
                      >
                          <v-expansion-panel-header>ID: {{product.productId}}</v-expansion-panel-header>
                          <p class="text-left ml-6">Name: {{product.productName}}</p>
                          <p class="text-left ml-6">Price: {{product.productPrice}}</p>
                      </v-expansion-panel>
                  </v-expansion-panels>
                </v-expansion-panel-content>
            </v-expansion-panel>
          </v-expansion-panels>
      </v-row>
      <v-row justify="center" class="mb-6 mt-12">
          <h2>Other users using the same IP's</h2>
      </v-row>
      <v-row justify="center">
          <v-expansion-panels accordion>
            <v-expansion-panel
              v-for="(key,i) in usedIPS"
              :key="i"
            >
                <v-expansion-panel-header class="font-weight-bold">{{key[0]}}</v-expansion-panel-header>
                <v-expansion-panel-content>
                  Used in/by:
                  <v-expansion-panels inset>
                      <v-expansion-panel
                        v-for="(info,i) in positions[key[1]]"
                        :key="i"
                      >
                          <v-expansion-panel-header> Transaction ID: {{info.transactionId}}</v-expansion-panel-header>
                          <p class="text-left ml-6">user ID: {{info.madeBy[0].userId}}</p>
                          <p class="text-left ml-6">user name: {{info.madeBy[0].userName}}</p>
                      </v-expansion-panel>
                  </v-expansion-panels>
                </v-expansion-panel-content>
            </v-expansion-panel>
          </v-expansion-panels>
      </v-row>

      <v-row justify="center" class="mb-6 mt-12">
          <h2>Best selling dishes of the day</h2>
      </v-row>

      <v-row>
        <v-col
          v-for="(rec,i) in recommendations"
          :key="i"
          class="d-flex child-flex"
          cols="4"
        >
        <v-card>
          <v-img
            :src="`https://picsum.photos/500/300?image=${i * 5 + 10}`"
            :lazy-src="`https://picsum.photos/10/6?image=${i * 5 + 10}`"
            aspect-ratio="1"
            class="grey lighten-2"
          >
            <template v-slot:placeholder>
              <v-row
                class="fill-height ma-0"
                align="center"
                justify="center"
              >
                <v-progress-circular
                  indeterminate
                  color="grey lighten-5"
                ></v-progress-circular>
              </v-row>
            </template>
          </v-img>
           <v-card-title class="title">
              {{rec.name}}
            </v-card-title>
        </v-card>
        </v-col>
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
        showInfo: false,
        userName: "",
        ID: 0,
        usedIPS: [],
        positions : [],
        transactions: [],
        otherUsersIp: [],
        recommendations: []
      }
    },
    methods: {
      async searchUser(){
        this.loading = true;
        const response = await fetch(`${process.env.VUE_APP_API_URL}user_info/${this.userId}`);
        if(response){
          const resp_data = await response.json();
          if (resp_data && resp_data.generalInfo.info){
            this.loading = false;
            this.showInfo = true;
            let data = await JSON.parse(JSON.stringify(resp_data))
            this.userName = data.generalInfo.info[0].userName;
            this.ID = data.generalInfo.info[0].userId;
            this.transactions = data.generalInfo.info[0].transactions
            this.filterArray(data.sameIps.second_stage)
            this.recommendations = data.recommendedProducts
            return
          }
          this.loading = false
          
        }
      },
      filterArray(inputArray){
        let ips = new Map();
        let positions = [];
        let counter_ips = -1;
        let pos;
        inputArray.forEach(tr => {
          pos = ips.get(tr.ip);
          if(tr.madeBy[0].userId != this.userId){
              if (pos>=0){
                positions[pos].push(tr)
                return
              }
              counter_ips++;
              ips.set(tr.ip, counter_ips);
              positions[counter_ips] = new Array();
              positions[counter_ips].push(tr);
          }
        });
        this.usedIPS = Array.from(ips);
        this.positions = positions;
      }

    }
  }
</script>