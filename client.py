#!/usr/bin/env python3

import requests
import random

for x in range(10):
    features = []
    for i in range(10):
        features.append(round(random.random(), 2))

    ploads = {'name':"Person #" + str(x), "features": features}

    r = requests.post('http://face-recognizer.dev/api/person/',json=ploads)

for x in range(1):
    features = []
    for x in range(10):
        features.append(random.random())

    ploads = {"features": features}

    r = requests.get('http://face-recognizer.dev/api/person/',json=ploads)
    print(r.text)
