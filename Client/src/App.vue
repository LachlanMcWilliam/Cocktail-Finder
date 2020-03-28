<template>
  <v-app>
    <v-content>
      <Navbar />
      <Home />
    </v-content>
  </v-app>
</template>

<script>
import Home from "./views/Home";
import Navbar from "./components/Navbar";

export default {
  name: "App",

  components: {
    Navbar,
    Home
  }
};
</script>
