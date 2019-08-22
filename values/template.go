package values

const Template = `
pullCredentials:
  name: {{ .PullCredentials.Name }}
  username: {{ .PullCredentials.Username }}
  password: {{ .PullCredentials.Password }}

tls:
  default: []

consent-service:
  image:
    pullSecrets:
      - name: {{ .PullCredentials.Name }}

  serviceMonitor:
    enabled: {{ .ServiceMonitor.Enabled }}
  
  ingress:
    enabled: {{ .Ingress.Enabled }}
    path: /
    hosts:
      - {{ .ConsentService.URI }}
    tls:
      - hosts:
        - {{ .ConsentService.URI }}

  route:
    enabled: {{ .Route.Enabled }} 

  env:
    {{- if .ConsentService.OAuth.Google.Enabled }}
    google_oauth2_client_id: {{ .ConsentService.OAuth.Google.ClientID }}
    google_oauth2_client_secret: {{ .ConsentService.OAuth.Google.ClientSecret }}
    {{- end }}
    hydra_client_secret: {{ .Hydra.ClientSecret }}
    service_uri: {{ .ConsentService.URI }}
    admin_user: {{ .ConsentService.AdminUser }}
  mongodb:
    mongodbUsername: {{ .MongoDB.Username }} 
    mongodbPassword: {{ .MongoDB.Password }}
    mongodbRootPassword: {{ .MongoDB.RootPassword }}
    mongodbDatabase: {{ .MongoDB.Database }}
    mongodbEndpoint: {{ .MongoDB.Endpoint | quote }}

mongodb:
  mongodbUsername: {{ .MongoDB.Username }} 
  mongodbPassword: {{ .MongoDB.Password }}
  mongodbRootPassword: {{ .MongoDB.RootPassword }}
  mongodbDatabase: {{ .MongoDB.Database }}

postgresql:
  postgresqlUsername: {{ .Postgres.Username }}
  postgresqlPassword: {{ .Postgres.Password }}
  postgresqlDatabase: {{ .Postgres.Database }} 

hydra:
  postgresql:
    postgresqlUsername: {{ .Postgres.Username }}
    postgresqlPassword: {{ .Postgres.Password }}
    postgresqlDatabase: {{ .Postgres.Database }} 
    postgresqlEndpoint: {{ .Postgres.Endpoint | quote }} 
  config:
    system:
      secret: {{ .Hydra.ClientSecret }}
      cookiesecret: {{ .Hydra.CookieSecret }}

  ingress:
    enabled: {{ .Ingress.Enabled }}
    path: /
    hosts:
      - {{ .Hydra.URI }}
    tls:
      - hosts:
        - {{ .Hydra.URI }}
  
  route:
    enabled: {{ .Route.Enabled }}

  env:
    oauth2_issuer_url: {{ .Hydra.URI }}
    oauth2_consent_url: {{ .ConsentService.URI }}/v1/consent
    oauth2_login_url: {{ .ConsentService.URI }}/v1/login
    oauth2_logout_redirect_url: {{ .Frontend.URI }}
    callback_urls: {{ .Frontend.URI -}},{{- .Frontend.RevokeURI }}
    hydra_frontend_secret: {{ .Hydra.FrontendSecret }}
    hydra_consent_secret: {{ .Hydra.ConsentSecret }} 
  
token-service:
  serviceMonitor:
    enabled: {{ .ServiceMonitor.Enabled }}

  ingress:
    enabled: {{ .Ingress.Enabled }}
    hosts:
      - {{ .TokenService.URI }}
    tls:
      - hosts:
        - {{ .TokenService.URI }}

  route: 
    enabled: {{ .Route.Enabled }} 

  env:
    service_uri: {{ .TokenService.URI }}
  image:
    pullSecrets:
      - name: {{ .PullCredentials.Name }}
  mongodb:
    mongodbUsername: {{ .MongoDB.Username }}
    mongodbPassword: {{ .MongoDB.Password }}
    mongodbRootPassword: {{ .MongoDB.RootPassword }} 
    mongodbDatabase: {{ .MongoDB.Database }} 
    mongodbEndpoint: {{ .MongoDB.Endpoint | quote }}
  rsaSecretName: {{ .TokenService.RSASecretName }} 

frontend:
  image:
    pullSecrets:
      - name: {{ .PullCredentials.Name }}
  
  ingress:
    enabled: {{ .Ingress.Enabled }} 
    hosts:
      - {{ .Frontend.URI }}
    tls:
      - hosts:
        - {{ .Frontend.URI }}

  route:
    enabled: {{ .Route.Enabled }} 

  env:
    public_host_library: {{ .CloudLibraryServer.URI }}
    public_host_key_service: {{ .KeyService.URI }}
    public_host_build_service: {{ .RemoteBuildServer.URI }} 
    public_host_consent_service: {{ .ConsentService.URI }}
    public_host_token_service: {{ .TokenService.URI }}
    public_host_hydra: {{ .Hydra.URI }}
    public_host_front_end: {{ .Frontend.URI }}

key-service:
  image:
    pullSecrets: 
      - name: {{ .PullCredentials.Name }}

  mongodb:
    mongodbHost: {{ .MongoDB.Endpoint | quote }} 
    mongodbUsername: {{ .MongoDB.Username }} 
    mongodbDatabase: {{ .MongoDB.Database }} 
  secrets:
    mongodb:
      mongodbPassword: {{ .MongoDB.Password }}
      mongodbRootPassword: {{ .MongoDB.RootPassword }}

  serviceMonitor:
    enabled: {{ .ServiceMonitor.Enabled }}
   
  ingress:
    enabled: {{ .Ingress.Enabled }} 
    hosts:
      - {{ .KeyService.URI }}
    tls:
      - hosts:
        - {{ .KeyService.URI }}

  route:
    enabled: {{ .Route.Enabled }} 


cloud-library-server:
  image:
    pullSecrets:  
      - name: {{ .PullCredentials.Name }}

  serviceMonitor:
    enabled: {{ .ServiceMonitor.Enabled }}

  ingress:
    enabled: {{ .Ingress.Enabled }}
    hosts:
      - {{ .CloudLibraryServer.URI }}
    tls:
      - hosts:
        - {{ .CloudLibraryServer.URI }} 

  route:
    enabled: {{ .Route.Enabled }}

  s3:
    endpoint: {{ .S3.Endpoint | quote }}
    bucket: {{ .S3.Bucket }}
    access_key: {{ .S3.AccessKey }}
    secret_key: {{ .S3.SecretKey }}

  mongodb:
    mongodbUsername: {{ .MongoDB.Username }}
    mongodbPassword: {{ .MongoDB.Password }} 
    mongodbRootPassword: {{ .MongoDB.RootPassword }}
    mongodbDatabase: {{ .MongoDB.Database }}
    mongodbEndpoint: {{ .MongoDB.Endpoint | quote }}

  rabbitmq:
    username: {{ .RabbitMQ.Username }}
    password: {{ .RabbitMQ.Password }} 

minio:
  accessKey: {{ .S3.AccessKey }}
  secretKey: {{ .S3.SecretKey }}

  ingress:
    enabled: {{ .Ingress.Enabled }}
    hosts:
      - {{ .Minio.URI }}
    tls:
      - hosts:
        - {{ .Minio.URI }}
  
  route:
    enabled: {{ .Route.Enabled }} 

rabbitmq:
  rabbitmq:
    username: {{ .RabbitMQ.Username }}
    password: {{ .RabbitMQ.Password }}

cloud-library-pam:
  rabbitmq:
    password: {{ .RabbitMQ.Password }}
    username: {{ .RabbitMQ.Username }}
    pullSecrets:
      - name: {{ .PullCredentials.Name }}
  mongodb:
    mongodbUsername: {{ .MongoDB.Username }}
    mongodbPassword: {{ .MongoDB.Password }}
    mongodbRootPassword: {{ .MongoDB.RootPassword }} 
    mongodbDatabase: {{ .MongoDB.Database }} 
    mongodbEndpoint: {{ .MongoDB.Endpoint | quote }}
  s3:
    endpoint: {{ .S3.Endpoint | quote }} 
    bucket: {{ .S3.Bucket }} 
    access_key: {{ .S3.AccessKey }}
    secret_key: {{ .S3.SecretKey }} 

cloud-library-cronjobs:
  purger:
    image:
      pullSecrets: 
        - name: {{ .PullCredentials.Name }}
    s3:
      endpoint:  {{ .S3.Endpoint | quote }}
      bucket: {{ .S3.Bucket }}
      access_key: {{ .S3.AccessKey }} 
      secret_key: {{ .S3.SecretKey }}
    mongodb:
      mongodbUsername: {{ .MongoDB.Username }} 
      mongodbPassword: {{ .MongoDB.Password }} 
      mongodbRootPassword: {{ .MongoDB.RootPassword }}
      mongodbDatabase: {{ .MongoDB.Database }} 
      mongodbEndpoint: {{ .MongoDB.Endpoint | quote }}
  cleaner:
    image:
      pullSecret:
        - name: {{ .PullCredentials.Name }}
    s3:
      endpoint:  {{ .S3.Endpoint | quote }}
      bucket: {{ .S3.Bucket }}
      access_key: {{ .S3.AccessKey }} 
      secret_key: {{ .S3.SecretKey }}
    mongodb:
      mongodbUsername: {{ .MongoDB.Username }} 
      mongodbPassword: {{ .MongoDB.Password }} 
      mongodbRootPassword: {{ .MongoDB.RootPassword }}
      mongodbDatabase: {{ .MongoDB.Database }} 
      mongodbEndpoint: {{ .MongoDB.Endpoint | quote }}

remote-build-server:
  image:
    pullSecrets: 
      - name: {{ .PullCredentials.Name }}

  ingress:
    enabled: {{ .Ingress.Enabled }} 
    hosts:
      - {{ .RemoteBuildServer.URI }} 
    tls:
      - hosts:
        - {{ .RemoteBuildServer.URI }}
  
  route:
    enabled: {{ .Route.Enabled }} 

  serviceMonitor:
    enabled: {{ .ServiceMonitor.Enabled }} 
  
  rabbitmq:
    username: {{ .RabbitMQ.Username }} 
    password: {{ .RabbitMQ.Password }}
  
  redis:
    password: {{ .Redis.Password }}

  mongodb:
    mongodbUsername: {{ .MongoDB.Username }}
    mongodbPassword: {{ .MongoDB.Password }} 
    mongodbRootPassword: {{ .MongoDB.RootPassword }} 
    mongodbDatabase: {{ .MongoDB.Database }}
    mongodbEndpoint: {{ .MongoDB.Endpoint | quote }}

remote-build-manager:
  serviceMonitor:
    enabled: {{ .ServiceMonitor.Enabled }}
  rabbitmq:
    username: {{ .RabbitMQ.Username }}
    password: {{ .RabbitMQ.Password }} 

  ingress:
    enabled: {{ .Ingress.Enabled }}
    hosts:
      - {{ .RemoteBuildManager.URI }}
    tls:
      - hosts:
        - {{ .RemoteBuildManager.URI }}
  
  route:
    enabled: {{ .Route.Enabled }}

  image:
    pullSecret:
      - name: {{ .PullCredentials.Name }}
  mongodb:
    mongodbUsername: {{ .MongoDB.Username }}
    mongodbPassword: {{ .MongoDB.Password }} 
    mongodbRootPassword: {{ .MongoDB.RootPassword }}
    mongodbDatabase: {{ .MongoDB.Database }} 
    mongodbEndpoint: {{ .MongoDB.Endpoint | quote }}
  env:
    kubevirt_image_pull_secret: {{ .PullCredentials.Name }}
  rsaSecretName: {{ .TokenService.RSASecretName }} 

remote-build-jim:
  serviceMonitor:
    enabled: {{ .ServiceMonitor.Enabled }}
  rabbitmq:
    username: {{ .RabbitMQ.Username }}
    password: {{ .RabbitMQ.Password }} 
  image:
    pullSecrets:
      - name: {{ .PullCredentials.Name }}
  mongodb:
    mongodbUsername: {{ .MongoDB.Username }}
    mongodbPassword: {{ .MongoDB.Password }}
    mongodbRootPassword: {{ .MongoDB.RootPassword }} 
    mongodbDatabase: {{ .MongoDB.Database }} 
    mongodbEndpoint: {{ .MongoDB.Endpoint | quote }}

remote-build-cronjobs:
  purger:
    mongodb:
        mongodbUsername: {{ .MongoDB.Username }}
        mongodbPassword: {{ .MongoDB.Password }}
        mongodbRootPassword: {{ .MongoDB.RootPassword }} 
        mongodbDatabase: {{ .MongoDB.Database }} 
        mongodbEndpoint: {{ .MongoDB.Endpoint | quote }}

redis:
  password: {{ .Redis.Password }}
`
