<template>
  <v-container>
    <v-row>
      <v-autocomplete
        v-model="userIngredients"
        :items="possibleIngredients"
        outlined
        chips
        deletable-chips
        label="Ingredients"
        multiple
      ></v-autocomplete>
    </v-row>
    <v-row>
      <v-btn block @click="emitSearchEvent">search</v-btn>
    </v-row>
  </v-container>
</template>

<script>
import ingredients from "../utils/ingredients";
export default {
  data: () => ({
    possibleIngredients: ingredients,
    userIngredients: []
  }),
  methods: {
    emitSearchEvent: function() {
      //Fire when the search button is pressed
      //@arg The argument is an array of ingredients
      this.$emit("search", this.userIngredients);
    }
  }
};
</script>

<style>
</style>