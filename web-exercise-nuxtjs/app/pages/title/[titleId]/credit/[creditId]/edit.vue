<template>
    <div class="container">
        <h2 class="text-center mt-5 mb-3">Edit Credit</h2>
        <div v-if="successMessage" class="alert alert-success">{{ successMessage }}</div>
        <div class="card">
            <div class="card-header d-flex justify-content-end">
                <NuxtLink class="btn btn-outline-info float-right" :to="`/title/${titleId}/credit`">View All Credits
                </NuxtLink>
            </div>
            <div v-if="pending">
                Loading...
            </div>
            <div v-else-if="error">
                <p>Error Code: {{ error.statusCode }}</p>
                <p>Error Message: {{ error.message }}</p>
            </div>
            <div v-else class="card-body">
                <CreditForm :credit="credits.data" @submitForm="editCredit" />
            </div>
        </div>
    </div>
</template>

<script setup>


const { api_url: API_URL } = useRuntimeConfig().public

// Get id from route params
const { titleId, creditId } = useRoute().params

const { pending, data: credits, error } = useFetch(`${API_URL}/titles/${titleId}/credits/${creditId}`)

// Define reactive references for display message alert
const successMessage = ref('');

// Edit a credit and display success message before navigating
const editCredit = async (credit) => {
    await useFetch(`${API_URL}/titles/${titleId}/credits/${creditId}`, {
        method: 'PUT',
        body: credit
    })
    successMessage.value = 'credit edit succesfully'
    setTimeout(() => {
        successMessage.value = ''
        navigateTo(`/title/${titleId}/credit`)
    }, 2000);
}

</script>