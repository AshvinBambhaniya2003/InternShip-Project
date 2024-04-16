<template>
    <div class="container">
        <h2 class="text-center mt-5 mb-3">Create New Project</h2>
        <div v-if="successMessage" class="alert alert-success">{{ successMessage }}</div>
        <div class="card">
            <div class="card-header">
                <NuxtLink class="btn btn-outline-info float-right" to="/title">View All Movie/Show
                </NuxtLink>
            </div>
            <div class="card-body">
                <Form @submitForm="addTitle" />
            </div>
        </div>
    </div>
</template>

<script setup>

// Get API_URL from .env
const API_URL = useRuntimeConfig().public.api_url

// Define reactive references for display message alert
const successMessage = ref('');

// Add a new movie/show and display success message before navigating
const addTitle = async (titleData) => {
    await useFetch(`${API_URL}/titles`, {
        method: 'POST',
        body: titleData
    })

    successMessage.value = 'movie/show added successfully'
    setTimeout(() => {
        successMessage.value = ''
        navigateTo('/title')
    }, 2000);
}

</script>