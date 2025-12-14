# The "local" environment represents our local testings.
env "local" {
  #url = "postgres://postgres_user:postgres_password@host:port/db?sslmode=disable" 
  url = "" 
  migration {
    #dir = "atlas://project_name"
    dir = ""
  }
}