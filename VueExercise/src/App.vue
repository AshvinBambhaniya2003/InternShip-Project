<template>
  <nav class="navbar navbar-expand-lg bg-body-tertiary" data-bs-theme="dark">
    <div class="container">
      <a class="navbar-brand" href="#">Location Identifier App</a>
      <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarSupportedContent"
        aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"></span>
      </button>
      <div class="collapse navbar-collapse" id="navbarSupportedContent">
        <ul class="navbar-nav me-auto mb-2 mb-lg-0">
          <li class="nav-item">
            <a class="nav-link" aria-current="page" @click="activeTab = 'home'" href="#">Home</a>
          </li>
        </ul>
        <div class="d-flex" role="search" v-if="activeTab === 'home'">
          <input type="search" class="form-control me-2" v-model="ipAddress" placeholder="Search" aria-label="Search">
          <button class="btn btn-outline-primary" @click="search">Search</button>
        </div>
      </div>
    </div>
  </nav>
  <div class="container mt-5 text-center">
    <div v-if="activeTab === 'home'">
      <div v-if="loading" class="alert alert-info">Loading...</div>
      <div v-if="errorMessage" class="alert alert-danger">{{ errorMessage }}</div>

      <div v-if="geolocation" class="mb-4">
        <h2>{{ geolocation.city }}, {{ geolocation.country_name }}</h2>
        <p>Latitude: {{ geolocation.latitude }}</p>
        <p>Longitude: {{ geolocation.longitude }}</p>
        <div>
          <Map :latitude="geolocation.latitude" :longitude="geolocation.longitude" />
        </div>
        <button class="btn btn-primary btn-sm mt-2" @click="SaveData">Save Data</button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { onMounted, ref } from 'vue';
import Map from './components/Map.vue'
const GEOLOCATION_API_URL = import.meta.env.VITE_GEOLOCATION_API_URL;
const API_KEY = import.meta.env.VITE_API_KEY;
const JSON_SERVER_URL = import.meta.env.VITE_JSON_SERVER_URL;

// Fetches geolocation information for a given IP address using the IPGeolocation API.
async function fetchGeolocation(ip) {
  try {
    const geolocationInfo = await fetch(`${GEOLOCATION_API_URL}?apiKey=${API_KEY}&ip=${ip}`);
    if (!geolocationInfo.ok) {
      throw new Error('Network response was not ok');
    }
    return await geolocationInfo.json();
  } catch (error) {
    throw new Error('Error fetching geolocation information');
  }
}

// Fetches data for a given IP address from a json server, returning a Promise that resolves to the JSON response if successful, and rethrows any errors.
function fetchData(ipAddress) {
  return fetch(`${JSON_SERVER_URL}/${ipAddress}`)
    .then(geolocationInfo => {
      if (geolocationInfo.ok) {
        return geolocationInfo.json();
      } else {
        throw new Error('Network response was not ok.');
      }
    })
    .catch(error => {
      throw error; // Rethrow the error to be caught by the caller
    });
}

const activeTab = ref('home');
const ipAddress = ref('');
const geolocation = ref(null);
const loading = ref(false);
const errorMessage = ref('');

// Asynchronously performs a search based on the entered IP address, updating loading state and handling errors.
const search = async () => {
  loading.value = true;
  try {
    if (ipAddress.value === '') {
      geolocation.value = await fetchGeolocation(ipAddress.value);
      errorMessage.value = '';
    } else {
      await fetchData(ipAddress.value)
        .then(geolocationInfo => {
          geolocation.value = geolocationInfo
          errorMessage.value = '';
        })
        .catch(async (error) => {
          geolocation.value = await fetchGeolocation(ipAddress.value);
          errorMessage.value = '';
        });
    }

  } catch (error) {
    geolocation.value = null;
    errorMessage.value = error.message;
  }
  loading.value = false;
};


onMounted(() => {
  search()
})

// Save geolocation data to the json file
const SaveData = () => {
  fetchData(geolocation.value.ip)
    .then(geolocationInfo => {
      // Handle successful response
      console.log('Data fetched successfully:', geolocationInfo);
    })
    .catch(error => {
      geolocation.value.id = geolocation.value.ip
      fetch(`${JSON_SERVER_URL}`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(geolocation.value)
      }).catch(error => console.error('Error:', error));
    });
};

</script>