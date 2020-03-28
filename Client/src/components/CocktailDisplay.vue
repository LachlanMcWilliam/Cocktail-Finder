<template>
  <v-card max-width="400" ma-5>
    <v-img height="400px" v-bind:src="cocktail.thumbnail"></v-img>
    <v-card-title class="black--text align-center">{{ cocktail.name }}</v-card-title>
    <v-card-actions class="text-center">
      <CocktailInfoDisplay>
        <v-img slot="thumbnail" v-bind:src="cocktail.thumbnail"></v-img>
        <v-card-title slot="title" primary-title>{{cocktail.name}}</v-card-title>
        <v-card-text slot="instructions" class="pt-0 pb-0">{{cocktail.instructions}}</v-card-text>
        <div slot="ingredients" v-for="ingredient in cocktail.ingredients" :key="ingredient.name">
          <v-card-text class="pt-0 pb-0">{{ingredient.name}}: {{ingredient.measure}}</v-card-text>
        </div>
      </CocktailInfoDisplay>
    </v-card-actions>
  </v-card>
</template>

<script>
import CocktailInfoDisplay from "./CocktailInfoDialog";
//This component displays the thumbnail and name of the cocktail and contains a "CocktailInfoDisplay" componant as a child componant
export default {
  name: "CocktailDisplay",
  components: {
    CocktailInfoDisplay
  },
  props: {
    //The cocktail object that is to be displayed
    cocktail: {
      type: Object,
      required: true
    }
  }
};
</script>
