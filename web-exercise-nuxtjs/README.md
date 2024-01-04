# Netflix Front-end App

- Add Wireframe PDF
- For APIs - Use your developed backend apis

## Management UI with Nuxt3

Look at the [Nuxt 3 documentation](https://nuxt.com/docs/getting-started/introduction) to learn more.


## Installation

- Node.js - v18.0.0 or newer

## Setup

```bash
# clone repo
git clone https://git.pride.improwised.dev/Onboarding-2024/Ashvin-Bambhaniya.git
```

## Build Setup

For UI everything is under /web-exercise-nuxtjs/app/ folder, so go to /web-exercise-nuxtjs/app/ folder.

* Go to `/web-exercise-nuxtjs/app/` folder. 

* copy `.env.example` to `.env`. (skip if already done)

## .env.example

* copy `.env.example` to `.env`.
```
MODE=development
BASE_URL=
```
* `MODE`: This will indicate application state.
* `BASE_URL`: You can specify your application URL.

* run following commands

``` bash
# Make sure to install the dependencies
$ npm install 
```

## Development Server

Start the development server on http://localhost:3000

```bash
npm run dev
```

## Included Packages

- Bootstrap - for style
- Pinia - for State management
- Eslint - for on commit 
- Prettier - for on commit
- Vitest - for testcases
- Husky - for pre-commit hooks
- Commitlint - for github commit lint




## Store

This directory contains your Pinia Store files.

for use pinia first define pinia in nuxt config file

More information about the usage of this directory in [documentation]( https://pinia.vuejs.org/ssr/nuxt.html).

* Go to `/app/store` folder, this folder includes,

    * `actions`: Includes all the requests of project.
    * `getters`: pinia allows us to define "getters" in the store. You can think of them as computed properties for stores.
    *  `state`: Includes all the variable of the project.



## Pinia Testcases
This directory includes the test cases. *
More information to write test cases you can visit this [documentation](https://vitest.dev/guide/)
### for testcase run
```
npm run test
```

## Other Configuration file

Please check husky folder/files,.eslintrc.js,.prettierrc,tsconfig.ts,vitest.config.ts,plugins folder/files

## JSON Server

[Documentation](https://github.com/typicode/json-server)

json server run
```
npx json-server --watch data/db.json --port 5000
```

## Run via Docker

```docker-compose up --build```

It will run nuxt application and json-server for fake REST API (https://www.npmjs.com/package/json-server)


## Notes:

As per your project requirement you can change and remove unnecessary functionalities

## For more info check below issue

* [Web Excercise - NuxtJs](https://git.pride.improwised.dev/Onboarding-2024/Ashvin-Bambhaniya/issues/4)

