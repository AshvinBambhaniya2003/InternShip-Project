<template>
    <div class="container-fluid">
        <h5 class="text-center my-3">Details of Movie/Show</h5>
        <div v-if="titlePending">
            Loading...
        </div>
        <div v-else-if="titleError">
            <p>Error Code: {{ titleError.statusCode }}</p>
            <p>Error Message: {{ titleError.message }}</p>
        </div>
        <div v-else class="card">
            <div class="card-header text-center">
                <h5>{{ (title.title).toUpperCase() }}</h5>
            </div>
            <div class="card-body">
                <div class="row">
                    <div class="col-md-6">
                        <p class="card-text"><strong>Type:</strong> {{ title.type }}</p>
                        <p class="card-text"><strong>Description:</strong> {{ title.description }}</p>
                        <p class="card-text"><strong>Age Certification:</strong> {{ title.age_certification }}</p>
                        <p class="card-text"><strong>Genres:</strong> {{ title.genres }}</p>
                        <p class="card-text"><strong>Production Countries:</strong> {{ title.production_countries }}
                        </p>
                        <p class="card-text"><strong>Release Year:</strong> {{ title.release_year }}</p>
                    </div>
                    <div class="col-md-6">

                        <p class="card-text"><strong>Runtime:</strong> {{ title.runtime }}</p>
                        <p class="card-text"><strong>Seasons:</strong> {{ title.seasons }}</p>
                        <p class="card-text"><strong>IMDB ID:</strong> {{ title.imdb_id }}</p>
                        <p class="card-text"><strong>IMDB Score:</strong> {{ title.imdb_score }}</p>
                        <p class="card-text"><strong>IMDB Votes:</strong> {{ title.imdb_votes }}</p>
                        <p class="card-text"><strong>TMDB Popularity:</strong> {{ title.tmdb_popularity }}</p>
                        <p class="card-text"><strong>TMDB Score:</strong> {{ title.tmdb_score }}</p>
                    </div>
                </div>
            </div>
        </div>

        <div class="container-fluid text-center pt-3">
            <h5>Actors and Directors</h5>
            <div v-if="crditsPending">
                Loading...
            </div>
            <div v-else-if="crditsError">
                <p>Error Code: {{ crditsError.statusCode }}</p>
                <p>Error Message: {{ crditsError.message }}</p>
            </div>
            <div v-else-if="credits.length === 0">
                No any Credits
            </div>
            <div v-else class="row justify-content-center">
                <div class="col-md-2 my-3" v-for="credit in credits" :key="credit.id">
                    <div class="card cardhover">
                        <img @click="HandleClick(credit.id)"
                            src="https://t4.ftcdn.net/jpg/05/42/36/11/360_F_542361185_VFRJWpR2FH5OiAEVveWO7oZnfSccZfD3.jpg"
                            class="card-img-top" alt="...">
                        <div class="card-body">
                            <h5 class="card-title">{{ credit.name }}</h5>
                            <p class="card-text">Character: {{ credit.character }}</p>
                            <p class="card-text">Role: {{ credit.role }}</p>
                        </div>
                    </div>
                </div>
            </div>
            <div v-if="credits && credits.length !== 0" class="text-center mb-2">
                <NuxtLink :to="`/title/${titleId}/credit`" class="btn btn-outline-light btn-sm mx-1">All Credits
                </NuxtLink>
            </div>
            <div v-else class="text-center mt-2">
                <NuxtLink :to="`/title/${titleId}/credit/add`" class="btn btn-outline-light btn-sm mx-1">Add New Credit
                </NuxtLink>
            </div>
        </div>
    </div>

</template>

<script setup>

const HandleClick = (creditId) => {
    navigateTo(`/title/${titleId}/credit/${creditId}`)
}

// Get id from route params
const { titleId } = useRoute().params

// Get API_URL from .env
const { api_url: API_URL } = useRuntimeConfig().public

// Fetch data from the API using useFetch hook 
const { pending: titlePending, data: { value: { data: title } }, error: titleError } = await useFetch(`${API_URL}/titles/${titleId}`)

// Fetch credits data for a specific title from the API, limiting the response to the first 6 credits
const { pending: crditsPending, data: credits, error: crditsError } = useFetch(`${API_URL}/titles/${titleId}/credits`, {
    transform: (credits) => {
        if (credits.data) {
            return credits.data.slice(0, 6)
        }
        return []
    }
})

</script>