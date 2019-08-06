package values

const Template = `
pullCredentials:
  # name of the credential secret to create on your cluster
  name: {{ .PullCredentials.Name }}
  username: {{ .PullCredentials.Username }}
  password: {{ .PulLCredentials.Password }}

# options default global ingress tls configuration
tls:
  default: []

# overrides for the consent-service subchart
consent-service:
  image:
    pullSecrets:
      - name: {{ .PullCredentials.Name }}

  # enables prometheus service monitoring
  serviceMonitor:
    enabled: {{ .ServiceMonitor.Enabled }}
  
  # expose consent-service using Ingress objects.
  ingress:
    enabled: {{ .Ingress.Enabled }}
    path: /
    hosts:
      - {{ .ConsentService.URI }}
    tls:
      - hosts:
        - {{ .ConsentService.URI }}
      # specify the secret containing the TLS private information
      # secretName: singularity-enterprise-localhost

  # expose consent-service using OpenShift Routes (only for OpenShift Clusters)
  route:
    enabled: {{ .Route.Enabled }} 

  env:
    # Populate these values with credentials from each identity provider
    {{- if .ConsentService.OAuth.Google.Enabled }}
    google_oauth2_client_id: {{ .ConsentService.OAuth.Google.ClientID }}
    google_oauth2_client_secret: {{ .ConsentService.OAuth.Google.ClientSecret }}
    {{- end }}
    hydra_client_secret: {{ .Hydra.ClientSecret }}
    service_uri: {{ .ConsentService.URI }}
    # name of an user (not yet created) that will eventually be coerced into being an administrator
    admin_user: {{ .ConsentService.AdminUser }}
  # mongodb connection infor for consent-service. Must match the values specified for mongodb if using mongodb on Kubernetes
  mongodb:
    mongodbUsername: {{ .MongoDB.Username }} 
    mongodbPassword: {{ .MongoDB.Password }}
    mongodbRootPassword: {{ .Mongodb.RootPassword }}
    mongodbDatabase: {{ .MongoDB.Database }}
    mongodbEndpoint: {{ .MongoDB.Endpoint }}

# override mongodb subchart values
mongodb:
  mongodbUsername: {{ .MongoDB.Username }} 
  mongodbPassword: {{ .MongoDB.Password }}
  mongodbRootPassword: {{ .Mongodb.RootPassword }}
  mongodbDatabase: {{ .MongoDB.Database }}

# override postgresql subchart values
postgresql:
  postgresqlUsername: {{ .Postgres.Username }}
  postgresqlPassword: {{ .Postgres.Password }}
  postgresqlDatabase: {{ .Postgres.Database }} 

# override hydra subchart values
hydra:
  postgresql:
    postgresqlUsername: {{ .Postgres.Username }}
    postgresqlPassword: {{ .Postgres.Password }}
    postgresqlDatabase: {{ .Postgres.Database }} 
    postgresqlEndpoint: {{ .Postgres.Endpoint }}  # defaults to {{ .Release.Name }}-postgresql
  config:
    system:
      secret:  {{ .Hydra.ClientSecret }}
      cookiesecret:  {{ .Hydra.CookieSecret }}

  # expose hydra using Ingress
  ingress:
    enabled: {{ .Ingress.Enabled }}
    path: /
    hosts:
      - {{ .Hydra.URI }}
    tls:
      - hosts:
        - {{ .Hydra.URI }}
      # secretName: singularity-enterprise-localhost
  
  # expose hydra using OpenShift Routes (only on OpenShift)
  route:
    enabled: {{ .Route.Enabled }}

  env:
    # URL pointing to hydra
    oauth2_issuer_url: {{ .Hydra.URI }}
    # URL pointing to the consent-service consent url
    oauth2_consent_url: {{ .Hydra.ConsentURL }} 
    # URL pointing to the consent-service login url
    oauth2_login_url: {{ .Hydra.LoginURL }} 
    # URL pointing to the frontend for logouts
    oauth2_logout_redirect_url: {{ .Frontend.URI }}
    # callbacks pointing to the frontend for OAuth callbacks
    callback_urls: {{ .Frontend.URI -}},{{- .Frontend.RevokeURI }}
    hydra_frontend_secret: {{ .Hydra.FrontendSecret }}
    hydra_consent_secret: {{ .Hydra.ConsentSecret }} 
  
# override token-service subchart values
token-service:
  # enables prometheus monitoring
  serviceMonitor:
    enabled: {{ .ServiceMonitor.Enabled }}

  # expose token-service with Ingress
  ingress:
    enabled: {{ .Ingress.Enabled }}
    hosts:
      - {{ .TokenService.URI }}
    tls:
      - hosts:
        - {{ .TokenService.URI }}
      # secretName: singularity-enterprise-localhost

  # expose token-service with OpenShift Routes (only on OpenShift)
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
    mongodbEndpoint: {{ .MongoDB.Endpoint }}
  rsaSecretName: {{ .TokenService.RSASecretName }} 
  # rsa_token:

# overrides for the Angular 2 frontend sub-chart
frontend:
  image:
    pullSecrets:
      - name: {{ .PullCredentials.Name }}
  
  # expose frontend with Ingress
  ingress:
    enabled: {{ .Ingress.Enabled }} 
    hosts:
      - {{ .Frontend.URI }}
    tls:
      - hosts:
        - {{ .Frontend.URI }}
      # secretName: singularity-enterprise-localhost

  # expose frontend with OpenShift Routes
  route:
    enabled: {{ .Route.Enabled }} 

  env:
    # uri to cloud library
    public_host_library: {{ .CloudLibrary.URI }}
    # uri to key service
    public_host_key_service: {{ .KeyService.URI }}
    # uri to remote build server
    public_host_build_service: {{ .RemoteBuildServer.URI }}
    # uri to consent service
    public_host_consent_service: {{ .ConsentService.URI }}
    # uri to token service
    public_host_token_service: {{ .TokenService.URI }}
    # uri to hydra
    public_host_hydra: {{ .Hydra.URI }}
    # uri to frontend
    public_host_front_end: {{ .Frontend.URI }}

# overrides for key-service subchart
key-service:
  image:
    pullSecrets: 
      - name: {{ .PullCredentials.Name }}

  mongodb:
    mongodbHost: {{ .MongoDB.Endpoint }} 
    mongodbUsername: {{ .MongoDB.Username }} 
    mongodbDatabase: {{ .MongoDB.Database }} 
  secrets:
    mongodb:
      mongodbPassword: {{ .MongoDB.Password }}
      mongodbRootPassword: {{ .Mongodb.RootPassword }}

  # enables prometheus service monitoring
  serviceMonitor:
    enabled: {{ .ServiceMonitor.Enabled }}
   
  # expose key-service using Ingress
  ingress:
    enabled: {{ .Ingress.Enabled }} 
    hosts:
      - {{ .KeyService.URI }}
    tls:
      - hosts:
        - {{ .KeyService.URI }}
      # secretName: singularity-enterprise-localhost

  # expose key-service with OpenShift Routes
  route:
    enabled: {{ .Route.Enabled }} 


cloud-library-server:
  image:
    pullSecrets:  
      - name: {{ .PullCredentials.Name }}

  # enable prometheus service monitoring
  serviceMonitor:
    enabled: {{ .ServiceMonitor.Enabled }}

  # expose cloud library server with Ingress
  ingress:
    enabled: {{ .Ingress.Enabled }}
    hosts:
      - {{ .CloudLibraryServer.URI }}
    tls:
      - hosts:
        - {{ .CloudLibraryServer.URI }} 
      # secretName: singularity-enterprise-localhost

  # expose cloud library server with OpenShift Routes
  route:
    enabled: {{ .Route.Enabled }}

  s3:
    # s3 bucket location. Value of "local" uses in cluster minio instance
    endpoint: {{ .S3.Endpoint }}
    bucket: {{ .S3.Bucket }}
    access_key: {{ .S3.AccessKey }}
    secret_key: {{ .S3.SecretKey }}

  mongodb:
    mongodbUsername: {{ .MongoDB.Username }}
    mongodbPassword: {{ .MongoDB.Password }} 
    mongodbRootPassword: {{ .MongoDB.RootPassword }}
    mongodbDatabase: {{ .MongoDB.Database }}
    mongodbEndpoint: {{ .MongoDB.Endpoint }}

  rabbitmq:
    username: {{ .RabbitMQ.Username }}
    password: {{ .RabbitMQ.Password }} 

# overrides for minio subchart
minio:
  accessKey: {{ .S3.AccessKey }}
  secretKey: {{ .S3.SecretKey }}

  ingress:
    enabled: {{ .Ingress.Enabled }}
    hosts:
      - {{ .Minio.URI }}
    tls:
      - hosts:
        - {{ Minio.URI }}
      # secretName: singularity-enterprise-localhost
  
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
    mongodbEndpoint: {{ .MongoDB.Endpoint }}
  s3:
    endpoint: {{ .S3.Endpoint }} 
    bucket: {{ .S3.Bucket }} 
    access_key: {{ .S3.AccessKey }}
    secret_key: {{ .S3.SecretKey }} 

cloud-library-cronjobs:
  purger:
    image:
      pullSecret: 
        - name: {{ .PullCredentials.Name }}
    s3:
      endpoint:  {{ .S3.Endpoint }}
      bucket: {{ .S3.Bucket }}
      access_key: {{ .S3.AccessKey }} 
      secret_key: {{ .S3.SecretKey }}
    mongodb:
      mongodbUsername: {{ .MongoDB.Username }} 
      mongodbPassword: {{ .MongoDB.Password }} 
      mongodbRootPassword: {{ .MongoDB.RootPassword }}
      mongodbDatabase: {{ .MongoDB.Database }} 
      mongodbEndpoint: {{ .MongoDB.Endpoint }}
  cleaner:
    image:
      # cloud-library_server container also includes 'purge'
      pullSecret:
        - name: {{ .PullCredentials.Name }}
    s3:
      endpoint:  {{ .S3.Endpoint }}
      bucket: {{ .S3.Bucket }}
      access_key: {{ .S3.AccessKey }} 
      secret_key: {{ .S3.SecretKey }}
    mongodb:
      mongodbUsername: {{ .MongoDB.Username }} 
      mongodbPassword: {{ .MongoDB.Password }} 
      mongodbRootPassword: {{ .MongoDB.RootPassword }}
      mongodbDatabase: {{ .MongoDB.Database }} 
      mongodbEndpoint: {{ .MongoDB.Endpoint }}

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
      # secretName: singularity-enterprise-localhost
  
  route:
    enabled: {{ .Route.Enabled }} 

  serviceMonitor:
    enabled: {{ .ServiceMonitor.Enabled }} 

  env:
    # auth_uri: "https://auth.lvh.me/token"
    # consent_uri: "https://auth.lvh.me/consent"
    # library_uri: "https://library.lvh.me"
    # manager_uri: "wss://build.lvh.me"
    # server_uri: "https://build.lvh.me"
    # web_key_uri: "https://hydra.lvh.me/.well-known/jwks.json"
  
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
    mongodbEndpoint: {{ .MongoDB.Endpoint }}

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
      # secretName: singularity-enterprise-localhost
  
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
    mongodbEndpoint: {{ .MongoDB.Endpoint }}
  env:
    # auth_uri: "https://token.lvh.me"
    # consent_uri: "https://auth.lvh.me"
    # library_uri: "https://library.lvh.me"
    # manager_uri: "https://manager.lvh.me"
    # server_uri: "https://build.lvh.me"
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
    mongodbEndpoint: {{ .MongoDB.Endpoint }}

remote-build-cronjobs:
  purger:
    mongodb:
        mongodbUsername: {{ .MongoDB.Username }}
        mongodbPassword: {{ .MongoDB.Password }}
        mongodbRootPassword: {{ .MongoDB.RootPassword }} 
        mongodbDatabase: {{ .MongoDB.Database }} 
        mongodbEndpoint: {{ .MongoDB.Endpoint }}

redis:
  password: {{ .Redis.Password }}
`
