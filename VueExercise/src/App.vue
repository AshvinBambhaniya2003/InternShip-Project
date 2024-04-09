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
          <li class="nav-item">
            <a class="nav-link" @click="activeTab = 'detail'" href="#">Detail</a>
          </li>
          <li class="nav-item">
            <a class="nav-link" @click="activeTab = 'timezone'" href="#">Timezone</a>
          </li>
          <li class="nav-item">
            <a class="nav-link" @click="activeTab = 'astronomy'" href="#">Astronomy</a>
          </li>
          <li class="nav-item">
            <a class="nav-link" @click="activeTab = 'multiIpSearch'" href="#">MultiIpSearch</a>
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
        <button class="btn btn-primary btn-sm mt-2" @click="SaveData(geolocation)">Save Data</button>
      </div>
    </div>

    <div v-if="activeTab === 'detail'" class="mb-4">
      <div v-if="geolocation">
        <div class="card">
          <div class="card-header">
            <h2>{{ geolocation.city }}, {{ geolocation.country_name }}</h2>
          </div>
          <div class="card-body">
            <div class="row">
              <div class="col-md-6">
                <div class="mb-3">
                  <h3>Location Details</h3>
                  <p><strong>Continent:</strong> {{ geolocation.continent_name }}</p>
                  <p><strong>Country:</strong> {{ geolocation.country_name }}</p>
                  <p><strong>Country Code:</strong> {{ geolocation.country_code2 }}</p>
                  <p><strong>State/Province:</strong> {{ geolocation.state_prov }}</p>
                  <p><strong>City:</strong> {{ geolocation.city }}</p>
                  <p><strong>Zipcode:</strong> {{ geolocation.zipcode }}</p>
                  <p><strong>Latitude:</strong> {{ geolocation.latitude }}</p>
                  <p><strong>Longitude:</strong> {{ geolocation.longitude }}</p>
                </div>
              </div>
              <div class="col-md-6">
                <div class="mb-3">
                  <h3>Organization</h3>
                  <p><strong>ISP:</strong> {{ geolocation.isp }}</p>
                  <p><strong>Organization:</strong> {{ geolocation.organization }}</p>
                </div>
                <div class="mb-3">
                  <h3>Currency</h3>
                  <p><strong>Code:</strong> {{ geolocation.currency.code }}</p>
                  <p><strong>Name:</strong> {{ geolocation.currency.name }}</p>
                  <p><strong>Symbol:</strong> {{ geolocation.currency.symbol }}</p>
                </div>
                <div class="mb-3">
                  <img :src="geolocation.country_flag" alt="Flag" class="img-fluid">
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div v-if="activeTab === 'timezone'" class="mb-4">
      <h2>Timezone Tab</h2>
      <div v-if="timezone">
        <div>
          <p><strong>Timezone:</strong> {{ timezone.timezone }}</p>
          <p><strong>Offset:</strong> {{ timezone.timezone_offset }}</p>
          <p><strong>Offset with DST:</strong> {{ timezone.timezone_offset_with_dst }}</p>
          <p><strong>Date:</strong> {{ timezone.date }}</p>
          <p><strong>Time:</strong> {{ timezone.time_24 }}</p>
          <p><strong>Is DST:</strong> {{ timezone.is_dst }}</p>
        </div>
      </div>
    </div>

    <div v-if="activeTab === 'astronomy'" class="mb-4">
      <h2>Astronomy Tab</h2>
      <div v-if="astronomy">
        <p>Moon Rise: {{ astronomy.moonrise }}</p>
        <p>Moon Set: {{ astronomy.moonset }}</p>
        <p>Sunrise: {{ astronomy.sunrise }}</p>
        <p>Sunset: {{ astronomy.sunset }}</p>
        <p>Day Length: {{ astronomy.day_length }}</p>
      </div>
    </div>

    <div v-if="activeTab === 'multiIpSearch'">
      <div class="container mt-5">
        <div class="mb-3">
          <label for="ipAddresses">Enter IP Addresses (comma-separated):</label>
          <input class="form-control" id="ipAddresses" v-model="multipleIPs" />
        </div>
        <button class="btn btn-primary" @click="searchMultipleIPs">Search</button>
      </div>
      <div v-if="loading" class="alert alert-info">Loading...</div>
      <div v-if="errorMessage" class="alert alert-danger">{{ errorMessage }}</div>
      <div v-if="locations.length > 0" class="row mt-5">
        <div class="col-md-6">
          <div class="mb-4">
            <Map :locations="locations" />
          </div>
        </div>
        <div class="col-md-6">
          <div class="mb-4">
            <div v-for="(location, index) in locations" :key="index" class="card mb-3">
              <div class="card-header">
                {{ location.city }}, {{ location.country_name }}
                <button class="btn btn-primary btn-sm float-end" @click="SaveData(location)">Pin</button>
              </div>
              <div class="card-body">
                <p>Latitude: {{ location.latitude }}</p>
                <p>Longitude: {{ location.longitude }}</p>
              </div>
            </div>
          </div>
        </div>
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
    const geolocationInfo = await fetch(`${GEOLOCATION_API_URL}/ipgeo?apiKey=${API_KEY}&ip=${ip}`);
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

