import VueRouter from "vue-router";
import Home from "../views/Home.vue";
import Upload from "../views/Upload.vue";


const routes = [
  {
    path: "/",
    name: "Home",
    component: Home
  },
  {
    path: "/upload",
    name: "Upload",
    component: Upload
  },
  /*{
    path: "/users",
    name: "Home",
    component: Home
  },
  {
    path: "/user_info ",
    name: "Home",
    component: Home
  }*/
];

const router = new VueRouter({
  routes
});

export default router;
