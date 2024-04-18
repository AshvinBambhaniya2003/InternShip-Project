<template>
    <div class="container d-flex justify-content-center mt-4">
        <div v-if="pending">
            Loading...
        </div>
        <div v-else-if="error">
            <p>Error Code: {{ error.statusCode }}</p>
            <p>Error Message: {{ error.message }}</p>
        </div>
        <div v-else class="card text-center">
            <img src="https://t4.ftcdn.net/jpg/05/42/36/11/360_F_542361185_VFRJWpR2FH5OiAEVveWO7oZnfSccZfD3.jpg"
                class="card-img-top" alt="" />
            <div class="card-body">
                <ul class="list-group list-group-flush">
                    <li class="list-group-item text-light credit-li">
                        <h5 class="card-title">{{ credit.name }}</h5>
                    </li>
                    <li class="list-group-item text-light credit-li"><strong>Character:</strong>
                        {{ credit.character }}
                    </li>
                    <li class="list-group-item text-light credit-li"><strong>Role:</strong> {{ credit.role }}
                    </li>
                </ul>
            </div>
        </div>
    </div>
</template>

<script setup>

// Get API_URL from .env
const { api_url: API_URL } = useRuntimeConfig().public

// Get id from route params
const { titleId, creditId } = useRoute().params

// Fetch data from the API using useFetch hook 
const { pending, data: { value: { data: credit } }, error } = await useFetch(`${API_URL}/titles/${titleId}/credits/${creditId}`)

</script>

<style scoped>
.credit-li {
    background: #2a2b2e;
}
</style>