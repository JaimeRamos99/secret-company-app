import Vue from "vue"
import VueRouter from "vue-router";
import Home from "../views/Home.vue";
import Upload from "../views/Upload.vue";
import Users from "../views/Users.vue";
import Info from "../views/User_info.vue"

Vue.use(VueRouter)
const routes = [
  {
    path: "/",
    name: "home",
    component: Home
  },
  {
    path: "/upload",
    name: "upload",
    component: Upload
  },
  {
    path: "/users",
    name: "users",
    component: Users
  },
  {
    path: "/user_info",
    name: "user_info",
    component: Info
  }
];

const router = new VueRouter({
  mode: 'history',
  routes
});

export default router;
