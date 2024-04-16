<template>
    <form v-on:submit.prevent="handleSubmit" class="row g-3">

        <!-- Title -->
        <div class="col-md-12 mb-2">
            <label for="title" class="form-label">Title</label>
            <input type="text" v-model="title.title" class="form-control" id="title" name="title" required>
        </div>

        <!-- Description -->
        <div class="col-md-12 mb-2">
            <label for="description" class="form-label">Description</label>
            <input type="text" v-model="title.description" class="form-control" id="description" name="description"
                required>
        </div>

        <!-- Type -->
        <div class="col-md-6 mb-2">
            <label for="type" class="form-label">Type</label>
            <select class="form-select " v-model="title.type" id="type" name="type" required>
                <option value="">Select type</option>
                <option value="MOVIE">Movie</option>
                <option value="SHOW">Show</option>
            </select>
        </div>

        <!-- Age Certification -->
        <div class="col-md-6 mb-2">
            <label for="ageCertification" class="form-label">Age Certification</label>
            <select class="form-select" v-model="title.age_certification" id="ageCertification" name="ageCertification"
                required>
                <option value="">Select age certification</option>
                <option value="G">G</option>
                <option value="PG">PG</option>
                <option value="PG-13">PG-13</option>
                <option value="R">R</option>
                <option value="NC-17">NC-17</option>
                <option value="Not Rated">Not Rated</option>
            </select>
        </div>

        <!-- Production Country -->
        <div class="col-md-6 mb-2">
            <label for="productionCountry" class="form-label">Production Country</label>
            <input type="text" v-model="title.production_countries" class="form-control" id="productionCountry"
                name="productionCountry" required>
        </div>

        <!-- Genres -->
        <div class="col-md-6 mb-2">
            <label for="genres" class="form-label">Genres</label>
            <input type="text" v-model="title.genres" class="form-control" id="genres" name="genres" required>
        </div>

        <!-- Release Year -->
        <div class="col-md-4 mb-2">
            <label for="releaseYear" class="form-label">Release Year</label>
            <input type="number" v-model="title.release_year" class="form-control" id="releaseYear" name="releaseYear"
                required>
        </div>

        <!-- Runtime -->
        <div class="col-md-4 mb-2">
            <label for="runtime" class="form-label">Runtime</label>
            <input type="number" v-model="title.runtime" class="form-control" id="runtime" name="runtime" required>
        </div>

        <!-- Seasons -->
        <div class="col-md-4 mb-2">
            <label for="seasons" class="form-label">Seasons</label>
            <input type="number" v-model="title.seasons" class="form-control" id="seasons" name="seasons">
        </div>

        <!-- IMDB ID -->
        <div class="col-md-4 mb-2">
            <label for="imdbID" class="form-label">IMDB ID</label>
            <input type="text" v-model="title.imdb_id" class="form-control" id="imdbID" name="imdbID" required>
        </div>

        <!-- IMDB Score -->
        <div class="col-md-4 mb-2">
            <label for="imdbScore" class="form-label">IMDB Score</label>
            <input type="number" v-model="title.imdb_score" class="form-control" id="imdbScore" name="imdbScore"
                step="0.01" required>
        </div>

        <!-- IMDB Votes -->
        <div class="col-md-4 mb-2">
            <label for="imdbVotes" class="form-label">IMDB Votes</label>
            <input type="number" v-model="title.imdb_votes" class="form-control" id="imdbVotes" name="imdbVotes"
                step="0.01" required>
        </div>

        <!-- TMDB Popularity -->
        <div class="col-md-6 mb-2">
            <label for="tmdbPopularity" class="form-label">TMDB Popularity</label>
            <input type="number" v-model="title.tmdb_popularity" class="form-control" id="tmdbPopularity"
                name="tmdbPopularity" step="0.01" required>
        </div>

        <!-- TMDB Score -->
        <div class="col-md-6 mb-2">
            <label for="tmdbScore" class="form-label">TMDB Score</label>
            <input type="number" v-model="title.tmdb_score" class="form-control" id="tmdbScore" name="tmdbScore"
                step="0.01" required>
        </div>

        <div class="col-md-12 mb-2">
            <button type="submit" class="btn btn-primary">Submit</button>
        </div>
    </form>
</template>

<script setup>

// Define reactive reference for title
const title = reactive({})

// Define props
const props = defineProps({
    titleData: {
        type: Object
    }
})

// Assign titleData to title if available
if (props.titleData) {
    Object.assign(title, props.titleData)
}

// Define emits for editTitle and addTitle
const emit = defineEmits(['submitForm'])

// Handle form submission, Emit appropriate event based on whether titleData is available, Reset title object
const handleSubmit = () => {
    emit('submitForm', title)
}

</script>