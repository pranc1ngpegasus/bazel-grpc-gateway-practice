${ENV}:
  dialect: mysql
  datasource: ${DATABASE_USER}:${DATABASE_PASSWORD}@tcp(${DATABASE_HOST})/${DATABASE_NAME}?parseTime=true&charset=utf8mb4
  dir: ${DATABASE_MIGRATION_DIRECTORY}
