<template>
  <div class="container text-center mt-4">
    <h3>Trending Movies</h3>
  </div>

  <div class="container text-center my-3">
    <div v-if="pending">
      Loading...
    </div>
    <div v-else-if="error">
      <p>Error Code: {{ error.statusCode }}</p>
      <p>Error Message: {{ error.message }}</p>
    </div>
    <div v-else class="row justify-content-center">
      <div class="col-md-4 my-3" v-for="title in titleData.data" :key="title.id">
        <div class="card">
          <img @click="handleClick(title.id)"
            src="https://img.freepik.com/free-photo/view-3d-cinema-elements_23-2150720822.jpg" class="card-img-top"
            alt="...">
          <div class="card-body">
            <h5 class="card-title">{{ title.title }}</h5>
            <p class="card-text">{{ title.description }}</p>
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

const { pending, data: titleData, error } = useFetch(`${API_URL}/titles`)

</script>