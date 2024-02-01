# Use the official PHP 7.4 image as base
FROM php:8.1-fpm

# Install dependencies
RUN apt-get update && apt-get install -y \
	build-essential \
	locales \
	zip \
	vim \
	unzip \
	git \
	curl \
	libonig-dev \
	libzip-dev \
	libpng-dev \
	nginx \
	supervisor

# Clear cache
RUN apt-get clean && rm -rf /var/lib/apt/lists/*

# Install extensions
RUN docker-php-ext-install pdo_mysql mbstring zip exif pcntl
RUN docker-php-ext-install gd

# Set working directory
WORKDIR /var/www

# Remove the default nginx index page
RUN rm -rf /var/www/html

# Copy existing application directory
COPY . /var/www

# Change ownership of the application directory to www-data
RUN chown -R www-data:www-data /var/www
COPY nginx.conf /etc/nginx/nginx.conf

# Install composer
COPY --from=composer:1.10.20 /usr/bin/composer /usr/bin/composer

# Change current user to www
USER www-data

# Run composer install
RUN composer install --optimize-autoloader

# Supervisord
COPY supervisord.conf /etc/supervisor/conf.d/supervisord.conf

CMD ["supervisord"]
