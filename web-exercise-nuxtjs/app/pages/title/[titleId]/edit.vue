<template>
    <div class="container">
        <h2 class="text-center mt-5 mb-3">Edit Movie/Show</h2>
        <div v-if="successMessage" class="alert alert-success">{{ successMessage }}</div>
        <div class="card">
            <div class="card-header">
                <NuxtLink class="btn btn-outline-info float-right" to="/title">View All Movie/Show</NuxtLink>
            </div>
            <div v-if="pending">
                Loading...
            </div>
            <div v-else-if="error">
                <p>Error Code: {{ error.statusCode }}</p>
                <p>Error Message: {{ error.message }}</p>
            </div>
            <div v-else class="card-body">
                <Form :titleData="titles.data" @submitForm="editTitle" />
            </div>
        </div>
    </div>
</template>

<script setup>

// Get API_URL from .env
const { api_url: API_URL } = useRuntimeConfig().public

// Get id from route params
const { titleId } = useRoute().params

// Define reactive references for display message alert
const successMessage = ref('');

// Fetch data from the API using useFetch hook
const { pending, data: titles, error } = useFetch(`${API_URL}/titles/${titleId}`)

// Edit a movie/show and display success message before navigating
const editTitle = async (titleData) => {
    await useFetch(`${API_URL}/titles/${titleData.id}`, {
        method: 'PUT',
        body: titleData
    })
    successMessage.value = 'movie/show edit succesfully'
    setTimeout(() => {
        successMessage.value = ''
        navigateTo('/title')
    }, 2000);
}

</script>
