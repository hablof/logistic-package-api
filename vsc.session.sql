SELECT table_catalog, table_schema, table_name, privilege_type
FROM   information_schema.table_privileges
WHERE  grantee = 'docker';

-- GRANT ALL PRIVILEGES ON DATABASE "logistic_package_api" to docker;


-- Select * from "pg_users";