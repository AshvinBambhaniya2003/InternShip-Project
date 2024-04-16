<template>
    <h2 class="text-center mt-5 mb-3">Details of Movie/Show</h2>
    <div class="container">
        <div v-if="pending">
            Loading...
        </div>
        <div v-else-if="error">
            <p>Error Code: {{ error.statusCode }}</p>
            <p>Error Message: {{ error.message }}</p>
        </div>
        <div v-else class="card">
            <div class="card-header text-center">
                <h5>{{ (titleData.title).toUpperCase() }}</h5>
            </div>
            <div class="card-body text-center">
                <p class="card-text"><strong>Type:</strong> {{ titleData.type }}</p>
                <p class="card-text"><strong>Description:</strong> {{ titleData.description }}</p>
                <p class="card-text"><strong>Release Year:</strong> {{ titleData.release_year }}</p>
                <p class="card-text"><strong>Age Certification:</strong> {{ titleData.age_certification }}</p>
                <p class="card-text"><strong>Runtime:</strong> {{ titleData.runtime }}</p>
                <p class="card-text"><strong>Genres:</strong> {{ titleData.genres }}</p>
                <p class="card-text"><strong>Production Countries:</strong> {{ titleData.production_countries }}
                </p>
                <p class="card-text"><strong>Seasons:</strong> {{ titleData.seasons }}</p>
                <p class="card-text"><strong>IMDB ID:</strong> {{ titleData.imdb_id }}</p>
                <p class="card-text"><strong>IMDB Score:</strong> {{ titleData.imdb_score }}</p>
                <p class="card-text"><strong>IMDB Votes:</strong> {{ titleData.imdb_votes }}</p>
                <p class="card-text"><strong>TMDB Popularity:</strong> {{ titleData.tmdb_popularity }}</p>
                <p class="card-text"><strong>TMDB Score:</strong> {{ titleData.tmdb_score }}</p>
            </div>

        </div>
    </div>

</template>

<script setup>

// Get id from route params
const { titleId } = useRoute().params

// Get API_URL from .env
const { api_url: API_URL } = useRuntimeConfig().public

// Fetch data from the API using useFetch hook 
const { pending, data: { value: { data: titleData } }, error } = await useFetch(`${API_URL}/titles/${titleId}`)

</script>