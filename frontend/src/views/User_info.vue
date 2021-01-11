<template>
  <div class="user_info">
    <Info/>
  </div>
</template>

<script>
// @ is an alias to /src
//import Home from "@/components/Home.vue";
import Info from "@/components/UserInfo.vue";
export default {
  name: "User_info",
  components: {
    Info,
  },
};
</script>