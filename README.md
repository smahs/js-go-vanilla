js-go-vanilla
---

A vanilla Go and JS server client hybrid boilerplate, not based on any framework.

Setup
---
Install make, Go and Node. Use your distro's package manager for this.
Then install JS dependencies using `npm`.
 
```
npm install
```

Build
---
The `make` utility is used in this project.
```
make build
```

Backend
---
This runs only the backend.
```
make backend
```

Frontend
---
This runs only the frontend.
```
make frontend
```
Internally `npm` is used for running the frontend task. See `package.json` for more options.
