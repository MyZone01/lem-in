#   🐜 LEM-IN
## DESCRIPTION

This project is a simulation of an ant farm. It is written in Go and aims to find the quickest way to get a number of ants across a colony, from the start room to the end room.

## Project Overview

An ant farm is created with links and rooms. The ants are placed on one side and they need to find the exit. The goal is to bring them to the end room with as few moves as possible. The shortest path is not necessarily the simplest. 

The program displays the results in a specific format on the standard output:

```console
number_of_ants
the_rooms
the_links

Lx-y Lz-w Lr-o ...
```

Where `x, z, r` represent the ants numbers and `y, w, o` represent the rooms names.

## Installation

This project requires Go 1.16 or later. To install, clone this repository and build using Go:

```bash
git clone https://learn.zone01dakar.sn/git/papgueye/lem-in.git
cd lem-in
go build
```

## Usage

To use the program, pass the path of a text file containing the ant farm description as an argument. The format of the file should be as described in the Project Overview.

```bash
./lem-in path/to/your/file.txt
```

## Project Structure

This project makes use of several key data structures: Room, Tunnel, Ant, and AntFarm. 

- Room: Represents each room and contains the room name and its coordinates.
- Link: Represents each tunnel and contains pointers to the two rooms it connects.
- Ant: Represents each ant, containing an identifier and its current room.
- AntFarm: Represents the entire ant farm and contains a list of all ants, rooms, and links, and pointers to the start and end rooms.

Key functions:


##  AUTHORS
+   Pape Alassane Ba
+   Pape Gondia Gueye
+   Louis Sébastien Malack
+   Serigne Saliou Mbacké Mbaye

## LIVESHARE LINK
https://prod.liveshare.vsengsaas.visualstudio.com/join?50A8CA4313BB2D1694205170CDFCC8FF587F