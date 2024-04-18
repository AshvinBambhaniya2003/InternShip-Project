<template>
  <div id="carouselExampleAutoplaying" class="carousel slide" data-bs-ride="carousel">
    <div class="carousel-inner">

      <div class="carousel-item active">
        <img
          src="https://static1.colliderimages.com/wordpress/wp-content/uploads/2024/02/godzilla-x-kong-the-new-empire-poster.jpg"
          class="object-fit-cover d-block w-100" alt="Movie Poster" height="540px">
        <div class="carousel-caption d-none d-md-block">
          <h5>Godzilla x Kong: The New Empire</h5>
          <p>Godzilla and the almighty Kong face a colossal threat hidden deep within the planet, challenging their very
            existence and the survival of the human race.</p>
        </div>
      </div>

      <div class="carousel-item">
        <img
          src="https://images.bauerhosting.com/empire/2023/12/alex-garland-civil-war.jpg?ar=16%3A9&fit=crop&crop=top&auto=format&w=1440&q=80"
          class="object-fit-cover d-block w-100" alt="Movie Poster" height="540px">
        <div class="carousel-caption d-none d-md-block">
          <h5>Civil War</h5>
          <p>SA journey across a dystopian future America, following a team of military-embedded journalists as they
            race against time to reach DC before rebel factions descend upon the White House</p>
        </div>
      </div>

      <div class="carousel-item">
        <img src="https://deadline.com/wp-content/uploads/2024/01/MCDDUPA_WB041-e1708041579639.jpg?w=681&h=383&crop=1"
          class="d-block w-100" alt="Movie Poster" height="540px">
        <div class="carousel-caption d-none d-md-block">
          <h5>Dune: Part Two</h5>
          <p>Paul Atreides unites with Chani and the Fremen while seeking revenge against the conspirators who destroyed
            his family.</p>
        </div>
      </div>

      <div class="carousel-item">
        <img
          src="https://variety.com/wp-content/uploads/2024/02/Screen-Shot-2024-02-14-at-5.49.10-PM.png?w=1000&h=563&crop=1"
          class="d-block w-100" alt="Movie Poster" height="540px">
        <div class="carousel-caption d-none d-md-block">
          <h5>Joker: Folie à Deux</h5>
          <p>Failed comedian Arthur Fleck meets the love of his life, Harley Quinn, while incarcerated at Arkham State
            Hospital.</p>
        </div>
      </div>

      <div class="carousel-item">
        <img src="https://images.thedirect.com/media/article_full/deadpool-3-set-photos.jpg" class="d-block w-100"
          alt="Movie Poster" height="540px">
        <div class="carousel-caption d-none d-md-block">
          <h5>Deadpool & Wolverine</h5>
          <p>Wolverine is recovering from his injuries when he crosses paths with the loudmouth, Deadpool. They team up
            to defeat a common enemy.</p>
        </div>
      </div>

      <div class="carousel-item">
        <img src="https://i.ytimg.com/vi/g8zxiB5Qhsc/maxresdefault.jpg" class="d-block w-100" alt="Movie Poster"
          height="540px">
        <div class="carousel-caption d-none d-md-block">
          <h5>Monkey Man</h5>
          <p>A young man ekes out a meager living in an underground fight club where, night after night, wearing a
            gorilla mask, he's beaten bloody by more popular fighters for cash.</p>
        </div>
      </div>

      <div class="carousel-item">
        <img
          src="https://images.ctfassets.net/rporu91m20dc/2a3wESrWweWSZkTbNrXo2/2c29f6de40a685e412886db7edfeb096/FO76-Wastelanders-SteelDawn_Screens_01-SteelDawn-THUMB1.jpg"
          class="d-block w-100" alt="Movie Poster" height="540px">
        <div class="carousel-caption d-none d-md-block">
          <h5>Fall Out</h5>
          <p>In a future, post-apocalyptic Los Angeles brought about by nuclear decimation, citizens must live in
            underground bunkers to protect themselves from radiation, mutants and bandits.</p>
        </div>
      </div>

      <div class="carousel-item">
        <img src="https://static1.srcdn.com/wordpress/wp-content/uploads/2021/12/Attack-on-Titan-recap-header.jpg"
          class="d-block w-100" alt="Movie Poster" height="540px">
        <div class="carousel-caption d-none d-md-block">
          <h5>Attack on Titan</h5>
          <p>Humankind constructs lofty walls to safeguard itself from immense creatures, the Titans. It battles to
            regain its land captured by the Titans which is not easy.</p>
        </div>
      </div>

      <div class="carousel-item">
        <img src="https://img10.hotstar.com/image/upload/f_auto/sources/r1/cms/prod/5631/1655631-i-96a171bab2b4"
          class="d-block w-100" alt="Movie Poster" height="540px">
        <div class="carousel-caption d-none d-md-block">
          <h5>Shōgun</h5>
          <p>Set in Japan in the year 1600, Lord Yoshii Toranaga is fighting for his life as his enemies on the Council
            of Regents unite against him, when a mysterious European ship is found marooned in a nearby fishing village.
          </p>
        </div>
      </div>

      <div class="carousel-item">
        <img
          src="https://c4.wallpaperflare.com/wallpaper/582/58/501/breaking-bad-walter-white-jesse-pinkman-wallpaper-preview.jpg"
          class="d-block w-100 object-fit-cover" alt="Movie Poster" height="540px">
        <div class="carousel-caption d-none d-md-block">
          <h5>Breaking Bad</h5>
          <p>Walter White, a chemistry teacher, discovers that he has cancer and decides to get into the meth-making
            business to repay his medical debts. His priorities begin to change when he partners with Jesse.</p>
        </div>
      </div>

    </div>
    <button class="carousel-control-prev" type="button" data-bs-target="#carouselExampleAutoplaying"
      data-bs-slide="prev">
      <span class="carousel-control-prev-icon" aria-hidden="true"></span>
      <span class="visually-hidden">Previous</span>
    </button>
    <button class="carousel-control-next" type="button" data-bs-target="#carouselExampleAutoplaying"
      data-bs-slide="next">
      <span class="carousel-control-next-icon" aria-hidden="true"></span>
      <span class="visually-hidden">Next</span>
    </button>
  </div>

  <div class="text-light">
    <div class="container-fluid text-center pt-3">
      <h5>Trending Movies</h5>
      <div v-if="moviespending">
        Loading...
      </div>
      <div v-else-if="moviesError">
        <p>Error Code: {{ moviesError.statusCode }}</p>
        <p>Error Message: {{ moviesError.message }}</p>
      </div>
      <div v-else class="row justify-content-center">
        <div class="col-md-2 my-3" v-for="title in movies" :key="title.id">
          <div class="card cardhover">
            <img @click="handleClick(title.id)"
              src="https://img.freepik.com/free-photo/view-3d-cinema-elements_23-2150720822.jpg" class="card-img-top"
              alt="Movie Poster">
            <div class="card-body">
              <h5 class="card-title">{{ title.title }}</h5>
              <p class="card-text">{{ (title.description).slice(0, 120) }}...</p>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div class="container-fluid text-center pt-3">
      <h5>Trending Show</h5>
      <div v-if="showspending">
        Loading...
      </div>
      <div v-else-if="showsError">
        <p>Error Code: {{ showsError.statusCode }}</p>
        <p>Error Message: {{ showsError.message }}</p>
      </div>
      <div v-else class="row justify-content-center">
        <div class="col-md-2 my-3" v-for="title in shows" :key="title.id">
          <div class="card cardhover">
            <img @click="handleClick(title.id)"
              src="https://img.freepik.com/free-photo/view-3d-cinema-elements_23-2150720822.jpg"
              class="object-fit-cover card-img-top" alt="Movie Poster">
            <div class="card-body">
              <h5 class="card-title">{{ title.title }}</h5>
              <p class="card-text">{{ title.description.slice(0, 120) }}...</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>

// Get API_URL from .env
const { api_url: API_URL } = useRuntimeConfig().public

const handleClick = (id) => {
  navigateTo(`title/${id}`)
}

// Fetch movie data from the API, limiting the response to the first 6 movies
const { pending: moviespending, data: movies, error: moviesError } = await useFetch(`${API_URL}/titles?title_type=MOVIE`, {
  transform: (movies) => {
    return movies.data.slice(0, 6)
  }
})

// Fetch TV show data from the API, limiting the response to the first 6 shows
const { pending: showspending, data: shows, error: showsError } = useFetch(`${API_URL}/titles?title_type=SHOW`, {
  transform: (shows) => {
    return shows.data.slice(0, 6)
  }
})


</script>

<style>

body {
  background: rgb(4, 39, 67);
  color: white;
}

.card {
  background: #2a2b2e;
}

.cardhover:hover {
  background: black;
  cursor: pointer;
  transform: scale(1.05);
}
</style>