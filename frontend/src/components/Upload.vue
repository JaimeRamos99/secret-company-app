<template>
  <v-row align="center" justify="center">
    <v-col class="my-10 px-1" cols="4" xs="12" sm="6" md="4" lg="3" xl="2">
      <v-date-picker v-model="date"></v-date-picker>
      <v-row class="my-10 px-1" align="center" justify="center">
        <v-btn
          :loading="loading3"
          :disabled="loading3"
          color="blue"
          class="ma-2 white--text"
          @click="loader = 'loading3'"
        >
          Upload data
          <v-icon right dark>
            mdi-cloud-upload
          </v-icon>
        </v-btn>

        <v-snackbar v-model="snackbar">
          {{ textSnackbar }}

          <template v-slot:action="{ attrs }">
            <v-btn
              color="pink"
              text
              v-bind="attrs"
              @click="snackBarOff"
            >
              Close
            </v-btn>
          </template>
       </v-snackbar>

      </v-row>
    </v-col>
  </v-row>
  
</template>

<script>
export default {
  data: () => ({
    date: new Date().toISOString().substr(0, 10),
        loader: null,
        loading3: false,
        snackbar: false,
        textSnackbar: `Hello, I'm a snackbar`,
  }),
    watch: {
        async loader () {
          const l = this.loader
          this[l] = !this[l]

          if(this.loader){
            await this.uploadData()
          }
          
          this[l] = false
          this.loader = null
        },
    },
  methods: {
    async uploadData() {
      let data = { date: this.date };
      const settings = {
        method: "POST",
        redirect: "follow",
        mode: "cors",
        body: JSON.stringify(data),
        headers: {
          "Content-Type": "application/json",
        },
      };
      const response = await fetch("http://localhost:81/upload", settings);
      const resp_data = await response.json();
      this.snackbar = true
      this.textSnackbar = resp_data.message
    },
    snackBarOff(){
      this.snackbar = false
    }
  },
};
</script>

<style>
  .custom-loader {
    animation: loader 1s infinite;
    display: flex;
  }
  @-moz-keyframes loader {
    from {
      transform: rotate(0);
    }
    to {
      transform: rotate(360deg);
    }
  }
  @-webkit-keyframes loader {
    from {
      transform: rotate(0);
    }
    to {
      transform: rotate(360deg);
    }
  }
  @-o-keyframes loader {
    from {
      transform: rotate(0);
    }
    to {
      transform: rotate(360deg);
    }
  }
  @keyframes loader {
    from {
      transform: rotate(0);
    }
    to {
      transform: rotate(360deg);
    }
  }
</style>