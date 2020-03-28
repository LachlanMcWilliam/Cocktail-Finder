//import axios from 'axios';

const cocktails = {
  state: {
    cocktails: [
      {
        id: 1,
        name: "Screwdriver",
      },
      {
        id: 2,
        name: "Esspresso Martini",
      },
    ],
  },
  mutations: {},
  actions: {},
  getters: {
    allCocktails: state => state.cocktails,
  },
};

export default cocktails;
