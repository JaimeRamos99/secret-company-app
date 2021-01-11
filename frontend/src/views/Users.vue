<template>
  <div class="users">
    <Users/>
  </div>
</template>

<script>
// @ is an alias to /src
//import Home from "@/components/Home.vue";
import Users from "@/components/ListUsers.vue";
export default {
  name: "users",
  components: {
    Users,
  },
};
</script>