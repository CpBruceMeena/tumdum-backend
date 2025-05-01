#!/bin/bash

# Create uploads directory if it doesn't exist
mkdir -p uploads

# Function to generate a colored image with text
generate_image() {
    local output_file=$1
    local text=$2
    local bg_color=$3
    local text_color=$4
    local width=$5
    local height=$6

    magick -size ${width}x${height} xc:${bg_color} \
        -gravity center \
        -pointsize 40 \
        -fill ${text_color} \
        -annotate 0 "${text}" \
        "uploads/${output_file}"
}

# Generate restaurant logos (square images)
for i in {1..5}; do
    generate_image "restaurant_logo_${i}.jpg" "Restaurant ${i} Logo" "#FF5733" "white" 400 400
done

# Generate restaurant cover images (wide images)
for i in {1..5}; do
    generate_image "restaurant_cover_${i}.jpg" "Restaurant ${i} Cover" "#33FF57" "white" 1200 400
done

# Generate dish images (square images)
for i in {1..25}; do
    generate_image "dish_${i}.jpg" "Dish ${i}" "#3357FF" "white" 400 400
done

echo "Generated all test images in the uploads directory" 