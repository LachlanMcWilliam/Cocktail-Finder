# cocktail_finder

## Project setup

```
npm install
```

### Compiles and hot-reloads for development

```
npm run serve
```

### Compiles and minifies for production

```
npm run build
```

## Components

### Home

This view contains the functionality to search for drinks using ingredients inputed into the `IngredientInput` component. The cocktails are displayed using the `CocktailDisplay` component.

#### Methods

<!-- @vuese:Home:methods:start -->

| Method    | Description                                                                                                                                                                                                | Parameters                                                                                         |
| --------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| getDrinks | used to itterate through a list of ingredients and find the cocktails that can be made for them, it also has the functionality to filter out cocktails the program already has in memory out of the search | The argument for this method is an Array of Strings consisting of the ingredients you want scanned |

### IngredientInput

This component is a `<v-autofill>` component that takes in a list of ingredients that the user picks from the list of ingredients that are on file to ensure there are no out of bounds inputs.

#### Events

<!-- @vuese:IngredientInput:events:start -->

| Event Name | Description                            | Parameters                              |
| ---------- | -------------------------------------- | --------------------------------------- |
| search     | Fire when the search button is pressed | The argument is an array of ingredients |

<!-- @vuese:IngredientInput:events:end -->

### CocktailDisplay

This component is a `<v-card>` used to display the thumbail and name of a cocktail, it has a `CocktailInfoDisplay` component nested within it to provide the full information about the cocktail ie. how to make it and the ingredients involved.

#### Props

<!-- @vuese:CocktailDisplay:props:start -->

| Name     | Description                                 | Type     | Required | Default |
| -------- | ------------------------------------------- | -------- | -------- | ------- |
| cocktail | The cocktail object that is to be displayed | `Object` | `true`   | -       |

<!-- @vuese:CocktailDisplay:props:end -->

### CocktailInfoDisplay

This is a `<v-dialog>` with a `<v-btn>` activator (this means that it will apear as a button when you place it onto a view), it uses slots to put the information about the cocktail into the component.

#### Slots

<!-- @vuese:CocktailInfoDisplay:slots:start -->

| Name         | Description                           | Default Slot Content |
| ------------ | ------------------------------------- | -------------------- |
| thumbnail    | Cocktail Thumbnail                    | -                    |
| title        | Cocktail Title                        | -                    |
| instructions | Instructions how to make the cocktail | -                    |
| ingredients  | Cocktail ingredients and measures     | -                    |
