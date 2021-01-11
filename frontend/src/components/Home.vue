<template>
  <div>
    <v-parallax height="200" src="https://cdn.vuetifyjs.com/images/parallax/material2.jpg"></v-parallax>
    <v-container>
        <v-row class="text-center">
          <v-col cols="12">
            <v-img
              :src="require('../assets/logo.svg')"
              class="my-3"
              contain
              height="200"
            />
          </v-col>

          <v-col class="mb-4">
            <p class="subheading font-weight-regular">
              This test was performed by
              <a href="https://github.com/JaimeRamos99" target="_blank"
                >Jaime Ramos Rossetes</a
              >
            </p>
          </v-col>

          <v-col class="mb-5" cols="12">
            <h2 class="headline font-weight-bold mb-3">
              Explore the funcionalities
            </h2>
              <div class="headline text-xs-center pa-5"></div>

                <v-row class="content-center">
                  <v-col cols="12">
                    <v-btn color="blue" class="ma-2" value="load" to="/upload">
                      <span>Upload data</span>
                      <v-icon>cloud_upload</v-icon>
                    </v-btn>
                    <v-btn color="blue" class="ma-2" value="users" to="users">
                      <span>Users</span>
                      <v-icon>supervised_user_circle</v-icon>
                    </v-btn>

                    <v-btn color="blue" class="ma-2" value="user-info" to="/user_info">
                      <span>User info</span>
                      <v-icon>info</v-icon>
                    </v-btn>
                  </v-col>
                </v-row>
          </v-col>
        </v-row>
      </v-container>
  </div>
 
</template>

<script>
export default {
  data() {
    return {
      bottomNav: "recent",
    };
  },
};
</script>
