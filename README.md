# My Go Project
This is a Go project linked to GitHub.


# Run docker image of postgres
docker run --name postgres-db -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=@Proxima_ps12@ -e POSTGRES_DB=deliveryapp -p 5432:5432 -d postgres
