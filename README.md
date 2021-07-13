# face-recognizer

This face recognition service get features (fixed size vector of 256) of a face which represent attruibtes of
the face.
The service returns as an output 3 (configurable number) best matches persons that simillar to this given
vector.
The face recognition is being done by comparing a single vector to all other exists vectors in the database
by dotproduct action.

## Getting Started

### Installing

* Add the following line to your hosts file:
  127.0.0.1 face-recognizer.dev
* cd into your projects folder
* git clone https://github.com/levikobi/face-recognizer.git

### Executing program

* cd /face-recognizer
* run the following command in the terminal
```
$ skaffold dev
```
* in a second terminal run the client's script
```
$ ./client.py
```

## Author

Kobi Levi
https://www.linkedin.com/in/levikobi/
