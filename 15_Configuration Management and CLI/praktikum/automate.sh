#!/bin/bash

# Check the number of arguments
if [ $# -ne 3 ]; then
    echo "Usage: ./automate <folder_name> <username_facebook> <username_linkedin>"
    exit 1
fi

# Main folder name
NOW=$(date +"%c")
folder_name="$1 at $NOW"

# Create the folder structure
mkdir -p "$folder_name/about_me/personal"
mkdir -p "$folder_name/about_me/professional"
mkdir -p "$folder_name/my_friends"
mkdir -p "$folder_name/my_system_info"

# Create facebook.txt and linkedin.txt files
echo "https://facebook.com/$2" > "$folder_name/about_me/personal/facebook.txt"
echo "https://linkedin.com/in/$3" > "$folder_name/about_me/professional/linkedin.txt"

# Fetch the list of friends from the GitHub Gist link and save it to list_of_my_friends.txt
curl -sL "https://gist.githubusercontent.com/tegarimansyah/e91f335753ab2c7fb12815779677e914/raw/94864388379fecee450fde26e3e73bfb2bcda194/list%2520of%2520my%2520friends.txt" > "$folder_name/my_friends/list_of_my_friends.txt"

# Create about_this_laptop.txt file
about_this_laptop="Nama User: $(whoami)\n$(uname -a)"
echo -e "$about_this_laptop" > "$folder_name/my_system_info/about_this_laptop.txt"

# Create internet_connection.txt file (example content)
ping -c 3 google.com > "$folder_name/my_system_info/internet_connection.txt"

echo "Folder and file structure has been successfully created in $folder_name"
