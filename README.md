# Cocktail Finder API

### Table of Contents

You're sections headers will be used to reference location of destination.

- [About The Project](#about-the-project)
- [Description](#description)
- [TODO](#todo)
- [How To Use](#how-to-use)
- [License](#license)
- [Author Info](#author-info)

---

## About The Project

This was a project Ive wanted to do for a while now but I never set aside the time to do so. I created a similar webapp in my second year of university, but due to my lack of experience the final result was very slow and rather clunky. So Ive wanted to redo it for quite some time, and yesterday I decided to sit down and do it so I could practice my Vuejs skills and learn Golang along the way.

If you have Vuese cli installed then you can run `vuese serve --open` to view them in a web view.

---

## TODO:

- [x] Basic web front end (take in ingredients and display drinks)
- [x] Basic api functionality (Able to get the possible drinks)
- [x] Basic api concurrency
- [ ] Able to ignore certain ingredients as conditions (eg. water, ice)
- [ ] Able to search for a specific drink and view ingredients
- [ ] Web session functionality
- [ ] User logins and drink saving
- [ ] Fix whatever mess this project will become so it makes at least some sense

[Back To The Top](#read-me-template)

---

## Description

The webapp in its current form takes a list of ingredients from the user, it then sends them on to an intermediary API to process the ingredients and pull the drinks they can make from [thecocktaildb](https://www.thecocktaildb.com/) API
[Back To The Top](#read-me-template)

#### Technologies

Client built with:

- JavaScript
- VueJS
- Axios
- Vuetify

API built with:

- Golang
- Martini
- Martini Contrib CORS

[Back To The Top](#read-me-template)

---

## How To Use

You will need Golang and Node.js installed on your machine to use the app, to start the API run `go run main.go` in the root directory (it uses port 3000), you can run `go install .` to create an .exe file you can run instead. Then CD into the client directory and run `npm install` then `npm run serve` to start the web app.

[Back To The Top](#read-me-template)

---

## License

MIT License

Copyright (c) 2020 Lachlan McWilliam

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

[Back To The Top](#read-me-template)

---

## Author Info

- Linkedin - [Lachlan McWilliam](https://www.linkedin.com/in/lachlan-mcwilliam/)

[Back To The Top](#read-me-template)
