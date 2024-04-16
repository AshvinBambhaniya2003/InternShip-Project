<template>
    <form v-on:submit.prevent="handleSubmit" class="row g-3">

        <div class="col-md-6 mb-2">
            <label for="person_id" class="form-label">Person Id</label>
            <input type="number" v-model="credit.person_id" class="form-control" id="person_id" name="person_id"
                required>
        </div>

        <div class="col-md-6 mb-2">
            <label for="name" class="form-label">Name</label>
            <input type="text" v-model="credit.name" class="form-control" id="name" name="name" required>
        </div>
        <div class="col-md-6 mb-2">
            <label for="character" class="form-label">Character</label>
            <input type="text" v-model="credit.character" class="form-control" id="character" name="character" required>
        </div>

        <div class="col-md-6 mb-2">
            <label for="role" class="form-label">Role</label>
            <select class="form-select " v-model="credit.role" id="role" name="role" required>
                <option value="">Select role</option>
                <option value="DIRECTOR">DIRECTOR</option>
                <option value="ACTOR">ACTOR</option>
            </select>
        </div>

        <div class="col-md-12 mb-2">
            <button type="submit" class="btn btn-primary">Submit</button>
        </div>
    </form>
</template>

<script setup>

// Define reactive reference for title
const credit = reactive({})

// Define props
const props = defineProps({
    credit: {
        type: Object
    }
})

// Assign titleData to title if available
if (props.credit) {
    Object.assign(credit, props.credit)
}

// Define emits for editCredit and addCredit
const emit = defineEmits(['submitForm'])

// Handle form submission, Emit appropriate event 
const handleSubmit = () => {
    emit('submitForm', credit)
}

</script>