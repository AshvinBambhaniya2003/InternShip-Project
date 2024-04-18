<template>
    <div class="container">
        <h5 class="text-center my-3">List Credits</h5>
        <div v-if="deleteMessage" class="alert alert-success">{{ deleteMessage }}</div>
        <div class="card">
            <div class="card-header d-flex justify-content-between align-items-center">
                <NuxtLink :to="`/title/${titleId}/credit/add`" class="btn btn-outline-primary">Create New Credits
                </NuxtLink>
            </div>
            <div class="card-body">
                <table class="table table-bordered text-center text-light">
                    <thead>
                        <tr>
                            <th>Name</th>
                            <th>Character</th>
                            <th>Role</th>
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
                        <tr v-for="credit in creditsData.data" :key="credit.id">
                            <td>{{ credit.name }}</td>
                            <td>{{ credit.character }}</td>
                            <td>{{ credit.role }}</td>
                            <td>
                                <NuxtLink :to="`/title/${credit.title_id}/credit/${credit.id}`"
                                    class="btn btn-outline-light btn-sm mx-1">Detail
                                </NuxtLink>
                                <NuxtLink :to="`/title/${credit.title_id}/credit/${credit.id}/edit`"
                                    class="btn btn-outline-primary btn-sm mx-1">
                                    Edit
                                </NuxtLink>
                                <button @click="deleteCredit(credit.id)" class="btn btn-outline-danger btn-sm mx-1">
                                    Delete
                                </button>
                            </td>
                        </tr>
                    </tbody>
                </table>
            </div>
        </div>
    </div>
</template>

<script setup>

// Function to handle deletion of a title by ID
const deleteCredit = async (creditId) => {
    await useFetch(`${API_URL}/titles/${titleId}/credits/${creditId}`, {
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

// Get id from route params
const { titleId } = useRoute().params

// Define reactive references for the title object and delete messages
const deleteMessage = ref('');

// Fetch data from the API using useFetch hook 
const { refresh, pending, data: creditsData, error } = useFetch(`${API_URL}/titles/${titleId}/credits`)

</script>