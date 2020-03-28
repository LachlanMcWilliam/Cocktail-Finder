<template>
  <v-content>
    <v-container>
      <v-row>
        <IngredientInput v-on:search="getDrinks($event)" />
      </v-row>
      <v-row align="center" justify="center">
        <div v-for="cocktail in cocktails" :key="cocktail.idDrink">
          <CocktailDisplay v-bind:cocktail="cocktail" class="cocktailDisplay" />
        </div>
      </v-row>
    </v-container>
  </v-content>
</template>

<script>
import axios from "axios";

// @ is an alias to /src
import CocktailDisplay from "@/components/CocktailDisplay.vue";
import IngredientInput from "@/components/IngredientInput.vue";

export default {
  name: "Home",
  components: {
    CocktailDisplay,
    IngredientInput
  },
  data() {
    return {
      cocktails: [],
      ingredients: []
    };
  },
  methods: {
    /**
     * @vuese
     * used to itterate through a list of ingredients and find the cocktails that can be made for them, it also has the functionality to filter out cocktails the program already has in memory out of the search
     * @arg The argument for this method is an Array of Strings consisting of the ingredients you want scanned
     */
    getDrinks: async function(ing) {
      //eslint-disable-next-line
      console.table(ing);
      await axios
        .post("http://lvh.me:3000/", { ingredients: ing })
        .then(res => {
          //eslint-disable-next-line
          console.table(res.data);
          this.cocktails = res.data;
        })
        //eslint-disable-next-line
        .catch(err => console.log(err));
    }
  }
};
</script>

<style scoped>
.cocktailDisplay {
  margin: 20px;
}
</style>
