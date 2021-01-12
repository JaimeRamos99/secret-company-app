<template>
  <div>
    <v-parallax height="200" src="https://cdn.vuetifyjs.com/images/parallax/material2.jpg"></v-parallax>
    <div class="text-center">
          <v-row class="my-10 px-1" align="center" justify="center">
            <v-col  cols="2" xs="12" sm="8" md="4" lg="4" xl="4">
              <v-pagination v-model="page" :length="length"></v-pagination>
            </v-col> 
          </v-row>
        <div>
          <v-col>
              <v-row align="center" justify="center">
                <v-list three-line >
                    <template v-for="(usr, index) in visiblePages">

                        <v-divider
                        v-if="usr.divider"
                        :key="index"
                        :inset="usr.inset"
                        ></v-divider>

                        <v-list-item
                        v-else
                        :key="usr.title"
                        >
                        <v-list-item-avatar>
                            <v-img :src="avatar"></v-img>
                        </v-list-item-avatar>

                        <v-list-item-content>
                            <v-list-item-title v-html="'Nombre:  ' + usr.userName"></v-list-item-title>
                            <v-list-item-title v-html="'ID:  ' + usr.userId"></v-list-item-title>
                            <v-list-item-subtitle v-html="'Age: ' + usr.userAge"></v-list-item-subtitle>
                        </v-list-item-content>
                        </v-list-item>
                    </template>
                </v-list>
              </v-row>
          </v-col>   
        </div>
      </div>
  </div>  
</template>

<script>
  export default {
    data: () => ({
      page:1,
      length:0,
      perPage:5,
      users:[],
      selected:[],
      avatar: 'https://cdn.vuetifyjs.com/images/lists/1.jpg',
    }),
    methods: {
        async fetchUsers(){
            const response = await fetch(`${process.env.VUE_APP_API_URL}users`);
            const resp_data = await response.json();
            if (resp_data){
              this.users = resp_data.Users
              console.log(resp_data.Users)
              if(resp_data.Users){
                this.length = parseInt(Math.ceil(resp_data.Users.length/5))
                let left = this.perPage*(this.page-1)
                let right = left +  this.perPage
                this.selected = this.users.slice(left,right)
              }
            }
        },
    },
    computed: {
        visiblePages () {
            let left = this.perPage*(this.page - 1)
            let right = left + this.perPage
            return this.users.slice(left,right)
        }
    },
    mounted(){
      console.log("hhh")
      this.fetchUsers()
    },
  }
</script>

<style>
</style>