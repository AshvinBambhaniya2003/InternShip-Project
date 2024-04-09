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
    },
    locations: {
        type: Array,
        required: true
    }
})

const map = ref();
const mapContainer = ref();
const marker = ref()
const markers = ref([]);
const paths = ref([]);

onMounted(() => {
    if (props.latitude) {
        map.value = L.map(mapContainer.value).setView([props.latitude, props.longitude], 13);
        const content = `
                <div class="container">
                    <p>${props.latitude}, ${props.longitude}</p>
                </div>
                    `;

        marker.value = L.marker([props.latitude, props.longitude]).addTo(map.value);
        marker.value.bindPopup(content)
    } else {
        map.value = L.map(mapContainer.value).setView([props.locations[0].latitude, props.locations[0].longitude], 13);
    }

    L.tileLayer('https://tile.openstreetmap.org/{z}/{x}/{y}.png', {
        maxZoom: 19,
        attribution: '&copy; <a href="http://www.openstreetmap.org/copyright">OpenStreetMap</a>'
    }).addTo(map.value);

    if (props.locations) {
        drawMarkers(props.locations);
        drawPaths(props.locations);
    }
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

watch(() => props.locations, (newLocations) => {
    drawMarkers(newLocations);
    drawPaths(newLocations);
});

const drawMarkers = (locations) => {
    map.value.setView([locations[0].latitude, locations[0].longitude]);
    markers.value.forEach(marker => marker.remove());
    markers.value = [];
    locations.forEach(location => {
        const marker = L.marker([location.latitude, location.longitude]).addTo(map.value);
        const content = `
                <div class="container">
                    <p>${location.latitude}, ${location.longitude}</p>
                </div>
                    `;
        marker.bindPopup(content)
        markers.value.push(marker);
    });
};

const drawPaths = (locations) => {
    paths.value.forEach(path => path.remove());
    paths.value = [];
    if (locations.length >= 2) {
        for (let i = 0; i < locations.length - 1; i++) {
            const latLng1 = [locations[i].latitude, locations[i].longitude];
            const latLng2 = [locations[i + 1].latitude, locations[i + 1].longitude];
            const path = L.polyline([latLng1, latLng2], { color: 'blue' }).addTo(map.value);
            paths.value.push(path);
        }
    }
};

</script>

<style></style>
