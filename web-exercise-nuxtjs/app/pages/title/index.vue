<template>
    <div class="container">
        <h5 class="text-center my-3">List Movie/Show</h5>
        <div v-if="deleteMessage" class="alert alert-success">{{ deleteMessage }}</div>
        <div class="card">
            <div class="card-header d-flex justify-content-between align-items-center">
                <NuxtLink to="/title/add" class="btn btn-outline-primary">Create New Movie/Show</NuxtLink>
                <div class="row">
                    <div class="col-4">
                        <input type="text" v-model="title.title" class="form-control" id="title" name="title"
                            placeholder="Enter title">
                    </div>
                    <div class="col-5">
                        <select class="form-control" v-model="title.type" id="type" name="type">
                            <option value="">Select type</option>
                            <option value="MOVIE">Movie</option>
                            <option value="SHOW">Show</option>
                        </select>
                    </div>
                    <div class="col-3">
                        <button @click="filterTitle" class="btn btn-primary">Search</button>
                    </div>
                </div>
            </div>
            <div class="card-body">
                <table class="table table-bordered text-center text-light">
                    <thead>
                        <tr>
                            <th>Name</th>
                            <th>type</th>
                            <th>Age Certification</th>
                            <th>Seasons</th>
                            <th>Release Year</th>
                            <th>Runtime</th>
                            <th>IMDB Score</th>
                            <th width="240px">Action</th>
                        </tr>
                    </thead>
                    <div v-if="pending">
                        Loading...
                    </div>
                    <div v-else-if="error">
                        <p>Error Code: {{ error.statusCode }}</p>
                        <p>Error Message: {{ error.message }}</p>
                    </div>
                    <tbody v-else>
                        <tr v-for="title in titleData.data" :key="title.id">
                            <td>{{ title.title }}</td>
                            <td>{{ title.type }}</td>
                            <td>{{ title.age_certification }}</td>
                            <td>{{ title.seasons }}</td>
                            <td>{{ title.release_year }}</td>
                            <td>{{ title.runtime }}</td>
                            <td>{{ title.imdb_score }}</td>
                            <td>
                                <NuxtLink :to="`/title/${title.id}`" class="btn btn-outline-light btn-sm mx-1">Detail
                                </NuxtLink>
                                <NuxtLink :to="`/title/${title.id}/edit`" class="btn btn-outline-primary btn-sm mx-1">
                                    Edit
                                </NuxtLink>
                                <button @click="deleteTitle(title.id)" class="btn btn-outline-danger btn-sm mx-1">
                                    Delete
                                </button>
                            </td>
                        </tr>
                    </tbody>
                </table>
            </div>
            <div v-if="pending">
                Loading...
            </div>
            <div v-else-if="error">
                <p>Error Code: {{ error.statusCode }}</p>
                <p>Error Message: {{ error.message }}</p>
            </div>
            <div v-else aria-label="Page navigation">
                <ul class="pagination justify-content-center">
                    <li class="page-item" :class="{ disabled: page === 1 }">
                        <NuxtLink :to="{ path: route.path, query: { ...route.query, page: page - 1 } }"
                            class="page-link">Previous</NuxtLink>
                    </li>
                    <li class="page-item"
                        :class="{ disabled: (titleData.data === null) || (titleData.data.length < 10) }">
                        <NuxtLink :to="{ path: route.path, query: { ...route.query, page: page + 1 } }"
                            class="page-link">Next</NuxtLink>
                    </li>
                </ul>
            </div>
        </div>
    </div>
</template>

<script setup>

// Function to filter titles based on title name and type
const filterTitle = async () => {
    const query = {}
    if (title.value.title) {
        query.title_name = title.value.title
    }
    if (title.value.type) {
        query.title_type = title.value.type
    }
    navigateTo({ path: '/title', query })
}

// Function to handle deletion of a title by ID
const deleteTitle = async (id) => {

    await useFetch(`${API_URL}/titles/${id}`, {
        method: 'DELETE'
    })
    deleteMessage.value = 'Delete succesfully'

    setTimeout(() => {
        deleteMessage.value = ''
    }, 2000);
    refresh()
}

// Get API_URL from .env
const { api_url: API_URL } = useRuntimeConfig().public

// Importing the useRoute function from Nuxt.js to access route parameters
const route = useRoute()

// Compute the current page, title name, and title type from the route parameters
const page = computed(() => Number(route.query.page) || 1)
const title_name = computed(() => (route.query.title_name) || '')
const title_type = computed(() => (route.query.title_type) || '')

// Define reactive references for the title object and delete messages
const title = ref({ type: "" })
const deleteMessage = ref('');

// Fetch data from the API using useFetch hook with appropriate query parameters
const { refresh, data: titleData, error, pending } = useFetch(`${API_URL}/titles`, {
    query: { page, title_name, title_type }
})

</script>