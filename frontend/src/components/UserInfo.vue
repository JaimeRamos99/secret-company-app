<template>
    <v-container>
      <v-row>
        <v-col
          cols="12"
          sm="6"
        >
         <v-text-field
            v-model="message"
            label="Outlined"
            outlined
            clearable
          ></v-text-field>
          <v-btn
            elevation="2"
            large
          >Search user info</v-btn>
        </v-col>
      </v-row>
    </v-container>
</template>

<script>
  export default {
    data () {
      return {
        message: 'Enter the user ID',
      }
    },
  }
</script>