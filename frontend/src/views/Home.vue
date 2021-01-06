<template>
  <div class="home">
    <Upload />
  </div>
</template>

<script>
// @ is an alias to /src
//import Home from "@/components/Home.vue";
import Upload from "@/components/Upload.vue";
export default {
  name: "Home",
  components: {
    /*Home*/ Upload,
  },
};
</script>