// Fetches timezone information for a given location using the IPGeolocation API.
async function fetchTimezone(location) {
  try {
    const response = await fetch(`${GEOLOCATION_API_URL}/timezone?apiKey=${API_KEY}&location=${location}`);
    if (!response.ok) {
      throw new Error('Network response was not ok');
    }
    return await response.json();
  } catch (error) {
    throw new Error('Error fetching timezone information');
  }
}

// Fetches astronomy information for a given location using the IPGeolocation API.
async function fetchAstronomy(location) {
  try {
    const response = await fetch(`${GEOLOCATION_API_URL}/astronomy?apiKey=${API_KEY}&location=${location}`);
    if (!response.ok) {
      throw new Error('Network response was not ok');
    }
    return await response.json();
  } catch (error) {
    throw new Error('Error fetching astronomy information');
  }
}

const activeTab = ref('home');
const ipAddress = ref('');
const geolocation = ref(null);
const loading = ref(false);
const errorMessage = ref('');
const timezone = ref(null)
const astronomy = ref(null);
const locations = ref([]);
const multipleIPs = ref('');

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
    timezone.value = await fetchTimezone(geolocation.value.city);
    astronomy.value = await fetchAstronomy(geolocation.value.city);

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
const SaveData = (geolocationData) => {
  fetchData(geolocationData.ip)
    .then(geolocationInfo => {
      // Handle successful response
      console.log('Data fetched successfully:', geolocationInfo);
    })
    .catch(error => {
      geolocationData.id = geolocationData.ip
      fetch(`${JSON_SERVER_URL}`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(geolocationData)
      }).catch(error => console.error('Error:', error));
    });
};

// performs a search based on the entered IP addresses, updating loading state and handling errors.
const searchMultipleIPs = async () => {
  errorMessage.value = ''
  const ipAddresses = multipleIPs.value.split(',');

  // Validate IP addresses
  const validIPs = ipAddresses.filter(ip => {
    if (!isValidIPAddress(ip.trim())) {
      errorMessage.value = 'Please enter valid ip';
    }
    return isValidIPAddress(ip.trim())
  })

  if (errorMessage.value) {
    return
  }

  loading.value = true;
  const promises = validIPs.map(async (ip) => {
    return await fetchData(ip)
      .then(multipleIPsGeolocationInfo => {
        return multipleIPsGeolocationInfo
      })
      .catch(async (error) => {
        return await fetchGeolocation(ip.trim())
      });
  })
  try {
    const results = await Promise.all(promises);
    errorMessage.value = '';
    locations.value = results;
  } catch (error) {
    locations.value = null;
    errorMessage.value = error.message;
    console.error('Error fetching geolocation information:', error);
  }
  loading.value = false;
};

// Function to validate IP address
const isValidIPAddress = (ip) => {
  const ipAddressPattern = /^(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}$/;
  return ipAddressPattern.test(ip);
};

</script>