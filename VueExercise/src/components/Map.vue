<template>
    <div ref="mapContainer" class="mapContainer"></div>
</template>

<script setup>
import { onMounted, ref, watch } from 'vue';
import L from 'leaflet'
import '../assets/map.css'

const props = defineProps({
    latitude: {
        type: String,
        required: true
    },
    longitude: {
        type: String,
        required: true
    }
})

const map = ref();
const mapContainer = ref();
const marker = ref()

onMounted(() => {
    map.value = L.map(mapContainer.value).setView([props.latitude, props.longitude], 13);
    L.tileLayer('https://tile.openstreetmap.org/{z}/{x}/{y}.png', {
        maxZoom: 19,
        attribution: '&copy; <a href="http://www.openstreetmap.org/copyright">OpenStreetMap</a>'
    }).addTo(map.value);

    const content = `
                <div class="container">
                    <p>${props.latitude}, ${props.longitude}</p>
                </div>
                    `;

    marker.value = L.marker([props.latitude, props.longitude]).addTo(map.value);
    marker.value.bindPopup(content)
});

watch([() => props.latitude, () => props.longitude], ([newLat, newLon]) => {
    if (map.value) {
        const content = `
                <div class="container">
                    <p>${newLat}, ${newLon}</p>
                </div>
                    `;
        map.value.setView([newLat, newLon]);
        if (marker.value) {
            marker.value.setLatLng([newLat, newLon]);
            marker.value.bindPopup(content)
        } else if (props.showMarker) {
            marker.value = L.marker([newLat, newLon]).addTo(map.value);
            marker.value.bindPopup(content)
        }
    }
});

</script>

<style></style>
