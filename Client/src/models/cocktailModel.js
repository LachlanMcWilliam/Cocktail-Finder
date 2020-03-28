function Cocktail(
  id,
  name,
  alcoholic,
  glass,
  instructions,
  thumbnail,
  ingredients,
  measures
) {
  this.idDrink = id;
  this.name = name;
  this.ingredientsAndMeasures = createIngredientsAndMeasureString(
    ingredients,
    measures
  );
  this.ingredients = ingredients;
  this.alcoholic = alcoholic;
  this.glass = glass;
  this.instructions = instructions;
  this.thumbnail = thumbnail;
}

export default Cocktail;

function createIngredientsAndMeasureString(ingredients, measures) {
  let output = [];
  for (let i = 0; i < ingredients.length; i++) {
    let measure = "";
    if (measures[i]) {
      measure = measures[i];
    } else {
      measure = "To preferance";
    }

    output.push({
      index: i,
      name: ingredients[i],
      amount: ingredients[i] + ": " + measure,
    });
  }
  return output;
}
