<template>
    <div class="container">
        <h2>Geolocation Data</h2>
        <input type="text" v-model="searchQuery" placeholder="Search..." class="form-control mb-3" />
        <div v-if="loading" class="alert alert-info">Loading...</div>
        <div v-if="errorMessage" class="alert alert-danger">{{ errorMessage }}</div>

        <div v-if="geolocationData.length === 0">
            No Geolocation Data Found.
        </div>

        <div v-else>
            <table class="table table-striped-columns table-bordered">
                <thead>
                    <tr>
                        <th>City</th>
                        <th>Country</th>
                        <th>State/Province</th>
                        <th>Latitude</th>
                        <th>Longitude</th>
                        <th>ISP</th>
                        <th>Organization</th>
                    </tr>
                </thead>

                <tbody v-for="location in filteredPinnedLocations " :key="location.id">
                    <tr>
                        <td>{{ location.city }}</td>
                        <td>{{ location.country_name }}</td>
                        <td>{{ location.state_prov }}</td>
                        <td>{{ location.latitude }}</td>
                        <td>{{ location.longitude }}</td>
                        <td>{{ location.isp }}</td>
                        <td>{{ location.organization }}</td>
                    </tr>
                </tbody>
            </table>
        </div>
    </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue';
const JSON_SERVER_URL = import.meta.env.VITE_JSON_SERVER_URL;

// Fetches data from a json server, returning a Promise that resolves to the JSON response if successful, and rethrows any errors.
async function fetchGeolocationData() {
    try {
        const geolocationInfo = await fetch(`${JSON_SERVER_URL}`);
        if (!geolocationInfo.ok) {
            throw new Error('Network response was not ok');
        }
        return await geolocationInfo.json();
    } catch (error) {
        throw new Error('Error fetching geolocation information', error);
    }

}

const searchQuery = ref('');
const geolocationData = ref([])
const loading = ref(false);
const errorMessage = ref('');

// Filter geolocationData based on the search query
const filteredPinnedLocations = computed(() => {
    return geolocationData.value.filter(location => {
        const search = searchQuery.value.toLowerCase();
        return (
            location.city.toLowerCase().includes(search) ||
            location.country_name.toLowerCase().includes(search)
        );
    });
});

// List all Geolocation locations asynchronously and update loading state
const listGeolocationData = async () => {
    loading.value = true;
    try {
        geolocationData.value = await fetchGeolocationData()
    } catch (error) {
        errorMessage.value = error.message;
    }
    loading.value = false;
}

onMounted(() => {
    listGeolocationData();
})

</script>