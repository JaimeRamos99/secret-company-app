<template>
  <v-row align="center" justify="center">
    <v-col class="my-10 px-1" cols="4" xs="12" sm="6" md="4" lg="3" xl="2">
      <v-date-picker v-model="date"></v-date-picker>
      <v-row class="my-10 px-1" align="center" justify="center">
        <v-btn color="blue" class="ma-2 white--text" @click="uploadData">
          Upload data
          <v-icon right dark> mdi-cloud-upload </v-icon>
        </v-btn>
      </v-row>
    </v-col>
  </v-row>
</template>

<script>
export default {
  data: () => ({
    date: new Date().toISOString().substr(0, 10),
    loader: null,
    loading: false,
  }),

  methods: {
    async uploadData() {
      let data = { date: this.date };
      const settings = {
        method: "POST",
        redirect: "follow",
        mode: "cors",
        body: JSON.stringify(data),
        headers: {
          //Accept: "application/json",
          "Content-Type": "application/json",
        },
      };
      const response = await fetch("http://localhost:81/upload", settings);
      //const resp_data = await response.json();
      console.log("http://localhost:81/uploaad");
      console.log(response);
    },
  },
};
</script>