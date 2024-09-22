#!/bin/bash

set -e

if [ -f /etc/debian_version ]; then
    echo "Detected Debian-based system. Installing dependencies..."

    sudo apt-get update
    sudo apt-get install -y \
        libgl1-mesa-dev \
        libglfw3-dev \
        libxcursor-dev \
        libxrandr-dev \
        libxinerama-dev \
        libxi-dev \
        libxxf86vm-dev

elif [ -f /etc/fedora-release ]; then
    echo "Detected Fedora-based system. Installing dependencies..."

    sudo dnf install -y \
        mesa-libGL-devel \
        glfw-devel \
        libXcursor-devel \
        libXrandr-devel \
        libXinerama-devel \
        libXi-devel \
        libXxf86vm-devel

else
    echo "Unsupported operating system. Please install the dependencies manually."
    exit 1
fi

echo "Dependencies installed successfully."
