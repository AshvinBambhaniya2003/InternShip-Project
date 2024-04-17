<template>
    <div class="container">
        <h2 class="text-center mt-5 mb-3">Create New Credit</h2>
        <div v-if="successMessage" class="alert alert-success">{{ successMessage }}</div>
        <div class="card">
            <div class="card-header d-flex justify-content-end">
                <NuxtLink class="btn btn-outline-info float-right" :to="`/title/${titleId}/credit`">View All Credits
                </NuxtLink>
            </div>
            <div class="card-body">
                <CreditForm @submitForm="addCredit" />
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

// Add a new credits and display success message before navigating
const addCredit = async (credit) => {
    await useFetch(`${API_URL}/titles/${titleId}/credits`, {
        method: 'POST',
        body: credit
    })

    successMessage.value = 'credit added successfully'
    setTimeout(() => {
        successMessage.value = ''
        navigateTo(`/title/${titleId}/credit`)
    }, 2000);
}

</script>