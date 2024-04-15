<template>
    <nav class="navbar bg-dark navbar-dark navbar-expand-lg bg-body-tertiary">
        <div class="container">
            <a class="navbar-brand">NETFLIX</a>
            <button class="navbar-toggler" type="button" data-bs-toggle="collapse"
                data-bs-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false"
                aria-label="Toggle navigation">
                <span class="navbar-toggler-icon"></span>
            </button>
            <div class="collapse navbar-collapse" id="navbarSupportedContent">
                <ul class="navbar-nav me-auto mb-2 mb-lg-0">
                    <li class="nav-item">
                        <NuxtLink to="/" class="nav-link">Home</NuxtLink>
                    </li>
                    <li class="nav-item">
                        <NuxtLink to="/title" class="nav-link">List Movie/Show</NuxtLink>
                    </li>
                </ul>
            </div>
        </div>
    </nav>

    <slot />

</template>