## Skeleton Service

##### Description

This source is example of base service go. 

##### How to Use

To create new service based on skeleton service follow this steps:

1. Clone this repo with `git clone https://git.lifewood.dev/scm/lr/skeleton-service.git your-service`

2. Remove entire git directory with `cd your-service && rm -rf .git`

3. Reinitialize git and bind to new repository 

   ```bash
   cd your-service
   git init
   git add --all
   git commit -m "Initial Commit"
   git remote add origin https://git.lifewood.dev/scm/lr/your-service.git
   git push -u origin HEAD:master
   ```

4. Happy code

##### Dependencies

This source depends on `common-service` install dependencies with  `go get -u git.lifewood.dev/common-service@0.0.1`

##### Help and Assistances

Contact me at husnan@dmk.co.id