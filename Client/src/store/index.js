import Vue from "vue";
import Vuex from "vuex";
import axios from "axios";
import cocktails from "./modules/cocktails";

Vue.use(Vuex, axios);

export default new Vuex.Store({
  modules: {
    cocktails,
  },
});
