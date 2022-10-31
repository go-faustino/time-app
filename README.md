# Simple time application

1. Create an application on Go/Python that serve HTTP requests at the next routes:

   - 8080:/
   - 8080:/health

   / - returns local time in New-York, Berlin, Tokyo in HTML format\
   /health - return status code 200, in JSON format

2. Dockerize your app.
3. Provision your infra using Terraform:

   - VPC
   - Subnets
   - SGs
   - Routing tables
   - EKS cluster
   - ECR

4. Create and deploy Helm3 chart for your application.
5. Ensure that you have an access to the application endpoint ext
6. Provide an external http link
7. Upload your code to the Github public repository and provide link to us (put your code to the app, terraform, helm folders accordingly)
